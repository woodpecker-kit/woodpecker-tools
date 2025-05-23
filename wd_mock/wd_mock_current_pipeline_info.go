package wd_mock

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
)

var defaultCurrentPipelineInfo = setDefaultCurrentPipelineInfo()

type CurrentPipelineInfoOption func(*wd_info.CurrentPipelineInfo)

func setDefaultCurrentPipelineInfo() *wd_info.CurrentPipelineInfo {
	return &wd_info.CurrentPipelineInfo{
		CiPipelineNumber:        "10",
		CiPipelineParent:        "0",
		CiPipelineEvent:         "push",
		CiPipelineUrl:           "https://woodpecker.domain.com/repos/2/pipeline/10",
		CiPipelineForgeUrl:      "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/9c764dd487bce596c5c0402478fabde5f0344983",
		CiPipelineDeployTarget:  "",
		CiPipelineStatus:        wd_info.BuildStatusSuccess,
		CiPipelineCreated:       1705658141,
		CiPipelineCreatedT:      wd_flag.FormatTimeUTCBySetting(1705658141),
		CiPipelineStarted:       1705658156,
		CiPipelineStartedT:      wd_flag.FormatTimeUTCBySetting(1705658156),
		CiPipelineDurationHuman: wd_flag.DistanceBetweenTimestampSecondHuman(1705658141, 1705658166),
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

// Deprecated: remove at woodpecker server v3.0.0
func WithCiPipelineFiles(files string) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		//o.CiPipelineFiles = files
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
		o.CiPipelineCreatedT = wd_flag.FormatTimeUTCBySetting(created)
		o.CiPipelineDurationHuman = wd_flag.DistanceBetweenTimestampSecondHuman(int64(o.CiPipelineCreated), wd_flag.GetNowTimestampSecond())
	}
}

func WithCiPipelineStarted(started uint64) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		o.CiPipelineStarted = started
		o.CiPipelineStartedT = wd_flag.FormatTimeUTCBySetting(started)
	}
}

// Deprecated: remove at woodpecker server v3.0.0
func WithCiPipelineFinished(finished uint64) CurrentPipelineInfoOption {
	return func(o *wd_info.CurrentPipelineInfo) {
		//o.CiPipelineFinished = finished
		//o.CiPipelineFinishedT = wd_flag.FormatTimeUTCBySetting(finished)
		//o.CiPipelineDurationHuman = wd_flag.DistanceBetweenTimestampSecondHuman(int64(o.CiPipelineStarted), int64(o.CiPipelineFinished))
	}
}

func NewCurrentPipelineInfoFull() *wd_info.CurrentPipelineInfo {
	return setDefaultCurrentPipelineInfo()
}
