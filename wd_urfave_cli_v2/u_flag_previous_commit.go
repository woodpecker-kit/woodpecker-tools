package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

// nolint: dupl
func previousCommitFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitSha,
			Usage:   "Provides the previous commit sha",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitSha},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitRef,
			Usage:   "Provides the previous commit ref",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitRef},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitRefSpec,
			Usage:   "Provides the previous commit ref spec",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitRefSpec},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitBranch,
			Usage:   "Provides the previous commit branch",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitBranch},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitSourceBranch,
			Usage:   "Provides the previous commit source branch",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitSourceBranch},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitTargetBranch,
			Usage:   "Provides the previous commit target branch",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitTargetBranch},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitUrl,
			Usage:   "Provides the previous commit url",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitUrl},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitMessage,
			Usage:   "Provides the previous commit message",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitMessage},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitAuthor,
			Usage:   "Provides the previous commit author username",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitAuthor},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitAuthorEmail,
			Usage:   "Provides the previous commit author email",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitAuthorEmail},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliPreviousCommitAuthorAvatar,
			Usage:   "Provides the previous commit author avatar",
			EnvVars: []string{wd_flag.EnvKeyPreviousCiCommitAuthorAvatar},
			Hidden:  true,
		},
	}
}
