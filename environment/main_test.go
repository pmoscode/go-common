package environment

import (
	"testing"
)

func TestStringEnv(t *testing.T) {
	t.Setenv("keyStr", "something")
	t.Setenv("keyInt", "12")
	t.Setenv("keyBool1", "true")
	t.Setenv("keyBool2", "false")

	strTest := GetEnv("keyStr", "nothing")
	intTest := GetEnvInt("keyInt", 0)
	bool1Test := GetEnvBool("keyBool1", false)
	bool2Test := GetEnvBool("keyBool2", true)

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
}
func TestStringDefaultEnv(t *testing.T) {
	t.Setenv("keyInt", "notAnInt")
	t.Setenv("keyBool1", "3")
	t.Setenv("keyBool2", "blah")

	strTest := GetEnv("keyStr", "nothing")
	intTest := GetEnvInt("keyInt", 50)
	bool1Test := GetEnvBool("keyBool1", false)
	bool2Test := GetEnvBool("keyBool2", true)

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
}
