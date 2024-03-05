package wd_info

const (
	// from: https://woodpecker-ci.org/docs/usage/terminiology#pipeline-events

	// EventPipelinePush
	//  A push event is triggered when a commit is pushed to a branch.
	EventPipelinePush = "push"

	// EventPipelinePullRequest
	//  A pull request event is triggered when a pull request is opened or a new commit is pushed to it.
	EventPipelinePullRequest = "pull_request"

	// EventPipelinePullRequestClose
	//  A pull request closed event is triggered when a pull request is closed or merged.
	EventPipelinePullRequestClose = "pull_request_closed"

	// EventPipelineTag
	//  A tag event is triggered when a tag is pushed.
	EventPipelineTag = "tag"

	// EventPipelineManual
	//  A manual event is triggered when a user manually triggers a pipeline.
	EventPipelineManual = "manual"

	// EventPipelineCron
	//  A cron event is triggered when a cron job is executed.
	EventPipelineCron = "cron"

	// EventPipelineRelease
	//  A release event is triggered when a release, pre-release or draft is created. (You can apply further filters using evaluate with environment variables.)
	EventPipelineRelease = "release"
)
