package wd_flag

const (
	// NameCliCurrentCommitCiCommitSha
	//  Provides the current commit SHA
	NameCliCurrentCommitCiCommitSha = "current_commit.ci_commit_sha"

	// EnvKeyCurrentCommitCiCommitSha
	//  Provides the current commit SHA
	EnvKeyCurrentCommitCiCommitSha = "CI_COMMIT_SHA"

	// NameCliCurrentCommitCiCommitRef
	//  Provides the current commit ref
	NameCliCurrentCommitCiCommitRef = "current_commit.ci_commit_ref"

	// EnvKeyCurrentCommitCiCommitRef
	//  Provides the current commit ref
	EnvKeyCurrentCommitCiCommitRef = "CI_COMMIT_REF"

	// NameCliCurrentCommitCiCommitRefSpec
	//  Provides the current commit ref spec
	NameCliCurrentCommitCiCommitRefSpec = "current_commit.ci_commit_refspec"

	// EnvKeyCurrentCommitCiCommitRefSpec
	//  Provides the current commit ref spec
	EnvKeyCurrentCommitCiCommitRefSpec = "CI_COMMIT_REFSPEC"

	// NameCliCurrentCommitCiCommitBranch
	//  Provides the current commit branch (equals target branch for pull requests)
	NameCliCurrentCommitCiCommitBranch = "current_commit.ci_commit_branch"

	// EnvKeyCurrentCommitCiCommitBranch
	//  Provides the current commit branch (equals target branch for pull requests)
	EnvKeyCurrentCommitCiCommitBranch = "CI_COMMIT_BRANCH"

	// NameCliCurrentCommitCiCommitSourceBranch
	//  Provides the current commit source branch (empty if event is not pull_request)
	NameCliCurrentCommitCiCommitSourceBranch = "current_commit.ci_commit_source_branch"

	// EnvKeyCurrentCommitCiCommitSourceBranch
	//  Provides the current commit source branch (empty if event is not pull_request)
	EnvKeyCurrentCommitCiCommitSourceBranch = "CI_COMMIT_SOURCE_BRANCH"

	// NameCliCurrentCommitCiCommitTargetBranch
	//  Provides the current commit target branch (empty if event is not pull_request)
	NameCliCurrentCommitCiCommitTargetBranch = "current_commit.ci_commit_target_branch"

	// EnvKeyCurrentCommitCiCommitTargetBranch
	//  Provides the current commit target branch (empty if event is not pull_request)
	EnvKeyCurrentCommitCiCommitTargetBranch = "CI_COMMIT_TARGET_BRANCH"

	// NameCliCurrentCommitCiCommitTag
	//  Provides the current commit tag (empty if event is not tag)
	NameCliCurrentCommitCiCommitTag = "current_commit.ci_commit_tag"

	// EnvKeyCurrentCommitCiCommitTag
	//  Provides the current commit tag (empty if event is not tag)
	EnvKeyCurrentCommitCiCommitTag = "CI_COMMIT_TAG"

	// NameCliCurrentCommitCiCommitPullRequest
	//  Provides the current commit pull request (empty if event is not pull_request)
	NameCliCurrentCommitCiCommitPullRequest = "current_commit.ci_commit_pull_request"

	// EnvKeyCurrentCommitCiCommitPullRequest
	//  Provides the current commit pull request (empty if event is not pull_request)
	EnvKeyCurrentCommitCiCommitPullRequest = "CI_COMMIT_PULL_REQUEST"

	// NameCliCurrentCommitCiCommitPullRequestLabels
	//  Provides the current labels assigned to pull request (empty if event is not pull_request)
	NameCliCurrentCommitCiCommitPullRequestLabels = "current_commit.ci_commit_pull_request_labels"

	// EnvKeyCurrentCommitCiCommitPullRequestLabels
	//  Provides the current labels assigned to pull request (empty if event is not pull_request)
	EnvKeyCurrentCommitCiCommitPullRequestLabels = "CI_COMMIT_PULL_REQUEST_LABELS"

	// NameCliCurrentCommitCiCommitMessage
	//  Provides the current commit message
	NameCliCurrentCommitCiCommitMessage = "current_commit.ci_commit_message"

	// EnvKeyCurrentCommitCiCommitMessage
	//  Provides the current commit message
	EnvKeyCurrentCommitCiCommitMessage = "CI_COMMIT_MESSAGE"

	// NameCliCurrentCommitCiCommitAuthor
	//  Provides the current commit author
	NameCliCurrentCommitCiCommitAuthor = "current_commit.ci_commit_author"

	// EnvKeyCurrentCommitCiCommitAuthor
	//  Provides the current commit author
	EnvKeyCurrentCommitCiCommitAuthor = "CI_COMMIT_AUTHOR"

	// NameCliCurrentCommitCiCommitAuthorEmail
	//  Provides the current commit author email
	NameCliCurrentCommitCiCommitAuthorEmail = "current_commit.ci_commit_author_email"

	// EnvKeyCurrentCommitCiCommitAuthorEmail
	//  Provides the current commit author email
	EnvKeyCurrentCommitCiCommitAuthorEmail = "CI_COMMIT_AUTHOR_EMAIL"

	// NameCliCurrentCommitCiCommitAuthorAvatar
	//  Provides the current commit author avatar
	NameCliCurrentCommitCiCommitAuthorAvatar = "current_commit.ci_commit_author_avatar"

	// EnvKeyCurrentCommitCiCommitAuthorAvatar
	//  Provides the current commit author avatar
	EnvKeyCurrentCommitCiCommitAuthorAvatar = "CI_COMMIT_AUTHOR_AVATAR"
)
