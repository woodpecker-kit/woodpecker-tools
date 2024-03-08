package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/env_mock"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestFastMock(t *testing.T) {
	// mock FastMock
	type args struct {
		info wd_info.WoodpeckerInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "statusSuccess", // testdata/TestFastMock/sample.golden
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

			// do FastMock
			if tc.wantErr != nil {
				return
			}
			// verify FastMock
			g.AssertJson(t, t.Name(), tc.args.info)
		})
	}
}

func TestWoodPeckerEnvMock(t *testing.T) {
	// mock WoodPeckerEnvMock

	type mockArgs struct {
		workspace        string
		status           string
		ciWorkflowNumber string
	}

	tests := []struct {
		name    string
		args    mockArgs
		wantErr bool
	}{
		{
			name: "sample",
			args: mockArgs{
				//workspace: "",
				status:           wd_info.BuildStatusCreated,
				ciWorkflowNumber: "1",
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
				wd_mock.WithCurrentWorkflowInfo(
					wd_mock.WithCiWorkflowNumber(tc.args.ciWorkflowNumber),
				),
			)
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				return
			}

			env_mock.MockEnvByStruct(*gotResult)

			t.Logf("~> TestWoodPeckerEnvMock at env: \n%s", env_kit.FindEnv4PrintAsPrefixSortJust("CI_", 36))

			// verify NewWoodpeckerInfo
			assert.Equal(t, tc.args.ciWorkflowNumber, gotResult.CurrentInfo.CurrentWorkflowInfo.CiWorkflowNumber)
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}

func TestNewWoodpeckerInfo(t *testing.T) {
	// mock NewWoodpeckerInfo

	type mockArgs struct {
		workspace string
		status    string
	}

	tests := []struct {
		name    string
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
