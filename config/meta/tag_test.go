package meta

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type tagTestStruct struct {
	Named       string `env:"name=MY_VAR"`
	WithSelf    string `env:"self"`
	EmptyTag    string `env:""`
	NoTag       string
	Prefixed    string `env:"prefix=PRE"`
	Cutoff      string `env:"prefix=CUT,cutoff"`
	CutoffTrue  string `env:"prefix=CUT2,cutoff=true"`
	CutoffFalse string `env:"prefix=CUT3,cutoff=false"`
	WithDefault string `env:"name=DEF_VAR,default=hello"`
	NameAndSelf string `env:"name=PRIORITY,self"`
}

func fieldByName(name string) reflect.StructField {
	t := reflect.TypeOf(tagTestStruct{})
	f, _ := t.FieldByName(name)
	return f
}

func TestNewTagMeta(t *testing.T) {
	tag := NewTagMeta()
	assert.NotNil(t, tag)
	assert.Equal(t, Kind(None), tag.Kind())
}

func TestParseNoTag(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("NoTag"))
	require.NoError(t, err)
	assert.False(t, found)
	assert.Equal(t, Kind(None), tag.Kind())
}

func TestParseNameTag(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("Named"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, Kind(Name), tag.Kind())
	assert.Equal(t, "MY_VAR", tag.name)
}

func TestParseSelfTag(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("WithSelf"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, Kind(Name), tag.Kind())
	assert.Equal(t, "WITH_SELF", tag.name)
}

func TestParseEmptyTag(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("EmptyTag"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, Kind(Name), tag.Kind())
	assert.Equal(t, "EMPTY_TAG", tag.name)
}

func TestParsePrefixTag(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("Prefixed"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, Kind(Prefix), tag.Kind())
	assert.Equal(t, "PRE", tag.prefix)
	assert.False(t, tag.cutoff)
}

func TestParseCutoffImplicit(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("Cutoff"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.True(t, tag.cutoff)
}

func TestParseCutoffTrue(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("CutoffTrue"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.True(t, tag.cutoff)
}

func TestParseCutoffFalse(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("CutoffFalse"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.False(t, tag.cutoff)
}

func TestParseWithDefault(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("WithDefault"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "DEF_VAR", tag.name)
	assert.Equal(t, "hello", tag.defaultValue)
}

func TestParseNamePrioritySelf(t *testing.T) {
	tag := NewTagMeta()
	found, err := tag.Parse(fieldByName("NameAndSelf"))
	require.NoError(t, err)
	assert.True(t, found)
	assert.Equal(t, "PRIORITY", tag.name)
}

func TestValueAsString(t *testing.T) {
	t.Setenv("STR_TEST", "myvalue")
	tag := &Tag{name: "STR_TEST", defaultValue: "fallback"}
	assert.Equal(t, "myvalue", tag.ValueAsString())
}

func TestValueAsStringDefault(t *testing.T) {
	tag := &Tag{name: "STR_TEST_NONEXISTENT", defaultValue: "fallback"}
	assert.Equal(t, "fallback", tag.ValueAsString())
}

func TestValueAsInt(t *testing.T) {
	t.Setenv("INT_TEST", "42")
	tag := &Tag{name: "INT_TEST", defaultValue: ""}
	val, err := tag.ValueAsInt()
	require.NoError(t, err)
	assert.Equal(t, int64(42), val)
}

func TestValueAsIntDefault(t *testing.T) {
	tag := &Tag{name: "INT_TEST_NONEXISTENT", defaultValue: "10"}
	val, err := tag.ValueAsInt()
	require.NoError(t, err)
	assert.Equal(t, int64(10), val)
}

func TestValueAsIntInvalidDefault(t *testing.T) {
	tag := &Tag{name: "INT_TEST_NONEXISTENT", defaultValue: "notanint"}
	_, err := tag.ValueAsInt()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not parse default value")
}

func TestValueAsBool(t *testing.T) {
	t.Setenv("BOOL_TEST", "true")
	tag := &Tag{name: "BOOL_TEST", defaultValue: ""}
	val, err := tag.ValueAsBool()
	require.NoError(t, err)
	assert.True(t, val)
}

func TestValueAsBoolDefault(t *testing.T) {
	tag := &Tag{name: "BOOL_TEST_NONEXISTENT", defaultValue: "true"}
	val, err := tag.ValueAsBool()
	require.NoError(t, err)
	assert.True(t, val)
}

func TestValueAsBoolInvalidDefault(t *testing.T) {
	tag := &Tag{name: "BOOL_TEST_NONEXISTENT", defaultValue: "notabool"}
	_, err := tag.ValueAsBool()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "could not parse default value")
}

func TestValueAsMap(t *testing.T) {
	t.Setenv("MAP_one", "1")
	t.Setenv("MAP_two", "2")
	tag := &Tag{prefix: "MAP", cutoff: true}
	result := tag.ValueAsMap()
	assert.Len(t, result, 2)
	assert.Equal(t, "1", result["one"])
	assert.Equal(t, "2", result["two"])
}

func TestValueAsMapNoCutoff(t *testing.T) {
	t.Setenv("NOCUT_one", "a")
	tag := &Tag{prefix: "NOCUT", cutoff: false}
	result := tag.ValueAsMap()
	assert.Len(t, result, 1)
	_, hasFullKey := result["NOCUT_one"]
	assert.True(t, hasFullKey)
}

type badCutoffStruct struct {
	Bad string `env:"prefix=X,cutoff=invalid"`
}

func TestParseCutoffInvalidValue(t *testing.T) {
	tag := NewTagMeta()
	f := reflect.TypeOf(badCutoffStruct{})
	field, _ := f.FieldByName("Bad")
	found, err := tag.Parse(field)
	assert.True(t, found)
	require.Error(t, err)
}
