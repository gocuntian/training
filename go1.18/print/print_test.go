package print_test

import (
	"bytes"
	"github.com/gocuntian/training/go1.18/print"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPrintAnything(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	print.PrintAnything[string](buf, "hello world!")
	want := "hello world!\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
