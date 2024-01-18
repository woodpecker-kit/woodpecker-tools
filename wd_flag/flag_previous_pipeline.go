package wd_flag

const (
	// NameCliPreviousPipelineNumber
	//  Provides the previous pipeline number
	NameCliPreviousPipelineNumber = "previous_pipeline.number"

	// EnvKeyPreviousCiPipelineNumber
	//  Provides the previous pipeline number
	EnvKeyPreviousCiPipelineNumber = "CI_PREV_PIPELINE_NUMBER"

	// NameCliPreviousPipelineParent
	//  Provides the previous pipeline number of parent pipeline
	NameCliPreviousPipelineParent = "previous_pipeline.parent"

	// EnvKeyPreviousCiPipelineParent
	//  Provides the previous pipeline number of parent pipeline
	EnvKeyPreviousCiPipelineParent = "CI_PREV_PIPELINE_PARENT"

	// NameCliPreviousPipelineEvent
	//  Provides the previous pipeline event (push, pull_request, tag, deployment)
	NameCliPreviousPipelineEvent = "previous_pipeline.event"

	// EnvKeyPreviousCiPipelineEvent
	//  Provides the previous pipeline event (push, pull_request, tag, deployment)
	EnvKeyPreviousCiPipelineEvent = "CI_PREV_PIPELINE_EVENT"

	// NameCliPreviousPipelineUrl
	//  Provides the previous pipeline link in CI
	NameCliPreviousPipelineUrl = "previous_pipeline.url"

	// EnvKeyPreviousCiPipelineUrl
	//  Provides the previous pipeline link in CI
	EnvKeyPreviousCiPipelineUrl = "CI_PREV_PIPELINE_URL"

	// NameCliPreviousPipelineForgeUrl
	//  Provides the previous pipeline link in forge
	NameCliPreviousPipelineForgeUrl = "previous_pipeline.forge_url"

	// EnvKeyPreviousCiPipelineForgeUrl
	//  Provides the previous pipeline link in forge
	EnvKeyPreviousCiPipelineForgeUrl = "CI_PREV_PIPELINE_FORGE_URL"

	// NameCliPreviousPipelineDeployTarget
	//  Provides the previous pipeline deploy target for deployment events (ie production)
	NameCliPreviousPipelineDeployTarget = "previous_pipeline.deploy_target"

	// EnvKeyPreviousCiPipelineDeployTarget
	//  Provides the previous pipeline deploy target for deployment events (ie production)
	EnvKeyPreviousCiPipelineDeployTarget = "CI_PREV_PIPELINE_DEPLOY_TARGET"

	// NameCliPreviousPipelineStatus
	//  Provides the previous pipeline status (success, failure)
	NameCliPreviousPipelineStatus = "previous_pipeline.status"

	// EnvKeyPreviousCiPipelineStatus
	//  Provides the previous pipeline status (success, failure)
	EnvKeyPreviousCiPipelineStatus = "CI_PREV_PIPELINE_STATUS"

	// NameCliPreviousPipelineCreated
	//  Provides the previous pipeline created UNIX timestamp
	NameCliPreviousPipelineCreated = "previous_pipeline.created"

	// EnvKeyPreviousCiPipelineCreated
	//  Provides the previous pipeline created UNIX timestamp
	EnvKeyPreviousCiPipelineCreated = "CI_PREV_PIPELINE_CREATED"

	// NameCliPreviousPipelineStarted
	//  Provides the previous pipeline started UNIX timestamp
	NameCliPreviousPipelineStarted = "previous_pipeline.started"

	// EnvKeyPreviousCiPipelineStarted
	//  Provides the previous pipeline started UNIX timestamp
	EnvKeyPreviousCiPipelineStarted = "CI_PREV_PIPELINE_STARTED"

	// NameCliPreviousPipelineFinished
	//  Provides the previous pipeline finished UNIX timestamp
	NameCliPreviousPipelineFinished = "previous_pipeline.finished"

	// EnvKeyPreviousCiPipelineFinished
	//  Provides the previous pipeline finished UNIX timestamp
	EnvKeyPreviousCiPipelineFinished = "CI_PREV_PIPELINE_FINISHED"
)
