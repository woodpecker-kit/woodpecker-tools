package wd_mock_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/wd_mock"
	"testing"
)

func TestNewCiForgeInfo(t *testing.T) {
	// mock NewCiForgeInfo
	type args struct {
		forgeType string
		forgeUrl  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "gitea", // testdata/TestNewCiForgeInfo/sample.golden
			args: args{
				forgeType: "gitea",
				forgeUrl:  "https://gitea.domain.com",
			},
		},
		{
			name: "github", // testdata/TestNewCiForgeInfo/sample.golden
			args: args{
				forgeType: "github",
				forgeUrl:  "https://github.com",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do NewCiForgeInfo
			gotResult := wd_mock.NewCiForgeInfo(
				wd_mock.WithCiForgeType(tc.args.forgeType),
				wd_mock.WithCiForgeUrl(tc.args.forgeUrl),
			)
			assert.Equal(t, tc.wantErr, gotResult == nil)
			if tc.wantErr {
				return
			}
			// verify NewCiForgeInfo
			g.AssertJson(t, t.Name(), gotResult)
		})
	}
}
