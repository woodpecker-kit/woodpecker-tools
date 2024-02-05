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
		name         string
		root         string
		fileName     string
		envMap       map[string]string
		wantSavePath string
		wantErr      error
	}{
		{
			name:     "sample",
			root:     filepath.Join(testBaseFolderPath, transferSampleDirName, "sample"),
			fileName: env_transfer.DefaultWriterFileName,
			envMap: map[string]string{
				"foo": "bar",
				"bar": "foo",
			},
			wantSavePath: filepath.Join(testBaseFolderPath, transferSampleDirName, "sample", env_transfer.DefaultWriterFileName),
		},
		{
			name:     "some_empty_key",
			root:     filepath.Join(testBaseFolderPath, transferSampleDirName, "some_empty_key"),
			fileName: env_transfer.DefaultWriterFileName,
			envMap: map[string]string{
				"foo": "bar",
				"":    "foo",
			},
			wantSavePath: filepath.Join(testBaseFolderPath, transferSampleDirName, "some_empty_key", env_transfer.DefaultWriterFileName),
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
			wantSavePath: filepath.Join(testBaseFolderPath, transferSampleDirName, "specific_symbol_key", env_transfer.DefaultWriterFileName),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ColoredDiff),
			)

			// do WriteEnv2File
			dataFilePath, gotErr := env_transfer.WriteEnv2File(tc.root, tc.fileName, tc.envMap)
			assert.Equal(t, tc.wantErr, gotErr)
			if tc.wantErr != nil {
				return
			}
			assert.Equal(t, tc.wantSavePath, dataFilePath)
			readEnvMap, readErr := env_transfer.OverloadEnvFromFile(tc.root, tc.fileName)
			if readErr != nil {
				t.Logf("readErr: %v", readErr)
				return
			}
			// verify WriteEnv2File
			g.AssertJson(t, t.Name(), readEnvMap)
		})
	}
}

const transferEnvByKeysDirName = "transfer_env_by_keys"

func TestEnvByKeys(t *testing.T) {
	// mock EnvByKeys
	tests := []struct {
		name          string
		root          string
		fileName      string
		wantSaveErr   bool
		addEnv        map[string]string
		wantAddErr    bool
		removeEnv     []string
		wantRemoveErr bool
	}{
		{
			name:     "sample", // testdata/TestEnvByKeys/sample.golden
			root:     filepath.Join(testBaseFolderPath, transferEnvByKeysDirName, "sample"),
			fileName: env_transfer.DefaultWriterFileName,
			addEnv: map[string]string{
				"foo": "bar",
				"bar": "baz",
			},
			removeEnv: []string{"bar"},
		},
		{
			name:     "add_err", // testdata/TestEnvByKeys/sample.golden
			root:     filepath.Join(testBaseFolderPath, transferEnvByKeysDirName, "add_err"),
			fileName: env_transfer.DefaultWriterFileName,
			addEnv: map[string]string{
				`foo"one"`: "bar",
				"bar":      "baz",
			},
			wantAddErr:    true,
			removeEnv:     []string{"bar"},
			wantRemoveErr: false,
		},
		{
			name:     "add_empty", // testdata/TestEnvByKeys/sample.golden
			root:     filepath.Join(testBaseFolderPath, transferEnvByKeysDirName, "add_empty"),
			fileName: env_transfer.DefaultWriterFileName,
			addEnv: map[string]string{
				"": "bar",
			},
			wantAddErr:    true,
			removeEnv:     []string{"bar"},
			wantRemoveErr: true,
		},
		{
			name:     "remove_err", // testdata/TestEnvByKeys/sample.golden
			root:     filepath.Join(testBaseFolderPath, transferEnvByKeysDirName, "remove_err"),
			fileName: env_transfer.DefaultWriterFileName,
			addEnv: map[string]string{
				"foo": "bar",
				"bar": "baz",
			},
			wantAddErr:    false,
			removeEnv:     []string{`bar_one`},
			wantRemoveErr: true,
		},
		{
			name:     "remove_empty", // testdata/TestEnvByKeys/sample.golden
			root:     filepath.Join(testBaseFolderPath, transferEnvByKeysDirName, "remove_empty"),
			fileName: env_transfer.DefaultWriterFileName,
			addEnv: map[string]string{
				"foo": "bar",
				"bar": "baz",
			},
			wantAddErr:    false,
			removeEnv:     []string{""},
			wantRemoveErr: true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := goldie.New(t,
				goldie.WithDiffEngine(goldie.ClassicDiff),
			)

			// do EnvByKeys

			if len(tc.addEnv) > 0 {
				for key, v := range tc.addEnv {
					errAdd := env_transfer.AddOrCoverEnvByKey(key, v)
					assert.Equal(t, tc.wantAddErr, errAdd != nil)
					if tc.wantAddErr {
						t.Logf("want errAdd: %v", errAdd)
						return
					}
				}
			}

			if len(tc.removeEnv) > 0 {
				for _, key := range tc.removeEnv {
					errRemove := env_transfer.RemoveEnvByKey(key)
					assert.Equal(t, tc.wantRemoveErr, errRemove != nil)
					if tc.wantRemoveErr {
						t.Logf("want errRemove: %v", errRemove)
						return
					}
				}
			}
			_, saveEnvMap, gotErr := env_transfer.SaveEnv2File(tc.root, tc.fileName)
			assert.Equal(t, tc.wantSaveErr, gotErr != nil)
			if tc.wantSaveErr {
				t.Logf("want gotErr: %v", gotErr)
				return
			}
			readEnvMap, readErr := env_transfer.OverloadEnvFromFile(tc.root, tc.fileName)
			assert.Equal(t, readErr != nil, tc.wantSaveErr)
			if tc.wantSaveErr {
				t.Logf("readErr: %v", readErr)
				return
			}
			// verify EnvByKeys
			assert.Equal(t, saveEnvMap, readEnvMap)
			g.AssertJson(t, t.Name(), readEnvMap)
		})
	}
}
