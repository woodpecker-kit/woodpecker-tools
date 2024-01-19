package wd_mock

import "github.com/woodpecker-kit/woodpecker-tools/wd_info"

var defaultRepositoryInfo = setDefaultRepositoryInfo()

type RepositoryInfoOption func(*wd_info.RepositoryInfo)

func setDefaultRepositoryInfo() *wd_info.RepositoryInfo {
	return &wd_info.RepositoryInfo{
		CIRepo:              "woodpecker-kit/guidance-woodpecker-agent",
		CIRepoOwner:         "woodpecker-kit",
		CIRepoName:          "guidance-woodpecker-agent",
		CIRepoRemoteID:      "151",
		CIRepoScm:           "git",
		CIRepoURL:           "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent",
		CIRepoCloneURL:      "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent.git",
		CIRepoCloneSshUrl:   "git@gitea.domain.com:woodpecker-kit/guidance-woodpecker-agent.git",
		CIRepoDefaultBranch: "main",
		CIRepoPrivate:       true,
		CIRepoTrusted:       false,
	}
}

func NewRepositoryInfo(opts ...RepositoryInfoOption) (opt *wd_info.RepositoryInfo) {
	opt = defaultRepositoryInfo
	for _, o := range opts {
		o(opt)
	}
	defaultRepositoryInfo = setDefaultRepositoryInfo()
	return
}

func WithCIRepo(ciRepo string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepo = ciRepo
	}
}

func WithCIRepoOwner(ciRepoOwner string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoOwner = ciRepoOwner
	}
}

func WithCIRepoName(ciRepoName string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoName = ciRepoName
	}
}

func WithCIRepoRemoteID(ciRepoRemoteID string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoRemoteID = ciRepoRemoteID
	}
}

func WithCIRepoScm(ciRepoScm string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoScm = ciRepoScm
	}
}

func WithCIRepoURL(ciRepoURL string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoURL = ciRepoURL
	}
}

func WithCIRepoCloneURL(ciRepoCloneURL string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoCloneURL = ciRepoCloneURL
	}
}

func WithCIRepoCloneSshUrl(ciRepoCloneSshUrl string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoCloneSshUrl = ciRepoCloneSshUrl
	}
}

func WithCIRepoDefaultBranch(ciRepoDefaultBranch string) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoDefaultBranch = ciRepoDefaultBranch
	}
}

func WithCIRepoPrivate(ciRepoPrivate bool) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoPrivate = ciRepoPrivate
	}
}

func WithCIRepoTrusted(ciRepoTrusted bool) RepositoryInfoOption {
	return func(o *wd_info.RepositoryInfo) {
		o.CIRepoTrusted = ciRepoTrusted
	}
}

func NewRepositoryInfoFull() *wd_info.RepositoryInfo {
	return setDefaultRepositoryInfo()
}
