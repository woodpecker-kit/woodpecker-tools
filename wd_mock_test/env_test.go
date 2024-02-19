package wd_mock_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/woodpecker-kit/woodpecker-tools/env_mock"
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

	assert.True(t, env_mock.FetchOsEnvBool(keyEnvDebug, false))
	assert.Equal(t, 2, env_mock.FetchOsEnvInt(keyEnvCiNum, 0))
	assert.Equal(t, "foo", env_mock.FetchOsEnvStr(keyEnvCiKey, ""))
	envArray := env_mock.FetchOsEnvArray(keyEnvs)
	assert.Nil(t, envArray)

	setEnvStr(t, keyEnvs, "foo, bar,My ")

	envArray = env_mock.FetchOsEnvArray(keyEnvs)

	assert.NotNil(t, envArray)
	assert.Equal(t, "foo", envArray[0])
	assert.Equal(t, "bar", envArray[1])
	assert.Equal(t, "My", envArray[2])
}
