package command

import "strings"

type StatusCommand struct {
	Meta
}

func (c *StatusCommand) Run(args []string) int {

	return 0

}

func (c *StatusCommand) Synopsis() string {
	return ""
}

func (c *StatusCommand) Help() string {
	helpText := `Shows Information about deployment

`
	return strings.TrimSpace(helpText)
}
