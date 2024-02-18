package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewPreviousCommitInfo(t *testing.T) {
	// mock NewPreviousCommitInfo
	type args struct {
		ciPreviousCommitSha          string
		ciPreviousCommitRef          string
		ciPreviousCommitRefSpec      string
		ciPreviousCommitBranch       string
		ciPreviousCommitSourceBranch string
		ciPreviousCommitTargetBranch string
		ciPreviousCommitUrl          string
		ciPreviousCommitMessage      string
		ciPreviousCommitAuthor       string
		ciPreviousCommitAuthorEmail  string
		ciPreviousCommitAuthorAvatar string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewPreviousCommitInfo/sample.golden
			args: args{
				ciPreviousCommitSha:          "b2542f37e23645b2032fcf9f180e6940ef6915ac",
				ciPreviousCommitRef:          "refs/heads/main",
				ciPreviousCommitRefSpec:      "",
				ciPreviousCommitBranch:       "main",
				ciPreviousCommitSourceBranch: "",
				ciPreviousCommitTargetBranch: "",
				ciPreviousCommitUrl:          "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/b2542f37e23645b2032fcf9f180e6940ef6915ac",
				ciPreviousCommitMessage:      "test: try send env between steps",
				ciPreviousCommitAuthor:       "sinlov",
				ciPreviousCommitAuthorEmail:  "sinlov@mail.localhost",
				ciPreviousCommitAuthorAvatar: "",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewPreviousCommitInfo
			gotResult := wd_mock.NewPreviousCommitInfo(
				wd_mock.WithCiPreviousCommitSha(tc.args.ciPreviousCommitSha),
				wd_mock.WithCiPreviousCommitRef(tc.args.ciPreviousCommitRef),
				wd_mock.WithCiPreviousCommitRefSpec(tc.args.ciPreviousCommitRefSpec),
				wd_mock.WithCiPreviousCommitBranch(tc.args.ciPreviousCommitBranch),
				wd_mock.WithCiPreviousCommitSourceBranch(tc.args.ciPreviousCommitSourceBranch),
				wd_mock.WithCiPreviousCommitTargetBranch(tc.args.ciPreviousCommitTargetBranch),
				wd_mock.WithCiPreviousCommitUrl(tc.args.ciPreviousCommitUrl),
				wd_mock.WithCiPreviousCommitMessage(tc.args.ciPreviousCommitMessage),
				wd_mock.WithCiPreviousCommitAuthor(tc.args.ciPreviousCommitAuthor),
				wd_mock.WithCiPreviousCommitAuthorEmail(tc.args.ciPreviousCommitAuthorEmail),
				wd_mock.WithCiPreviousCommitAuthorAvatar(tc.args.ciPreviousCommitAuthorAvatar),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewPreviousCommitInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
