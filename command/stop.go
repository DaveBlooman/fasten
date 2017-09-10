package command

import (
	"strings"

	"github.com/DaveBlooman/fasten/config"
	"github.com/DaveBlooman/fasten/connect"
	"github.com/DaveBlooman/fasten/output"
)

type StopCommand struct {
	Meta
}

func (c *StopCommand) Run(args []string) int {
	fastenConfig, err := config.LoadFastenFile()
	if err != nil {
		output.Error("opening config file")
		return 1
	}

	fastenSSH, err := connect.SetupSSH(fastenConfig)
	if err != nil {
		output.Error("Unable to setup SSH")
		return 1
	}

	for _, service := range fastenConfig.Applications {
		fastenSSH.StopService(service.Lang, fastenConfig.InstallDir+"/"+service.Name)
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
