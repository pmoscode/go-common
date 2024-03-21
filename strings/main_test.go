package strings

import (
	"fmt"
	"testing"
)

func TestSimplePrint(t *testing.T) {
	testObj := struct {
		Key1 string
		Key2 string
		Key3 string
	}{
		Key1: "val1",
		Key2: "val2",
		Key3: "val3",
	} //"{\"key1\":\"val1\",\"key2\":\"val2\",\"key3\":\"val3\"}"

	result := PrettyPrint(testObj)

	lines := countRune(result, '\n') + 1

	if lines != 5 {
		fmt.Println(result)
		t.Fatal("Line count mismatch: ", lines, " should be ", 5)
	}
}
func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
