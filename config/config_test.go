package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type TestConfigSimpleData struct {
	GmdDataStr   string             `env:"name=DATA_ONE_STR"`
	ButtDataStr1 string             `env:""`
	ButtDataStr2 string             `env:"self"`
	GmdDataInt   int                `env:"name=DATA_ONE_INT"`
	GmdDataMap1  map[string]string  `env:"prefix=GMD,cutoff=false"`
	GmdDataMap2  map[string]float32 `env:"prefix=DAT,cutoff=true"`
	GmdDataMap3  map[string]string  `env:"prefix=CUT"`
}

type TestConfigComplexData struct {
	GmdDataStr   string             `env:"name=DATA_ONE_STR,default=default string"`
	ButtDataStr1 string             `env:"name=BUTT_DATA,self"`
	ButtDataStr2 string             `env:"self"`
	GmdDataInt   int                `env:"name=DATA_ONE_INT"`
	GmdNoDataInt int                `env:"name=DATA_NO_INT"`
	GmdDataMap1  map[string]string  `env:"prefix=GMD,cutoff=false"`
	GmdDataMap2  map[string]float32 `env:"prefix=DAT,cutoff"`
	GmdData3     string             `env:"name=NO_DAT_FIRST,prefix=CUT"`
}

func TestResolveTagsSimple(t *testing.T) {
	t.Setenv("DATA_ONE_STR", "Everything")
	t.Setenv("DATA_ONE_INT", "42")
	t.Setenv("GMD_STRING_1", "one")
	t.Setenv("GMD_STRING_2", "two")
	t.Setenv("GMD_STRING_3", "three")
	t.Setenv("CUT_STRING_1", "two")
	t.Setenv("CUT_STRING_2", "three")
	t.Setenv("DAT_INT_1", "1")
	t.Setenv("DAT_INT_2", "2")
	t.Setenv("DAT_INT_3", "3")
	t.Setenv("DAT_INT_4", "4")
	t.Setenv("BUTT_DATA_STR_1", "Karl Ranseier")
	t.Setenv("BUTT_DATA_STR_2", "ist tot")

	testData := TestConfigSimpleData{}

	err := LoadFromEnvironment(&testData)
	require.NoError(t, err)

	assert.Equal(t, "Everything", testData.GmdDataStr)
	assert.Equal(t, "Karl Ranseier", testData.ButtDataStr1)
	assert.Equal(t, "ist tot", testData.ButtDataStr2)
	assert.Equal(t, 42, testData.GmdDataInt)
	assert.Len(t, testData.GmdDataMap1, 3)
	assert.Len(t, testData.GmdDataMap2, 4)
	assert.Len(t, testData.GmdDataMap3, 2)
}

func TestResolveTagsComplex(t *testing.T) {
	t.Setenv("BUTT_DATA", "Karl Ranseier")
	t.Setenv("BUTT_DATA_STR_2", "ist tot")
	t.Setenv("DATA_ONE_INT", "42")
	t.Setenv("GMD_STRING_1", "one")
	t.Setenv("GMD_STRING_2", "two")
	t.Setenv("GMD_STRING_3", "three")
	t.Setenv("DAT_INT_1", "1")
	t.Setenv("DAT_INT_2", "2")
	t.Setenv("DAT_INT_3", "3")
	t.Setenv("DAT_INT_4", "4")
	t.Setenv("NO_DAT_FIRST", "no prefix")

	testData := TestConfigComplexData{}

	err := LoadFromEnvironment(&testData)
	require.NoError(t, err)

	assert.Equal(t, "default string", testData.GmdDataStr)
	assert.Equal(t, "Karl Ranseier", testData.ButtDataStr1)
	assert.Equal(t, "ist tot", testData.ButtDataStr2)
	assert.Equal(t, 42, testData.GmdDataInt)
	assert.Equal(t, 0, testData.GmdNoDataInt)
	assert.Equal(t, "no prefix", testData.GmdData3)
	assert.Len(t, testData.GmdDataMap1, 3)
	assert.Len(t, testData.GmdDataMap2, 4)
}

