package main

import (
	"fmt"
	"os"
	"time"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func main() {
	sortingSliceSizes := []int{10, 100, 1000, 10_000, 100_000, 1_000_000}
	runsPerSize := 1

	results := make([]*data.Evaluation, 0, len(sortingSliceSizes)*5)
	exec := NewExecutor(data.Random, runsPerSize, sortingSliceSizes)

	results = append(results, exec.Run(sorting.InsertionSort)...)
	results = append(results, exec.Run(sorting.RadixSort)...)
	results = append(results, exec.Run(sorting.AmericanFlagSort)...)
	results = append(results, exec.Run(sorting.AmericanFlagSortParallel)...)
	results = append(results, exec.Run(sorting.StandartSort)...)

	formater := data.NewFormater(results)
	fmt.Println(formater.TableString())

	output := fmt.Sprintf("%s\n", formater.String())
	err := persitsResults(output, exec.GetGeneratorName())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}

func persitsResults(output string, generatorName string) error {
	fileName := fmt.Sprintf("%s_%s.txt", generatorName, time.Now().Format("2006-01-02_15-04-05"))
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 6644)
	if err != nil {
		return fmt.Errorf("Error opening file: %v\n", err)
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Error closing file: %v\n", closeErr)
			os.Exit(1)
		}
	}()

	_, err = file.WriteString(output)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v\n", err)
	}

	return nil
}
