package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterSimple(t *testing.T) {
	a, b, c, d := "3", "2", "45", "30"
	testItems := []*string{&a, &b, &c, &d}

	filterObj := NewFilter(testItems)
	filtered := filterObj.Filter(func(val string) bool {
		return val == "2"
	})

	assert.Len(t, *filtered, 1)
	assert.Equal(t, "2", (*filtered)[0])
}

type Complex struct {
	Name string
	Calc int
	Loc  int
}

func TestFilterComplex(t *testing.T) {
	testItems := []*Complex{
		{Name: "Test 1", Calc: 5, Loc: 15},
		{Name: "Test 2", Calc: 2, Loc: 30},
		{Name: "Test 3", Calc: 12, Loc: 26},
		{Name: "Test 4", Calc: 20, Loc: 30},
	}

	filterObj := NewFilter(testItems)
	filtered := filterObj.Filter(func(val Complex) bool {
		return val.Loc == 30
	})

	assert.Len(t, *filtered, 2)
	assert.Equal(t, 30, (*filtered)[0].Loc)
}

func TestMultipleFilterComplex(t *testing.T) {
	testItems := []*Complex{
		{Name: "Test 1", Calc: 5, Loc: 15},
		{Name: "Test 2", Calc: 2, Loc: 30},
		{Name: "Test 3", Calc: 12, Loc: 26},
		{Name: "Test 4", Calc: 2, Loc: 30},
		{Name: "Test 5", Calc: 20, Loc: 30},
	}

	filterObj := NewFilter(testItems)
	filtered := filterObj.Filter(
		func(val Complex) bool { return val.Loc == 30 },
		func(val Complex) bool { return val.Calc == 2 },
		func(val Complex) bool { return val.Name == "Test 4" },
	)

	assert.Len(t, *filtered, 1)
	assert.Equal(t, "Test 4", (*filtered)[0].Name)
}
