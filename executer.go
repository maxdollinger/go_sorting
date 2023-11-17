package main

import (
	"log"
	"runtime"
	"sort"
	"time"

	"github.com/maxdolliger/timesort/data"
)

const (
	RANDOM      string = "random"
	NEAR_SORTED string = "near_sorted"
	SORTED      string = "sorted"
	REVERSED    string = "reversed"
	DBL_SORTED  string = "sorted_doubled"
)

const (
	STANDARD_SORT      string = "standard"
	RADIX_SORT         string = "radix"
	RADIX_SORT_INPLACE string = "radix_inplace"
)

type Executor struct {
	runs         int
	inputSizes   []int
	distribution string
	method       string
}

func NewExecutor(runs int, sizes []int) *Executor {
	return &Executor{
		runs:       runs,
		inputSizes: sizes,
	}
}

func (e *Executor) UseGenerator(gen string) *Executor {
	e.distribution = gen
	return e
}

func (e *Executor) generateSlice(n int) []time.Time {

	switch e.distribution {
	case RANDOM:
		return data.Random(n)
	case NEAR_SORTED:
		return data.NearSorted(n)
	case SORTED:
		return data.Sorted(n)
	case REVERSED:
		return data.Reversed(n)
	case DBL_SORTED:
		return data.SortedDoubled(n)
	}

	return []time.Time{}
}

func (e *Executor) UseSortingFn(fn string) *Executor {
	e.method = fn
	return e
}

func (e *Executor) sort(s []time.Time) []time.Time {

	return s
}

func (e *Executor) Run() []*data.Evaluation {

	results := make([]*data.Evaluation, len(e.inputSizes))

	for i, size := range e.inputSizes {
		results[i] = data.NewEvaluation(size, e.distribution, e.method)

		for j := 0; j < e.runs; j++ {

			timeSlice := e.generateSlice(size)

			start := time.Now()
			timeSlice = e.sort(timeSlice)
			t := time.Since(start)

			results[i].AddMemorySnapshot()
			results[i].AddExecTime(t)

			if !sort.SliceIsSorted(timeSlice, func(i, j int) bool {
				return timeSlice[i].Unix() < timeSlice[j].Unix()
			}) {
				log.Fatalf("%s,%s,%v: not sorted", results[i].Method, results[i].Distribution, size)
			}

			runtime.GC()
		}
	}

	return results
}
