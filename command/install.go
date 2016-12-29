package command

import (
	"strings"

	"github.com/DaveBlooman/fasten/connect"
)

type InstallCommand struct {
	Meta
}

func (c *InstallCommand) Run(args []string) int {
	err := connect.Run("/Users/dblooman/go/src/github.com/DaveBlooman/fasten/libraries/ruby.sh", "/Users/dblooman/Dropbox/Important/daves-aws.pem")

	if err != nil {
		return 1
	}

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
