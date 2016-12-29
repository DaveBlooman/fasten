package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestInstallCommand_implement(t *testing.T) {
	var _ cli.Command = &InstallCommand{}
}
