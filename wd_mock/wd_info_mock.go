package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var (
	defaultOptionWoodpeckerInfo = setDefaultOptionWoodpeckerInfo()
)

func setDefaultOptionWoodpeckerInfo() *wd_info.WoodpeckerInfo {
	return &wd_info.WoodpeckerInfo{
		BasicInfo:      *NewBasicInfoFull(),
		CiSystemInfo:   *NewCiSystemInfoFull(),
		CiForgeInfo:    *NewCiForgeInfoFull(),
		RepositoryInfo: *NewRepositoryInfoFull(),
		CurrentInfo: wd_info.CurrentInfo{
			CurrentCommitInfo:   *NewCurrentCommitInfoFull(),
			CurrentPipelineInfo: *NewCurrentPipelineInfoFull(),
			CurrentWorkflowInfo: *NewCurrentWorkflowInfoFull(),
			CurrentStepInfo:     *NewCurrentStepInfoFull(),
		},
		PreviousInfo: wd_info.PreviousInfo{
			PreviousCommitInfo:   *NewPreviousCommitInfoFull(),
			PreviousPipelineInfo: *NewPreviousPipelineInfoFull(),
		},
	}
}

type WoodpeckerInfoOption func(*wd_info.WoodpeckerInfo)

// NewWoodpeckerInfo
//
//	use as
//
//	changeWoodpeckerInfo := NewWoodpeckerInfo(
//	    WithStatus(""),
//	)
func NewWoodpeckerInfo(opts ...WoodpeckerInfoOption) (opt *wd_info.WoodpeckerInfo) {
	opt = defaultOptionWoodpeckerInfo
	for _, o := range opts {
		o(opt)
	}
	defaultOptionWoodpeckerInfo = setDefaultOptionWoodpeckerInfo()
	return
}

func WithCiWorkspace(workspace string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.BasicInfo.CIWorkspace = workspace
	}
}

func WithCurrentPipelineStatus(status string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus = status
	}
}

func WithCiSystemInfo(opts ...CiSystemInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CiSystemInfo = *NewCiSystemInfo(opts...)
	}
}

func WithCiForgeInfo(opts ...CiForgeInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CiForgeInfo = *NewCiForgeInfo(opts...)
	}
}

func WithRepositoryInfo(opts ...RepositoryInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.RepositoryInfo = *NewRepositoryInfo(opts...)
	}
}

func WithCurrentCommitInfo(opts ...CurrentCommitInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentCommitInfo = *NewCurrentCommitInfo(opts...)
	}
}

func WithCurrentPipelineInfo(opts ...CurrentPipelineInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentPipelineInfo = *NewCurrentPipelineInfo(opts...)
	}
}

func WithCurrentWorkflowInfo(opts ...CurrentWorkflowInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentWorkflowInfo = *NewCurrentWorkflowInfo(opts...)
	}
}

func WithCurrentStepInfo(opts ...CurrentStepInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentStepInfo = *NewCurrentStepInfo(opts...)
	}
}

func WithPreviousCommitInfo(opts ...PreviousCommitInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.PreviousInfo.PreviousCommitInfo = *NewPreviousCommitInfo(opts...)
	}
}

func WithPreviousPipelineInfo(opts ...PreviousPipelineInfoOption) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.PreviousInfo.PreviousPipelineInfo = *NewPreviousPipelineInfo(opts...)
	}
}
