package greater_test

import (
	"fmt"
	"github.com/gocuntian/training/go1.18/greater"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type MyInt int

func (m MyInt) Greater(v MyInt) bool {
	fmt.Println(m)
	fmt.Println(v)
	return m > v
}

func TestIsGreaterTrue(t *testing.T) {
	t.Parallel()
	want := true
	got := greater.IsGreater(MyInt(2), MyInt(1))
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsGreaterFalse(t *testing.T) {
	t.Parallel()
	want := false
	got := greater.IsGreater(MyInt(1), MyInt(2))
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
