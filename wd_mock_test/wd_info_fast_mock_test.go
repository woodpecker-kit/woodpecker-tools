package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestWoodpeckerInfoFastMock(t *testing.T) {
	// mock WoodpeckerInfoFastMock
	type args struct {
		info wd_info.WoodpeckerInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "worksSpace", // testdata/TestWoodpeckerInfoFastMock/worksSpace.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.FastWorkSpace("foo"),
				),
			},
		},
		{
			name: "statusSuccess", // testdata/TestWoodpeckerInfoFastMock/statusSuccess.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.FastCurrentStatus(wd_info.BuildStatusSuccess),
				),
			},
		},
		{
			name: "statusFailure", // testdata/TestWoodpeckerInfoFastMock/statusFailure.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.FastCurrentStatus(wd_info.BuildStatusFailure),
				),
			},
		},
		{
			name: "fastTag", // testdata/TestWoodpeckerInfoFastMock/fastTag.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.FastTag("v1.0.0", "new tag"),
				),
			},
		},
		{
			name: "fastPullRequest", // testdata/TestWoodpeckerInfoFastMock/fastPullRequest.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.FastPullRequest("1", "new pr", "feature-support", "main", "main"),
				),
			},
		},
		{
			name: "fastPullRequestClose", // testdata/TestWoodpeckerInfoFastMock/fastPullRequestClose.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.FastPullRequestClose("1", "new pr", "feature-support", "main", "main"),
				),
			},
		},
		{
			name: "fastPushCommitBranch", // testdata/TestWoodpeckerInfoFastMock/fastPushCommitBranch.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.FastPushCommitBranch("feature", "87366f44f57be21a42468dbb8bbca32cc695324b", "new commit"),
				),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do WoodpeckerInfoFastMock
			if tc.wantErr {
				return
			}
			// verify WoodpeckerInfoFastMock
			g.AssertJson(t, t.Name(), tc.args.info)
		})
	}
}
