package env_transfer_test

import (
	"fmt"
	"github.com/sebdah/goldie/v2"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/env_transfer"
	"path/filepath"
	"testing"
)

func TestWriteEnv2File(t *testing.T) {
	// mock WriteEnv2File
	tests := []struct {
		name     string
		root     string
		fileName string
		envMap   map[string]string
		wantErr  error
	}{
		{
			name:     "sample",
			root:     filepath.Join(testBaseFolderPath, transferSampleDirName, "sample"),
			fileName: env_transfer.DefaultWriterFileName,
			envMap: map[string]string{
				"foo": "bar",
				"bar": "foo",
			},
		},
		{
			name:     "some_empty_key",
			root:     filepath.Join(testBaseFolderPath, transferSampleDirName, "some_empty_key"),
			fileName: env_transfer.DefaultWriterFileName,
			envMap: map[string]string{
				"foo": "bar",
				"":    "foo",
			},
		},
		{
			name:     "full_empty_key",
			root:     filepath.Join(testBaseFolderPath, transferSampleDirName, "full_empty_key"),
			fileName: env_transfer.DefaultWriterFileName,
			envMap:   map[string]string{},
			wantErr:  fmt.Errorf("envData is empty"),
		},
		{
			name:     "specific_symbol_key",
			root:     filepath.Join(testBaseFolderPath, transferSampleDirName, "specific_symbol_key"),
			fileName: env_transfer.DefaultWriterFileName,
			envMap: map[string]string{
				"dot":           "foo.bar",
				"quotation":     `foo"bar`,
				"eq":            "foo=bar",
				"french_quotes": `foo:bar`,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ColoredDiff),
			)

			// do WriteEnv2File
			envData, gotErr := env_transfer.WriteEnv2File(tc.root, tc.fileName, tc.envMap)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			readEnv, readErr := readFileAsByte(envData)
			if readErr != nil {
				t.Fatal(readErr)
			}
			// verify WriteEnv2File
			g.Assert(t, t.Name(), readEnv)
		})
	}
}
