package environment

import (
	"fmt"
	"testing"
)

func TestStringEnv(t *testing.T) {
	t.Setenv("keyStr", "something")
	t.Setenv("keyInt", "12")
	t.Setenv("keyBool1", "true")
	t.Setenv("keyBool2", "false")
	t.Setenv("keyFloat32", "31.5")
	t.Setenv("keyFloat64", "34.5")

	strTest := GetEnv("keyStr", "nothing")
	intTest := GetEnvInt("keyInt", 0)
	bool1Test := GetEnvBool("keyBool1", false)
	bool2Test := GetEnvBool("keyBool2", true)
	float32Test := GetEnvFloat32("keyFloat32", 11.4)
	float64Test := GetEnvFloat64("keyFloat64", 12.4)

	if strTest != "something" {
		t.Fatal("Expected string not set: ", "something")
	}
	if intTest != 12 {
		t.Fatal("Expected int not set: ", 12)
	}
	if bool1Test != true {
		t.Fatal("Expected bool not set: ", true)
	}
	if bool2Test != false {
		t.Fatal("Expected bool not set: ", false)
	}
	if float32Test != 31.5 {
		t.Fatal("Expected float32 not set: ", 31.5)
	}
	if float64Test != 34.5 {
		t.Fatal("Expected float64 not set: ", 34.5)
	}
}

func TestStringDefaultEnv(t *testing.T) {
	t.Setenv("keyInt", "notAnInt")
	t.Setenv("keyBool1", "3")
	t.Setenv("keyBool2", "blah")
	t.Setenv("keyFloat32", "true")
	t.Setenv("keyFloat64", "23fb")
	t.Setenv("keyFloat64int", "34")

	strTest := GetEnv("keyStr", "nothing")
	intTest := GetEnvInt("keyInt", 50)
	bool1Test := GetEnvBool("keyBool1", false)
	bool2Test := GetEnvBool("keyBool2", true)
	float32Test := GetEnvFloat32("keyFloat32", 10.5)
	float64Test := GetEnvFloat64("keyFloat64", 20.5)
	float64intTest := GetEnvFloat64("keyFloat64int", 30.5)

	if strTest != "nothing" {
		t.Fatal("Expected string not set: ", "nothing")
	}
	if intTest != 50 {
		t.Fatal("Expected int not set: ", 50)
	}
	if bool1Test != false {
		t.Fatal("Expected bool not set: ", false)
	}
	if bool2Test != true {
		t.Fatal("Expected bool not set: ", true)
	}
	if float32Test != 10.5 {
		t.Fatal("Expected float32 not set: ", 10.5)
	}
	if float64Test != 20.5 {
		t.Fatal("Expected float64 not set: ", 20.5)
	}
	if float64intTest != 34 {
		t.Fatal("Expected float64 not set: ", 34)
	}
}

func TestEnvMap(t *testing.T) {
	t.Setenv("TEST_one", "one")
	t.Setenv("TEST_two", "two")
	t.Setenv("TEST_three", "3")
	t.Setenv("TESTING_invalid", "4")
	t.Setenv("four", "4")

	envMap := GetEnvMap("TEST", true)

	if len(envMap) != 3 {
		t.Fatal("Expected set len == 3 got: ", len(envMap))
	}

	fmt.Println(envMap)
}
