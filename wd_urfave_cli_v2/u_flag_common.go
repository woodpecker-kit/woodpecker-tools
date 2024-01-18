package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

// CommonFlag
// Other modules also have flags
func CommonFlag() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    wd_flag.NameCliPluginDebug,
			Usage:   "Provides the debug flag. This value is true when the is open debug mode",
			EnvVars: []string{wd_flag.EnvKeyPluginDebug},
		},
	}
}

// CommonRequestBasicFlag
//
//	do plugin request flag.
func CommonRequestBasicFlag() []cli.Flag {
	return []cli.Flag{
		&cli.UintFlag{
			Name:    wd_flag.NameCliPluginTimeoutSecond,
			Usage:   "do request timeout setting second.",
			Hidden:  true,
			Value:   10,
			EnvVars: []string{wd_flag.EnvKeyPluginTimeoutSecond},
		},

		&cli.UintFlag{
			Name:    wd_flag.NameCliPluginUploadTimeoutSecond,
			Usage:   "do request upload timeout setting second",
			Hidden:  true,
			Value:   300,
			EnvVars: []string{wd_flag.EnvKeyPluginUploadTimeoutSecond},
		},
	}
}
