package wd_info_shot

import _ "github.com/woodpecker-kit/woodpecker-tools/wd_info"

type (
	// Repo repo base info
	Repo struct {
		// Scm
		//	must is: git hg
		// from wd_info.RepositoryInfo member CIRepoScm
		Scm string
		// Link
		// from wd_info.RepositoryInfo member CIRepoURL
		Link string
		// FullName
		//  repository full name
		// from wd_info.RepositoryInfo member CIRepo
		FullName string
		// ShortName
		// short name
		// from wd_info.RepositoryInfo member CIRepoName
		ShortName string
		// OwnerName
		//  repo owner name
		// from wd_info.RepositoryInfo member CIRepoOwner
		OwnerName string
		// HttpUrl
		// from wd_info.RepositoryInfo member CIRepoCloneURL
		HttpUrl string
		// SshUrl
		// from wd_info.RepositoryInfo member CIRepoCloneSshUrl
		SshUrl string
		// Host
		// this from HttpUrl host
		// from wd_info.RepositoryInfo member CiRepoHost
		Host string
		// Hostname
		// this from HttpUrl hostname
		// from wd_info.RepositoryInfo member CiRepoHostname
		Hostname string
		// Port
		// this from HttpUrl port
		// from wd_info.RepositoryInfo member CiRepoPort
		Port string
	}

	System struct {
		// Version
		// ci system version
		// from wd_info.CiSystemInfo member CiSystemVersion
		Version string

		// Url
		// ci system url
		// from wd_info.CiSystemInfo member CiSystemUrl
		Url string

		// Hostname
		// ci system host name
		// from wd_info.CiSystemInfo member CiSystemHost
		Hostname string

		// Name
		// the CI system name
		// from wd_info.CiSystemInfo member CiSystemName
		Name string

		// RunnerBackend
		// Possible values are auto-detect, docker, local or kubernetes
		// from wd_info.CiSystemInfo member WoodpeckerBackend
		RunnerBackend string

		// RunnerMachine
		// name of the CI machine
		// from wd_info.CiSystemInfo member CiMachine
		RunnerMachine string

		// RunnerHostName
		// host name of the CI machine
		// from wd_info.CiSystemInfo member WoodpeckerAgentHostName
		RunnerHostName string

		// RunnerPlatform
		//  Provides ci system platform as: linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64
		//  The syntax of the platform is GOOS/GOARCH from https://go.dev/doc/install/source
		// from wd_info.CiSystemInfo member CiSystemPlatform
		RunnerPlatform string
	}

	Build struct {
		// WorkSpace
		// runner workspace
		// from wd_info.BasicInfo member CIWorkspace
		WorkSpace string
		// Status
		// from wd_info.CurrentPipelineInfo CiPipelineStatus
		Status string
		// Number
		// from wd_info.CurrentPipelineInfo CiPipelineNumber
		Number string
		// LinkCi
		//	just ci url
		// from wd_info.CurrentPipelineInfo CiPipelineLink
		LinkCi string
		// LinkCommit
		//	just git forge url
		// from wd_info.CurrentPipelineInfo CiPipelineForgeUrl
		LinkCommit string
		// Event
		// from wd_info.CurrentPipelineInfo CiPipelineEvent
		Event string
		// Sha
		// from wd_info.CurrentCommitInfo CiCommitSha
		Sha string
		// Ref
		// from wd_info.CurrentCommitInfo CiCommitRef
		Ref string
		// RefSpec
		// from wd_info.CurrentCommitInfo CiCommitRefSpec
		RefSpec string
		// Message
		// from wd_info.CurrentCommitInfo CiCommitMessage
		Message string
		// Branch
		// from wd_info.CurrentCommitInfo CiCommitBranch
		CommitBranch string
		// Tag
		// from wd_info.CurrentCommitInfo CiCommitTag
		Tag string
		// PR
		// from wd_info.CurrentCommitInfo CiCommitPullRequest
		PR string
		// SourceBranch
		//	only in PR has
		// from wd_info.CurrentCommitInfo CiCommitSourceBranch
		SourceBranch string
		// TargetBranch
		//	only in PR has
		// from wd_info.CurrentCommitInfo CiCommitTargetBranch
		TargetBranch string
		// Created
		// from wd_info.CurrentPipelineInfo CiPipelineCreated
		Created uint64
		// CreatedAt
		// from wd_info.CurrentPipelineInfo CiPipelineCreatedT
		CreatedAt string
		// Started
		// from wd_info.CurrentPipelineInfo CiPipelineStarted
		Started uint64
		// StartedAt
		// from wd_info.CurrentPipelineInfo CiPipelineStartedT
		StartedAt string
		// Finished
		// from wd_info.CurrentPipelineInfo CiPipelineFinished
		Finished uint64
		// FinishedAt
		// from wd_info.CurrentPipelineInfo CiPipelineFinishedT
		FinishedAt string
		// DurationHuman
		//	This value is calculated by subtracting the pipeline created timestamp from the build finished timestamp.
		// from wd_info.CurrentPipelineInfo CiPipelineDurationHuman
		DurationHuman string
	}

	Commit struct {
		Link         string
		Sha          string
		Ref          string
		RefSpec      string
		CommitBranch string
		Message      string
		Tag          string
		PR           string
		SourceBranch string
		TargetBranch string
		PreRelease   bool
		CommitAuthor CommitAuthor
	}

	Stage struct {
		// Name
		// from wd_info.CurrentStepInfo CiStepName
		Name string
		// NumStep
		// from wd_info.CurrentStepInfo CiStepNumber
		NumStep string
		// Status
		// from wd_info.CurrentStepInfo CiStepStatus
		Status string
		// Started
		// from wd_info.CurrentStepInfo CiStepStarted
		Started uint64
		// StartedAt
		// from wd_info.CurrentStepInfo CiStepStartedT
		StartedAt string
		// Finished
		// from wd_info.CurrentStepInfo CiStepFinished
		Finished uint64
		// FinishedAt
		// from wd_info.CurrentStepInfo CiStepFinishedT
		FinishedAt string
		// DurationStepHuman
		//	This value is calculated by subtracting the step created timestamp from the build finished timestamp.
		// from wd_info.CurrentStepInfo CiStepDurationHuman
		DurationStepHuman string
	}

	PrevBuild struct {
		// Status
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineStatus
		Status string
		// Number
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineNumber
		Number string
		// LinkCi
		//	just ci url
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineUrl
		LinkCi string
		// LinkCommit
		//	just git forge url
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineForgeUrl
		LinkCommit string
		// Event
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineEvent
		Event string
		// Sha
		// from wd_info.PreviousCommitInfo CiPreviousCommitSha
		Sha string
		// Ref
		// from wd_info.PreviousCommitInfo CiPreviousCommitRef
		Ref string
		// RefSpec
		// from wd_info.PreviousCommitInfo CiPreviousCommitRefSpec
		RefSpec string
		// Message
		// from wd_info.PreviousCommitInfo CiPreviousCommitMessage
		Message string
		// Branch
		// from wd_info.PreviousCommitInfo CiPreviousCommitBranch
		CommitBranch string
		// SourceBranch
		//	only in PR has
		// from wd_info.PreviousCommitInfo CiPreviousCommitSourceBranch
		SourceBranch string
		// TargetBranch
		//	only in PR has
		// from wd_info.PreviousCommitInfo CiPreviousCommitTargetBranch
		TargetBranch string
		// Created
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineCreated
		Created uint64
		// CreatedAt
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineCreatedT
		CreatedAt string
		// Started
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineStarted
		Started uint64
		// StartedAt
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineStartedT
		StartedAt string
		// Finished
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineFinished
		Finished uint64
		// FinishedAt
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineFinishedT
		FinishedAt string
		// DurationHuman
		//	This value is calculated by subtracting the pipeline created timestamp from the build finished timestamp.
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineDurationHuman
		DurationHuman string
		// DeployTarget
		// from wd_info.PreviousPipelineInfo CiPreviousPipelineDeployTarget
		DeployTarget string

		CommitAuthor CommitAuthor
	}

	WoodpeckerInfoShort struct {
		Repo      Repo
		System    System
		Build     Build
		Commit    Commit
		Stage     Stage
		PrevBuild PrevBuild
	}
)

// CommitAuthor commit author info
type CommitAuthor struct {
	//  Username
	Username string //   the author username for the current commit
	//  Email
	Email string //   providers the author email for the current commit
	//  Avatar
	Avatar string //   providers the author avatar for the current commit
}
