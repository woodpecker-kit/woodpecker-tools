package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func currentPipelineFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineFiles,
			Usage:   "Provides the current pipeline files just like [\".woodpecker/.build.yml\"]",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineFiles},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineNumber,
			Usage:   "Provides the current pipeline number",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineNumber},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineParent,
			Usage:   "Provides number of parent pipeline",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineParent},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineEvent,
			Usage:   "Provides the current pipeline event (push, pull_request, tag, deployment, cron, manual)",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineEvent},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineUrl,
			Usage:   "Provides the current link to the web UI for the pipeline",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineUrl},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineForgeUrl,
			Usage:   "Provides the current link to the forge's web UI for the commit(s) or tag that triggered the pipeline",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineForgeUrl},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineDeployTarget,
			Usage:   "Provides the current pipeline deploy target for deployment events (ie production)",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineDeployTarget},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentPipelineStatus,
			Usage:   "Provides the current pipeline status (success, failure)",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineStatus},
			Hidden:  true,
		},

		&cli.Uint64Flag{
			Name:    wd_flag.NameCliCurrentPipelineCreated,
			Usage:   "Provides the current pipeline created timestamp",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineCreated},
			Hidden:  true,
		},

		&cli.Uint64Flag{
			Name:    wd_flag.NameCliCurrentPipelineStarted,
			Usage:   "Provides the current pipeline started timestamp",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineStarted},
			Hidden:  true,
		},

		&cli.Uint64Flag{
			Name:    wd_flag.NameCliCurrentPipelineFinished,
			Usage:   "Provides the current pipeline finished timestamp",
			EnvVars: []string{wd_flag.EnvKeyCurrentPipelineFinished},
			Hidden:  true,
		},
	}
}
