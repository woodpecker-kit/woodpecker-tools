package wd_info

type (
	// BasicInfo
	//  woodpecker basic info.
	BasicInfo struct {
		// CI
		//  Provides the ci flag. This value is `woodpecker` when the is run in woodpecker
		//  by env: CI
		CI string

		// CIWorkspace
		//  Provides the Path of the workspace where source code gets cloned to
		//  by env: CI_WORKSPACE
		CIWorkspace string
	}

	// CiSystemInfo
	//  woodpecker ci system info.
	CiSystemInfo struct {
		// CiSystemName
		//  Provides name of the CI system: woodpecker
		//  by env: CI_SYSTEM_NAME
		CiSystemName string

		// CiSystemUrl
		//  Provides link to CI system
		//  by env: CI_SYSTEM_URL
		CiSystemUrl string

		// CiSystemHost
		//  Provides hostname of CI server
		//  by env: CI_SYSTEM_HOST
		CiSystemHost string

		// CiSystemVersion
		//  Provides version of CI server
		//  by env: CI_SYSTEM_VERSION
		CiSystemVersion string
	}

	// CiForgeInfo
	//  woodpecker ci forge info.
	CiForgeInfo struct {
		// CiForgeType
		//  Provides name of forge (gitea, github, ...)
		//  by env: CI_FORGE_TYPE
		CiForgeType string

		// CiForgeUrl
		//  Provides root URL of configured forge
		//  by env: CI_FORGE_URL
		CiForgeUrl string
	}

	// RepositoryInfo
	//  woodpecker repository info.
	RepositoryInfo struct {
		// CIRepo
		//  Provides the repository repo flag. This value is `<owner>/<name>` when the is run in woodpecker
		//  by env: CI_REPO
		CIRepo string

		// CIRepoOwner
		//  Provides the repository owner flag. This value is `<owner>` when the is run in woodpecker
		//  by env: CI_REPO_OWNER
		CIRepoOwner string

		// CIRepoName
		//  Provides the repository name flag. This value is `<name>` when the is run in woodpecker
		//  by env: CI_REPO_NAME
		CIRepoName string

		// CIRepoRemoteID
		//  Provides the repository remote ID, is the UID it has in the forge.
		//  by env: CI_REPO_REMOTE_ID
		CIRepoRemoteID string

		// CIRepoScm
		//  Provides the repository scm flag. repository SCM (git).
		//  by env: CI_REPO_SCM
		CIRepoScm string

		// CIRepoURL
		//  Provides the repository url flag. repository URL.
		//  by env: CI_REPO_URL
		CIRepoURL string

		// CIRepoCloneURL
		//  Provides the repository clone url flag. repository clone URL.
		//  by env: CI_REPO_CLONE_URL
		CIRepoCloneURL string

		// CIRepoCloneSshUrl
		//  Provides the repository clone ssh url flag. repository clone ssh URL.
		//  by env: CI_REPO_CLONE_SSH_URL
		CIRepoCloneSshUrl string

		// CIRepoDefaultBranch
		//  Provides the repository default branch flag. repository default branch.
		//  by env: CI_REPO_DEFAULT_BRANCH
		CIRepoDefaultBranch string

		// CIRepoPrivate
		//  Provides the repository private flag. repository private.
		//  by env: CI_REPO_PRIVATE
		CIRepoPrivate bool

		// CIRepoTrusted
		//  Provides the repository trusted flag. repository trusted.
		//  by env: CI_REPO_TRUSTED
		CIRepoTrusted bool
	}

	// CurrentCommitInfo
	//  woodpecker current commit info.
	CurrentCommitInfo struct {
		// CiCommitSha
		//  Provides the current commit SHA
		//  by env: CI_COMMIT_SHA
		CiCommitSha string

		// CiCommitRef
		//  Provides the current commit ref
		//  by env: CI_COMMIT_REF
		CiCommitRef string

		// CiCommitRefSpec
		//  Provides the current commit ref spec
		//  by env: CI_COMMIT_REFSPEC
		CiCommitRefSpec string

		// CiCommitBranch
		//  Provides the current commit branch (equals target branch for pull requests)
		//  by env: CI_COMMIT_BRANCH
		CiCommitBranch string

		// CiCommitSourceBranch
		//  Provides the current commit source branch (empty if event is not pull_request)
		//  by env: CI_COMMIT_SOURCE_BRANCH
		CiCommitSourceBranch string

		// CiCommitTargetBranch
		//  Provides the current commit target branch (empty if event is not pull_request)
		//  by env: CI_COMMIT_TARGET_BRANCH
		CiCommitTargetBranch string

		// CiCommitTag
		//  Provides the current commit tag (empty if event is not tag)
		//  by env: CI_COMMIT_TAG
		CiCommitTag string

		// CiCommitPullRequest
		//  Provides the current commit pull request (empty if event is not pull_request)
		//  by env: CI_COMMIT_PULL_REQUEST
		CiCommitPullRequest string

		// CiCommitPullRequestLabels
		//  Provides the current labels assigned to pull request (empty if event is not pull_request)
		//  by env: CI_COMMIT_PULL_REQUEST_LABELS
		CiCommitPullRequestLabels string

		// CiCommitMessage
		//  Provides the current commit message
		//  by env: CI_COMMIT_MESSAGE
		CiCommitMessage string

		// CiCommitAuthorName
		//  Provides the current commit author name
		//  by env: CI_COMMIT_AUTHOR_NAME
		CiCommitAuthorName string

		// CiCommitAuthorEmail
		//  Provides the current commit author email
		//  by env: CI_COMMIT_AUTHOR_EMAIL
		CiCommitAuthorEmail string

		// CiCommitAuthorAvatar
		//  Provides the current commit author avatar
		//  by env: CI_COMMIT_AUTHOR_AVATAR
		CiCommitAuthorAvatar string
	}

	// CurrentPipelineInfo
	//  woodpecker current pipeline info.
	CurrentPipelineInfo struct {
		// CiPipelineNumber
		//  Provides the current pipeline number
		//  by env: CI_PIPELINE_NUMBER
		CiPipelineNumber string

		// CiPipelineParent
		//  Provides the current pipeline parent
		//  by env: CI_PIPELINE_PARENT
		CiPipelineParent string

		// CiPipelineEvent
		//  Provides the current pipeline event
		//  by env: CI_PIPELINE_EVENT
		CiPipelineEvent string

		// CiPipelineUrl
		//  Provides the current pipeline url
		//  by env: CI_PIPELINE_URL
		CiPipelineUrl string

		// CiPipelineForgeUrl
		//  Provides the current pipeline forge url
		//  by env: CI_PIPELINE_FORGE_URL
		CiPipelineForgeUrl string

		// CiPipelineDeployTarget
		//  Provides the current pipeline deploy target
		//  by env: CI_PIPELINE_DEPLOY_TARGET
		CiPipelineDeployTarget string

		// CiPipelineStatus
		//  Provides the current pipeline status
		//  by env: CI_PIPELINE_STATUS
		CiPipelineStatus string

		// CiPipelineCreated
		//  Provides the current pipeline created
		//  by env: CI_PIPELINE_CREATED
		CiPipelineCreated uint64

		// CiPipelineStarted
		//  Provides the current pipeline started
		//  by env: CI_PIPELINE_STARTED
		CiPipelineStarted uint64

		// CiPipelineFinished
		//  Provides the current pipeline finished
		//  by env: CI_PIPELINE_FINISHED
		CiPipelineFinished uint64
	}

	// CurrentWorkflowInfo
	//  woodpecker current workflow info.
	CurrentWorkflowInfo struct {
		// CiWorkflowName
		//  Provides the current workflow name
		//  by env: CI_WORKFLOW_NAME
		CiWorkflowName string
	}

	// CurrentStepInfo
	//  woodpecker current step info.
	CurrentStepInfo struct {
		// CiStepName
		//  Provides the current step name
		//  by env: CI_STEP_NAME
		CiStepName string

		// CiStepNumber
		//  Provides the current step number
		//  by env: CI_STEP_NUMBER
		CiStepNumber string

		// CiStepStatus
		//  Provides the current step status (success, failure)
		//  by env: CI_STEP_STATUS
		CiStepStatus string

		// CiStepStarted
		//  Provides step started UNIX timestamp
		//  by env: CI_STEP_STARTED
		CiStepStarted uint64

		// CiStepFinished
		//  Provides the step finished UNIX timestamp
		//  by env: CI_STEP_FINISHED
		CiStepFinished uint64

		// CiStepUrl
		//  Provides the current URL to step in UI
		//  by env: CI_STEP_URL
		CiStepUrl string
	}

	// CurrentInfo
	//  woodpecker current info.
	CurrentInfo struct {
		CurrentCommitInfo CurrentCommitInfo

		CurrentPipelineInfo CurrentPipelineInfo

		CurrentWorkflowInfo CurrentWorkflowInfo

		CurrentStepInfo CurrentStepInfo
	}

	// PreviousCommitInfo
	//  woodpecker previous commit info.
	PreviousCommitInfo struct {
		// CiPreviousCommitSha
		//  Provides the previous commit SHA
		//  by env: CI_PREV_COMMIT_SHA
		CiPreviousCommitSha string

		// CiPreviousCommitRef
		//  Provides the previous commit ref
		//  by env: CI_PREV_COMMIT_REF
		CiPreviousCommitRef string

		// CiPreviousCommitRefSpec
		//  Provides the previous commit ref spec
		//  by env: CI_PREV_COMMIT_REFSPEC
		CiPreviousCommitRefSpec string

		// CiPreviousCommitBranch
		//  Provides the previous commit branch
		//  by env: CI_PREV_COMMIT_BRANCH
		CiPreviousCommitBranch string

		// CiPreviousCommitSourceBranch
		//  Provides the previous commit source branch
		//  by env: CI_PREV_COMMIT_SOURCE_BRANCH
		CiPreviousCommitSourceBranch string

		// CiPreviousCommitTargetBranch
		//  Provides the previous commit target branch
		//  by env: CI_PREV_COMMIT_TARGET_BRANCH
		CiPreviousCommitTargetBranch string

		// CiPreviousCommitUrl
		//  Provides the previous commit url
		//  by env: CI_PREV_COMMIT_URL
		CiPreviousCommitUrl string

		// CiPreviousCommitMessage
		//  Provides the previous commit message
		//  by env: CI_PREV_COMMIT_MESSAGE
		CiPreviousCommitMessage string

		// CiPreviousCommitAuthor
		//  Provides the previous commit author username
		//  by env: CI_PREV_COMMIT_AUTHOR
		CiPreviousCommitAuthor string

		// CiPreviousCommitAuthorEmail
		//  Provides the previous commit author email
		//  by env: CI_PREV_COMMIT_AUTHOR_EMAIL
		CiPreviousCommitAuthorEmail string

		// CiPreviousCommitAuthorAvatar
		//  Provides the previous commit author avatar
		//  by env: CI_PREV_COMMIT_AUTHOR_AVATAR
		CiPreviousCommitAuthorAvatar string
	}

	// PreviousPipelineInfo
	//  woodpecker previous pipeline info.
	PreviousPipelineInfo struct {
		// CiPreviousPipelineNumber
		//  Provides the previous pipeline number
		//  by env: CI_PREV_PIPELINE_NUMBER
		CiPreviousPipelineNumber string

		// CiPreviousPipelineParent
		//  Provides the previous pipeline number of parent pipeline
		//  by env: CI_PREV_PIPELINE_PARENT
		CiPreviousPipelineParent string

		// CiPreviousPipelineEvent
		//  Provides the previous pipeline event (push, pull_request, tag, deployment)
		//  by env: CI_PREV_PIPELINE_EVENT
		CiPreviousPipelineEvent string

		// CiPreviousPipelineUrl
		//  Provides the previous pipeline link in CI
		//  by env: CI_PREV_PIPELINE_URL
		CiPreviousPipelineUrl string

		// CiPreviousPipelineForgeUrl
		//  Provides the previous pipeline link in forge
		//  by env: CI_PREV_PIPELINE_FORGE_URL
		CiPreviousPipelineForgeUrl string

		// CiPreviousPipelineDeployTarget
		//  Provides the previous pipeline deploy target for deployment events (ie production)
		//  by env: CI_PREV_PIPELINE_DEPLOY_TARGET
		CiPreviousPipelineDeployTarget string

		// CiPreviousPipelineStatus
		//  Provides the previous pipeline status (success, failure)
		//  by env: CI_PREV_PIPELINE_STATUS
		CiPreviousPipelineStatus string

		// CiPreviousPipelineCreated
		//  Provides the previous pipeline created
		//  by env: CI_PREV_PIPELINE_CREATED
		CiPreviousPipelineCreated uint64

		// CiPreviousPipelineStarted
		//  Provides the previous pipeline started
		//  by env: CI_PREV_PIPELINE_STARTED
		CiPreviousPipelineStarted uint64

		// CiPreviousPipelineFinished
		//  Provides the previous pipeline finished
		//  by env: CI_PREV_PIPELINE_FINISHED
		CiPreviousPipelineFinished uint64
	}

	// PreviousInfo
	//  woodpecker previous info.
	PreviousInfo struct {
		PreviousCommitInfo PreviousCommitInfo

		PreviousPipelineInfo PreviousPipelineInfo
	}

	WoodpeckerInfo struct {
		BasicInfo BasicInfo

		CiSystemInfo CiSystemInfo

		CiForgeInfo CiForgeInfo

		RepositoryInfo RepositoryInfo

		CurrentInfo CurrentInfo

		PreviousInfo PreviousInfo
	}
)
