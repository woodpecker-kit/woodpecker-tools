package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewRepositoryInfo(t *testing.T) {
	// mock NewRepositoryInfo
	type args struct {
		ciRepo              string
		ciRepoOwner         string
		ciRepoName          string
		ciRepoRemoteID      string
		ciRepoScm           string
		ciRepoURL           string
		ciRepoCloneURL      string
		ciRepoCloneSshUrl   string
		ciRepoDefaultBranch string
		ciRepoPrivate       bool
		ciRepoTrusted       bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewRepositoryInfo/sample.golden
			args: args{
				ciRepo:              "github.com",
				ciRepoOwner:         "woodpecker-kit",
				ciRepoName:          "guidance-woodpecker-agent",
				ciRepoRemoteID:      "151",
				ciRepoScm:           "git",
				ciRepoURL:           "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent",
				ciRepoCloneURL:      "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent.git",
				ciRepoCloneSshUrl:   "git@gitea.domain.com:woodpecker-kit/guidance-woodpecker-agent.git",
				ciRepoDefaultBranch: "main",
				ciRepoPrivate:       true,
				ciRepoTrusted:       false,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewRepositoryInfo
			gotResult := wd_mock.NewRepositoryInfo(
				wd_mock.WithCIRepo(tc.args.ciRepo),
				wd_mock.WithCIRepoOwner(tc.args.ciRepoOwner),
				wd_mock.WithCIRepoName(tc.args.ciRepoName),
				wd_mock.WithCIRepoRemoteID(tc.args.ciRepoRemoteID),
				wd_mock.WithCIRepoScm(tc.args.ciRepoScm),
				wd_mock.WithCIRepoURL(tc.args.ciRepoURL),
				wd_mock.WithCIRepoCloneURL(tc.args.ciRepoCloneURL),
				wd_mock.WithCIRepoCloneSshUrl(tc.args.ciRepoCloneSshUrl),
				wd_mock.WithCIRepoDefaultBranch(tc.args.ciRepoDefaultBranch),
				wd_mock.WithCIRepoPrivate(tc.args.ciRepoPrivate),
				wd_mock.WithCIRepoTrusted(tc.args.ciRepoTrusted),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewRepositoryInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
