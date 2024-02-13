package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultCurrentWorkflowInfo = setDefaultCurrentWorkflowInfo()

type CurrentWorkflowInfoOption func(*wd_info.CurrentWorkflowInfo)

func setDefaultCurrentWorkflowInfo() *wd_info.CurrentWorkflowInfo {
	return &wd_info.CurrentWorkflowInfo{
		CiWorkflowName:   "build",
		CiWorkflowNumber: "1",
	}
}

func NewCurrentWorkflowInfo(opts ...CurrentWorkflowInfoOption) (opt *wd_info.CurrentWorkflowInfo) {
	opt = defaultCurrentWorkflowInfo
	for _, o := range opts {
		o(opt)
	}
	defaultCurrentWorkflowInfo = setDefaultCurrentWorkflowInfo()
	return
}

func WithCiWorkflowName(name string) CurrentWorkflowInfoOption {
	return func(o *wd_info.CurrentWorkflowInfo) {
		o.CiWorkflowName = name
	}
}

func WithCiWorkflowNumber(number string) CurrentWorkflowInfoOption {
	return func(o *wd_info.CurrentWorkflowInfo) {
		o.CiWorkflowNumber = number
	}
}

func NewCurrentWorkflowInfoFull() *wd_info.CurrentWorkflowInfo {
	return setDefaultCurrentWorkflowInfo()
}
