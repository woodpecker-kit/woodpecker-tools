package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewBasicInfo(t *testing.T) {
	// mock NewBasicInfo
	type optsArgs struct {
		ci        string
		workspace string
	}
	tests := []struct {
		name     string
		optsArgs optsArgs
		wantErr  bool
	}{
		{
			name: "sample", // testdata/TestNewBasicInfo/sample.golden
			optsArgs: optsArgs{
				ci:        "sample",
				workspace: "sample",
			},
		},
		{
			name: "woodpecker", // testdata/TestNewBasicInfo/sample.golden
			optsArgs: optsArgs{
				ci:        "woodpecker",
				workspace: "woodpecker",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewBasicInfo
			gotResult := wd_mock.NewBasicInfo(
				wd_mock.WithCI(tc.optsArgs.ci),
				wd_mock.WithCIWorkspace(tc.optsArgs.workspace),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewBasicInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
