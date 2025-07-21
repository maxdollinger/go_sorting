package sorting_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/maxdolliger/go_sorting/data"
	"github.com/maxdolliger/go_sorting/sorting"
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
