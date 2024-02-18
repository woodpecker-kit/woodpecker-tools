package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewCurrentWorkflowInfo(t *testing.T) {
	// mock NewCurrentWorkflowInfo
	type args struct {
		ciWorkflowName   string
		ciWorkflowNumber string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sample", // testdata/TestNewCurrentWorkflowInfo/sample.golden
			args: args{
				ciWorkflowName:   "build",
				ciWorkflowNumber: "1",
			},
		},
		{
			name: "two", // testdata/TestNewCurrentWorkflowInfo/two.golden
			args: args{
				ciWorkflowName:   "notify",
				ciWorkflowNumber: "2",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewCurrentWorkflowInfo
			gotResult := wd_mock.NewCurrentWorkflowInfo(
				wd_mock.WithCiWorkflowName(tc.args.ciWorkflowName),
				wd_mock.WithCiWorkflowNumber(tc.args.ciWorkflowNumber),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewCurrentWorkflowInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
