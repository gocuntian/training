package funcmap_test

import (
	"github.com/gocuntian/training/go1.18/funcmap"
	"github.com/google/go-cmp/cmp"
	"testing"
	"unicode"
)

func TestFuncMapIntInt_Apply(t *testing.T) {
	t.Parallel()
	fm := funcmap.FuncMap[int, int]{
		"double": func(i int) int {
			return i * 2
		},
		"addOne": func(i int) int {
			return i + 1
		},
	}
	want := 4
	got := fm.Apply("double", 2)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFuncMapRuneBool_Apply(t *testing.T) {
	t.Parallel()
	fm := funcmap.FuncMap[rune, bool]{
		"upper": unicode.IsUpper,
		"lower": unicode.IsLower,
	}
	if ok := fm.Apply("upper", 'A'); !ok {
		t.Fatal("upper('A'): want true, got false")
	}

}
