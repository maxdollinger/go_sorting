package sorting_test

import (
	"sort"
	"testing"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func TestAmericanFlagSort(t *testing.T) {

	s := data.Random(100000)

	sorting.AmericanFlagSort(s)

	isSorted := sort.SliceIsSorted(s, func(i, j int) bool {
		return s[i].Unix() < s[j].Unix()
	})

	if !isSorted {
		t.Fatal("not sorted")
	}
}

func TestAmericanFlagSortP(t *testing.T) {

	s := data.Random(100000)

	sorting.AmericanFlagSortP(s)

	isSorted := sort.SliceIsSorted(s, func(i, j int) bool {
		return s[i].Unix() < s[j].Unix()
	})

	if !isSorted {
		t.Fatal("not sorted")
	}
}
