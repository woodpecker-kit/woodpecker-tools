package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func ciSystemFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliCiSystemName,
			Usage:   "Provides the ci system name. This value is `woodpecker` when the is run in woodpecker",
			EnvVars: []string{wd_flag.EnvKeyCiSystemName},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCiSystemUrl,
			Usage:   "Provides  link to CI system",
			EnvVars: []string{wd_flag.EnvKeyCiSystemUrl},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCiSystemHost,
			Usage:   "Provides hostname of CI server",
			EnvVars: []string{wd_flag.EnvKeyCiSystemHost},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCiSystemVersion,
			Usage:   "Provides version of CI server",
			EnvVars: []string{wd_flag.EnvKeyCiSystemVersion},
			Hidden:  true,
		},
	}
}
