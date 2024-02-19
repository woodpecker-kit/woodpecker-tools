package wd_info

import (
	"fmt"
	"github.com/Masterminds/semver/v3"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

const WoodpeckerInfoSupportVersion = "2.3.0"

// CiSystemVersionCheck
// this function is used to check the ci system version by wd_info.WoodpeckerInfoSupportVersion
//
//	if the ci system version is empty, return error
//	if the ci system version is not supported, return error
//	if version is less than wd_info.WoodpeckerInfoSupportVersion, return error
//	if the ci system version is supported, return nil
func CiSystemVersionCheck(info WoodpeckerInfo) error {

	if info.CiSystemInfo.CiSystemVersion == "" {
		return fmt.Errorf("woodpeacker ci system version is empty, please set env %s", wd_flag.EnvKeyCiSystemVersion)
	}

	targetVersion, errTargetVersion := semver.NewVersion(info.CiSystemInfo.CiSystemVersion)
	if errTargetVersion != nil {
		return fmt.Errorf("can not parse target version: %s err: %v", info.CiSystemInfo.CiSystemVersion, errTargetVersion)
	}

	checkVersion, _ := semver.NewConstraint(fmt.Sprintf(">= %s", WoodpeckerInfoSupportVersion))

	validateOk, errors := checkVersion.Validate(targetVersion)
	if !validateOk {
		return fmt.Errorf("woodpeacker ci system version: %s not support, err: %v", info.CiSystemInfo.CiSystemVersion, errors)
	}

	return nil
}
