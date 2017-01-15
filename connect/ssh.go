package connect

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/DaveBlooman/fasten/appmeta"
	"github.com/DaveBlooman/fasten/languages"
	"github.com/DaveBlooman/fasten/output"

	"golang.org/x/crypto/ssh"
)

type SSHClient struct {
	Config     *ssh.ClientConfig
	Host       string
	Port       int
	Key        string
	Connection *ssh.Client
}

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

func (s *SSHClient) SSHConfig() error {

	s.Config = &ssh.ClientConfig{
		User: "ec2-user",
		Auth: []ssh.AuthMethod{
			PublicKeyFile(s.Key),
		},
	}

	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.Host, s.Port), s.Config)
	if err != nil {
		fmt.Printf("Failed to dial: %s", err)
		return err
	}

	s.Connection = connection

	return nil
}

func (s *SSHClient) Run(script []byte, language, appstack string, installCommand languages.Install) error {
	fileCopy := s.CopyFile(script, language+"_install.sh", "/tmp")
	if fileCopy != nil {
		return fileCopy
	}

	resp := installSoftware(s.Connection, language, installCommand)
	if resp != nil {
		return resp
	}

	return nil

}

func (s *SSHClient) CopyFile(contentsBytes []byte, filename, destination string) error {
	session, err := s.Connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return err
	}

	fmt.Println(destination)
	fmt.Println(filename)

	defer session.Close()

	session.Stderr = os.Stderr
	session.Stdout = os.Stdout

	contents := string(contentsBytes)

	go func() {
		w, err := session.StdinPipe()
		if err != nil {
			fmt.Println("err")
		}
		defer w.Close()
		fmt.Fprintln(w, "C0644", len(contents), filename)
		fmt.Fprint(w, contents)
		fmt.Fprint(w, "\x00")
	}()
	if err := session.Run("/usr/bin/scp -trp " + destination + "/" + filename); err != nil {
		output.Error("Error copying installation files to server")
	}
	return nil
}

func installSoftware(connection *ssh.Client, language string, install languages.Install) error {
	session, err := connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return err
	}

	installFile := language + "_install.sh"
	time := time.Now()

	software := map[string]string{"Language": language, "Install File": installFile, "time": time.String()}
	output.Banner("Installation...", software)

	// ** Implement verbose flag here
	// session.Stderr = os.Stderr
	// session.Stdout = os.Stdout

	err = session.Run("sudo -E sh /tmp/" + installFile)

	if err != nil {
		output.Error("Software Installation Failed")
		return err
	}

	session.Close()

	session2, err := connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return err
	}

	err = session2.Run(install.PreCommand)

	if err != nil {
		fmt.Println(err.Error())
		output.Error("Software Pre-Installation Failed")
		return err
	}
	session2.Close()

	output.Standard("Successfully installed " + language + "\n\n")

	return nil
}

func (s *SSHClient) MakeDir(destination string) error {

	session, err := s.Connection.NewSession()
	if err != nil {
		fmt.Println(err)
	}

	session.Stderr = os.Stderr
	err = session.Run("mkdir -p " + destination)
	if err != nil {
		output.Standard("unable to make destination")
		return err
	}

	return nil
}

func (s *SSHClient) InstallDeps(install languages.Install, language, path string) error {
	session, err := s.Connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return err
	}

	// ** Implement verbose flag here
	// session.Stderr = os.Stderr
	// session.Stdout = os.Stdout
	output.Standard("Installing dependencies, this may take some time...")

	defer session.Close()

	err = session.Run("cd " + path + " && " + install.Command)

	if err != nil {
		output.Error("Dependecies Installation Failed")
		return err
	}

	return nil
}

func (s *SSHClient) StartApplication(app appmeta.Application, stack appmeta.AppStack, path string) error {

	checkSession, err := s.Connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return err
	}

	err = checkSession.Run(fmt.Sprintf("if [[ -a \"%s/%s.pid\" ]]; then exit 1; else exit 0; fi", path, app.Lang))
	checkSession.Close()
	if err != nil {

		stopSession, err := s.Connection.NewSession()
		if err != nil {
			fmt.Printf("Failed to create session: %s", err)
			return err
		}

		// stopSession.Stdout = os.Stdout
		// stopSession.Stderr = os.Stderr

		err = stopSession.Run(fmt.Sprintf("cd %s ; cat %s.pid | xargs kill", path, app.Lang))

		if err != nil {
			output.Error("Failed to stop application")
			return err
		}

		output.Standard("Stopping Application")

		// Give time for application to stop
		time.Sleep(3 * time.Second)

		stopSession.Close()
	}

	if app.PreCommand != "" {
		preSession, err := s.Connection.NewSession()
		if err != nil {
			output.Error("Failed to create SSH session")
			return err
		}

		err = preSession.Run(fmt.Sprintf("cd %s && %s", path, app.PreCommand))
		if err != nil {
			output.Error("Pre-command failed to run")
			return err
		}

		preSession.Close()
	}

	startSession, err := s.Connection.NewSession()
	if err != nil {
		output.Error("Failed to create SSH session")
		return err
	}

	err = startSession.Run(fmt.Sprintf("cd %s ; nohup %s >/dev/null 2>&1 >> app.log & echo $! > %s/%s.pid", path, app.RunCommand, path, app.Lang))

	if err != nil {
		output.Error("Application start failed")
		return err
	}

	startSession.Close()

	output.Standard("Starting Application")

	return nil
}
