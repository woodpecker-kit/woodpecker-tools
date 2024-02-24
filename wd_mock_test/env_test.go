package wd_mock_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/env_kit"
	"testing"
)

func TestEnvKeys(t *testing.T) {
	// mock EnvKeys
	const keyEnvs = "ENV_KEYS"
	t.Logf("~> mock EnvKeys")

	setEnvBool(t, keyEnvDebug, true)

	setEnvInt64(t, keyEnvCiNum, 2)

	setEnvStr(t, keyEnvCiKey, "foo")

	// do EnvKeys
	t.Logf("~> do EnvKeys")

	// verify EnvKeys

	assert.True(t, env_kit.FetchOsEnvBool(keyEnvDebug, false))
	assert.Equal(t, 2, env_kit.FetchOsEnvInt(keyEnvCiNum, 0))
	assert.Equal(t, "foo", env_kit.FetchOsEnvStr(keyEnvCiKey, ""))
	envArray := env_kit.FetchOsEnvArray(keyEnvs)
	assert.Nil(t, envArray)

	setEnvStr(t, keyEnvs, "foo, bar,My ")

	envArray = env_kit.FetchOsEnvArray(keyEnvs)

	assert.NotNil(t, envArray)
	assert.Equal(t, "foo", envArray[0])
	assert.Equal(t, "bar", envArray[1])
	assert.Equal(t, "My", envArray[2])
	t.Logf("~> print EnvKeys prefix as [ CI_ ]:\n%s\n", env_kit.FindAllEnvByPrefix4Print("CI_"))
	t.Logf("~> print All Env as sort and just:\n%s\n", env_kit.FindAllEnv4PrintAsSortJust(24))
}
