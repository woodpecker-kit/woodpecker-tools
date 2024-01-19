package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultCiForgeInfo = setDefaultCiForgeInfo()

type CiForgeInfoOption func(*wd_info.CiForgeInfo)

func setDefaultCiForgeInfo() *wd_info.CiForgeInfo {
	return &wd_info.CiForgeInfo{
		CiForgeType: "gitea",
		CiForgeUrl:  "https://gitea.domain.com",
	}
}

func NewCiForgeInfo(opts ...CiForgeInfoOption) (opt *wd_info.CiForgeInfo) {
	opt = defaultCiForgeInfo
	for _, o := range opts {
		o(opt)
	}
	defaultCiForgeInfo = setDefaultCiForgeInfo()
	return
}

func WithCiForgeType(forgeType string) CiForgeInfoOption {
	return func(o *wd_info.CiForgeInfo) {
		o.CiForgeType = forgeType
	}
}

func WithCiForgeUrl(forgeUrl string) CiForgeInfoOption {
	return func(o *wd_info.CiForgeInfo) {
		o.CiForgeUrl = forgeUrl
	}
}

func NewCiForgeInfoFull() *wd_info.CiForgeInfo {
	return setDefaultCiForgeInfo()
}
