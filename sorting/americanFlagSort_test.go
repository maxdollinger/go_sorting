package sorting_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func TestAmericanFlagSort(t *testing.T) {
	s := data.Random(10)

	fmt.Println(s)

	sorting.AmericanFlagSort(s)

	fmt.Println(s)

	isSorted := sort.SliceIsSorted(s, func(i, j int) bool {
		return s[i].SortValue() < s[j].SortValue()
	})

	if !isSorted {
		t.Fatal("not sorted")
	}
}
