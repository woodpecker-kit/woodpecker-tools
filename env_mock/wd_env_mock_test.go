package env_mock

import (
	"github.com/sinlov-go/unittest-kit/env_kit"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestMockEnvByStruct(t *testing.T) {
	// mock MockEnvByStruct

	type MockConfig struct {
		KeyFoo string `mock_env_key:"PLUGIN_KEY_FOO" mock_env_default:"foo"`
		KeyBar string `mock_env_key:"PLUGIN_KEY_BAR"`
	}

	type MockCase struct {
		PluginDebug bool `mock_env_key:"PLUGIN_DEBUG"`
		TimeOut     uint `mock_env_key:"PLUGIN_TIMEOUT"`

		MockConfig MockConfig
	}

	tests := []struct {
		name    string
		c       interface{}
		wantRes map[string]string
		wantErr bool
	}{
		{
			name: "sample",
			c: MockCase{
				PluginDebug: true,
				TimeOut:     10,

				MockConfig: MockConfig{
					KeyFoo: "",
					KeyBar: "bar",
				},
			},
			wantRes: map[string]string{
				"PLUGIN_DEBUG":   "true",
				"PLUGIN_TIMEOUT": "10",
				"PLUGIN_KEY_FOO": "foo",
				"PLUGIN_KEY_BAR": "bar",
			},
		},

		{
			name: "ptr",
			c: &MockCase{
				PluginDebug: true,
				TimeOut:     10,

				MockConfig: MockConfig{
					KeyFoo: "",
					KeyBar: "bar",
				},
			},
			wantRes: map[string]string{
				"PLUGIN_DEBUG":   "true",
				"PLUGIN_TIMEOUT": "10",
				"PLUGIN_KEY_FOO": "foo",
				"PLUGIN_KEY_BAR": "bar",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			// do MockEnvByStruct
			MockEnvByStruct(tc.c)

			// verify MockEnvByStruct
			assert.False(t, tc.wantErr)
			if tc.wantErr {
				return
			}
			//assert.Equal(t, tc.wantRes, gotResult)
			for key, val := range tc.wantRes {
				osEnvStr := env_kit.FetchOsEnvStr(key, "")
				assert.Equal(t, val, osEnvStr)
			}

			for _, e := range os.Environ() {
				if strings.Index(e, "PLUGIN") == 0 {
					t.Logf("env: %s\n", e)
				}
			}
		})
	}
}

func TestPanicMockEnvCanNotSupportDouble(t *testing.T) {
	// mock TestPanicMockEnvCanNotSupportDouble
	errString := "can not support double ptr, please check your input"

	type MockConfig struct {
		KeyFoo string `mock_env_key:"PLUGIN_KEY_FOO" mock_env_default:"foo"`
		KeyBar string `mock_env_key:"PLUGIN_KEY_BAR"`
	}

	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicMockEnvCanNotSupportDouble
		target := &MockConfig{
			KeyFoo: "",
			KeyBar: "bar",
		}
		MockEnvByStruct(&target)
	}) {
		// verify TestPanicMockEnvCanNotSupportDouble
		t.Fatalf("TestPanicMockEnvCanNotSupportDouble should panic")
	}
}
