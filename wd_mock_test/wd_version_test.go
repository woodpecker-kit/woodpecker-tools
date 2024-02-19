package wd_mock

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

const (
	mockVersionGreater = "2.3.1"
)

func mockWoodpeckerInfoByVersion(version string) wd_info.WoodpeckerInfo {
	info := wd_mock.NewWoodpeckerInfo(
		wd_mock.WithCurrentPipelineStatus(wd_info.BuildStatusCreated),
		wd_mock.WithCiForgeInfo(),
		wd_mock.WithRepositoryInfo(),
		wd_mock.WithCurrentCommitInfo(),
		wd_mock.WithCurrentPipelineInfo(),
		wd_mock.WithCurrentStepInfo(),
		wd_mock.WithPreviousCommitInfo(),
		wd_mock.WithPreviousPipelineInfo(),
		wd_mock.WithCiSystemInfo(
			wd_mock.WithCiSystemVersion(version),
		),
	)
	return *info
}

func TestCiSystemVersionCheck(t *testing.T) {
	// mock CiSystemVersionCheck
	type args struct {
		info wd_info.WoodpeckerInfo
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "empty",
			args: args{
				info: mockWoodpeckerInfoByVersion(""),
			},
			wantErr: fmt.Errorf("woodpeacker ci system version is empty, please set env CI_SYSTEM_VERSION"),
		},
		{
			name: "not support",
			args: args{
				info: mockWoodpeckerInfoByVersion("Semantic Versioning"),
			},
			wantErr: fmt.Errorf("can not parse target version: Semantic Versioning err: Invalid Semantic Version"),
		},
		{
			name: "less",
			args: args{
				info: mockWoodpeckerInfoByVersion("1.0.0"),
			},
			wantErr: fmt.Errorf("woodpeacker ci system version: 1.0.0 not support, err: [1.0.0 is less than %s]", wd_info.WoodpeckerInfoSupportVersion),
		},
		{
			name: "equal",
			args: args{
				info: mockWoodpeckerInfoByVersion(wd_info.WoodpeckerInfoSupportVersion),
			},
		},
		{
			name: "greater",
			args: args{
				info: mockWoodpeckerInfoByVersion(mockVersionGreater),
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do CiSystemVersionCheck
			gotErr := wd_info.CiSystemVersionCheck(tc.args.info)

			// verify CiSystemVersionCheck
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
		})
	}
}
