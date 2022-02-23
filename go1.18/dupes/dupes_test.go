package dupes_test

import (
	"github.com/gocuntian/training/go1.18/dupes"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDupes(t *testing.T) {
	t.Parallel()
	s := []int{1, 2, 3, 4, 5, 1}
	want := true
	got := dupes.Dupes(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
