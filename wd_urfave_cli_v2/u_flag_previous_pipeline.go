package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func previousPipelineFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousPipelineNumber,
			Usage:   "Provides the previous pipeline number",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineNumber},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousPipelineParent,
			Usage:   "Provides the previous pipeline number of parent pipeline",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineParent},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousPipelineEvent,
			Usage:   "Provides the previous pipeline event (push, pull_request, tag, deployment)",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineEvent},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousPipelineUrl,
			Usage:   "Provides the previous pipeline link in CI",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineUrl},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousPipelineForgeUrl,
			Usage:   "Provides the previous pipeline link in forge",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineForgeUrl},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousPipelineDeployTarget,
			Usage:   "Provides the previous pipeline deploy target for deployment events (ie production)",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineDeployTarget},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousPipelineStatus,
			Usage:   "Provides the previous pipeline status (success, failure)",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineStatus},
			Hidden:  true,
		},

		&cli.Uint64Flag{
			Name:    wd_flag.NameCliPreviousPipelineCreated,
			Usage:   "Provides the previous pipeline created",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineCreated},
			Hidden:  true,
		},

		&cli.Uint64Flag{
			Name:    wd_flag.NameCliPreviousPipelineStarted,
			Usage:   "Provides the previous pipeline started",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineStarted},
			Hidden:  true,
		},

		&cli.Uint64Flag{
			Name:    wd_flag.NameCliPreviousPipelineFinished,
			Usage:   "Provides the previous pipeline finished",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiPipelineFinished},
			Hidden:  true,
		},
	}
}
