package main

import (
	"log"
	"runtime"

	"github.com/maxdolliger/timesort/data"
)

func main() {

	out, err := NewOutput("results_std.txt")
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}

	executor := NewExecutor(500, []int{1000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000})

	testSortingFn(STANDARD_SORT, executor, out)
	// testSortingFn(RADIX_SORT_INPLACE, executor, out)
}

func testSortingFn(fn string, e *Executor, o *Output) {
	e.UseSortingFn(fn)
	o.Write(data.NewFormater(e.UseGenerator(RANDOM).Run()))
	o.Write(data.NewFormater(e.UseGenerator(NEAR_SORTED).Run()))
	o.Write(data.NewFormater(e.UseGenerator(SORTED).Run()))
	o.Write(data.NewFormater(e.UseGenerator(REVERSED).Run()))
	o.Write(data.NewFormater(e.UseGenerator(DBL_SORTED).Run()))
	runtime.GC()
}
