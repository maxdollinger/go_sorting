package sorting_test

import (
	"sort"
	"testing"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func TestRadixSort(t *testing.T) {

	s := data.Random(10000)
	sorting.RadixSort(s)

	if !sort.SliceIsSorted(s, func(i, j int) bool {
		return s[i].Unix() < s[j].Unix()
	}) {
		t.Log(s)
		t.Fatalf("not sorted")
	}
}
