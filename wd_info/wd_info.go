package wd_info

type (
	// BasicInfo
	//  woodpecker basic info.
	BasicInfo struct {
		// CI
		//  Provides the ci flag. This value is `woodpecker` when the is run in woodpecker
		//  by env: CI
		CI string `mock_env_key:"CI" mock_env_default:"woodpecker"`

		// CIWorkspace
		//  Provides the Path of the workspace where source code gets cloned to
		//  by env: CI_WORKSPACE
		CIWorkspace string `mock_env_key:"CI_WORKSPACE"`
	}

	// CiSystemInfo
	//  woodpecker ci system info.
	CiSystemInfo struct {

		// WoodpeckerBackend
		//  Provides the backend flag. This value is type when the is run in woodpecker, default is empty.
		//  Possible values are auto-detect, docker, local or kubernetes
		//  by env: WOODPECKER_BACKEND
		WoodpeckerBackend string `mock_env_key:"WOODPECKER_BACKEND" mock_env_default:"docker"`

		// WoodpeckerAgentHostName
		//  Provides the configures the agent hostname. default is empty.
		//  by env: WOODPECKER_HOST_NAME
		WoodpeckerAgentHostName string `mock_env_key:"WOODPECKER_HOST_NAME"`

		// WoodpeckerFilterLabels
		//  Configures labels to filter pipeline pick up.
		//  Use a list of key-value pairs like key=value,second-key=*. * can be used as a wildcard.
		//  By default, agents provide three additional labels platform=os/arch, hostname=my-agent and repo=*
		//  most this will be empty
		//  by env: WOODPECKER_FILTER_LABELS
		WoodpeckerFilterLabels []string `mock_env_key:"WOODPECKER_FILTER_LABELS"`

		// CiMachine
		//  Provides name of the CI machine
		//  by env: CI_MACHINE
		CiMachine string `mock_env_key:"CI_MACHINE"`

		// CiSystemPlatform
		//  Provides ci system platform as: linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64
		//  The syntax of the platform is GOOS/GOARCH from https://go.dev/doc/install/source
		//  by env: CI_SYSTEM_PLATFORM
		CiSystemPlatform string `mock_env_key:"CI_SYSTEM_PLATFORM"`

		// CiSystemName
		//  Provides name of the CI system: woodpecker
		//  by env: CI_SYSTEM_NAME
		CiSystemName string `mock_env_key:"CI_SYSTEM_NAME" mock_env_default:"woodpecker"`

		// CiSystemUrl
		//  Provides link to CI system
		//  by env: CI_SYSTEM_URL
		CiSystemUrl string `mock_env_key:"CI_SYSTEM_URL"`

		// CiSystemHost
		//  Provides hostname of CI server
		//  by env: CI_SYSTEM_HOST
		CiSystemHost string `mock_env_key:"CI_SYSTEM_HOST"`

		// CiSystemVersion
		//  Provides version of CI server
		//  by env: CI_SYSTEM_VERSION
		CiSystemVersion string `mock_env_key:"CI_SYSTEM_VERSION"`
	}

	// CiForgeInfo
	//  woodpecker ci forge info.
	CiForgeInfo struct {
		// CiForgeType
		//  Provides name of forge (gitea, github, ...)
		//  by env: CI_FORGE_TYPE
		CiForgeType string `mock_env_key:"CI_FORGE_TYPE"`

		// CiForgeUrl
		//  Provides root URL of configured forge
		//  by env: CI_FORGE_URL
		CiForgeUrl string `mock_env_key:"CI_FORGE_URL"`
	}

	// RepositoryInfo
	//  woodpecker repository info.
	RepositoryInfo struct {
		// CIRepo
		//  Provides the repository repo flag. This value is `<owner>/<name>` when the is run in woodpecker
		//  by env: CI_REPO
		CIRepo string `mock_env_key:"CI_REPO"`

		// CIRepoOwner
		//  Provides the repository owner flag. This value is `<owner>` when the is run in woodpecker
		//  by env: CI_REPO_OWNER
		CIRepoOwner string `mock_env_key:"CI_REPO_OWNER"`

		// CIRepoName
		//  Provides the repository name flag. This value is `<name>` when the is run in woodpecker
		//  by env: CI_REPO_NAME
		CIRepoName string `mock_env_key:"CI_REPO_NAME"`

		// CIRepoRemoteID
		//  Provides the repository remote ID, is the UID it has in the forge.
		//  by env: CI_REPO_REMOTE_ID
		CIRepoRemoteID string `mock_env_key:"CI_REPO_REMOTE_ID"`

		// CIRepoScm
		//  Provides the repository scm flag. repository SCM (git).
		//  by env: CI_REPO_SCM
		CIRepoScm string `mock_env_key:"CI_REPO_SCM" mock_env_default:"git"`

		// CIRepoURL
		//  Provides the repository url flag. repository URL.
		//  by env: CI_REPO_URL
		CIRepoURL string `mock_env_key:"CI_REPO_URL"`

		// CIRepoCloneURL
		//  Provides the repository clone url flag. repository clone URL.
		//  by env: CI_REPO_CLONE_URL
		CIRepoCloneURL string `mock_env_key:"CI_REPO_CLONE_URL"`

		// CIRepoCloneSshUrl
		//  Provides the repository clone ssh url flag. repository clone ssh URL.
		//  by env: CI_REPO_CLONE_SSH_URL
		CIRepoCloneSshUrl string `mock_env_key:"CI_REPO_CLONE_SSH_URL"`

		// CIRepoDefaultBranch
		//  Provides the repository default branch flag. repository default branch.
		//  by env: CI_REPO_DEFAULT_BRANCH
		CIRepoDefaultBranch string `mock_env_key:"CI_REPO_DEFAULT_BRANCH"`

		// CIRepoPrivate
		//  Provides the repository private flag. repository private.
		//  by env: CI_REPO_PRIVATE
		CIRepoPrivate bool `mock_env_key:"CI_REPO_PRIVATE"`

		// CIRepoTrusted
		//  Provides the repository trusted flag. repository trusted.
		//  by env: CI_REPO_TRUSTED
		CIRepoTrusted bool `mock_env_key:"CI_REPO_TRUSTED"`

		// CIRepoHost
		//  Provides the repository host, parse from env `CI_REPO_CLONE_URL`
		CiRepoHost string

		// CIRepoHostname
		//  Provides the repository hostname, parse from env `CI_REPO_CLONE_URL`
		CiRepoHostname string

		// CIRepoPort
		//  Provides the repository port, parse from env `CI_REPO_CLONE_URL`
		CiRepoPort string
	}

	// CurrentCommitInfo
	//  woodpecker current commit info.
	CurrentCommitInfo struct {
		// CiCommitSha
		//  Provides the current commit SHA
		//  by env: CI_COMMIT_SHA
		CiCommitSha string `mock_env_key:"CI_COMMIT_SHA"`

		// CiCommitRef
		//  Provides the current commit ref
		//  by env: CI_COMMIT_REF
		CiCommitRef string `mock_env_key:"CI_COMMIT_REF"`

		// CiCommitRefSpec
		//  Provides the current commit ref spec
		//  by env: CI_COMMIT_REFSPEC
		CiCommitRefSpec string `mock_env_key:"CI_COMMIT_REFSPEC"`

		// CiCommitUrl
		//  Provides the current commit url
		//  by env: CI_COMMIT_URL
		CiCommitUrl string `mock_env_key:"CI_COMMIT_URL"`

		// CiCommitBranch
		//  Provides the current commit branch (equals target branch for pull requests)
		//  by env: CI_COMMIT_BRANCH
		CiCommitBranch string `mock_env_key:"CI_COMMIT_BRANCH"`

		// CiCommitSourceBranch
		//  Provides the current commit source branch (empty if event is not pull_request)
		//  by env: CI_COMMIT_SOURCE_BRANCH
		CiCommitSourceBranch string `mock_env_key:"CI_COMMIT_SOURCE_BRANCH"`

		// CiCommitTargetBranch
		//  Provides the current commit target branch (empty if event is not pull_request)
		//  by env: CI_COMMIT_TARGET_BRANCH
		CiCommitTargetBranch string `mock_env_key:"CI_COMMIT_TARGET_BRANCH"`

		// CiCommitTag
		//  Provides the current commit tag (empty if event is not tag)
		//  by env: CI_COMMIT_TAG
		CiCommitTag string `mock_env_key:"CI_COMMIT_TAG"`

		// CiCommitPullRequest
		//  Provides the current commit pull request (empty if event is not pull_request)
		//  by env: CI_COMMIT_PULL_REQUEST
		CiCommitPullRequest string `mock_env_key:"CI_COMMIT_PULL_REQUEST"`

		// CiCommitPullRequestLabels
		//  Provides the current labels assigned to pull request (empty if event is not pull_request)
		//  by env: CI_COMMIT_PULL_REQUEST_LABELS
		CiCommitPullRequestLabels string `mock_env_key:"CI_COMMIT_PULL_REQUEST_LABELS"`

		// CiCommitMessage
		//  Provides the current commit message
		//  by env: CI_COMMIT_MESSAGE
		CiCommitMessage string `mock_env_key:"CI_COMMIT_MESSAGE" string_line_2_raw:"ci_commit_message"`

		// CiCommitAuthor
		//  Provides the current commit author name
		//  by env: CI_COMMIT_AUTHOR
		CiCommitAuthor string `mock_env_key:"CI_COMMIT_AUTHOR" string_line_2_raw:"ci_commit_author"`

		// CiCommitAuthorEmail
		//  Provides the current commit author email
		//  by env: CI_COMMIT_AUTHOR_EMAIL
		CiCommitAuthorEmail string `mock_env_key:"CI_COMMIT_AUTHOR_EMAIL" string_line_2_raw:"ci_commit_author_email"`

		// CiCommitAuthorAvatar
		//  Provides the current commit author avatar
		//  by env: CI_COMMIT_AUTHOR_AVATAR
		CiCommitAuthorAvatar string `mock_env_key:"CI_COMMIT_AUTHOR_AVATAR"`

		// CiCommitPreRelease
		//  Provides the current commit release is a pre-release (empty if event is not release)
		//  support at woodpecker 2.3.0
		CiCommitPreRelease bool `mock_env_key:"CI_COMMIT_PRE_RELEASE"`
	}

	// CurrentPipelineInfo
	//  woodpecker current pipeline info.
	CurrentPipelineInfo struct {
		// CiPipelineFiles
		//  Provides the current pipeline files, like [".woodpecker/.build.yml"]
		//  by env: CI_PIPELINE_FILES
		CiPipelineFiles string `mock_env_key:"CI_PIPELINE_FILES"`

		// CiPipelineNumber
		//  Provides the current pipeline number
		//  by env: CI_PIPELINE_NUMBER
		CiPipelineNumber string `mock_env_key:"CI_PIPELINE_NUMBER"`

		// CiPipelineParent
		//  Provides the current pipeline parent
		//  by env: CI_PIPELINE_PARENT
		CiPipelineParent string `mock_env_key:"CI_PIPELINE_PARENT"`

		// CiPipelineEvent
		//  Provides the current pipeline event
		//  by env: CI_PIPELINE_EVENT
		CiPipelineEvent string `mock_env_key:"CI_PIPELINE_EVENT"`

		// CiPipelineUrl
		//  Provides the current pipeline url
		//  by env: CI_PIPELINE_URL
		CiPipelineUrl string `mock_env_key:"CI_PIPELINE_URL"`

		// CiPipelineForgeUrl
		//  Provides the current pipeline forge url
		//  by env: CI_PIPELINE_FORGE_URL
		CiPipelineForgeUrl string `mock_env_key:"CI_PIPELINE_FORGE_URL"`

		// CiPipelineDeployTarget
		//  Provides the current pipeline deploy target
		//  by env: CI_PIPELINE_DEPLOY_TARGET
		CiPipelineDeployTarget string `mock_env_key:"CI_PIPELINE_DEPLOY_TARGET"`

		// CiPipelineStatus
		//  Provides the current pipeline status
		//  by env: CI_PIPELINE_STATUS
		CiPipelineStatus string `mock_env_key:"CI_PIPELINE_STATUS"`

		// CiPipelineCreated
		//  Provides the current pipeline created
		//  by env: CI_PIPELINE_CREATED
		CiPipelineCreated uint64 `mock_env_key:"CI_PIPELINE_CREATED"`

		// CiPipelineCreatedT
		// Provides the current pipeline created, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_PIPELINE_CREATED
		CiPipelineCreatedT string

		// CiPipelineStarted
		//  Provides the current pipeline started
		//  by env: CI_PIPELINE_STARTED
		CiPipelineStarted uint64 `mock_env_key:"CI_PIPELINE_STARTED"`

		// CiPipelineStartedT
		// Provides the current pipeline started, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_PIPELINE_STARTED
		CiPipelineStartedT string

		// CiPipelineFinished
		//  Provides the current pipeline finished
		//  by env: CI_PIPELINE_FINISHED
		CiPipelineFinished uint64 `mock_env_key:"CI_PIPELINE_FINISHED"`

		// CiPipelineFinishedT
		// Provides the current pipeline finished, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_PIPELINE_FINISHED
		CiPipelineFinishedT string

		// CiPipelineDurationHuman
		//  Provides the total pipeline time in seconds.
		//  This value is calculated by subtracting the pipeline created timestamp from the build finished timestamp.
		CiPipelineDurationHuman string
	}

	// CurrentWorkflowInfo
	//  woodpecker current workflow info.
	CurrentWorkflowInfo struct {
		// CiWorkflowName
		//  Provides the current workflow name
		//  by env: CI_WORKFLOW_NAME
		CiWorkflowName string `mock_env_key:"CI_WORKFLOW_NAME" string_line_2_raw:"ci_workflow_name"`

		// CiWorkflowNumber
		//  Provides the current workflow number, this will start from 1
		//  by env: CI_WORKFLOW_NUMBER
		CiWorkflowNumber string `mock_env_key:"CI_WORKFLOW_NUMBER" mock_env_default:"1"`
	}

	// CurrentStepInfo
	//  woodpecker current step info.
	CurrentStepInfo struct {
		// CiStepName
		//  Provides the current step name
		//  by env: CI_STEP_NAME
		CiStepName string `mock_env_key:"CI_STEP_NAME" string_line_2_raw:"ci_step_name"`

		// CiStepNumber
		//  Provides the current step number
		//  by env: CI_STEP_NUMBER
		CiStepNumber string `mock_env_key:"CI_STEP_NUMBER"`

		// CiStepStatus
		//  Provides the current step status (success, failure)
		//  by env: CI_STEP_STATUS
		CiStepStatus string `mock_env_key:"CI_STEP_STATUS"`

		// CiStepStarted
		//  Provides step started UNIX timestamp
		//  by env: CI_STEP_STARTED
		CiStepStarted uint64 `mock_env_key:"CI_STEP_STARTED"`

		// CiStepStartedT
		// Provides the current step started, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_STEP_STARTED
		CiStepStartedT string

		// CiStepFinished
		//  Provides the step finished UNIX timestamp
		//  by env: CI_STEP_FINISHED
		CiStepFinished uint64 `mock_env_key:"CI_STEP_FINISHED"`

		// CiStepFinishedT
		// Provides the current step finished, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_STEP_FINISHED
		CiStepFinishedT string

		// CiStepDurationHuman
		//  Provides the total step time in seconds.
		//  This value is calculated by subtracting the step created timestamp from the build finished timestamp.
		CiStepDurationHuman string

		// CiStepUrl
		//  Provides the current URL to step in UI
		//  by env: CI_STEP_URL
		CiStepUrl string `mock_env_key:"CI_STEP_URL"`
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
		CiPreviousCommitSha string `mock_env_key:"CI_PREV_COMMIT_SHA"`

		// CiPreviousCommitRef
		//  Provides the previous commit ref
		//  by env: CI_PREV_COMMIT_REF
		CiPreviousCommitRef string `mock_env_key:"CI_PREV_COMMIT_REF"`

		// CiPreviousCommitRefSpec
		//  Provides the previous commit ref spec
		//  by env: CI_PREV_COMMIT_REFSPEC
		CiPreviousCommitRefSpec string `mock_env_key:"CI_PREV_COMMIT_REFSPEC"`

		// CiPreviousCommitBranch
		//  Provides the previous commit branch
		//  by env: CI_PREV_COMMIT_BRANCH
		CiPreviousCommitBranch string `mock_env_key:"CI_PREV_COMMIT_BRANCH"`

		// CiPreviousCommitSourceBranch
		//  Provides the previous commit source branch
		//  by env: CI_PREV_COMMIT_SOURCE_BRANCH
		CiPreviousCommitSourceBranch string `mock_env_key:"CI_PREV_COMMIT_SOURCE_BRANCH"`

		// CiPreviousCommitTargetBranch
		//  Provides the previous commit target branch
		//  by env: CI_PREV_COMMIT_TARGET_BRANCH
		CiPreviousCommitTargetBranch string `mock_env_key:"CI_PREV_COMMIT_TARGET_BRANCH"`

		// CiPreviousCommitUrl
		//  Provides the previous commit url
		//  by env: CI_PREV_COMMIT_URL
		CiPreviousCommitUrl string `mock_env_key:"CI_PREV_COMMIT_URL"`

		// CiPreviousCommitMessage
		//  Provides the previous commit message
		//  by env: CI_PREV_COMMIT_MESSAGE
		CiPreviousCommitMessage string `mock_env_key:"CI_PREV_COMMIT_MESSAGE" string_line_2_raw:"ci_prev_commit_message"`

		// CiPreviousCommitAuthor
		//  Provides the previous commit author username
		//  by env: CI_PREV_COMMIT_AUTHOR
		CiPreviousCommitAuthor string `mock_env_key:"CI_PREV_COMMIT_AUTHOR" string_line_2_raw:"ci_prev_commit_author"`

		// CiPreviousCommitAuthorEmail
		//  Provides the previous commit author email
		//  by env: CI_PREV_COMMIT_AUTHOR_EMAIL
		CiPreviousCommitAuthorEmail string `mock_env_key:"CI_PREV_COMMIT_AUTHOR_EMAIL" string_line_2_raw:"ci_prev_commit_author_email"`

		// CiPreviousCommitAuthorAvatar
		//  Provides the previous commit author avatar
		//  by env: CI_PREV_COMMIT_AUTHOR_AVATAR
		CiPreviousCommitAuthorAvatar string `mock_env_key:"CI_PREV_COMMIT_AUTHOR_AVATAR"`
	}

	// PreviousPipelineInfo
	//  woodpecker previous pipeline info.
	PreviousPipelineInfo struct {
		// CiPreviousPipelineNumber
		//  Provides the previous pipeline number
		//  by env: CI_PREV_PIPELINE_NUMBER
		CiPreviousPipelineNumber string `mock_env_key:"CI_PREV_PIPELINE_NUMBER"`

		// CiPreviousPipelineParent
		//  Provides the previous pipeline number of parent pipeline
		//  by env: CI_PREV_PIPELINE_PARENT
		CiPreviousPipelineParent string `mock_env_key:"CI_PREV_PIPELINE_PARENT"`

		// CiPreviousPipelineEvent
		//  Provides the previous pipeline event (push, pull_request, tag, deployment)
		//  by env: CI_PREV_PIPELINE_EVENT
		CiPreviousPipelineEvent string `mock_env_key:"CI_PREV_PIPELINE_EVENT"`

		// CiPreviousPipelineUrl
		//  Provides the previous pipeline link in CI
		//  by env: CI_PREV_PIPELINE_URL
		CiPreviousPipelineUrl string `mock_env_key:"CI_PREV_PIPELINE_URL"`

		// CiPreviousPipelineForgeUrl
		//  Provides the previous pipeline link in forge
		//  by env: CI_PREV_PIPELINE_FORGE_URL
		CiPreviousPipelineForgeUrl string `mock_env_key:"CI_PREV_PIPELINE_FORGE_URL"`

		// CiPreviousPipelineDeployTarget
		//  Provides the previous pipeline deploy target for deployment events (ie production)
		//  by env: CI_PREV_PIPELINE_DEPLOY_TARGET
		CiPreviousPipelineDeployTarget string `mock_env_key:"CI_PREV_PIPELINE_DEPLOY_TARGET"`

		// CiPreviousPipelineStatus
		//  Provides the previous pipeline status (success, failure)
		//  by env: CI_PREV_PIPELINE_STATUS
		CiPreviousPipelineStatus string `mock_env_key:"CI_PREV_PIPELINE_STATUS"`

		// CiPreviousPipelineCreated
		//  Provides the previous pipeline created
		//  by env: CI_PREV_PIPELINE_CREATED
		CiPreviousPipelineCreated uint64 `mock_env_key:"CI_PREV_PIPELINE_CREATED"`

		// CiPreviousPipelineCreatedT
		// Provides the previous pipeline created, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_PREV_PIPELINE_CREATED
		CiPreviousPipelineCreatedT string

		// CiPreviousPipelineStarted
		//  Provides the previous pipeline started
		//  by env: CI_PREV_PIPELINE_STARTED
		CiPreviousPipelineStarted uint64 `mock_env_key:"CI_PREV_PIPELINE_STARTED"`

		// CiPreviousPipelineStartedT
		// Provides the previous pipeline started, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_PREV_PIPELINE_STARTED
		CiPreviousPipelineStartedT string

		// CiPreviousPipelineFinished
		//  Provides the previous pipeline finished
		//  by env: CI_PREV_PIPELINE_FINISHED
		CiPreviousPipelineFinished uint64 `mock_env_key:"CI_PREV_PIPELINE_FINISHED"`

		// CiPreviousPipelineFinishedT
		// Provides the previous pipeline finished, unix timestamp format by wd_flag.TimeFormatDefault
		//  by env: CI_PREV_PIPELINE_FINISHED
		CiPreviousPipelineFinishedT string

		// CiPreviousPipelineDurationHuman
		//  Provides the total pipeline time in seconds.
		//  This value is calculated by subtracting the pipeline created timestamp from the build finished timestamp.
		CiPreviousPipelineDurationHuman string
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
