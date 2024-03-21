package filter

import "testing"

func TestFilterSimple(t *testing.T) {
	a, b, c, d := "3", "2", "45", "30"
	testVal := "2"

	testItems := []*string{&a, &b, &c, &d}

	filterObj := NewFilter(testItems)

	filterFunc := func(val string) bool {
		return val == testVal
	}

	filtered := filterObj.Filter(filterFunc)

	if len(*filtered) != 1 {
		t.Fatal("Wrong filter count: ", len(*filtered), " should be ", 1)
	}

	if (*filtered)[0] != testVal {
		t.Fatal("filtered value not ", testVal, " -> ", (*filtered)[0])
	}
}

type Complex struct {
	Name string
	Calc int
	Loc  int
}

func TestFilterComplex(t *testing.T) {
	testItems := []*Complex{
		{
			Name: "Test 1",
			Calc: 5,
			Loc:  15,
		},
		{
			Name: "Test 2",
			Calc: 2,
			Loc:  30,
		},
		{
			Name: "Test 3",
			Calc: 12,
			Loc:  26,
		},
		{
			Name: "Test 4",
			Calc: 20,
			Loc:  30,
		},
	}
	testVal := 30

	filterObj := NewFilter(testItems)

	filterFunc := func(val Complex) bool {
		return val.Loc == testVal
	}

	filtered := filterObj.Filter(filterFunc)

	if len(*filtered) != 2 {
		t.Fatal("Wrong filter count: ", len(*filtered), " should be ", 2)
	}

	if (*filtered)[0].Loc != testVal {
		t.Fatal("filtered value not ", testVal, " -> ", (*filtered)[0].Loc)
	}
}

func TestMultipleFilterComplex(t *testing.T) {
	testItems := []*Complex{
		{
			Name: "Test 1",
			Calc: 5,
			Loc:  15,
		},
		{
			Name: "Test 2",
			Calc: 2,
			Loc:  30,
		},
		{
			Name: "Test 3",
			Calc: 12,
			Loc:  26,
		},
		{
			Name: "Test 4",
			Calc: 2,
			Loc:  30,
		},
		{
			Name: "Test 5",
			Calc: 20,
			Loc:  30,
		},
	}
	testLoc := 30
	testName := "Test 4"
	testCalc := 2

	filterObj := NewFilter(testItems)

	filterLocFunc := func(val Complex) bool {
		return val.Loc == testLoc
	}
	filterNameFunc := func(val Complex) bool {
		return val.Name == testName
	}
	filterCalcFunc := func(val Complex) bool {
		return val.Calc == testCalc
	}

	filtered := filterObj.Filter(filterLocFunc, filterCalcFunc, filterNameFunc)

	if len(*filtered) != 1 {
		t.Fatal("Wrong filter count: ", len(*filtered), " should be ", 1)
	}

	if (*filtered)[0].Name != testName {
		t.Fatal("filtered value not ", testName, " -> ", (*filtered)[0].Name)
	}
}
