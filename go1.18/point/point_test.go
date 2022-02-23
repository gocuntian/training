package point_test

import (
	"github.com/gocuntian/training/go1.18/point"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPointXY(t *testing.T) {
	t.Parallel()
	p := point.Point{
		X: 1,
		Y: 2,
	}
	want := 1
	got := point.GetX[point.Point](p)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
	want = 2
	got = point.GetY[point.Point](p)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
