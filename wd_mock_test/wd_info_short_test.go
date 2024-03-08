package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info_shot"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestParseWoodpeckerInfo2Shot(t *testing.T) {
	// mock ParseWoodpeckerInfo2Shot
	type args struct {
		info wd_info.WoodpeckerInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "sample", // testdata/TestParseWoodpeckerInfo2Shot/sample.golden
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusSuccess),
				),
			},
		},
		{
			name: "statusFailure",
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusFailure),
				),
			},
		},
		{
			name: "fastTag",
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.WithFastMockTag("v1.0.0", "new tag"),
				),
			},
		},
		{
			name: "fastPullRequest",
			args: args{
				info: *wd_mock.NewWoodpeckerInfo(
					wd_mock.WithFastMockPullRequest("1", "new pr", "feature-support", "main", "main"),
				),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do ParseWoodpeckerInfo2Shot
			gotResult := wd_info_shot.ParseWoodpeckerInfo2Shot(tc.args.info)
			if tc.wantErr != nil {
				return
			}
			// verify ParseWoodpeckerInfo2Shot
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
