package config

import (
	"log"
	"testing"
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
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", testData)

	driveTest(1, t, testData.GmdDataStr == "Everything", "Everything", testData.GmdDataStr)
	driveTest(2, t, testData.ButtDataStr1 == "Karl Ranseier", "Karl Ranseier", testData.ButtDataStr1)
	driveTest(3, t, testData.ButtDataStr2 == "ist tot", "ist tot", testData.ButtDataStr2)
	driveTest(4, t, testData.GmdDataInt == 42, 42, testData.GmdDataInt)

	driveTest(5, t, len(testData.GmdDataMap1) == 3, 3, len(testData.GmdDataMap1))
	driveTest(6, t, len(testData.GmdDataMap2) == 4, 4, len(testData.GmdDataMap2))
	driveTest(7, t, len(testData.GmdDataMap3) == 2, 2, len(testData.GmdDataMap3))
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
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%+v", testData)

	driveTest(1, t, testData.GmdDataStr == "default string", "default string", testData.GmdDataStr)
	driveTest(2, t, testData.ButtDataStr1 == "Karl Ranseier", "Karl Ranseier", testData.ButtDataStr1)
	driveTest(3, t, testData.ButtDataStr2 == "ist tot", "ist tot", testData.ButtDataStr2)
	driveTest(4, t, testData.GmdDataInt == 42, 42, testData.GmdDataInt)
	driveTest(5, t, testData.GmdNoDataInt == 0, 0, testData.GmdNoDataInt)
	driveTest(6, t, testData.GmdData3 == "no prefix", "no prefix", testData.GmdData3)

	driveTest(7, t, len(testData.GmdDataMap1) == 3, 3, len(testData.GmdDataMap1))
	driveTest(8, t, len(testData.GmdDataMap2) == 4, 4, len(testData.GmdDataMap2))
}

func driveTest(testLine int, t *testing.T, condition bool, expectedValue any, actualValue any) {
	if !condition {
		t.Fatalf("expected '%v' for test-line '%d', but got '%v'!", expectedValue, testLine, actualValue)
	}
}
