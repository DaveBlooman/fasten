package connect

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"
)

type SSHCommand struct {
	Path  string
	Env   []string
	Stdin io.Reader
	// Stdout io.Writer
	Stderr io.Writer
}

type SSHClient struct {
	Config *ssh.ClientConfig
	Host   string
	Port   int
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

func Run(command, key string) error {
	f, err := os.Open(command)
	if err != nil {
		return err
	}

	sshConfig := &ssh.ClientConfig{
		User: "ec2-user",
		Auth: []ssh.AuthMethod{
			PublicKeyFile(key),
		},
	}

	client := &SSHClient{
		Config: sshConfig,
		Host:   "52.210.250.28",
		Port:   22,
	}

	cmd := &SSHCommand{
		Path: "sudo -E sh /tmp/install.sh",
	}

	fmt.Printf("Running command: %s\n", cmd.Path)

	connection, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", client.Host, client.Port), client.Config)
	if err != nil {
		fmt.Printf("Failed to dial: %s", err)
		return err
	}

	fileCopy := copyFile(f, connection)
	if fileCopy != nil {
		return fileCopy
	}

	resp := installSoftware(connection)
	if resp != nil {
		return err
	}

	return nil

}

func copyFile(fileReader io.Reader, connection *ssh.Client) error {

	session, err := connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return err
	}

	defer session.Close()

	contents_bytes, _ := ioutil.ReadAll(fileReader)
	contents := string(contents_bytes)

	go func() {
		w, err := session.StdinPipe()
		if err != nil {
			fmt.Println("err")
		}
		defer w.Close()
		fmt.Fprintln(w, "C0644", len(contents), "install.sh")
		fmt.Fprint(w, contents)
		fmt.Fprint(w, "\x00")
	}()
	if err := session.Run("/usr/bin/scp -tr /tmp"); err != nil {
		panic("Failed to run: " + err.Error())
	}
	return nil
}

func installSoftware(connection *ssh.Client) error {
	session, err := connection.NewSession()
	if err != nil {
		fmt.Printf("Failed to create session: %s", err)
		return err
	}

	defer session.Close()
	output, err := session.Output("sudo -E sh /tmp/install.sh")
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(output))
		return err
	}

	fmt.Println(string(output))

	return nil
}
