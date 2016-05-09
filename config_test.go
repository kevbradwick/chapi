package chapi

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBaseUrl(t *testing.T) {
	defaultUrl := BaseUrl()

	// default url
	assert.Equal(t, defaultUrl, BaseUrl())

	// overriding
	SetBaseUrl("foobar")
	assert.Equal(t, "foobar", BaseUrl())

	// reset
	SetBaseUrl(defaultUrl)
}

func TestApiKey(t *testing.T) {
	defaultKey := ApiKey()

	// empty by default
	assert.Empty(t, ApiKey())

	// override
	SetApiKey("foobar")
	assert.Equal(t, "foobar", ApiKey())

	// reset
	SetApiKey(defaultKey)
}
