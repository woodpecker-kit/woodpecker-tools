package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func currentWorkflowFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCiWorkflowName,
			Usage:   "Provides workflow name",
			EnvVars: []string{wd_flag.EnvKeyCurrentCiWorkflowName},
			Hidden:  true,
		},
	}
}
