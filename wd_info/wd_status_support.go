package wd_info

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
)

var ( // pluginBuildStateSupport
	pluginBuildStateSupport = []string{
		BuildStatusCreated,
		BuildStatusRunning,
		BuildStatusSuccess,
		BuildStatusFailure,
		BuildStatusError,
		BuildStatusKilled,
	}
)

func BuildStatusSupport() []string {
	return pluginBuildStateSupport
}

func CheckBuildStatusSupport(status string) error {
	if !(string_tools.StringInArr(status, pluginBuildStateSupport)) {
		return fmt.Errorf("wd_info build status not support now [ %s ], must in %v", status, pluginBuildStateSupport)
	}
	return nil
}
