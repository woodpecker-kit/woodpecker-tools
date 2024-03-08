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
		workspace string
		status    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "sample", // testdata/TestParseWoodpeckerInfo2Shot/sample.golden
			args: args{
				//workspace: "",
				status: wd_info.BuildStatusSuccess,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			gotWoodPeckerInfo := wd_mock.NewWoodpeckerInfo(
				wd_mock.WithCiWorkspace(tc.args.workspace),
				wd_mock.WithCurrentPipelineStatus(tc.args.status),
			)
			// do ParseWoodpeckerInfo2Shot
			gotResult := wd_info_shot.ParseWoodpeckerInfo2Shot(*gotWoodPeckerInfo)
			if tc.wantErr != nil {
				return
			}
			// verify ParseWoodpeckerInfo2Shot
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
