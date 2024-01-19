package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultCurrentStepInfo = setDefaultCurrentStepInfo()

type CurrentStepInfoOption func(*wd_info.CurrentStepInfo)

func setDefaultCurrentStepInfo() *wd_info.CurrentStepInfo {
	return &wd_info.CurrentStepInfo{
		CiStepName:     "checkout",
		CiStepNumber:   "0",
		CiStepStatus:   "success",
		CiStepStarted:  1705658156,
		CiStepFinished: 1705658166,
		CiStepUrl:      "https://woodpecker.domain.com/repos/2/pipeline/10",
	}
}

func NewCurrentStepInfo(opts ...CurrentStepInfoOption) (opt *wd_info.CurrentStepInfo) {
	opt = defaultCurrentStepInfo
	for _, o := range opts {
		o(opt)
	}
	defaultCurrentStepInfo = setDefaultCurrentStepInfo()
	return
}

func WithCiStepName(stepName string) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepName = stepName
	}
}

func WithCiStepNumber(stepNumber string) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepNumber = stepNumber
	}
}

func WithCiStepStatus(stepStatus string) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepStatus = stepStatus
	}
}

func WithCiStepStarted(stepStarted uint64) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepStarted = stepStarted
	}
}

func WithCiStepFinished(stepFinished uint64) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepFinished = stepFinished
	}
}

func WithCiStepUrl(stepUrl string) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepUrl = stepUrl
	}
}

func NewCurrentStepInfoFull() *wd_info.CurrentStepInfo {
	return setDefaultCurrentStepInfo()
}
