package wd_mock

import (
	"fmt"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info_parse"
)

var (
	defaultOptionWoodpeckerInfo = setDefaultOptionWoodpeckerInfo()
)

func setDefaultOptionWoodpeckerInfo() *wd_info.WoodpeckerInfo {
	info := wd_info.WoodpeckerInfo{
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
	return &info
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

	changByDroneEnv(opt)
	return
}

func changByDroneEnv(info *wd_info.WoodpeckerInfo) {
	if info.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus == "" {
		relateCIPipelineStatus := env_kit.FetchOsEnvStr(wd_flag.RelatedEnvDroneCIPipelineStatus, "")
		// https://github.com/woodpecker-ci/woodpecker/pull/4193
		if relateCIPipelineStatus != "" {
			info.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus = relateCIPipelineStatus
		}
	}
}

func FastWorkSpace(workspace string) WoodpeckerInfoOption {
	return WithCiWorkspace(workspace)
}

func WithCiWorkspace(workspace string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.BasicInfo.CIWorkspace = workspace
	}
}

func FastCurrentStatus(status string) WoodpeckerInfoOption {
	return WithCurrentPipelineStatus(status)
}

func WithCurrentPipelineStatus(status string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus = status
	}
}

func FastTag(tagName string, msg string) WoodpeckerInfoOption {
	return WithFastMockTag(tagName, msg)
}

func WithFastMockTag(tagName string, msg string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentCommitInfo.CiCommitMessage = fmt.Sprintf("tag: %s, %s", tagName, msg)
		o.CurrentInfo.CurrentCommitInfo.CiCommitTag = tagName
		o.CurrentInfo.CurrentCommitInfo.CiCommitBranch = fmt.Sprintf("refs/tags/%s", tagName)
		o.CurrentInfo.CurrentCommitInfo.CiCommitRef = fmt.Sprintf("refs/tags/%s", tagName)
		o.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels = ""
		o.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent = wd_info.EventPipelineTag
	}
}

func FastPullRequest(prNumber, msg string, sourceBranch string, targetBranch string, commitBranch string) WoodpeckerInfoOption {
	return WithFastMockPullRequest(prNumber, msg, sourceBranch, targetBranch, commitBranch)
}

func WithFastMockPullRequest(prNumber, msg string, sourceBranch string, targetBranch string, commitBranch string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentCommitInfo.CiCommitMessage = fmt.Sprintf("pr: %s, %s", prNumber, msg)
		o.CurrentInfo.CurrentCommitInfo.CiCommitTag = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest = prNumber
		o.CurrentInfo.CurrentCommitInfo.CiCommitRef = fmt.Sprintf("refs/pull/%s/head", prNumber)
		o.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch = sourceBranch
		o.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch = targetBranch
		o.CurrentInfo.CurrentCommitInfo.CiCommitBranch = commitBranch
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels = ""
		o.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent = wd_info.EventPipelinePullRequest

		if o.RepositoryInfo.CIRepoURL != "" {
			forgeUrl := fmt.Sprintf("%s/pulls/%s", o.RepositoryInfo.CIRepoURL, prNumber)
			o.CurrentInfo.CurrentPipelineInfo.CiPipelineForgeUrl = forgeUrl
		}

	}
}

func FastPullRequestClose(prNumber, msg string, sourceBranch string, targetBranch string, commitBranch string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentCommitInfo.CiCommitMessage = fmt.Sprintf("pr close: %s, %s", prNumber, msg)
		o.CurrentInfo.CurrentCommitInfo.CiCommitTag = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest = prNumber
		o.CurrentInfo.CurrentCommitInfo.CiCommitRef = fmt.Sprintf("refs/pull/%s/head", prNumber)
		o.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch = sourceBranch
		o.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch = targetBranch
		o.CurrentInfo.CurrentCommitInfo.CiCommitBranch = commitBranch
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels = ""
		o.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent = wd_info.EventPipelinePullRequestClose

		if o.RepositoryInfo.CIRepoURL != "" {
			forgeUrl := fmt.Sprintf("%s/pulls/%s", o.RepositoryInfo.CIRepoURL, prNumber)
			o.CurrentInfo.CurrentPipelineInfo.CiPipelineForgeUrl = forgeUrl
		}

	}
}

func FastPushCommitBranch(commitBranch, sha, msg string) WoodpeckerInfoOption {
	return func(o *wd_info.WoodpeckerInfo) {
		o.CurrentInfo.CurrentCommitInfo.CiCommitMessage = fmt.Sprintf("push: %s", msg)
		o.CurrentInfo.CurrentCommitInfo.CiCommitSha = sha
		o.CurrentInfo.CurrentCommitInfo.CiCommitBranch = commitBranch
		o.CurrentInfo.CurrentCommitInfo.CiCommitRef = fmt.Sprintf("refs/heads/%s", commitBranch)
		o.CurrentInfo.CurrentCommitInfo.CiCommitTag = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch = ""
		o.CurrentInfo.CurrentCommitInfo.CiCommitPullRequestLabels = ""
		o.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent = wd_info.EventPipelinePush

		if o.RepositoryInfo.CIRepoURL != "" {
			forgeUrl := fmt.Sprintf("%s/commit/%s", o.RepositoryInfo.CIRepoURL, sha)
			o.CurrentInfo.CurrentPipelineInfo.CiPipelineForgeUrl = forgeUrl
		}

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
		_ = wd_info_parse.ParseRepositoryInfoByRepositoryInfo(&o.RepositoryInfo)
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
