package wd_flag

const (
	// NameCliCurrentPipelineFiles
	//  Provides the current pipeline files, like [".woodpecker/.build.yml"]
	NameCliCurrentPipelineFiles = "current_pipeline.files"

	// EnvKeyCurrentPipelineFiles
	//  Provides the current pipeline files, like [".woodpecker/.build.yml"]
	EnvKeyCurrentPipelineFiles = "CI_PIPELINE_FILES"

	// NameCliCurrentPipelineNumber
	//  Provides the current pipeline number
	NameCliCurrentPipelineNumber = "current_pipeline.number"

	// EnvKeyCurrentPipelineNumber
	//  Provides the current pipeline number
	EnvKeyCurrentPipelineNumber = "CI_PIPELINE_NUMBER"

	// NameCliCurrentPipelineParent
	//  Provides number of parent pipeline
	NameCliCurrentPipelineParent = "current_pipeline.parent"

	// EnvKeyCurrentPipelineParent
	//  Provides number of parent pipeline
	EnvKeyCurrentPipelineParent = "CI_PIPELINE_PARENT"

	// NameCliCurrentPipelineEvent
	//  Provides the current pipeline event (push, pull_request, tag, deployment, cron, manual)
	NameCliCurrentPipelineEvent = "current_pipeline.event"

	// EnvKeyCurrentPipelineEvent
	//  Provides the current pipeline event (push, pull_request, tag, deployment, cron, manual)
	EnvKeyCurrentPipelineEvent = "CI_PIPELINE_EVENT"

	// NameCliCurrentPipelineUrl
	//  Provides the current link to the web UI for the pipeline
	NameCliCurrentPipelineUrl = "current_pipeline.url"

	// EnvKeyCurrentPipelineUrl
	//  Provides the current link to the web UI for the pipeline
	EnvKeyCurrentPipelineUrl = "CI_PIPELINE_URL"

	// NameCliCurrentPipelineForgeUrl
	//  Provides the current link to the forge's web UI for the commit(s) or tag that triggered the pipeline
	NameCliCurrentPipelineForgeUrl = "current_pipeline.forge_url"

	// EnvKeyCurrentPipelineForgeUrl
	//  Provides the current link to the forge's web UI for the commit(s) or tag that triggered the pipeline
	EnvKeyCurrentPipelineForgeUrl = "CI_PIPELINE_FORGE_URL"

	// NameCliCurrentPipelineDeployTarget
	//  Provides the current pipeline deploy target for deployment events (ie production)
	NameCliCurrentPipelineDeployTarget = "current_pipeline.deploy_target"

	// EnvKeyCurrentPipelineDeployTarget
	//  Provides the current pipeline deploy target for deployment events (ie production)
	EnvKeyCurrentPipelineDeployTarget = "CI_PIPELINE_DEPLOY_TARGET"

	// NameCliCurrentPipelineStatus
	//  Provides the current pipeline status (success, failure)
	NameCliCurrentPipelineStatus = "current_pipeline.status"

	// EnvKeyCurrentPipelineStatus
	//  Provides the current pipeline status (success, failure)
	EnvKeyCurrentPipelineStatus = "CI_PIPELINE_STATUS"

	// NameCliCurrentPipelineCreated
	//  Provides the current pipeline created UNIX timestamp
	NameCliCurrentPipelineCreated = "current_pipeline.created"

	// EnvKeyCurrentPipelineCreated
	//  Provides the current pipeline created UNIX timestamp
	EnvKeyCurrentPipelineCreated = "CI_PIPELINE_CREATED"

	// NameCliCurrentPipelineStarted
	//  Provides the current pipeline started UNIX timestamp
	NameCliCurrentPipelineStarted = "current_pipeline.started"

	// EnvKeyCurrentPipelineStarted
	//  Provides the current pipeline started UNIX timestamp
	EnvKeyCurrentPipelineStarted = "CI_PIPELINE_STARTED"

	// NameCliCurrentPipelineFinished
	//  Provides the current pipeline finished UNIX timestamp
	NameCliCurrentPipelineFinished = "current_pipeline.finished"

	// EnvKeyCurrentPipelineFinished
	//  Provides the current pipeline finished UNIX timestamp
	EnvKeyCurrentPipelineFinished = "CI_PIPELINE_FINISHED"
)
