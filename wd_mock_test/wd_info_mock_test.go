package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewWoodpeckerInfo(t *testing.T) {
	// mock NewWoodpeckerInfo

	type mockArgs struct {
		workspace string
		status    string
	}

	tests := []struct {
		name    string
		c       wd_info.WoodpeckerInfo
		args    mockArgs
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewWoodpeckerInfo/sample.golden
			args: mockArgs{
				//workspace: "",
				status: wd_info.BuildStatusCreated,
			},
		},
		{
			name: "failure-status",
			args: mockArgs{
				//workspace: "",
				status: wd_info.BuildStatusFailure,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewWoodpeckerInfo
			gotResult := wd_mock.NewWoodpeckerInfo(
				wd_mock.WithCiWorkspace(tc.args.workspace),
				wd_mock.WithCurrentPipelineStatus(tc.args.status),
			)
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				return
			}
			// verify NewWoodpeckerInfo
			assert.Equal(t, tc.args.status, gotResult.CurrentInfo.CurrentPipelineInfo.CiPipelineStatus)
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
