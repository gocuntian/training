package contains_test

import (
	"fmt"
	"github.com/gocuntian/training/go1.18/contains"
	"testing"
	"unicode"
)

func positive(p int) bool {
	fmt.Println(p)
	return p > 0
}

func TestContainsFuncTrue(t *testing.T) {
	t.Parallel()
	input := []int{-2, 0, 1, -1, 5}
	if ok := contains.ContainsFunc(input, positive); !ok {
		t.Fatalf("%v: want true for 'contains positive', got false", input)
	}
}

func TestContainsFuncFalse(t *testing.T) {
	t.Parallel()
	input := []rune("hello")
	if ok := contains.ContainsFunc(input, unicode.IsUpper); ok {
		t.Fatalf("%q: want false for 'contains uppercase', got true", input)
	}
}
