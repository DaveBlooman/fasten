package command

import (
	"strings"
)

type DeployCommand struct {
	Meta
}

func (c *DeployCommand) Run(args []string) int {
	// Write your code here
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
