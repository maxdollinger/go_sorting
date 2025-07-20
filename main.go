package main

import (
	"fmt"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func main() {
	sortingSliceSizes := []int{10, 100, 1000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000}
	runsPerSize := 5
	executor := NewExecutor[data.SortableNumber](sorting.AmericanFlagSort, runsPerSize, sortingSliceSizes)

	fmt.Println("Sorting numbers...")
	results := executor.Run(data.Random)
	fmt.Printf("%s\n", data.NewFormater(results))
}
