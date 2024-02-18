package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewCurrentStepInfo(t *testing.T) {
	// mock NewCurrentStepInfo
	type args struct {
		ciStepName     string
		ciStepNumber   string
		ciStepStatus   string
		ciStepStarted  uint64
		ciStepFinished uint64
		ciStepUrl      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewCurrentStepInfo/sample.golden
			args: args{
				ciStepName:     "checkout",
				ciStepNumber:   "0",
				ciStepStatus:   "success",
				ciStepStarted:  1705658156,
				ciStepFinished: 1705658166,
				ciStepUrl:      "https://woodpecker.domain.com/repos/2/pipeline/10",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewCurrentStepInfo
			gotResult := wd_mock.NewCurrentStepInfo(
				wd_mock.WithCiStepName(tc.args.ciStepName),
				wd_mock.WithCiStepNumber(tc.args.ciStepNumber),
				wd_mock.WithCiStepStatus(tc.args.ciStepStatus),
				wd_mock.WithCiStepStarted(tc.args.ciStepStarted),
				wd_mock.WithCiStepFinished(tc.args.ciStepFinished),
				wd_mock.WithCiStepUrl(tc.args.ciStepUrl),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewCurrentStepInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
