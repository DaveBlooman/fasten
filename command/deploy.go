package command

import (
	"fmt"
	"strings"

	"github.com/DaveBlooman/fasten/config"
	"github.com/DaveBlooman/fasten/connect"
	"github.com/DaveBlooman/fasten/files"
	"github.com/DaveBlooman/fasten/fileutils"
	"github.com/DaveBlooman/fasten/languages"
	"github.com/DaveBlooman/fasten/output"

	yaml "gopkg.in/yaml.v2"
)

type DeployCommand struct {
	Meta
}

func (c *DeployCommand) Run(args []string) int {
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

	for _, app := range fastenConfig.Applications {

		installFile, err := files.Asset(fmt.Sprintf("libraries/%s/%s/install.yaml", fastenConfig.OS, app.Lang))
		if err != nil {
			output.Error("Installation file cannot be loaded")
		}

		var installCommand languages.Install
		err = yaml.Unmarshal([]byte(installFile), &installCommand)
		if err != nil {
			output.Error(err.Error())
		}

		appDestination := fastenConfig.InstallDir + "/" + app.Name

		err = fastenSSH.MakeDir(appDestination)
		if err != nil {
			output.Error(err.Error())
		}

		err = fileutils.BuildLocal(installCommand.BuildLocal, app)
		if err != nil {
			output.Error("Unable to build local components")
		}

		err = fastenSSH.StopService(app.Lang, appDestination)
		if err != nil {
			output.Error("Unable to stop service")
		}

		err = fileutils.CopyFiles(app, fastenConfig, appDestination, fastenSSH)
		if err != nil {
			output.Standard("Error copying files to the server")
			output.Error(err.Error())
		}

		err = fastenSSH.InstallDeps(installCommand, app.Lang, appDestination)
		if err != nil {
			output.Error("Unable to install application dependecies")
		}

		err = fastenSSH.StartApplication(app, fastenConfig, appDestination)
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
