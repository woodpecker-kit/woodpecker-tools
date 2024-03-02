package wd_mock

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
)

var defaultCurrentPipelineInfo = setDefaultCurrentPipelineInfo()

type CurrentPipelineInfoOption func(*wd_info.CurrentPipelineInfo)

func setDefaultCurrentPipelineInfo() *wd_info.CurrentPipelineInfo {
	return &wd_info.CurrentPipelineInfo{
		CiPipelineFiles:        `[".woodpecker/.build.yml"]`,
		CiPipelineNumber:       "10",
		CiPipelineParent:       "0",
		CiPipelineEvent:        "push",
		CiPipelineUrl:          "https://woodpecker.domain.com/repos/2/pipeline/10",
		CiPipelineForgeUrl:     "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/9c764dd487bce596c5c0402478fabde5f0344983",
		CiPipelineDeployTarget: "",
		CiPipelineStatus:       "success",
		CiPipelineCreated:      1705658141,
		CiPipelineCreatedT:     wd_flag.FormatTimeUTC(1705658141, wd_flag.TimeFormatDefault),
		CiPipelineStarted:      1705658156,
		CiPipelineStartedT:     wd_flag.FormatTimeUTC(1705658156, wd_flag.TimeFormatDefault),
		CiPipelineFinished:     1705658166,
		CiPipelineFinishedT:    wd_flag.FormatTimeUTC(1705658166, wd_flag.TimeFormatDefault),
	}
}

func NewCurrentPipelineInfo(opts ...CurrentPipelineInfoOption) (opt *wd_info.CurrentPipelineInfo) {
	opt = defaultCurrentPipelineInfo
	for _, o := range opts {
		o(opt)
	}
	defaultCurrentPipelineInfo = setDefaultCurrentPipelineInfo()
	return
}

func WithCiPipelineFiles(files string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineFiles = files
	}
}

func WithCiPipelineNumber(number string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineNumber = number
	}
}

func WithCiPipelineParent(parent string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineParent = parent
	}
}

func WithCiPipelineEvent(event string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineEvent = event
	}
}

func WithCiPipelineUrl(url string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineUrl = url
	}
}

func WithCiPipelineForgeUrl(forgeUrl string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineForgeUrl = forgeUrl
	}
}

func WithCiPipelineDeployTarget(deployTarget string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineDeployTarget = deployTarget
	}
}

func WithCiPipelineStatus(status string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineStatus = status
	}
}

func WithCiPipelineCreated(created uint64) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineCreated = created
		o.CiPipelineCreatedT = wd_flag.FormatTimeUTC(created, wd_flag.TimeFormatDefault)
	}
}

func WithCiPipelineStarted(started uint64) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineStarted = started
		o.CiPipelineStartedT = wd_flag.FormatTimeUTC(started, wd_flag.TimeFormatDefault)
	}
}

func WithCiPipelineFinished(finished uint64) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineFinished = finished
		o.CiPipelineFinishedT = wd_flag.FormatTimeUTC(finished, wd_flag.TimeFormatDefault)
	}
}

func NewCurrentPipelineInfoFull() *wd_info.CurrentPipelineInfo {
	return setDefaultCurrentPipelineInfo()
}
