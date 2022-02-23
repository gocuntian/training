package compose_test

import (
	"github.com/gocuntian/training/go1.18/compose"
	"testing"
)

func isOdd(p int) bool {
	return p%2 != 0
}

func next(p int) int {
	return p + 1
}

func TestComposeInt(t *testing.T) {
	t.Parallel()
	odd := compose.Compose(isOdd, next, 1)
	if odd {
		t.Fatal("isOdd(next(1)): want false, got true")
	}
}
