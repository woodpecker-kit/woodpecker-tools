package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func basicFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliCi,
			Usage:   "Provides the ci flag. This value is `woodpecker` when the is run in woodpecker",
			EnvVars: []string{wd_flag.EnvKeyNameCi},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCiWorkspace,
			Usage:   "Provides the Path of the workspace where source code gets cloned to",
			EnvVars: []string{wd_flag.EnvKeyNameCiWorkspace},
			Hidden:  true,
		},
	}
}
