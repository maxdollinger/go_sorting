package main

import (
	"log"

	"github.com/maxdolliger/timesort/data"
)

func main() {
	out, err := NewOutput("results.txt")
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}

	executor := NewExecutor(100, []int{1000, 10_000, 100_000, 1_000_000})

	testSortingFn(RADIX_SORT_INPLACE, executor, out)
	testSortingFn(STANDARD_SORT, executor, out)
}

func testSortingFn(fn string, e *Executor, o *Output) {
	e.UseSortingFn(fn)
	o.Write(data.NewFormater(e.UseGenerator(RANDOM).Run()))
	// o.Write(data.NewFormater(e.UseGenerator(NEAR_SORTED).Run()))
	o.Write(data.NewFormater(e.UseGenerator(SORTED).Run()))
	// o.Write(data.NewFormater(e.UseGenerator(REVERSED).Run()))
	// o.Write(data.NewFormater(e.UseGenerator(DBL_SORTED).Run()))
}
