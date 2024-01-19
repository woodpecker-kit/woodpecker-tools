package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultPreviousCommitInfo = setDefaultPreviousCommitInfo()

type PreviousCommitInfoOption func(*wd_info.PreviousCommitInfo)

func setDefaultPreviousCommitInfo() *wd_info.PreviousCommitInfo {
	return &wd_info.PreviousCommitInfo{
		CiPreviousCommitSha:          "b2542f37e23645b2032fcf9f180e6940ef6915ac",
		CiPreviousCommitRef:          "refs/heads/main",
		CiPreviousCommitRefSpec:      "",
		CiPreviousCommitBranch:       "main",
		CiPreviousCommitSourceBranch: "",
		CiPreviousCommitTargetBranch: "",
		CiPreviousCommitUrl:          "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/b2542f37e23645b2032fcf9f180e6940ef6915ac",
		CiPreviousCommitMessage:      "test: try send env between steps",
		CiPreviousCommitAuthor:       "sinlov",
		CiPreviousCommitAuthorEmail:  "sinlov@noreply.localhost",
		CiPreviousCommitAuthorAvatar: "",
	}
}

func NewPreviousCommitInfo(opts ...PreviousCommitInfoOption) (opt *wd_info.PreviousCommitInfo) {
	opt = defaultPreviousCommitInfo
	for _, o := range opts {
		o(opt)
	}
	defaultPreviousCommitInfo = setDefaultPreviousCommitInfo()
	return
}

func WithCiPreviousCommitSha(commitSha string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitSha = commitSha
	}
}

func WithCiPreviousCommitRef(commitRef string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitRef = commitRef
	}
}

func WithCiPreviousCommitRefSpec(commitRefSpec string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitRefSpec = commitRefSpec
	}
}

func WithCiPreviousCommitBranch(commitBranch string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitBranch = commitBranch
	}
}

func WithCiPreviousCommitSourceBranch(commitSourceBranch string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitSourceBranch = commitSourceBranch
	}
}

func WithCiPreviousCommitTargetBranch(commitTargetBranch string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitTargetBranch = commitTargetBranch
	}
}

func WithCiPreviousCommitUrl(commitUrl string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitUrl = commitUrl
	}
}

func WithCiPreviousCommitMessage(commitMessage string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitMessage = commitMessage
	}
}

func WithCiPreviousCommitAuthor(commitAuthor string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitAuthor = commitAuthor
	}
}

func WithCiPreviousCommitAuthorEmail(commitAuthorEmail string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitAuthorEmail = commitAuthorEmail
	}
}

func WithCiPreviousCommitAuthorAvatar(commitAuthorAvatar string) PreviousCommitInfoOption {
	return func(o *wd_info.PreviousCommitInfo) {
		o.CiPreviousCommitAuthorAvatar = commitAuthorAvatar
	}
}

func NewPreviousCommitInfoFull() *wd_info.PreviousCommitInfo {
	return setDefaultPreviousCommitInfo()
}
