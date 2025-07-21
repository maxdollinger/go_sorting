package main

import (
	"log"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

type (
	SortingFn[T sorting.Sortable]     func(s []T)
	DataGenerator[T sorting.Sortable] func(n int) []T
)

type Executor[T sorting.Sortable] struct {
	runs       int
	inputSizes []int
	sortingFn  SortingFn[T]
	generator  DataGenerator[T]
}

func NewExecutor[T sorting.Sortable](dataGenerator DataGenerator[T], runs int, sizes []int) *Executor[T] {
	return &Executor[T]{
		runs:       runs,
		inputSizes: sizes,
		generator:  dataGenerator,
	}
}

func (e *Executor[T]) Run(sortingFn SortingFn[T]) []*data.Evaluation {
	results := make([]*data.Evaluation, len(e.inputSizes))

	distribution := getFunctionName(e.generator)
	sortingMethod := getFunctionName(sortingFn)

	log.Printf("Starting runs for %s with %s data ...", sortingMethod, distribution)

	for i, size := range e.inputSizes {
		results[i] = data.NewEvaluation(size, distribution, sortingMethod)

		for j := 0; j < e.runs; j++ {

			dataToSort := e.generator(size)

			start := time.Now()
			sortingFn(dataToSort)
			execTime := time.Since(start)
			results[i].AddMemorySnapshot()
			results[i].AddExecTime(execTime)

			if !sort.SliceIsSorted(dataToSort, func(i, j int) bool {
				return dataToSort[i].SortValue() < dataToSort[j].SortValue()
			}) {
				log.Fatalf("%s,%s,%v: not sorted", results[i].Method, results[i].Distribution, size)
			}

			runtime.GC()
		}
	}

	log.Printf("DONE")

	return results
}

func (e *Executor[T]) GetGeneratorName() string {
	return getFunctionName(e.generator)
}

func getFunctionName(fn interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	if i := strings.LastIndex(name, "/"); i != -1 {
		name = name[i+1:]
	}

	if i := strings.Index(name, "."); i != -1 {
		name = name[i+1:]
	}

	if i := strings.Index(name, "["); i != -1 {
		name = name[:i]
	}

	return name
}
