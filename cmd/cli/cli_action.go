package cli

import (
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/pkgJson"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2/cli_exit_urfave"
)

var wdPlugin *Plugin

type (
	// Plugin plugin all config
	Plugin struct {
		Name    string
		Version string
		Debug   bool

		WoodpeckerInfo *wd_info.WoodpeckerInfo
	}
)

// GlobalBeforeAction
// do command Action before flag global.
func GlobalBeforeAction(c *cli.Context) error {
	isDebug := wd_urfave_cli_v2.IsBuildDebugOpen(c)
	if isDebug {
		wd_log.OpenDebug()
	}

	// bind woodpeckerInfo
	woodpeckerInfo := wd_urfave_cli_v2.UrfaveCliBindInfo(c)

	cliVersion := pkgJson.GetPackageJsonVersionGoStyle()
	wd_log.Debugf("cli version is %s\n", cliVersion)
	wd_log.Debugf("load woodpecker finish at repo link: %v\n", woodpeckerInfo.RepositoryInfo.CIRepoURL)

	plugin := Plugin{
		Name:    pkgJson.GetPackageJsonName(),
		Version: cliVersion,
		Debug:   isDebug,

		WoodpeckerInfo: &woodpeckerInfo,
	}

	wdPlugin = &plugin

	return nil
}

// GlobalAction
// do cli Action before flag.
func GlobalAction(c *cli.Context) error {
	if wdPlugin == nil {
		return cli_exit_urfave.ErrMsg("must success run GlobalBeforeAction then run GlobalAction")
	}

	if wdPlugin.Name == "" {
		return cli_exit_urfave.ErrMsg("missing name, please set name")
	}

	if wdPlugin.Version == "" {
		return cli_exit_urfave.ErrMsg("missing version, please set version")
	}

	if wdPlugin.WoodpeckerInfo == nil {
		return cli_exit_urfave.ErrMsg("missing woodpecker info, please set woodpecker info env")
	}

	if wdPlugin.Debug { // print now all env
		wd_log.Debugf("this plugin start with all env: %s", env_kit.FindAllEnv4PrintAsSortJust(36))
	}

	errCheckVersion := wd_info.CiSystemVersionMinimumSupport(*wdPlugin.WoodpeckerInfo)
	if errCheckVersion != nil {
		return cli_exit_urfave.Err(errCheckVersion)
	}

	return nil
}

// GlobalAfterAction
//
//	do command Action after flag global.
//
//nolint:golint,unused
func GlobalAfterAction(c *cli.Context) error {
	wd_log.Infof("=> finish run: %s, version: %s\n", wdPlugin.Name, wdPlugin.Version)
	return nil
}
