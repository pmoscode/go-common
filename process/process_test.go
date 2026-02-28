package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetExecutableName(t *testing.T) {
	name, err := GetExecutableName()
	require.NoError(t, err)
	assert.NotEmpty(t, name)
	assert.NotContains(t, name, "/", "should not contain path separator")
	assert.NotContains(t, name, "\\", "should not contain path separator")
}
