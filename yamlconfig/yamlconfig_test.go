package yamlconfig

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfig struct {
	Data TestConfigData `yaml:"someData"`
}

type TestConfigData struct {
	One string `yaml:"one"`
	Two int    `yaml:"two"`
}

func TestParseYaml(t *testing.T) {
	testData := `
someData:
  one: localhost
  two: 2
`
	c := TestConfig{}
	err := parseYaml([]byte(testData), &c)
	require.NoError(t, err)
	assert.Equal(t, "localhost", c.Data.One)
	assert.Equal(t, 2, c.Data.Two)
}

func TestParseYamlInvalid(t *testing.T) {
	err := parseYaml([]byte(`{{{invalid`), &TestConfig{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing YAML file")
}

func TestLoadConfigSuccess(t *testing.T) {
	content := `
someData:
  one: production
  two: 8080
`
	path := createTempYaml(t, content)

	var cfg TestConfig
	err := LoadConfig(path, &cfg)
	require.NoError(t, err)
	assert.Equal(t, "production", cfg.Data.One)
	assert.Equal(t, 8080, cfg.Data.Two)
}

func TestLoadConfigFileNotFound(t *testing.T) {
	err := LoadConfig("/nonexistent/file.yaml", &TestConfig{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error reading YAML file")
}

func TestLoadConfigInvalidYaml(t *testing.T) {
	path := createTempYaml(t, `{{{invalid yaml`)

	err := LoadConfig(path, &TestConfig{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "error parsing YAML file")
}

func createTempYaml(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")
	err := os.WriteFile(path, []byte(content), 0600)
	require.NoError(t, err)
	return path
}
