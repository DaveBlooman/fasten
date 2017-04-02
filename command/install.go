package command

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/DaveBlooman/fasten/appmeta"
	"github.com/DaveBlooman/fasten/connect"
	"github.com/DaveBlooman/fasten/files"
	"github.com/DaveBlooman/fasten/languages"
	"github.com/DaveBlooman/fasten/output"
	_ "github.com/jteeuwen/go-bindata"
)

type InstallCommand struct {
	Meta
}

func (c *InstallCommand) Run(args []string) int {

	var appStack appmeta.AppStack

	fastenFile, err := ioutil.ReadFile("fasten.yaml")
	if err != nil {
		output.Error("opening config file")
	}

	err = yaml.Unmarshal([]byte(fastenFile), &appStack)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	msg := map[string]string{}

	for i, app := range appStack.Applications {
		setupFile, err := files.Asset(fmt.Sprintf("libraries/%s/%s/%s.sh", appStack.OS, app.Lang, app.Lang))
		if err != nil {
			output.Error("language not found")
		}

		installFile, err := files.Asset(fmt.Sprintf("libraries/%s/%s/install.yaml", appStack.OS, app.Lang))
		if err != nil {
			output.Error("Installation file cannot be loaded")
		}

		var installCommand languages.Install
		err = yaml.Unmarshal([]byte(installFile), &installCommand)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		fastenSSH := connect.SSHClient{
			Host: appStack.IP,
			Port: 22,
			Key:  appStack.KeyPair,
			User: "ubuntu",
		}

		err = fastenSSH.SSHConfig()
		if err != nil {
			output.Standard("Error setting up SSH")
		}

		appDestination := appStack.InstallDir + "/" + app.Lang

		err = fastenSSH.Run(setupFile, app.Lang, appDestination, installCommand)
		if err != nil {
			return 1
		}

		msg["Language "+strconv.Itoa(i+1)] = app.Lang

	}

	msg["Operating System"] = appStack.OS
	msg["Server Address"] = appStack.IP

	output.Banner("Installation Complete", msg)

	return 0
}

func (c *InstallCommand) Synopsis() string {
	return ""
}

func (c *InstallCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
