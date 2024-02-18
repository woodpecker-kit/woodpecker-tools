package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewCiSystemInfo(t *testing.T) {
	// mock NewCiSystemInfo
	type args struct {
		ciSystemName    string
		ciSystemUrl     string
		ciSystemHost    string
		ciSystemVersion string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "woodpecker", // testdata/TestNewCiSystemInfo/woodpecker.golden
			args: args{
				ciSystemName:    "woodpecker",
				ciSystemUrl:     "https://woodpecker.domain.com",
				ciSystemHost:    "woodpecker.domain.com",
				ciSystemVersion: wd_info.WoodpeckerInfoSupportVersion,
			},
		},
		{
			name: "self-system", // testdata/TestNewCiSystemInfo/self-system.golden
			args: args{
				ciSystemName:    "self-system",
				ciSystemUrl:     "https://self-system.domain.com",
				ciSystemHost:    "self-system.domain.com",
				ciSystemVersion: "v1.0.0",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewCiSystemInfo
			gotResult := wd_mock.NewCiSystemInfo(
				wd_mock.WithCiSystemName(tc.args.ciSystemName),
				wd_mock.WithCiSystemUrl(tc.args.ciSystemUrl),
				wd_mock.WithCiSystemHost(tc.args.ciSystemHost),
				wd_mock.WithCiSystemVersion(tc.args.ciSystemVersion),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewCiSystemInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
