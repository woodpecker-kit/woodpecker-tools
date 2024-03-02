package wd_mock

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
)

var defaultPreviousPipelineInfo = setDefaultPreviousPipelineInfo()

type PreviousPipelineInfoOption func(*wd_info.PreviousPipelineInfo)

func setDefaultPreviousPipelineInfo() *wd_info.PreviousPipelineInfo {
	return &wd_info.PreviousPipelineInfo{
		CiPreviousPipelineNumber:       "9",
		CiPreviousPipelineParent:       "",
		CiPreviousPipelineEvent:        "push",
		CiPreviousPipelineUrl:          "https://woodpecker.domain.com/repos/2/pipeline/9",
		CiPreviousPipelineForgeUrl:     "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/b2542f37e23645b2032fcf9f180e6940ef6915ac",
		CiPreviousPipelineDeployTarget: "",
		CiPreviousPipelineStatus:       "success",
		CiPreviousPipelineCreated:      1705657931,
		CiPreviousPipelineCreatedT:     wd_flag.FormatTimeUTC(1705657931, wd_flag.TimeFormatDefault),
		CiPreviousPipelineStarted:      1705657942,
		CiPreviousPipelineStartedT:     wd_flag.FormatTimeUTC(1705657942, wd_flag.TimeFormatDefault),
		CiPreviousPipelineFinished:     1705657957,
		CiPreviousPipelineFinishedT:    wd_flag.FormatTimeUTC(1705657957, wd_flag.TimeFormatDefault),
	}
}

func NewPreviousPipelineInfo(opts ...PreviousPipelineInfoOption) (opt *wd_info.PreviousPipelineInfo) {
	opt = defaultPreviousPipelineInfo
	for _, o := range opts {
		o(opt)
	}
	defaultPreviousPipelineInfo = setDefaultPreviousPipelineInfo()
	return
}

func WithCiPreviousPipelineNumber(pipelineNumber string) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineNumber = pipelineNumber
	}
}

func WithCiPreviousPipelineParent(pipelineParent string) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineParent = pipelineParent
	}
}

func WithCiPreviousPipelineEvent(pipelineEvent string) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineEvent = pipelineEvent
	}
}

func WithCiPreviousPipelineUrl(pipelineUrl string) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineUrl = pipelineUrl
	}
}

func WithCiPreviousPipelineForgeUrl(pipelineForgeUrl string) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineForgeUrl = pipelineForgeUrl
	}
}

func WithCiPreviousPipelineDeployTarget(pipelineDeployTarget string) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineDeployTarget = pipelineDeployTarget
	}
}

func WithCiPreviousPipelineStatus(pipelineStatus string) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineStatus = pipelineStatus
	}
}

func WithCiPreviousPipelineCreated(pipelineCreated uint64) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineCreated = pipelineCreated
		o.CiPreviousPipelineCreatedT = wd_flag.FormatTimeUTC(pipelineCreated, wd_flag.TimeFormatDefault)
	}
}

func WithCiPreviousPipelineStarted(pipelineStarted uint64) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineStarted = pipelineStarted
		o.CiPreviousPipelineStartedT = wd_flag.FormatTimeUTC(pipelineStarted, wd_flag.TimeFormatDefault)
	}
}

func WithCiPreviousPipelineFinished(pipelineFinished uint64) PreviousPipelineInfoOption {
	return func(o *wd_info.PreviousPipelineInfo) {
		o.CiPreviousPipelineFinished = pipelineFinished
		o.CiPreviousPipelineFinishedT = wd_flag.FormatTimeUTC(pipelineFinished, wd_flag.TimeFormatDefault)
	}
}

func NewPreviousPipelineInfoFull() *wd_info.PreviousPipelineInfo {
	return setDefaultPreviousPipelineInfo()
}
