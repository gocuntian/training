package sequence_test

import (
	"github.com/gocuntian/training/go1.18/sequence"
	"testing"
)

func TestEmptyTrue(t *testing.T) {
	t.Parallel()
	s := sequence.Sequence[int]{}
	if ok := s.Empty(); !ok {
		t.Fatal("false for empty sequence")
	}
}

func TestEmptyFalse(t *testing.T) {
	t.Parallel()
	s := sequence.Sequence[string]{"a", "b", "c"}
	if ok := s.Empty(); ok {
		t.Fatal("true for non-empty sequence")
	}
}
