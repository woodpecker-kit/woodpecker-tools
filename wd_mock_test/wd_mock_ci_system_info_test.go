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
		woodpeckerBackend       string
		woodpeckerAgentHostName string
		woodpeckerFilterLabels  []string
		ciMachine               string
		ciSystemPlatform        string
		ciSystemName            string
		ciSystemUrl             string
		ciSystemHost            string
		ciSystemVersion         string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "woodpecker", // testdata/TestNewCiSystemInfo/woodpecker.golden
			args: args{
				woodpeckerBackend:       "docker",
				woodpeckerAgentHostName: "",
				woodpeckerFilterLabels:  nil,
				ciMachine:               "worker",
				ciSystemPlatform:        "linux/amd64",
				ciSystemName:            "woodpecker",
				ciSystemUrl:             "https://woodpecker.domain.com",
				ciSystemHost:            "woodpecker.domain.com",
				ciSystemVersion:         wd_info.WoodpeckerInfoSupportVersion,
			},
		},
		{
			name: "self-system", // testdata/TestNewCiSystemInfo/self-system.golden
			args: args{
				woodpeckerBackend:       "local",
				woodpeckerAgentHostName: "",
				woodpeckerFilterLabels:  nil,
				ciMachine:               "worker",
				ciSystemPlatform:        "linux/amd64",
				ciSystemName:            "self-system",
				ciSystemUrl:             "https://self-system.domain.com",
				ciSystemHost:            "self-system.domain.com",
				ciSystemVersion:         "v1.0.0",
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
				wd_mock.WithWoodpeckerBackend(tc.args.woodpeckerBackend),
				wd_mock.WithWoodpeckerAgentHostName(tc.args.woodpeckerAgentHostName),
				wd_mock.WithWoodpeckerFilterLabels(tc.args.woodpeckerFilterLabels),
				wd_mock.WithCiMachine(tc.args.ciMachine),
				wd_mock.WithCiSystemPlatform(tc.args.ciSystemPlatform),
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
