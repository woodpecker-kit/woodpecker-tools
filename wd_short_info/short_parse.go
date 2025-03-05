package wd_short_info

import (
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_flag"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
)

func ParseWoodpeckerInfo2Short(info wd_info.WoodpeckerInfo) WoodpeckerInfoShort {
	short := WoodpeckerInfoShort{
		CurrentWorkflow: CurrentWorkflow{
			Name:   info.CurrentInfo.CurrentWorkflowInfo.CiWorkflowName,
			Number: info.CurrentInfo.CurrentWorkflowInfo.CiWorkflowNumber,
		},
		Repo: Repo{
			Scm:       info.RepositoryInfo.CIRepoScm,
			Link:      info.RepositoryInfo.CIRepoURL,
			FullName:  info.RepositoryInfo.CIRepo,
			ShortName: info.RepositoryInfo.CIRepoName,
			OwnerName: info.RepositoryInfo.CIRepoOwner,
			HttpUrl:   info.RepositoryInfo.CIRepoCloneURL,
			SshUrl:    info.RepositoryInfo.CIRepoCloneSshUrl,
			Host:      info.RepositoryInfo.CiRepoHost,
			Hostname:  info.RepositoryInfo.CiRepoHostname,
			Port:      info.RepositoryInfo.CiRepoPort,
		},
		System: System{
			Version:  info.CiSystemInfo.CiSystemVersion,
			Url:      info.CiSystemInfo.CiSystemUrl,
			Hostname: info.CiSystemInfo.CiSystemHost,
			Name:     info.CiSystemInfo.CiSystemName,

			RunnerBackend:  info.CiSystemInfo.WoodpeckerBackend,
			RunnerMachine:  info.CiSystemInfo.CiMachine,
			RunnerHostName: info.CiSystemInfo.WoodpeckerAgentHostName,
			RunnerPlatform: info.CiSystemInfo.CiSystemPlatform,
		},
		Build: Build{
			WorkSpace:     info.BasicInfo.CIWorkspace,
			Status:        info.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus,
			Number:        info.CurrentInfo.CurrentPipelineInfo.CiPipelineNumber,
			LinkCi:        info.CurrentInfo.CurrentPipelineInfo.CiPipelineUrl,
			LinkCommit:    info.CurrentInfo.CurrentPipelineInfo.CiPipelineForgeUrl,
			Event:         info.CurrentInfo.CurrentPipelineInfo.CiPipelineEvent,
			Sha:           info.CurrentInfo.CurrentCommitInfo.CiCommitSha,
			Ref:           info.CurrentInfo.CurrentCommitInfo.CiCommitRef,
			RefSpec:       info.CurrentInfo.CurrentCommitInfo.CiCommitRefSpec,
			Message:       info.CurrentInfo.CurrentCommitInfo.CiCommitMessage,
			CommitBranch:  info.CurrentInfo.CurrentCommitInfo.CiCommitBranch,
			Tag:           info.CurrentInfo.CurrentCommitInfo.CiCommitTag,
			PR:            info.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest,
			SourceBranch:  info.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch,
			TargetBranch:  info.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch,
			Created:       info.CurrentInfo.CurrentPipelineInfo.CiPipelineCreated,
			CreatedAt:     info.CurrentInfo.CurrentPipelineInfo.CiPipelineCreatedT,
			Started:       info.CurrentInfo.CurrentPipelineInfo.CiPipelineStarted,
			StartedAt:     info.CurrentInfo.CurrentPipelineInfo.CiPipelineStartedT,
			DurationHuman: info.CurrentInfo.CurrentPipelineInfo.CiPipelineDurationHuman,
		},
		Commit: Commit{
			Link:         info.CurrentInfo.CurrentPipelineInfo.CiPipelineForgeUrl,
			Sha:          info.CurrentInfo.CurrentCommitInfo.CiCommitSha,
			Ref:          info.CurrentInfo.CurrentCommitInfo.CiCommitRef,
			RefSpec:      info.CurrentInfo.CurrentCommitInfo.CiCommitRefSpec,
			CommitBranch: info.CurrentInfo.CurrentCommitInfo.CiCommitBranch,
			Message:      info.CurrentInfo.CurrentCommitInfo.CiCommitMessage,
			Tag:          info.CurrentInfo.CurrentCommitInfo.CiCommitTag,
			PR:           info.CurrentInfo.CurrentCommitInfo.CiCommitPullRequest,
			SourceBranch: info.CurrentInfo.CurrentCommitInfo.CiCommitSourceBranch,
			TargetBranch: info.CurrentInfo.CurrentCommitInfo.CiCommitTargetBranch,
			PreRelease:   info.CurrentInfo.CurrentCommitInfo.CiCommitPreRelease,
			CommitAuthor: CommitAuthor{
				Username: info.CurrentInfo.CurrentCommitInfo.CiCommitAuthor,
				Email:    info.CurrentInfo.CurrentCommitInfo.CiCommitAuthorEmail,
				Avatar:   info.CurrentInfo.CurrentCommitInfo.CiCommitAuthorAvatar,
			},
		},
		Stage: Stage{
			Name:      info.CurrentInfo.CurrentStepInfo.CiStepName,
			NumStep:   info.CurrentInfo.CurrentStepInfo.CiStepNumber,
			Started:   info.CurrentInfo.CurrentStepInfo.CiStepStarted,
			StartedAt: info.CurrentInfo.CurrentStepInfo.CiStepStartedT,
			//Finished:          info.CurrentInfo.CurrentStepInfo.CiStepFinished,
			//FinishedAt:        info.CurrentInfo.CurrentStepInfo.CiStepFinishedT,
			NowTime:           info.CurrentInfo.CurrentStepInfo.CiStepNowTime,
			NowTimeT:          info.CurrentInfo.CurrentStepInfo.CiStepNowTimeT,
			DurationStepHuman: info.CurrentInfo.CurrentStepInfo.CiStepDurationHuman,
		},
		PrevBuild: PrevBuild{
			Status:        info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineStatus,
			Number:        info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineNumber,
			LinkCi:        info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineUrl,
			LinkCommit:    info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineForgeUrl,
			Event:         info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineEvent,
			Sha:           info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitSha,
			Ref:           info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitRef,
			RefSpec:       info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitRefSpec,
			Message:       info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitMessage,
			CommitBranch:  info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitBranch,
			SourceBranch:  info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitSourceBranch,
			TargetBranch:  info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitTargetBranch,
			Created:       info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineCreated,
			CreatedAt:     info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineCreatedT,
			Started:       info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineStarted,
			StartedAt:     info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineStartedT,
			Finished:      info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineFinished,
			FinishedAt:    info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineFinishedT,
			DurationHuman: info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineDurationHuman,
			DeployTarget:  info.PreviousInfo.PreviousPipelineInfo.CiPreviousPipelineDeployTarget,
			CommitAuthor: CommitAuthor{
				Username: info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitAuthor,
				Email:    info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitAuthorEmail,
				Avatar:   info.PreviousInfo.PreviousCommitInfo.CiPreviousCommitAuthorAvatar,
			},
		},
	}

	relatedShortInfo(&short)

	return short
}

func relatedShortInfo(w *WoodpeckerInfoShort) {
	if w.Build.Status == "" {
		// https://github.com/woodpecker-ci/woodpecker/pull/4193
		relateCIPipelineStatus := env_kit.FetchOsEnvStr(wd_flag.RelatedEnvDroneCIPipelineStatus, "")
		if relateCIPipelineStatus != "" {
			w.Build.Status = relateCIPipelineStatus
		}
	}
}