func TestLoadFromEnvironmentNonPointer(t *testing.T) {
	type Cfg struct {
		Val string `env:"name=X"`
	}
	err := LoadFromEnvironment(Cfg{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "pointer")
}

func TestLoadFromEnvironmentBoolField(t *testing.T) {
	type Cfg struct {
		Flag bool `env:"name=BOOL_CFG"`
	}
	t.Setenv("BOOL_CFG", "true")

	var cfg Cfg
	err := LoadFromEnvironment(&cfg)
	require.NoError(t, err)
	assert.True(t, cfg.Flag)
}

type unexportedFieldConfig struct {
	Public  string `env:"name=PUB_VAL"`
	private string `env:"name=PRIV_VAL"` //nolint:unused // test for unexported field handling
}

func TestLoadFromEnvironmentUnexportedField(t *testing.T) {
	t.Setenv("PUB_VAL", "public")
	t.Setenv("PRIV_VAL", "private")

	cfg := unexportedFieldConfig{}
	err := LoadFromEnvironment(&cfg)
	require.NoError(t, err)
	assert.Equal(t, "public", cfg.Public)
	// private field should not be set (unexported)
}

func TestLoadFromEnvironmentNoTags(t *testing.T) {
	type Cfg struct {
		Value string
	}
	cfg := Cfg{Value: "original"}
	err := LoadFromEnvironment(&cfg)
	require.NoError(t, err)
	assert.Equal(t, "original", cfg.Value, "fields without env tag should not be modified")
}

func TestLoadFromFileYaml(t *testing.T) {
	type FileCfg struct {
		Host string `yaml:"host" env:"name=YAML_FILE_HOST"`
		Port int    `yaml:"port" env:"name=YAML_FILE_PORT"`
	}
	content := `
host: filehost
port: 3000
`
	path := createTempConfigFile(t, "config.yaml", content)

	// Set env vars to match file values so LoadFromEnvironment doesn't overwrite
	t.Setenv("YAML_FILE_HOST", "filehost")
	t.Setenv("YAML_FILE_PORT", "3000")

	var cfg FileCfg
	err := LoadFromFile(path, &cfg, YAML)
	require.NoError(t, err)
	assert.Equal(t, "filehost", cfg.Host)
	assert.Equal(t, 3000, cfg.Port)
}

func TestLoadFromFileJson(t *testing.T) {
	type FileCfg struct {
		Host string `json:"host" env:"name=JSON_FILE_HOST"`
		Port int    `json:"port" env:"name=JSON_FILE_PORT"`
	}
	content := `{"host":"jsonhost","port":4000}`
	path := createTempConfigFile(t, "config.json", content)

	t.Setenv("JSON_FILE_HOST", "jsonhost")
	t.Setenv("JSON_FILE_PORT", "4000")

	var cfg FileCfg
	err := LoadFromFile(path, &cfg, JSON)
	require.NoError(t, err)
	assert.Equal(t, "jsonhost", cfg.Host)
	assert.Equal(t, 4000, cfg.Port)
}

func TestLoadFromFileNotFound(t *testing.T) {
	type Cfg struct {
		Val string `env:"name=X"`
	}
	var cfg Cfg
	err := LoadFromFile("/nonexistent/config.yaml", &cfg, YAML)
	require.Error(t, err)
}

func TestLoadFromFileEnvOverride(t *testing.T) {
	type FileCfg struct {
		Host string `yaml:"host" env:"name=OVERRIDE_HOST"`
	}
	content := `host: fromfile`
	path := createTempConfigFile(t, "config.yaml", content)

	t.Setenv("OVERRIDE_HOST", "fromenv")

	var cfg FileCfg
	err := LoadFromFile(path, &cfg, YAML)
	require.NoError(t, err)
	// Environment should override file value
	assert.Equal(t, "fromenv", cfg.Host)
}

func createTempConfigFile(t *testing.T, name, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, name)
	err := os.WriteFile(path, []byte(content), 0600)
	require.NoError(t, err)
	return path
}
