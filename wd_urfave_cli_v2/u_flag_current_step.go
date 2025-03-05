package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func currentStepFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCiStepName,
			Usage:   "Provides the current step name",
			EnvVars: []string{wd_flag.EnvKeyCurrentCiStepName},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCiStepNumber,
			Usage:   "Provides the current step number",
			EnvVars: []string{wd_flag.EnvKeyCurrentCiStepNumber},
			Hidden:  true,
		},

		// Deprecated: remove at woodpecker server 3.0.0
		//&cli.StringFlag{
		//	Name:    wd_flag.NameCliCurrentCiStepStatus,
		//	Usage:   "Provides the current step status (success, failure)",
		//	EnvVars: []string{wd_flag.EnvKeyCurrentCiStepStatus},
		//	Hidden:  true,
		//},

		&cli.Uint64Flag{
			Name:    wd_flag.NameCliCurrentCiStepStarted,
			Usage:   "Provides the current step started UNIX timestamp",
			EnvVars: []string{wd_flag.EnvKeyCurrentCiStepStarted},
			Hidden:  true,
		},

		// Deprecated: remove at woodpecker server 3.0.0
		//&cli.Uint64Flag{
		//	Name:    wd_flag.NameCliCurrentCiStepFinished,
		//	Usage:   "Provides the current step finished UNIX timestamp",
		//	EnvVars: []string{wd_flag.EnvKeyCurrentCiStepFinished},
		//	Hidden:  true,
		//},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCiStepUrl,
			Usage:   "Provides the current URL to step in UI",
			EnvVars: []string{wd_flag.EnvKeyCurrentCiStepUrl},
			Hidden:  true,
		},
	}
}
