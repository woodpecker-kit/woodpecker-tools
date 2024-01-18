package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/pkgJson"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_urfave_cli_v2/exit_cli"
	"runtime"
	"time"
)

const (
	defaultExitCode    = 1
	CopyrightStartYear = "2023"
)

func NewCliApp() *cli.App {
	name := pkgJson.GetPackageJsonName()
	exit_cli.ChangeDefaultExitCode(defaultExitCode)
	if name == "" {
		panic("package.json name is empty")
	}

	app := cli.NewApp()
	app.Version = pkgJson.GetPackageJsonVersionGoStyle()
	app.Name = name
	app.Usage = pkgJson.GetPackageJsonDescription()
	year := time.Now().Year()
	jsonAuthor := pkgJson.GetPackageJsonAuthor()
	app.Copyright = fmt.Sprintf("© %s-%d %s by: %s, run on %s %s",
		CopyrightStartYear, year, jsonAuthor.Name, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	author := &cli.Author{
		Name:  jsonAuthor.Name,
		Email: jsonAuthor.Email,
	}
	app.Authors = []*cli.Author{
		author,
	}

	app.Before = GlobalBeforeAction
	app.Action = GlobalAction
	app.After = GlobalAfterAction

	flags := wd_urfave_cli_v2.UrfaveCliAppendCliFlags(
		wd_urfave_cli_v2.WoodpeckerUrfaveCliFlags(),
		wd_urfave_cli_v2.CommonFlag(),
		wd_urfave_cli_v2.CommonRequestBasicFlag())

	app.Flags = flags
	app.HideHelpCommand = false

	return app
}
