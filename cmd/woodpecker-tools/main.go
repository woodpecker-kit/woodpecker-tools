//go:build !test

package main

import (
	_ "embed"
	resource "github.com/woodpecker-kit/woodpecker-tools"
	"github.com/woodpecker-kit/woodpecker-tools/cmd/cli"
	"github.com/woodpecker-kit/woodpecker-tools/pkgJson"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"os"
)

var buildID string

func init() {
	if buildID == "" {
		buildID = "unknown"
	}
}

func main() {
	wd_log.SetLogLineDeep(wd_log.DefaultExtLogLineMaxDeep)
	pkgJson.InitPkgJsonContent(resource.PackageJson)

	// init wd_template
	//wd_template.RegisterSettings(wd_template.DefaultHelpers)

	app := cli.NewCliApp(buildID)

	// app run as urfave
	if err := app.Run(os.Args); nil != err {
		wd_log.Warnf("run err: %v", err)
	}
}
