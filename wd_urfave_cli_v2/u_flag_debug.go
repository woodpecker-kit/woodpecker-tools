package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

// IsBuildDebugOpen
// when config or drone build open debug will open debug
func IsBuildDebugOpen(c *cli.Context) bool {
	return c.Bool(wd_flag.NameCliPluginDebug)
}
