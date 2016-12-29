package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDeployCommand_implement(t *testing.T) {
	var _ cli.Command = &DeployCommand{}
}
