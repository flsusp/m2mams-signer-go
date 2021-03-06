package kprovider

import (
	"github.com/flsusp/m2mams-signer-go/m2mams"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadKeyFromSpecificVar(t *testing.T) {
	fakeEnv := m2mams.NewFakeEnv()
	pkp := EnvironmentVariableKProvider{
		Environment: fakeEnv,
	}

	fakeEnv.Setenv("SOMECONTEXT_SOMEKEYPAIR_PK", validPrivateKey)

	key, err := pkp.LoadPrivateKey("", "somecontext", "somekeypair")
	assert.NoError(t, err)
	err = key.Validate()
	assert.NoError(t, err)
}

func TestLoadKeyFromGenericVar(t *testing.T) {
	fakeEnv := m2mams.NewFakeEnv()
	pkp := EnvironmentVariableKProvider{
		Environment: fakeEnv,
	}

	fakeEnv.Setenv("M2MAMS_PK", validPrivateKey)

	key, err := pkp.LoadPrivateKey("", "somecontext", "somekeypair")
	assert.NoError(t, err)
	err = key.Validate()
	assert.NoError(t, err)
}

func TestFailsLoadingKeyIfEnvVarsNotFound(t *testing.T) {
	fakeEnv := m2mams.NewFakeEnv()
	pkp := EnvironmentVariableKProvider{
		Environment: fakeEnv,
	}

	key, err := pkp.LoadPrivateKey("", "somecontext", "somekeypair")
	assert.Error(t, err)
	assert.Nil(t, key)
}
