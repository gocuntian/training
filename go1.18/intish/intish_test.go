package intish_test

import (
	"github.com/gocuntian/training/go1.18/intish"
	"github.com/google/go-cmp/cmp"
	"testing"
)

type MyInt int

func TestIsPositiveTrueFor1(t *testing.T) {
	t.Parallel()
	input := MyInt(1)
	want := true
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositiveFalseForMinus1(t *testing.T) {
	t.Parallel()
	input := MyInt(-1)
	want := false
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestIsPositiveFalseForZero(t *testing.T) {
	t.Parallel()
	input := MyInt(0)
	want := false
	got := intish.IsPositive(input)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
