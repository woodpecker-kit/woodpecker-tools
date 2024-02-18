package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewCurrentPipelineInfo(t *testing.T) {
	// mock NewCurrentPipelineInfo
	type args struct {
		ciPipelineNumber       string
		ciPipelineParent       string
		ciPipelineEvent        string
		ciPipelineUrl          string
		ciPipelineForgeUrl     string
		ciPipelineDeployTarget string
		ciPipelineStatus       string
		ciPipelineCreated      uint64
		ciPipelineStarted      uint64
		ciPipelineFinished     uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewCurrentPipelineInfo/sample.golden
			args: args{
				ciPipelineNumber:       "10",
				ciPipelineParent:       "0",
				ciPipelineEvent:        "push",
				ciPipelineUrl:          "https://woodpecker.domain.com/repos/2/pipeline/10",
				ciPipelineForgeUrl:     "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/9c764dd487bce596c5c0402478fabde5f0344983",
				ciPipelineDeployTarget: "",
				ciPipelineStatus:       "success",
				ciPipelineCreated:      1705658141,
				ciPipelineStarted:      1705658156,
				ciPipelineFinished:     1705658166,
			},
		},
		{
			name: "one", // testdata/TestNewCurrentPipelineInfo/one.golden
			args: args{
				ciPipelineNumber:       "1",
				ciPipelineParent:       "0",
				ciPipelineEvent:        "push",
				ciPipelineUrl:          "https://woodpecker.domain.com/repos/2/pipeline/1",
				ciPipelineForgeUrl:     "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/9c764dd487bce596c5c0402478fabde5f0345983",
				ciPipelineDeployTarget: "",
				ciPipelineStatus:       "success",
				ciPipelineCreated:      1705638141,
				ciPipelineStarted:      1705638156,
				ciPipelineFinished:     1705458166,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewCurrentPipelineInfo
			gotResult := wd_mock.NewCurrentPipelineInfo(
				wd_mock.WithCiPipelineNumber(tc.args.ciPipelineNumber),
				wd_mock.WithCiPipelineParent(tc.args.ciPipelineParent),
				wd_mock.WithCiPipelineEvent(tc.args.ciPipelineEvent),
				wd_mock.WithCiPipelineUrl(tc.args.ciPipelineUrl),
				wd_mock.WithCiPipelineForgeUrl(tc.args.ciPipelineForgeUrl),
				wd_mock.WithCiPipelineDeployTarget(tc.args.ciPipelineDeployTarget),
				wd_mock.WithCiPipelineStatus(tc.args.ciPipelineStatus),
				wd_mock.WithCiPipelineCreated(tc.args.ciPipelineCreated),
				wd_mock.WithCiPipelineStarted(tc.args.ciPipelineStarted),
				wd_mock.WithCiPipelineFinished(tc.args.ciPipelineFinished),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewCurrentPipelineInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
