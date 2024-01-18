package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func currentCommitFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitSha,
			Usage:   "Provides the current commit SHA",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitSha},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitRef,
			Usage:   "Provides the current commit ref",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitRef},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitRefSpec,
			Usage:   "Provides the current commit ref spec",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitRefSpec},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitBranch,
			Usage:   "Provides the current commit branch (equals target branch for pull requests)",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitBranch},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitSourceBranch,
			Usage:   "Provides the current commit source branch (empty if event is not pull_request)",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitSourceBranch},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitTargetBranch,
			Usage:   "Provides the current commit target branch (empty if event is not pull_request)",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitTargetBranch},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitTag,
			Usage:   "Provides the current commit tag (empty if event is not tag)",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitTag},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitPullRequest,
			Usage:   "Provides the current commit pull request (empty if event is not pull_request)",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitPullRequest},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitPullRequestLabels,
			Usage:   "Provides the current commit pull request labels (empty if event is not pull_request)",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitPullRequestLabels},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitMessage,
			Usage:   "Provides the current commit message",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitMessage},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitAuthor,
			Usage:   "Provides the current commit author",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitAuthor},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitAuthorEmail,
			Usage:   "Provides the current commit author email",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitAuthorEmail},
			Hidden:  true,
		},

		&cli.StringFlag{
			Name:    wd_flag.NameCliCurrentCommitCiCommitAuthorAvatar,
			Usage:   "Provides the current commit author avatar",
			EnvVars: []string{wd_flag.EnvKeyCurrentCommitCiCommitAuthorAvatar},
			Hidden:  true,
		},
	}
}
