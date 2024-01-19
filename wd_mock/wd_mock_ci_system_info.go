package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultCiSystemInfo = setDefaultCiSystemInfo()

type CiSystemInfoOption func(*wd_info.CiSystemInfo)

func setDefaultCiSystemInfo() *wd_info.CiSystemInfo {
	return &wd_info.CiSystemInfo{
		CiSystemName:    "woodpecker",
		CiSystemUrl:     "https://woodpecker.domain.com",
		CiSystemHost:    "woodpecker.domain.com",
		CiSystemVersion: "2.1.1",
	}
}

func NewCiSystemInfo(opts ...CiSystemInfoOption) (opt *wd_info.CiSystemInfo) {
	opt = defaultCiSystemInfo
	for _, o := range opts {
		o(opt)
	}
	defaultCiSystemInfo = setDefaultCiSystemInfo()
	return
}

func WithCiSystemName(name string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.CiSystemName = name
	}
}

func WithCiSystemUrl(url string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.CiSystemUrl = url
	}
}

func WithCiSystemHost(host string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.CiSystemHost = host
	}
}

func WithCiSystemVersion(version string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.CiSystemVersion = version
	}
}

func NewCiSystemInfoFull() *wd_info.CiSystemInfo {
	return setDefaultCiSystemInfo()
}
