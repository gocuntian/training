package main

import (
	"github.com/google/go-cmp/cmp"
	"golang.org/x/exp/slices"
	"log"
)

type intItem struct {
	s1   []int
	s2   []int
	want bool
}

func equal[T comparable](v1, v2 T) bool {
	return v1 == v2
}

func main() {
	var items = []intItem{
		{
			[]int{1},
			nil,
			false,
		},
		{
			[]int{},
			nil,
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			true,
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3, 4},
			false,
		},
	}

	for _, item := range items {
		got := slices.Equal(item.s1, item.s2)
		if !cmp.Equal(got, item.want) {
			log.Println("got=", got, "want=", item.want)
		}
		got = slices.EqualFunc(item.s1, item.s2, equal[int])
		if !cmp.Equal(got, item.want) {
			log.Fatalf("EqualFunc(%v, %v, equal[int]) = %t, want %t", item.s1, item.s2, got, item.want)
		}
		gotInt := slices.Compare(item.s1, item.s2)
		want := 0
		if !cmp.Equal(want, gotInt) {
			log.Fatalf("Compare(%v, %v) = %d", item.s1, item.s2, gotInt)
		}
	}

}
