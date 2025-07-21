package sorting_test

import (
	"sort"
	"testing"

	"github.com/maxdolliger/go_sorting/data"
	"github.com/maxdolliger/go_sorting/sorting"
)

func TestRadixSort(t *testing.T) {
	s := data.Random(10000)
	sorting.RadixSort(s)

	if !sort.SliceIsSorted(s, func(i, j int) bool {
		return s[i].SortValue() < s[j].SortValue()
	}) {
		t.Log(s)
		t.Fatalf("not sorted")
	}
}
