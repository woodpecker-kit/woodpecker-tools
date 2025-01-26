package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"testing"
)

func TestParseWoodpeckerInfo2Shot(t *testing.T) {
	// mock ParseWoodpeckerInfo2Short
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
			//tc.args.info.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus = ""

			// do ParseWoodpeckerInfo2Short
			gotResult := wd_short_info.ParseWoodpeckerInfo2Short(tc.args.info)
			if tc.wantErr != nil {
				return
			}
			errCheck := wd_info.CheckBuildStatusSupport(tc.args.info.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus)
			if errCheck != nil {
				t.Errorf("check build status error: %v", errCheck)
			}
			errParseCheck := wd_info.CheckBuildStatusSupport(gotResult.Build.Status)
			if errParseCheck != nil {
				t.Fatalf("check build status error: %v", errParseCheck)
			}

			// verify ParseWoodpeckerInfo2Short
			g.AssertJson(t, t.Name(), gotResult)

		})
	}
}
