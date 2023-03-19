package yamlconfig

import "testing"

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
	if err != nil {
		t.Fatal("Error thrown: ", err)
	}

	if c.Data.One != "localhost" {
		t.Fatal("Config: MqttHost is '", c.Data.One, "' not 'localhost'")
	}
}
