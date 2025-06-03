package wd_mock

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
)

var defaultCurrentStepInfo = setDefaultCurrentStepInfo()

type CurrentStepInfoOption func(*wd_info.CurrentStepInfo)

func setDefaultCurrentStepInfo() *wd_info.CurrentStepInfo {
	return &wd_info.CurrentStepInfo{
		CiStepName:          "checkout",
		CiStepNumber:        "0",
		CiStepStarted:       1705658156,
		CiStepStartedT:      wd_flag.FormatTimeUTCBySetting(1705658156),
		CiStepNowTime:       1705658166,
		CiStepNowTimeT:      wd_flag.FormatTimeUTCBySetting(1705658166),
		CiStepDurationHuman: wd_flag.DistanceBetweenTimestampSecondHuman(1705658156, 1705658166),
		CiStepUrl:           "https://woodpecker.domain.com/repos/2/pipeline/10",
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

// Deprecated: remove at woodpecker server 3.0.0.
func WithCiStepStatus(stepStatus string) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		// o.CiStepStatus = stepStatus
	}
}

func WithCiStepStarted(stepStarted uint64) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepStarted = stepStarted
		o.CiStepStartedT = wd_flag.FormatTimeUTCBySetting(stepStarted)
		o.CiStepDurationHuman = wd_flag.DistanceBetweenTimestampSecondHuman(
			int64(o.CiStepStarted),
			int64(o.CiStepNowTime),
		)
	}
}

// Deprecated: remove at woodpecker server 3.0.0, instead of WithCiStepNowTime.
func WithCiStepFinished(stepFinished uint64) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		// o.CiStepFinished = stepFinished
		// o.CiStepFinishedT = wd_flag.FormatTimeUTCBySetting(stepFinished)
		// o.CiStepDurationHuman = wd_flag.DistanceBetweenTimestampSecondHuman(int64(o.CiStepStarted), int64(o.CiStepFinished))
	}
}

func WithCiStepNowTime(stepNowTime uint64) CurrentStepInfoOption {
	return func(o *wd_info.CurrentStepInfo) {
		o.CiStepNowTime = stepNowTime
		o.CiStepDurationHuman = wd_flag.DistanceBetweenTimestampSecondHuman(
			int64(o.CiStepStarted),
			int64(o.CiStepNowTime),
		)
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
