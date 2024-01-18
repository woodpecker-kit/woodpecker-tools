package wd_flag

const (
	// NameCliRepositoryCiRepo
	//  Provides the repository repo flag. This value is `<owner>/<name>` when the is run in woodpecker
	NameCliRepositoryCiRepo = "repository.ci_repo"

	// EnvKeyRepositoryCiRepo
	//  Provides the repository repo flag. This value is `<owner>/<name>` when the is run in woodpecker
	EnvKeyRepositoryCiRepo = "CI_REPO"

	// NameCliRepositoryCiOwner
	//  Provides the repository owner flag. This value is `<owner>` when the is run in woodpecker
	NameCliRepositoryCiOwner = "repository.ci_repo_owner"

	// EnvKeyRepositoryCiOwner
	//  Provides the repository owner flag. This value is `<owner>` when the is run in woodpecker
	EnvKeyRepositoryCiOwner = "CI_REPO_OWNER"

	// NameCliRepositoryCiName
	//  Provides the repository name flag. This value is `<name>` when the is run in woodpecker
	NameCliRepositoryCiName = "repository.ci_repo_name"

	// EnvKeyRepositoryCiName
	//  Provides the repository name flag. This value is `<name>` when the is run in woodpecker
	EnvKeyRepositoryCiName = "CI_REPO_NAME"

	// NameCliRepositoryCiRemoteId
	//  Provides the repository remote ID, is the UID it has in the forge.
	NameCliRepositoryCiRemoteId = "repository.ci_remote_id"

	// EnvKeyRepositoryCiRemoteId
	//  Provides the repository remote ID, is the UID it has in the forge.
	EnvKeyRepositoryCiRemoteId = "CI_REPO_REMOTE_ID"

	// NameCliRepositoryCiScm
	//  Provides the repository scm flag. repository SCM (git).
	NameCliRepositoryCiScm = "repository.ci_repo_scm"

	// EnvKeyRepositoryCiScm
	//  Provides the repository scm flag. repository SCM (git).
	EnvKeyRepositoryCiScm = "CI_REPO_SCM"

	// NameCliRepositoryCiUrl
	//  Provides the repository url flag. repository URL.
	NameCliRepositoryCiUrl = "repository.ci_repo_url"

	// EnvKeyRepositoryCiUrl
	//  Provides the repository url flag. repository URL.
	EnvKeyRepositoryCiUrl = "CI_REPO_URL"

	// NameCliRepositoryCiCloneUrl
	//  Provides the repository clone url flag. repository clone URL.
	NameCliRepositoryCiCloneUrl = "repository.ci_repo_clone_url"

	// EnvKeyRepositoryCiCloneUrl
	//  Provides the repository clone url flag. repository clone URL.
	EnvKeyRepositoryCiCloneUrl = "CI_REPO_CLONE_URL"

	// NameCliRepositoryCiCloneSshUrl
	//  Provides the repository clone ssh url flag. repository clone SSH URL.
	NameCliRepositoryCiCloneSshUrl = "repository.ci_repo_clone_ssh_url"

	// EnvKeyRepositoryCiCloneSshUrl
	//  Provides the repository clone ssh url flag. repository clone SSH URL.
	EnvKeyRepositoryCiCloneSshUrl = "CI_REPO_CLONE_SSH_URL"

	// NameCliRepositoryCiDefaultBranch
	//  Provides the repository default branch flag. repository default branch.
	NameCliRepositoryCiDefaultBranch = "repository.ci_repo_default_branch"

	// EnvKeyRepositoryCiDefaultBranch
	//  Provides the repository default branch flag. repository default branch.
	EnvKeyRepositoryCiDefaultBranch = "CI_REPO_DEFAULT_BRANCH"

	// NameCliRepositoryCiPrivate
	//  Provides the repository private flag. repository is private.
	NameCliRepositoryCiPrivate = "repository.ci_repo_private"

	// EnvKeyRepositoryCiPrivate
	//  Provides the repository private flag. repository is private.
	EnvKeyRepositoryCiPrivate = "CI_REPO_PRIVATE"

	// NameCliRepositoryCiTrusted
	//  Provides the repository trusted flag. repository is trusted.
	NameCliRepositoryCiTrusted = "repository.ci_repo_trusted"

	// EnvKeyRepositoryCiTrusted
	//  Provides the repository trusted flag. repository is trusted.
	EnvKeyRepositoryCiTrusted = "CI_REPO_TRUSTED"
)
