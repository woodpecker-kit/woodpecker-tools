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

func TestCiSystemVersionMinimumSupport(t *testing.T) {
	// mock CiSystemVersionMinimumSupport
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
			wantErr: fmt.Errorf("semver version: 1.0.0 not support, err: [1.0.0 is less than %s]", wd_info.WoodpeckerInfoSupportVersion),
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

			// do CiSystemVersionMinimumSupport
			gotErr := wd_info.CiSystemVersionMinimumSupport(tc.args.info)

			// verify CiSystemVersionMinimumSupport
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
		})
	}
}

func TestCiSystemVersionConstraint(t *testing.T) {
	// mock CiSystemVersionConstraint
	type args struct {
		info           wd_info.WoodpeckerInfo
		minimumVersion string
		maximumVersion string
	}
	tests := []struct {
		name    string
		args    args
		wantErr error
	}{
		{
			name: "empty CI_SYSTEM_VERSION",
			args: args{
				info:           mockWoodpeckerInfoByVersion(""),
				minimumVersion: "2.0.0",
				maximumVersion: "3.0.0",
			},
			wantErr: fmt.Errorf("woodpeacker ci system version is empty, please set env CI_SYSTEM_VERSION"),
		},
		{
			name: "empty minimumVersion",
			args: args{
				info:           mockWoodpeckerInfoByVersion(wd_info.WoodpeckerInfoSupportVersion),
				minimumVersion: "",
				maximumVersion: "3.0.0",
			},
			wantErr: fmt.Errorf("minimum version is empty, please check"),
		},
		{
			name: "empty maximumVersion",
			args: args{
				info:           mockWoodpeckerInfoByVersion(wd_info.WoodpeckerInfoSupportVersion),
				minimumVersion: "2.0.0",
				maximumVersion: "",
			},
			wantErr: fmt.Errorf("maximum version is empty, please check"),
		},
		{
			name: "not support CI_SYSTEM_VERSION",
			args: args{
				info:           mockWoodpeckerInfoByVersion("Semantic Versioning"),
				minimumVersion: "2.0.0",
				maximumVersion: "3.0.0",
			},
			wantErr: fmt.Errorf("can not parse target version: Semantic Versioning err: Invalid Semantic Version"),
		},
		{
			name: "not support minimumVersion",
			args: args{
				info:           mockWoodpeckerInfoByVersion(wd_info.WoodpeckerInfoSupportVersion),
				minimumVersion: "Semantic Versioning",
				maximumVersion: "3.0.0",
			},
			wantErr: fmt.Errorf("can not parse constraint: <= 3.0.0, >= Semantic Versioning err: improper constraint: <= 3.0.0, >= Semantic Versioning"),
		},
		{
			name: "not support maximumVersion",
			args: args{
				info:           mockWoodpeckerInfoByVersion(wd_info.WoodpeckerInfoSupportVersion),
				minimumVersion: "2.0.0",
				maximumVersion: "Semantic Versioning",
			},
			wantErr: fmt.Errorf("can not parse constraint: <= Semantic Versioning, >= 2.0.0 err: improper constraint: <= Semantic Versioning, >= 2.0.0"),
		},
		{
			name: "less",
			args: args{
				info:           mockWoodpeckerInfoByVersion("1.0.0"),
				minimumVersion: "2.0.0",
				maximumVersion: "3.0.0",
			},
			wantErr: fmt.Errorf("semver version: 1.0.0 not support, err: [1.0.0 is less than 2.0.0]"),
		},
		{
			name: "constraint",
			args: args{
				info:           mockWoodpeckerInfoByVersion("2.3.0"),
				minimumVersion: "2.0.0",
				maximumVersion: "3.0.0",
			},
		},
		{
			name: "greater",
			args: args{
				info:           mockWoodpeckerInfoByVersion("2.11.0"),
				minimumVersion: "2.0.0",
				maximumVersion: "2.10.0",
			},
			wantErr: fmt.Errorf("semver version: 2.11.0 not support, err: [2.11.0 is greater than 2.10.0]"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do CiSystemVersionConstraint
			gotErr := wd_info.CiSystemVersionConstraint(tc.args.info, tc.args.minimumVersion, tc.args.maximumVersion)

			// verify CiSystemVersionConstraint
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
		})
	}
}
