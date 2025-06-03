package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info_parse"
)

// UrfaveCliBindInfo
//
//	woodpecker urfave cli bind info
//
// want change time format, please change wd_flag.SetTimeFormat before this method.
func UrfaveCliBindInfo(c *cli.Context) wd_info.WoodpeckerInfo {
	pipelineCreateAt := c.Uint64(wd_flag.NameCliCurrentPipelineCreated)
	pipelineCreateAtT := wd_flag.FormatTimeUTCBySetting(pipelineCreateAt)
	pipelineStartAt := c.Uint64(wd_flag.NameCliCurrentPipelineStarted)
	pipelineStartAtT := wd_flag.FormatTimeUTCBySetting(pipelineStartAt)
	pipelineDurationHuman := wd_flag.DistanceBetweenTimestampSecondHuman(
		int64(pipelineCreateAt),
		wd_flag.GetNowTimestampSecond(),
	)

	previousPipelineCreateAt := c.Uint64(wd_flag.NameCliPreviousPipelineCreated)
	previousPipelineCreateAtT := wd_flag.FormatTimeUTCBySetting(previousPipelineCreateAt)
	previousPipelineStartedAt := c.Uint64(wd_flag.NameCliPreviousPipelineStarted)
	previousPipelineStartedAtT := wd_flag.FormatTimeUTCBySetting(previousPipelineStartedAt)
	previousPipelineFinishedAt := c.Uint64(wd_flag.NameCliPreviousPipelineFinished)
	previousPipelineFinishedAtT := wd_flag.FormatTimeUTCBySetting(previousPipelineFinishedAt)
	previousPipelineDurationHuman := wd_flag.DistanceBetweenTimestampSecondHuman(
		int64(previousPipelineCreateAt),
		int64(previousPipelineFinishedAt),
	)

	stepStartedAt := c.Uint64(wd_flag.NameCliCurrentCiStepStarted)
	stepStartedAtT := wd_flag.FormatTimeUTCBySetting(stepStartedAt)
	stepNowTimeAt := wd_flag.GetNowTimestampSecondU64()
	stepNowTimeAtT := wd_flag.FormatTimeUTCBySetting(stepNowTimeAt)
	stepDurationHuman := wd_flag.DistanceBetweenTimestampSecondHuman(
		int64(stepStartedAt),
		int64(stepNowTimeAt),
	)

	currentInfo := wd_info.CurrentInfo{
		CurrentCommitInfo: wd_info.CurrentCommitInfo{
			CiCommitSha:          c.String(wd_flag.NameCliCurrentCommitCiCommitSha),
			CiCommitRef:          c.String(wd_flag.NameCliCurrentCommitCiCommitRef),
			CiCommitRefSpec:      c.String(wd_flag.NameCliCurrentCommitCiCommitRefSpec),
			CiCommitBranch:       c.String(wd_flag.NameCliCurrentCommitCiCommitBranch),
			CiCommitSourceBranch: c.String(wd_flag.NameCliCurrentCommitCiCommitSourceBranch),
			CiCommitTargetBranch: c.String(wd_flag.NameCliCurrentCommitCiCommitTargetBranch),
			CiCommitTag:          c.String(wd_flag.NameCliCurrentCommitCiCommitTag),
			CiCommitPullRequest:  c.String(wd_flag.NameCliCurrentCommitCiCommitPullRequest),
			CiCommitPullRequestLabels: c.String(
				wd_flag.NameCliCurrentCommitCiCommitPullRequestLabels,
			),
			CiCommitMessage:      c.String(wd_flag.NameCliCurrentCommitCiCommitMessage),
			CiCommitAuthor:       c.String(wd_flag.NameCliCurrentCommitCiCommitAuthor),
			CiCommitAuthorEmail:  c.String(wd_flag.NameCliCurrentCommitCiCommitAuthorEmail),
			CiCommitAuthorAvatar: c.String(wd_flag.NameCliCurrentCommitCiCommitAuthorAvatar),
			CiCommitPreRelease:   c.Bool(wd_flag.NameCliCurrentCommitCiCommitPreRelease),
		},

		CurrentPipelineInfo: wd_info.CurrentPipelineInfo{
			CiPipelineNumber:       c.String(wd_flag.NameCliCurrentPipelineNumber),
			CiPipelineParent:       c.String(wd_flag.NameCliCurrentPipelineParent),
			CiPipelineEvent:        c.String(wd_flag.NameCliCurrentPipelineEvent),
			CiPipelineUrl:          c.String(wd_flag.NameCliCurrentPipelineUrl),
			CiPipelineForgeUrl:     c.String(wd_flag.NameCliCurrentPipelineForgeUrl),
			CiPipelineDeployTarget: c.String(wd_flag.NameCliCurrentPipelineDeployTarget),
			// CiPipelineStatus:        c.String(wd_flag.NameCliCurrentPipelineStatus),
			CiPipelineCreated:       pipelineCreateAt,
			CiPipelineCreatedT:      pipelineCreateAtT,
			CiPipelineStarted:       pipelineStartAt,
			CiPipelineStartedT:      pipelineStartAtT,
			CiPipelineDurationHuman: pipelineDurationHuman,
		},

		CurrentWorkflowInfo: wd_info.CurrentWorkflowInfo{
			CiWorkflowName:   c.String(wd_flag.NameCliCurrentCiWorkflowName),
			CiWorkflowNumber: c.String(wd_flag.NameCliCurrentCiWorkflowNumber),
		},

		CurrentStepInfo: wd_info.CurrentStepInfo{
			CiStepName:          c.String(wd_flag.NameCliCurrentCiStepName),
			CiStepNumber:        c.String(wd_flag.NameCliCurrentCiStepNumber),
			CiStepStarted:       stepStartedAt,
			CiStepStartedT:      stepStartedAtT,
			CiStepNowTime:       stepNowTimeAt,
			CiStepNowTimeT:      stepNowTimeAtT,
			CiStepDurationHuman: stepDurationHuman,
			CiStepUrl:           c.String(wd_flag.NameCliCurrentCiStepUrl),
		},
	}

	previousInfo := wd_info.PreviousInfo{
		PreviousCommitInfo: wd_info.PreviousCommitInfo{
			CiPreviousCommitSha:          c.String(wd_flag.NameCliPreviousCommitSha),
			CiPreviousCommitRef:          c.String(wd_flag.NameCliPreviousCommitRef),
			CiPreviousCommitRefSpec:      c.String(wd_flag.NameCliPreviousCommitRefSpec),
			CiPreviousCommitBranch:       c.String(wd_flag.NameCliPreviousCommitBranch),
			CiPreviousCommitSourceBranch: c.String(wd_flag.NameCliPreviousCommitSourceBranch),
			CiPreviousCommitTargetBranch: c.String(wd_flag.NameCliPreviousCommitTargetBranch),
			CiPreviousCommitUrl:          c.String(wd_flag.NameCliPreviousCommitUrl),
			CiPreviousCommitMessage:      c.String(wd_flag.NameCliPreviousCommitMessage),
			CiPreviousCommitAuthor:       c.String(wd_flag.NameCliPreviousCommitAuthor),
			CiPreviousCommitAuthorEmail:  c.String(wd_flag.NameCliPreviousCommitAuthorEmail),
			CiPreviousCommitAuthorAvatar: c.String(wd_flag.EnvKeyCurrentCommitCiCommitAuthorAvatar),
		},

		PreviousPipelineInfo: wd_info.PreviousPipelineInfo{
			CiPreviousPipelineNumber:        c.String(wd_flag.NameCliPreviousPipelineNumber),
			CiPreviousPipelineParent:        c.String(wd_flag.NameCliPreviousPipelineParent),
			CiPreviousPipelineEvent:         c.String(wd_flag.NameCliPreviousPipelineEvent),
			CiPreviousPipelineUrl:           c.String(wd_flag.NameCliPreviousPipelineUrl),
			CiPreviousPipelineForgeUrl:      c.String(wd_flag.NameCliPreviousPipelineForgeUrl),
			CiPreviousPipelineDeployTarget:  c.String(wd_flag.NameCliPreviousPipelineDeployTarget),
			CiPreviousPipelineStatus:        c.String(wd_flag.NameCliPreviousPipelineStatus),
			CiPreviousPipelineCreated:       previousPipelineCreateAt,
			CiPreviousPipelineCreatedT:      previousPipelineCreateAtT,
			CiPreviousPipelineStarted:       previousPipelineStartedAt,
			CiPreviousPipelineStartedT:      previousPipelineStartedAtT,
			CiPreviousPipelineFinished:      previousPipelineFinishedAt,
			CiPreviousPipelineFinishedT:     previousPipelineFinishedAtT,
			CiPreviousPipelineDurationHuman: previousPipelineDurationHuman,
		},
	}

	wdInfo := wd_info.WoodpeckerInfo{
		BasicInfo: wd_info.BasicInfo{
			CI:          c.String(wd_flag.NameCliCi),
			CIWorkspace: c.String(wd_flag.NameCliCiWorkspace),
		},

		CiSystemInfo: wd_info.CiSystemInfo{
			WoodpeckerBackend:       c.String(wd_flag.NameCliWoodpeckerBackend),
			WoodpeckerAgentHostName: c.String(wd_flag.NameCliWoodpeckerHostName),
			WoodpeckerFilterLabels:  c.StringSlice(wd_flag.NameCliWoodpeckerFilterLabels),
			CiMachine:               c.String(wd_flag.NameCliCiMachine),
			CiSystemPlatform:        c.String(wd_flag.NameCliCiSystemPlatform),
			CiSystemName:            c.String(wd_flag.NameCliCiSystemName),
			CiSystemUrl:             c.String(wd_flag.NameCliCiSystemUrl),
			CiSystemHost:            c.String(wd_flag.NameCliCiSystemHost),
			CiSystemVersion:         c.String(wd_flag.NameCliCiSystemVersion),
		},

		CiForgeInfo: wd_info.CiForgeInfo{
			CiForgeType: c.String(wd_flag.NameCliCiForgeType),
			CiForgeUrl:  c.String(wd_flag.NameCliCiForgeUrl),
		},

		RepositoryInfo: wd_info.RepositoryInfo{
			CIRepo:              c.String(wd_flag.NameCliRepositoryCiRepo),
			CIRepoOwner:         c.String(wd_flag.NameCliRepositoryCiOwner),
			CIRepoName:          c.String(wd_flag.NameCliRepositoryCiName),
			CIRepoRemoteID:      c.String(wd_flag.NameCliRepositoryCiRemoteId),
			CIRepoScm:           c.String(wd_flag.NameCliRepositoryCiScm),
			CIRepoURL:           c.String(wd_flag.NameCliRepositoryCiUrl),
			CIRepoCloneURL:      c.String(wd_flag.NameCliRepositoryCiCloneUrl),
			CIRepoCloneSshUrl:   c.String(wd_flag.NameCliRepositoryCiCloneSshUrl),
			CIRepoDefaultBranch: c.String(wd_flag.NameCliRepositoryCiDefaultBranch),
			CIRepoPrivate:       c.Bool(wd_flag.NameCliRepositoryCiPrivate),
			CIRepoTrusted:       c.Bool(wd_flag.NameCliRepositoryCiTrusted),
		},
		CurrentInfo: currentInfo,

		PreviousInfo: previousInfo,
	}

	_ = wd_info_parse.ParseRepositoryInfoByWoodPeckerInfo(&wdInfo)

	return wdInfo
}
