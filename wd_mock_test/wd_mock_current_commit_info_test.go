package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewCurrentCommitInfo(t *testing.T) {
	// mock NewCurrentCommitInfo
	type args struct {
		ciCommitSha               string
		ciCommitRef               string
		ciCommitRefSpec           string
		ciCommitBranch            string
		ciCommitSourceBranch      string
		ciCommitTargetBranch      string
		ciCommitTag               string
		ciCommitPullRequest       string
		ciCommitPullRequestLabels string
		ciCommitMessage           string
		ciCommitAuthor            string
		ciCommitAuthorEmail       string
		ciCommitAuthorAvatar      string
		ciCommitPreRelease        bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewCurrentCommitInfo/sample.golden
			args: args{
				ciCommitSha:               "9c764dd487bce596c5c0402478fabde5f0344983",
				ciCommitRef:               "refs/heads/main",
				ciCommitRefSpec:           "",
				ciCommitBranch:            "main",
				ciCommitSourceBranch:      "",
				ciCommitTargetBranch:      "",
				ciCommitTag:               "",
				ciCommitPullRequest:       "",
				ciCommitPullRequestLabels: "",
				ciCommitMessage:           "test: use source to load env vars at before step",
				ciCommitAuthor:            "sinlov",
				ciCommitAuthorEmail:       "sinlov@noreply.localhost",
				ciCommitAuthorAvatar:      "foo",
				ciCommitPreRelease:        false,
			},
		},
		{
			name: "dev", // testdata/TestNewCurrentCommitInfo/dev.golden
			args: args{
				ciCommitSha:               "9c764dd487bce596c5c0402478fabde5f0344984",
				ciCommitRef:               "refs/heads/dev",
				ciCommitRefSpec:           "",
				ciCommitBranch:            "dev",
				ciCommitSourceBranch:      "",
				ciCommitTargetBranch:      "",
				ciCommitTag:               "",
				ciCommitPullRequest:       "",
				ciCommitPullRequestLabels: "",
				ciCommitMessage:           "test: use source to load env vars at before step",
				ciCommitAuthor:            "sinlov",
				ciCommitAuthorEmail:       "",
				ciCommitAuthorAvatar:      "foo",
				ciCommitPreRelease:        true,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewCurrentCommitInfo
			gotResult := wd_mock.NewCurrentCommitInfo(
				wd_mock.WithCiCommitSha(tc.args.ciCommitSha),
				wd_mock.WithCiCommitRef(tc.args.ciCommitRef),
				wd_mock.WithCiCommitRefSpec(tc.args.ciCommitRefSpec),
				wd_mock.WithCiCommitBranch(tc.args.ciCommitBranch),
				wd_mock.WithCiCommitSourceBranch(tc.args.ciCommitSourceBranch),
				wd_mock.WithCiCommitTargetBranch(tc.args.ciCommitTargetBranch),
				wd_mock.WithCiCommitTag(tc.args.ciCommitTag),
				wd_mock.WithCiCommitPullRequest(tc.args.ciCommitPullRequest),
				wd_mock.WithCiCommitPullRequestLabels(tc.args.ciCommitPullRequestLabels),
				wd_mock.WithCiCommitMessage(tc.args.ciCommitMessage),
				wd_mock.WithCiCommitAuthor(tc.args.ciCommitAuthor),
				wd_mock.WithCiCommitAuthorEmail(tc.args.ciCommitAuthorEmail),
				wd_mock.WithCiCommitAuthorAvatar(tc.args.ciCommitAuthorAvatar),
				wd_mock.WithCiCommitPreRelease(tc.args.ciCommitPreRelease),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewCurrentCommitInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
