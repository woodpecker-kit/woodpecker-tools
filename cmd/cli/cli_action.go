package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/pkgJson"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2"
	"log"
	"os"
	"strings"
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
		panic(fmt.Errorf("must success run GlobalBeforeAction then run GlobalAction"))
	}

	if wdPlugin.Debug { // print now all env
		for _, e := range os.Environ() {
			if strings.Index(e, "CI") == 0 {
				log.Println(e)
			}
		}
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
