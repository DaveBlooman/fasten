package main

import (
	"github.com/DaveBlooman/fasten/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{

		"init": func() (cli.Command, error) {
			return &command.InitCommand{
				Meta: *meta,
			}, nil
		},

		"deploy": func() (cli.Command, error) {
			return &command.DeployCommand{
				Meta: *meta,
			}, nil
		},

		"stop": func() (cli.Command, error) {
			return &command.StopCommand{
				Meta: *meta,
			}, nil
		},

		"install": func() (cli.Command, error) {
			return &command.InstallCommand{
				Meta: *meta,
			}, nil
		},
	}
}
