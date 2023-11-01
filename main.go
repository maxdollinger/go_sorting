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

	exec := NewExecutor(100, []int{1000, 10_000, 100_000})

	exec.UseSortingFn(STANDARD_SORT)
	out.Write(data.NewFormater(exec.UseGenerator(RANDOM).Run()))
	out.Write(data.NewFormater(exec.UseGenerator(NEAR_SORTED).Run()))
	out.Write(data.NewFormater(exec.UseGenerator(SORTED).Run()))
	out.Write(data.NewFormater(exec.UseGenerator(REVERSED).Run()))
	out.Write(data.NewFormater(exec.UseGenerator(DBL_SORTED).Run()))
}
