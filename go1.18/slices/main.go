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
	log.Println(v1, v2)
	return v1 == v2
}

func main() {
	TestEqualAndEqualFunc()
	TestClone()
	TestCompactAndCompactFunc()
}

func TestEqualAndEqualFunc() {
	log.Println("TestEqualAndEqualFunc")
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
			log.Printf("EqualFunc(%v, %v, equal[int]) = %t, want %t", item.s1, item.s2, got, item.want)
		}
		gotInt := slices.Compare(item.s1, item.s2)
		want := 0
		if !cmp.Equal(want, gotInt) {
			log.Printf("Compare(%v, %v) = %d", item.s1, item.s2, gotInt)
		}
	}
}

func TestClone() {
	log.Println("slice.Clone")
	s1 := []int{1, 2, 3}
	s2 := slices.Clone(s1)
	if ok := slices.Equal(s1, s2); !ok {
		log.Printf("Clone(%v) = %v, want %v", s1, s2, s1)
	}
	s1[0] = 4
	want := []int{1, 2, 3}
	if ok := slices.Equal(s2, want); !ok {
		log.Printf("Clone(%v) changed unexpectedly to %v", want, s2)
	}
	log.Println([]int(nil))
	if got := slices.Clone([]int(nil)); got != nil {
		log.Printf("Clone(nil) = %#v, want nil", got)
	}
	log.Println(s1[:0])
	log.Println(len(s1[:0]))
	if got := slices.Clone(s1[:0]); got == nil || len(got) != 0 {
		log.Printf("Clone(%v) = %#v, want %#v", s1[:0], got, s1[:0])
	}

}

func TestCompactAndCompactFunc() {
	s := []int{1, 2, 2, 3, 3, 4, 5}
	clone := slices.Clone(s)
	s[1] = 7
	log.Println(clone)
	clone = append(clone, 6)
	log.Println("<", clone)
	s2 := slices.Compact(clone)
	clone = append(clone, 7)
	log.Println(s)
	log.Println(s2)
	log.Println(clone)
	clone = append(clone, []int{7, 8}...)
	log.Println(clone)
	s3 := slices.CompactFunc(clone, equal[int])
	log.Println(s3)
	var ss []int
	ss = nil
	cloneSlice := slices.Clone(ss)
	log.Println(cloneSlice)
	ss1 := slices.Compact(cloneSlice)
	log.Println(ss1)
	ss2 := slices.CompactFunc(cloneSlice, equal[int])
	log.Println(ss2)
	if ss2 != nil {
		log.Println("nil")
	}
	if ss2[:0] == nil {
		log.Println("nil")
	}
}
