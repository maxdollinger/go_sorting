package sorting_test

import (
	"sort"
	"testing"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func TestRadixInplace(t *testing.T) {

	s := data.Random(10000)

	sorting.RadixSortInplace(s)

	isSorted := sort.SliceIsSorted(s, func(i, j int) bool {
		return s[i].Unix() < s[j].Unix()
	})

	if !isSorted {
		t.Fatal("not sorted")
	}
}
