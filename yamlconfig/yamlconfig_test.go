package yamlconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, err)
	assert.Equal(t, "localhost", c.Data.One)
	assert.Equal(t, 2, c.Data.Two)
}
