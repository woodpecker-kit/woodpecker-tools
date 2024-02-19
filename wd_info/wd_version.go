package wd_info

import (
	"fmt"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
)

const WoodpeckerInfoSupportVersion = "2.3.0"

// CiSystemVersionMinimumSupport
// this function is used to check the ci system version by wd_info.WoodpeckerInfoSupportVersion
//
//	if the ci system version is empty, return error
//	if the ci system version is not supported, return error
//	if version is less than wd_info.WoodpeckerInfoSupportVersion, return error
//	if the ci system version is supported, return nil
func CiSystemVersionMinimumSupport(info WoodpeckerInfo) error {

	if info.CiSystemInfo.CiSystemVersion == "" {
		return fmt.Errorf("woodpeacker ci system version is empty, please set env %s", wd_flag.EnvKeyCiSystemVersion)
	}

	return SemverVersionMinimumSupport(info.CiSystemInfo.CiSystemVersion, WoodpeckerInfoSupportVersion)
}

// CiSystemVersionConstraint
// this function is used to check the ci system version by constraint
//
//	if the ci system version is empty, return error
//	if the ci system version is not pass, return error
//	if the ci system version is pass, return nil
func CiSystemVersionConstraint(info WoodpeckerInfo, minimumVersion, maximumVersion string) error {

	if info.CiSystemInfo.CiSystemVersion == "" {
		return fmt.Errorf("woodpeacker ci system version is empty, please set env %s", wd_flag.EnvKeyCiSystemVersion)
	}

	return SemverVersionConstraint(info.CiSystemInfo.CiSystemVersion, minimumVersion, maximumVersion)
}
