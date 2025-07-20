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
	sortingFn[T sorting.Sortable]       func(s []T)
	randomGenerator[T sorting.Sortable] func(n int) []T
)

type Executor[T sorting.Sortable] struct {
	runs       int
	inputSizes []int
	method     sortingFn[T]
	generator  randomGenerator[T]
}

func NewExecutor[T sorting.Sortable](sortingFn sortingFn[T], runs int, sizes []int) *Executor[T] {
	return &Executor[T]{
		runs:       runs,
		inputSizes: sizes,
		method:     sortingFn,
	}
}

func (e *Executor[T]) Run(generateFn randomGenerator[T]) []*data.Evaluation {
	results := make([]*data.Evaluation, len(e.inputSizes))

	distribution := getFunctionName(generateFn)
	sortingMethod := getFunctionName(e.method)

	log.Printf("Distribution: %s", distribution)
	log.Printf("Sorting method: %s", sortingMethod)

	for i, size := range e.inputSizes {
		results[i] = data.NewEvaluation(size, distribution, sortingMethod)

		for j := 0; j < e.runs; j++ {

			dataToSort := generateFn(size)

			start := time.Now()
			e.method(dataToSort)
			execTime := time.Since(start)
			results[i].AddMemorySnapshot()
			results[i].AddExecTime(execTime)

			if !sort.SliceIsSorted(dataToSort, func(i, j int) bool {
				return dataToSort[i].SortValue() < dataToSort[j].SortValue()
			}) {
				log.Fatalf("%s,%s,%v: not sorted", results[i].Method, results[i].Distribution, size)
			}

			log.Printf("DONE -> size: %d run: %d t: %s \n", size, j+1, execTime)

			runtime.GC()
		}
	}

	return results
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
