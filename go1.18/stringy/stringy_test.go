package stringy_test

import (
	"bytes"
	"github.com/gocuntian/training/go1.18/stringy"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type greeting struct{}

func (greeting) String() string {
	return "Hello"
}

func TestStringify(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	stringy.Stringify[greeting](buf, greeting{})
	want := "Hello\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
