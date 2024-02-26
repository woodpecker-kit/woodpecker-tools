package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func ciSystemFlag() []cli.Flag {
	return []cli.Flag{

		&cli.StringFlag{
			Name:    wd_flag.NameCliWoodpeckerBackend,
			Usage:   " Provides the backend flag. This value is type when the is run in woodpecker, default is empty.",
			EnvVars: []string{wd_flag.EnvKeyWoodpeckerBackend},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliWoodpeckerHostName,
			Usage:   "Provides the hostname flag. This value is hostname when the is run in woodpecker, default is empty.",
			EnvVars: []string{wd_flag.EnvKeyWoodpeckerHostName},
			Hidden:  true,
		},
		&cli.StringSliceFlag{
			Name:    wd_flag.NameCliWoodpeckerFilterLabels,
			Usage:   "Use a list of key-value pairs like key=value,second-key=*. * can be used as a wildcard",
			EnvVars: []string{wd_flag.EnvKeyWoodpeckerFilterLabels},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCiMachine,
			Usage:   "Provides name of the CI machine",
			EnvVars: []string{wd_flag.EnvKeyCiMachine},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCiSystemPlatform,
			Usage:   "Provides ci system platform as: linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64",
			EnvVars: []string{wd_flag.EnvKeyCiSystemPlatform},
			Hidden:  true,
		},

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
