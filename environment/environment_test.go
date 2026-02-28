package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringEnv(t *testing.T) {
	t.Setenv("keyStr", "something")
	t.Setenv("keyInt", "12")
	t.Setenv("keyBool1", "true")
	t.Setenv("keyBool2", "false")
	t.Setenv("keyFloat32", "31.5")
	t.Setenv("keyFloat64", "34.5")

	assert.Equal(t, "something", GetEnv("keyStr", "nothing"))
	assert.Equal(t, 12, GetEnvInt("keyInt", 0))
	assert.Equal(t, true, GetEnvBool("keyBool1", false))
	assert.Equal(t, false, GetEnvBool("keyBool2", true))
	assert.Equal(t, float32(31.5), GetEnvFloat32("keyFloat32", 11.4))
	assert.Equal(t, 34.5, GetEnvFloat64("keyFloat64", 12.4))
}

func TestStringDefaultEnv(t *testing.T) {
	t.Setenv("keyInt", "notAnInt")
	t.Setenv("keyBool1", "3")
	t.Setenv("keyBool2", "blah")
	t.Setenv("keyFloat32", "true")
	t.Setenv("keyFloat64", "23fb")
	t.Setenv("keyFloat64int", "34")

	assert.Equal(t, "nothing", GetEnv("keyStr", "nothing"))
	assert.Equal(t, 50, GetEnvInt("keyInt", 50))
	assert.Equal(t, false, GetEnvBool("keyBool1", false))
	assert.Equal(t, true, GetEnvBool("keyBool2", true))
	assert.Equal(t, float32(10.5), GetEnvFloat32("keyFloat32", 10.5))
	assert.Equal(t, 20.5, GetEnvFloat64("keyFloat64", 20.5))
	assert.Equal(t, float64(34), GetEnvFloat64("keyFloat64int", 30.5))
}

func TestEnvMap(t *testing.T) {
	t.Setenv("TEST_one", "one")
	t.Setenv("TEST_two", "two")
	t.Setenv("TEST_three", "3")
	t.Setenv("TESTING_invalid", "4")
	t.Setenv("four", "4")

	envMap := GetEnvMap("TEST", true)

	assert.Len(t, envMap, 3)
}
