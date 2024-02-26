package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultCiSystemInfo = setDefaultCiSystemInfo()

type CiSystemInfoOption func(*wd_info.CiSystemInfo)

func setDefaultCiSystemInfo() *wd_info.CiSystemInfo {
	return &wd_info.CiSystemInfo{
		WoodpeckerBackend:       "docker",
		WoodpeckerAgentHostName: "",
		WoodpeckerFilterLabels:  nil,
		CiMachine:               "worker",
		CiSystemPlatform:        "linux/amd64",
		CiSystemName:            "woodpecker",
		CiSystemUrl:             "https://woodpecker.domain.com",
		CiSystemHost:            "woodpecker.domain.com",
		CiSystemVersion:         wd_info.WoodpeckerInfoSupportVersion,
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

func WithWoodpeckerBackend(backend string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.WoodpeckerBackend = backend
	}
}

func WithWoodpeckerAgentHostName(hostName string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.WoodpeckerAgentHostName = hostName
	}
}

func WithWoodpeckerFilterLabels(labels []string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.WoodpeckerFilterLabels = labels
	}
}

func WithCiMachine(machine string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.CiMachine = machine
	}
}

func WithCiSystemPlatform(platform string) CiSystemInfoOption {
	return func(o *wd_info.CiSystemInfo) {
		o.CiSystemPlatform = platform
	}

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
