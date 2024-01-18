package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

func repositoryFlag() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiRepo,
			Usage:   "Provides the repository repo flag. This value is `<owner>/<name>` when the is run in woodpecker",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiRepo},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiOwner,
			Usage:   "Provides the repository owner flag. This value is `<owner>` when the is run in woodpecker",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiOwner},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiName,
			Usage:   "Provides the repository name flag. This value is `<name>` when the is run in woodpecker",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiName},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiRemoteId,
			Usage:   "Provides the repository remote ID, is the UID it has in the forge.",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiRemoteId},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiScm,
			Usage:   "Provides the repository scm flag. repository SCM (git).",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiScm},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiUrl,
			Usage:   "Provides the repository url flag. repository URL.",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiUrl},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiCloneUrl,
			Usage:   "Provides the repository clone url flag. repository clone URL.",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiCloneUrl},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiCloneSshUrl,
			Usage:   "Provides the repository clone ssh url flag. repository clone ssh URL.",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiCloneSshUrl},
			Hidden:  true,
		},
		&cli.StringFlag{
			Name:    wd_flag.NameCliRepositoryCiDefaultBranch,
			Usage:   "Provides the repository default branch flag. repository default branch.",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiDefaultBranch},
			Hidden:  true,
		},
		&cli.BoolFlag{
			Name:    wd_flag.NameCliRepositoryCiPrivate,
			Usage:   "Provides the repository private flag. repository private.",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiPrivate},
			Hidden:  true,
		},
		&cli.BoolFlag{
			Name:    wd_flag.NameCliRepositoryCiTrusted,
			Usage:   "Provides the repository trusted flag. repository trusted.",
			EnvVars: []string{wd_flag.EnvKeyRepositoryCiTrusted},
			Hidden:  true,
		},
	}
}
