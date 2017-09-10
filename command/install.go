package command

import (
	"fmt"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/DaveBlooman/fasten/config"
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
	fastenConfig, err := config.LoadFastenFile()
	if err != nil {
		output.Error("opening config file")
		return 1
	}

	msg := map[string]string{}

	for i, app := range fastenConfig.Applications {
		setupFile, err := files.Asset(fmt.Sprintf("libraries/%s/%s/%s.sh", fastenConfig.OS, app.Lang, app.Lang))
		if err != nil {
			output.Error(fmt.Sprintf("libraries/%s/%s/%s.sh", fastenConfig.OS, app.Lang, app.Lang))
			output.Error("language not found")
		}

		installFile, err := files.Asset(fmt.Sprintf("libraries/%s/%s/install.yaml", fastenConfig.OS, app.Lang))
		if err != nil {
			output.Error("Installation file cannot be loaded")
		}

		var installCommand languages.Install
		err = yaml.Unmarshal([]byte(installFile), &installCommand)
		if err != nil {
			output.Error("Installation file cannot be loaded " + err.Error())
		}

		fastenSSH := connect.SSHClient{
			Host: fastenConfig.IP,
			Port: 22,
			Key:  fastenConfig.KeyPair,
			User: fastenConfig.SSHUser,
		}

		err = fastenSSH.SSHConfig()
		if err != nil {
			output.Standard("Error setting up SSH")
		}

		appDestination := fastenConfig.InstallDir + "/" + app.Lang

		err = fastenSSH.Run(setupFile, app.Lang, appDestination, installCommand)
		if err != nil {
			return 1
		}

		msg["Language "+strconv.Itoa(i+1)] = app.Lang

	}

	msg["Operating System"] = fastenConfig.OS
	msg["Server Address"] = fastenConfig.IP

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
