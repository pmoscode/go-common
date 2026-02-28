package config

import (
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
