package merge_test

import (
	"github.com/gocuntian/training/go1.18/merge"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMergeInt(t *testing.T) {
	t.Parallel()
	input := []map[int]bool{
		{
			1: false,
			2: true,
			3: false,
		},
		{
			2: false,
			3: true,
			4: true,
			5: false,
			6: true,
		},
	}

	want := map[int]bool{
		1: false,
		2: false,
		3: true,
		4: true,
		5: false,
		6: true,
	}
	got := merge.Merge(input...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestMergeString(t *testing.T) {
	t.Parallel()
	input := []map[string]any{
		{
			"a": nil,
		},
		{
			"b": "hello world",
			"c": 0,
			"a": 6 + 2i,
		},
	}
	want := map[string]any{
		"a": 6 + 2i,
		"b": "hello world",
		"c": 0,
	}
	got := merge.Merge(input...)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
