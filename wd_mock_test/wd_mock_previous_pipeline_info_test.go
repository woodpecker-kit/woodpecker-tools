package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewPreviousPipelineInfo(t *testing.T) {
	// mock NewPreviousPipelineInfo
	type args struct {
		ciPreviousPipelineNumber       string
		ciPreviousPipelineParent       string
		ciPreviousPipelineEvent        string
		ciPreviousPipelineUrl          string
		ciPreviousPipelineForgeUrl     string
		ciPreviousPipelineDeployTarget string
		ciPreviousPipelineStatus       string
		ciPreviousPipelineCreated      uint64
		ciPreviousPipelineStarted      uint64
		ciPreviousPipelineFinished     uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewPreviousPipelineInfo/sample.golden
			args: args{
				ciPreviousPipelineNumber:       "9",
				ciPreviousPipelineParent:       "",
				ciPreviousPipelineEvent:        "push",
				ciPreviousPipelineUrl:          "https://woodpecker.domain.com/repos/2/pipeline/9",
				ciPreviousPipelineForgeUrl:     "https://gitea.domain.com/woodpecker-kit/guidance-woodpecker-agent/commit/b2542f37e23645b2032fcf9f180e6940ef6915ac",
				ciPreviousPipelineDeployTarget: "",
				ciPreviousPipelineStatus:       "success",
				ciPreviousPipelineCreated:      1705657931,
				ciPreviousPipelineStarted:      1705657942,
				ciPreviousPipelineFinished:     1705657957,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewPreviousPipelineInfo
			gotResult := wd_mock.NewPreviousPipelineInfo(
				wd_mock.WithCiPreviousPipelineNumber(tc.args.ciPreviousPipelineNumber),
				wd_mock.WithCiPreviousPipelineParent(tc.args.ciPreviousPipelineParent),
				wd_mock.WithCiPreviousPipelineEvent(tc.args.ciPreviousPipelineEvent),
				wd_mock.WithCiPreviousPipelineUrl(tc.args.ciPreviousPipelineUrl),
				wd_mock.WithCiPreviousPipelineForgeUrl(tc.args.ciPreviousPipelineForgeUrl),
				wd_mock.WithCiPreviousPipelineDeployTarget(tc.args.ciPreviousPipelineDeployTarget),
				wd_mock.WithCiPreviousPipelineStatus(tc.args.ciPreviousPipelineStatus),
				wd_mock.WithCiPreviousPipelineCreated(tc.args.ciPreviousPipelineCreated),
				wd_mock.WithCiPreviousPipelineStarted(tc.args.ciPreviousPipelineStarted),
				wd_mock.WithCiPreviousPipelineFinished(tc.args.ciPreviousPipelineFinished),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewPreviousPipelineInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
