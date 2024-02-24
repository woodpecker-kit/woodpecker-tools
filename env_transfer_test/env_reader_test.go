package env_transfer_test

import (
	"github.com/sebdah/goldie/v2"
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/env_transfer"
	"path/filepath"
	"testing"
)

const transferReaderSampleDirName = "transfer_reader_sample"

func TestOverloadEnvFromFile(t *testing.T) {
	// mock OverloadEnvFromFile
	tests := []struct {
		name         string
		root         string
		fileName     string
		writeEnvMap  map[string]string
		checkEnvMap  map[string]string
		wantWriteErr error
		wantReadErr  bool
	}{
		{
			name:     "sample",
			root:     filepath.Join(testBaseFolderPath, transferReaderSampleDirName, "sample"),
			fileName: env_transfer.DefaultWriterFileName,
			writeEnvMap: map[string]string{
				"foo": "bar",
				"bar": "foo",
			},
			checkEnvMap: map[string]string{
				"foo": "bar",
				"bar": "foo",
			},
		},
		{
			name:     "specific_symbol_key_dot",
			root:     filepath.Join(testBaseFolderPath, transferReaderSampleDirName, "specific_symbol_key_dot"),
			fileName: env_transfer.DefaultWriterFileName,
			writeEnvMap: map[string]string{
				"dot": "foo.bar",
			},
			checkEnvMap: map[string]string{
				"dot": "foo.bar",
			},
			wantReadErr: false,
		},
		{
			name:     "specific_symbol_key_single_quotation",
			root:     filepath.Join(testBaseFolderPath, transferReaderSampleDirName, "specific_symbol_key_single_quotation"),
			fileName: env_transfer.DefaultWriterFileName,
			writeEnvMap: map[string]string{
				"quotation": `foo'bar`,
			},
			checkEnvMap: map[string]string{
				"quotation": `foo'bar`,
			},
			wantReadErr: false,
		},
		{
			name:     "specific_symbol_key_double_quotation",
			root:     filepath.Join(testBaseFolderPath, transferReaderSampleDirName, "specific_symbol_key_double_quotation"),
			fileName: env_transfer.DefaultWriterFileName,
			writeEnvMap: map[string]string{
				"quotation": `foo"bar`,
			},
			checkEnvMap: map[string]string{
				"quotation": `foo"bar`,
			},
			wantReadErr: true,
		},
		{
			name:     "specific_symbol_key_eq",
			root:     filepath.Join(testBaseFolderPath, transferReaderSampleDirName, "specific_symbol_key_eq"),
			fileName: env_transfer.DefaultWriterFileName,
			writeEnvMap: map[string]string{
				"eq": "foo=bar",
			},
			checkEnvMap: map[string]string{
				"eq": "foo=bar",
			},
			wantReadErr: false,
		},
		{
			name:     "specific_symbol_key_french_quotes",
			root:     filepath.Join(testBaseFolderPath, transferReaderSampleDirName, "specific_symbol_key_french_quotes"),
			fileName: env_transfer.DefaultWriterFileName,
			writeEnvMap: map[string]string{
				"french_quotes": `foo:bar`,
			},
			checkEnvMap: map[string]string{
				"french_quotes": `foo:bar`,
			},
			wantReadErr: false,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do WriteEnv2File
			_, gotErr := env_transfer.WriteEnv2File(tc.root, tc.fileName, tc.writeEnvMap)
			assert.Equal(t, tc.wantWriteErr, gotErr)
			if tc.wantWriteErr != nil {
				return
			}

			// do OverloadEnvFromFile
			loadEnvMap, readErr := env_transfer.OverloadEnvFromFile(tc.root, tc.fileName)
			assert.Equal(t, readErr != nil, tc.wantReadErr)
			if tc.wantReadErr {
				t.Logf("readErr: %v", readErr)
				return
			}
			for k, v := range tc.checkEnvMap {
				assert.Equal(t, loadEnvMap[k], v)
				assert.Equal(t, env_kit.FetchOsEnvStr(k, ""), v)
			}
			// verify OverloadEnvFromFile
			g.AssertJson(t, t.Name(), loadEnvMap)
		})
	}
}
