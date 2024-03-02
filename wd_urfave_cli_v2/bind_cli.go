package wd_urfave_cli_v2

import (
	"github.com/urfave/cli/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
)

// UrfaveCliBindInfo
//
//	woodpecker urfave cli bind info
func UrfaveCliBindInfo(c *cli.Context) wd_info.WoodpeckerInfo {

	pipelineCreateAt := c.Uint64(wd_flag.NameCliCurrentPipelineCreated)
	pipelineCreateAtT := wd_flag.FormatTimeUTC(pipelineCreateAt, wd_flag.TimeFormatDefault)
	pipelineStartAt := c.Uint64(wd_flag.NameCliCurrentPipelineStarted)
	pipelineStartAtT := wd_flag.FormatTimeUTC(pipelineStartAt, wd_flag.TimeFormatDefault)
	pipelineFinishAt := c.Uint64(wd_flag.NameCliCurrentPipelineFinished)
	pipelineFinishAtT := wd_flag.FormatTimeUTC(pipelineFinishAt, wd_flag.TimeFormatDefault)

	previousPipelineCreateAt := c.Uint64(wd_flag.NameCliPreviousPipelineCreated)
	previousPipelineCreateAtT := wd_flag.FormatTimeUTC(previousPipelineCreateAt, wd_flag.TimeFormatDefault)
	previousPipelineStartedAt := c.Uint64(wd_flag.NameCliPreviousPipelineStarted)
	previousPipelineStartedAtT := wd_flag.FormatTimeUTC(previousPipelineStartedAt, wd_flag.TimeFormatDefault)
	previousPipelineFinishedAt := c.Uint64(wd_flag.NameCliPreviousPipelineFinished)
	previousPipelineFinishedAtT := wd_flag.FormatTimeUTC(previousPipelineFinishedAt, wd_flag.TimeFormatDefault)

	stepStartedAt := c.Uint64(wd_flag.NameCliCurrentCiStepStarted)
	stepStartedAtT := wd_flag.FormatTimeUTC(stepStartedAt, wd_flag.TimeFormatDefault)
	stepFinishedAt := c.Uint64(wd_flag.NameCliCurrentCiStepFinished)
	stepFinishedAtT := wd_flag.FormatTimeUTC(stepFinishedAt, wd_flag.TimeFormatDefault)

	var currentInfo = wd_info.CurrentInfo{
		CurrentCommitInfo: wd_info.CurrentCommitInfo{
			CiCommitSha:               c.String(wd_flag.NameCliCurrentCommitCiCommitSha),
			CiCommitRef:               c.String(wd_flag.NameCliCurrentCommitCiCommitRef),
			CiCommitRefSpec:           c.String(wd_flag.NameCliCurrentCommitCiCommitRefSpec),
			CiCommitUrl:               c.String(wd_flag.NameCliCurrentCommitCiCommitUrl),
			CiCommitBranch:            c.String(wd_flag.NameCliCurrentCommitCiCommitBranch),
			CiCommitSourceBranch:      c.String(wd_flag.NameCliCurrentCommitCiCommitSourceBranch),
			CiCommitTargetBranch:      c.String(wd_flag.NameCliCurrentCommitCiCommitTargetBranch),
			CiCommitTag:               c.String(wd_flag.NameCliCurrentCommitCiCommitTag),
			CiCommitPullRequest:       c.String(wd_flag.NameCliCurrentCommitCiCommitPullRequest),
			CiCommitPullRequestLabels: c.String(wd_flag.NameCliCurrentCommitCiCommitPullRequestLabels),
			CiCommitMessage:           c.String(wd_flag.NameCliCurrentCommitCiCommitMessage),
			CiCommitAuthor:            c.String(wd_flag.NameCliCurrentCommitCiCommitAuthor),
			CiCommitAuthorEmail:       c.String(wd_flag.NameCliCurrentCommitCiCommitAuthorEmail),
			CiCommitAuthorAvatar:      c.String(wd_flag.NameCliCurrentCommitCiCommitAuthorAvatar),
			CiCommitPreRelease:        c.Bool(wd_flag.NameCliCurrentCommitCiCommitPreRelease),
		},

		CurrentPipelineInfo: wd_info.CurrentPipelineInfo{
			CiPipelineFiles:        c.String(wd_flag.NameCliCurrentPipelineFiles),
			CiPipelineNumber:       c.String(wd_flag.NameCliCurrentPipelineNumber),
			CiPipelineParent:       c.String(wd_flag.NameCliCurrentPipelineParent),
			CiPipelineEvent:        c.String(wd_flag.NameCliCurrentPipelineEvent),
			CiPipelineUrl:          c.String(wd_flag.NameCliCurrentPipelineUrl),
			CiPipelineForgeUrl:     c.String(wd_flag.NameCliCurrentPipelineForgeUrl),
			CiPipelineDeployTarget: c.String(wd_flag.NameCliCurrentPipelineDeployTarget),
			CiPipelineStatus:       c.String(wd_flag.NameCliCurrentPipelineStatus),
			CiPipelineCreated:      pipelineCreateAt,
			CiPipelineCreatedT:     pipelineCreateAtT,
			CiPipelineStarted:      pipelineStartAt,
			CiPipelineStartedT:     pipelineStartAtT,
			CiPipelineFinished:     pipelineFinishAt,
			CiPipelineFinishedT:    pipelineFinishAtT,
		},

		CurrentWorkflowInfo: wd_info.CurrentWorkflowInfo{
			CiWorkflowName:   c.String(wd_flag.NameCliCurrentCiWorkflowName),
			CiWorkflowNumber: c.String(wd_flag.NameCliCurrentCiWorkflowNumber),
		},

		CurrentStepInfo: wd_info.CurrentStepInfo{
			CiStepName:      c.String(wd_flag.NameCliCurrentCiStepName),
			CiStepNumber:    c.String(wd_flag.NameCliCurrentCiStepNumber),
			CiStepStatus:    c.String(wd_flag.NameCliCurrentCiStepStatus),
			CiStepStarted:   stepStartedAt,
			CiStepStartedT:  stepStartedAtT,
			CiStepFinished:  stepFinishedAt,
			CiStepFinishedT: stepFinishedAtT,
			CiStepUrl:       c.String(wd_flag.NameCliCurrentCiStepUrl),
		},
	}

	var previousInfo = wd_info.PreviousInfo{
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
			CiPreviousPipelineNumber:       c.String(wd_flag.NameCliPreviousPipelineNumber),
			CiPreviousPipelineParent:       c.String(wd_flag.NameCliPreviousPipelineParent),
			CiPreviousPipelineEvent:        c.String(wd_flag.NameCliPreviousPipelineEvent),
			CiPreviousPipelineUrl:          c.String(wd_flag.NameCliPreviousPipelineUrl),
			CiPreviousPipelineForgeUrl:     c.String(wd_flag.NameCliPreviousPipelineForgeUrl),
			CiPreviousPipelineDeployTarget: c.String(wd_flag.NameCliPreviousPipelineDeployTarget),
			CiPreviousPipelineStatus:       c.String(wd_flag.NameCliPreviousPipelineStatus),
			CiPreviousPipelineCreated:      previousPipelineCreateAt,
			CiPreviousPipelineCreatedT:     previousPipelineCreateAtT,
			CiPreviousPipelineStarted:      previousPipelineStartedAt,
			CiPreviousPipelineStartedT:     previousPipelineStartedAtT,
			CiPreviousPipelineFinished:     previousPipelineFinishedAt,
			CiPreviousPipelineFinishedT:    previousPipelineFinishedAtT,
		},
	}

	var wdInfo = wd_info.WoodpeckerInfo{
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
	return wdInfo
}
