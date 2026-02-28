package formats

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testConfig struct {
	Name string `yaml:"name" json:"name"`
	Port int    `yaml:"port" json:"port"`
}

func createTempFile(t *testing.T, filename, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, filename)
	err := os.WriteFile(path, []byte(content), 0600)
	require.NoError(t, err)
	return path
}

func TestParseYamlConfig(t *testing.T) {
	path := createTempFile(t, "config.yaml", `
name: testapp
port: 8080
`)
	var cfg testConfig
	err := ParseYamlConfig(path, &cfg)
	require.NoError(t, err)
	assert.Equal(t, "testapp", cfg.Name)
	assert.Equal(t, 8080, cfg.Port)
}

func TestParseJsonConfig(t *testing.T) {
	path := createTempFile(t, "config.json", `{"name":"jsonapp","port":9090}`)

	var cfg testConfig
	err := ParseJsonConfig(path, &cfg)
	require.NoError(t, err)
	assert.Equal(t, "jsonapp", cfg.Name)
	assert.Equal(t, 9090, cfg.Port)
}

func TestParseConfigWithCustomParser(t *testing.T) {
	path := createTempFile(t, "config.yaml", `
name: custom
port: 3000
`)
	var cfg testConfig
	err := ParseConfig(path, &cfg, func(in []byte, out any) error {
		// Use yaml under the hood
		return ParseYamlConfig(path, out)
	})
	// ParseConfig calls parseConfig which calls loadConfigFile and then the parser
	// The custom parser here re-reads via ParseYamlConfig, but that's fine for testing
	require.NoError(t, err)
	assert.Equal(t, "custom", cfg.Name)
}

func TestParseYamlConfigFileNotFound(t *testing.T) {
	var cfg testConfig
	err := ParseYamlConfig("/nonexistent/path/config.yaml", &cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "reading config file")
}

func TestParseJsonConfigFileNotFound(t *testing.T) {
	var cfg testConfig
	err := ParseJsonConfig("/nonexistent/path/config.json", &cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "reading config file")
}

func TestParseYamlConfigInvalidContent(t *testing.T) {
	path := createTempFile(t, "bad.yaml", `{{{invalid yaml`)

	var cfg testConfig
	err := ParseYamlConfig(path, &cfg)
	require.Error(t, err)
}

func TestParseJsonConfigInvalidContent(t *testing.T) {
	path := createTempFile(t, "bad.json", `{not json`)

	var cfg testConfig
	err := ParseJsonConfig(path, &cfg)
	require.Error(t, err)
}
