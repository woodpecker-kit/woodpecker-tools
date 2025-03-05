package wd_mock

import (
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
)

var defaultCurrentCommitInfo = setDefaultCurrentCommitInfo()

type CurrentCommitInfoOption func(*wd_info.CurrentCommitInfo)

func setDefaultCurrentCommitInfo() *wd_info.CurrentCommitInfo {
	return &wd_info.CurrentCommitInfo{
		CiCommitSha:               "9c764dd487bce596c5c0402478fabde5f0344983",
		CiCommitRef:               "refs/heads/main",
		CiCommitRefSpec:           "",
		CiCommitBranch:            "main",
		CiCommitSourceBranch:      "",
		CiCommitTargetBranch:      "",
		CiCommitTag:               "",
		CiCommitPullRequest:       "",
		CiCommitPullRequestLabels: "",
		CiCommitMessage:           "test: use source to load env vars at before step",
		CiCommitAuthor:            "sinlov",
		CiCommitAuthorEmail:       "sinlov@noreply.localhost",
		CiCommitAuthorAvatar:      "",
		CiCommitPreRelease:        false,
	}
}

func NewCurrentCommitInfo(opts ...CurrentCommitInfoOption) (opt *wd_info.CurrentCommitInfo) {
	opt = defaultCurrentCommitInfo
	for _, o := range opts {
		o(opt)
	}
	defaultCurrentCommitInfo = setDefaultCurrentCommitInfo()
	return
}

func WithCiCommitSha(commitSha string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitSha = commitSha
	}
}

func WithCiCommitRef(commitRef string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitRef = commitRef
	}
}

func WithCiCommitRefSpec(commitRefSpec string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitRefSpec = commitRefSpec
	}
}

// Deprecated: remove at woodpecker server 3.0.0, instead use WithCiPipelineForgeUrl
func WithCiCommitUrl(commitUrl string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		//o.CiCommitUrl = commitUrl
	}
}

func WithCiCommitBranch(commitBranch string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitBranch = commitBranch
	}
}

func WithCiCommitSourceBranch(commitSourceBranch string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitSourceBranch = commitSourceBranch
	}
}

func WithCiCommitTargetBranch(commitTargetBranch string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitTargetBranch = commitTargetBranch
	}
}

func WithCiCommitTag(commitTag string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitTag = commitTag
	}
}

func WithCiCommitPullRequest(commitPullRequest string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitPullRequest = commitPullRequest
	}
}

func WithCiCommitPullRequestLabels(commitPullRequestLabels string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitPullRequestLabels = commitPullRequestLabels
	}
}

func WithCiCommitMessage(commitMessage string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitMessage = commitMessage
	}
}

func WithCiCommitAuthor(commitAuthor string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitAuthor = commitAuthor
	}
}

func WithCiCommitAuthorEmail(commitAuthorEmail string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitAuthorEmail = commitAuthorEmail
	}
}

func WithCiCommitAuthorAvatar(commitAuthorAvatar string) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitAuthorAvatar = commitAuthorAvatar
	}
}

func WithCiCommitPreRelease(commitPreRelease bool) CurrentCommitInfoOption {
	return func(o *wd_info.CurrentCommitInfo) {
		o.CiCommitPreRelease = commitPreRelease
	}
}

func NewCurrentCommitInfoFull() *wd_info.CurrentCommitInfo {
	return setDefaultCurrentCommitInfo()
}
