package command

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/DaveBlooman/fasten/appmeta"
	"github.com/DaveBlooman/fasten/connect"
	"github.com/DaveBlooman/fasten/output"
	yaml "gopkg.in/yaml.v1"
)

type StopCommand struct {
	Meta
}

func (c *StopCommand) Run(args []string) int {

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
		User: "ubuntu",
	}

	err = fastenSSH.SSHConfig()
	if err != nil {
		output.Standard("Error setting up SSH")
	}
	for _, service := range appStack.Applications {
		fastenSSH.StopService(service.Lang, appStack.InstallDir+"/"+service.Name)
	}

	return 0

}

func (c *StopCommand) Synopsis() string {
	return ""
}

func (c *StopCommand) Help() string {
	helpText := `Shows Information about deployment

`
	return strings.TrimSpace(helpText)
}
