package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func ciForgeFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliCiForgeType,
			Usage:   "Provides name of forge (gitea, github, ...)",
			EnvVars: []string{wd_flag.EnvKeyCiForgeType},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCiForgeUrl,
			Usage:   "Provides root URL of configured forge",
			EnvVars: []string{wd_flag.EnvKeyCiForgeUrl},
			Hidden:  true,
		},
	}
}
