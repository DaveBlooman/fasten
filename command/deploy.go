package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/DaveBlooman/fasten/appmeta"
	"github.com/DaveBlooman/fasten/connect"
	"github.com/DaveBlooman/fasten/languages"
	"github.com/DaveBlooman/fasten/output"

	yaml "gopkg.in/yaml.v2"
)

type DeployCommand struct {
	Meta
}

func (c *DeployCommand) Run(args []string) int {

	var appStack appmeta.AppStack

	fastenFile, err := ioutil.ReadFile("fasten.yaml")
	if err != nil {
		output.Error("opening config file")
	}

	err = yaml.Unmarshal([]byte(fastenFile), &appStack)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fastenSSH := connect.SSHClient{
		Host: appStack.IP,
		Port: 22,
		Key:  appStack.KeyPair,
	}

	err = fastenSSH.SSHConfig()
	if err != nil {
		output.Standard("Error setting up SSH")
	}

	for _, app := range appStack.Applications {

		installFile, err := Asset(fmt.Sprintf("libraries/%s/%s/install.yaml", appStack.OS, app.Lang))
		if err != nil {
			output.Error("Installation file cannot be loaded")
		}

		var installCommand languages.Install
		err = yaml.Unmarshal([]byte(installFile), &installCommand)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		appDestination := appStack.InstallDir + "/" + app.Name

		err = fastenSSH.MakeDir(appDestination)
		if err != nil {
			fmt.Println(err)
		}

		err = copyFiles(app, appStack, appDestination, fastenSSH)
		if err != nil {
			fmt.Println(err)
		}

		err = fastenSSH.InstallDeps(installCommand, app.Lang, appDestination)
		if err != nil {
			output.Error("Unable to install application dependecies")
		}

		err = fastenSSH.StartApplication(app, appStack, appDestination)
		if err != nil {
			output.Error("Unable to start/stop application")
		}

		output.Success("Application: " + app.Name + " has been deployed successfully")
	}

	return 0
}

func (c *DeployCommand) Synopsis() string {
	return ""
}

func (c *DeployCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}

func copyFiles(app appmeta.Application, stack appmeta.AppStack, appDestination string, connection connect.SSHClient) error {

	fileList, _ := getFiles(app.Path)
	//
	// for _, dir := range dirlist {
	// 	dirName := strings.SplitAfter(dir, app.Path)[1]
	// 	err := connection.MakeDir(appDestination + dirName)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }

	for _, file := range fileList {
		// data, err := ioutil.ReadFile(file)

		fileName := strings.SplitAfter(file, app.Path)[1]

		if strings.HasPrefix(fileName, "/") {
			fileName = strings.TrimPrefix(fileName, "/")

			fmt.Println(fileName)
			fmt.Println(strings.Replace(fileName, app.Path, "", -1))
		}

		// if err != nil {
		// 	return err
		// }
		//
		// fileCopy := connection.CopyFile(data, fileName, appDestination)
		// if fileCopy != nil {
		// 	return fileCopy
		// }
	}

	return nil
}

func getFiles(filesDir string) ([]string, []string) {

	fileList := []string{}
	dirList := []string{}

	err := filepath.Walk(filesDir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatal("Error:", err)
		}
		if f.IsDir() {
			dirList = append(dirList, path)
			return nil
		}
		fileList = append(fileList, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	return fileList, dirList
}
