package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultBasicInfo = setDefaultBasicInfo()

type BasicInfoOption func(*wd_info.BasicInfo)

func setDefaultBasicInfo() *wd_info.BasicInfo {
	return &wd_info.BasicInfo{
		CI: "woodpecker",
	}
}

func NewBasicInfo(opts ...BasicInfoOption) (opt *wd_info.BasicInfo) {
	opt = defaultBasicInfo
	for _, o := range opts {
		o(opt)
	}
	defaultBasicInfo = setDefaultBasicInfo()
	return
}

func WithCI(ci string) BasicInfoOption {
	return func(o *wd_info.BasicInfo) {
		o.CI = ci
	}
}

func WithCIWorkspace(workspace string) BasicInfoOption {
	return func(o *wd_info.BasicInfo) {
		o.CIWorkspace = workspace
	}
}

func NewBasicInfoFull() *wd_info.BasicInfo {
	return setDefaultBasicInfo()
}
