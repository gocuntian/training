package stack_test

import (
	"github.com/gocuntian/training/go1.18/stack"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestStack_Push(t *testing.T) {
	t.Parallel()
	s := stack.Stack[int]{}
	s.Push(0)
	want := 1
	got := s.Len()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStack_PushPop(t *testing.T) {
	t.Parallel()
	s := stack.Stack[string]{}
	s.Push("a", "b")
	if s.Len() != 2 {
		t.Fatal("Push didn't add all values to stack")
	}
	s.Pop()
	want := "a"
	got, ok := s.Pop()
	if !ok {
		t.Fatal("Pop returned not ok on non-empty stack")
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestStackEmpty_Pop(t *testing.T) {
	t.Parallel()
	s := stack.Stack[int]{}
	_, ok := s.Pop()
	if ok {
		t.Fatal("Pop returned ok on empty stack")
	}
}
