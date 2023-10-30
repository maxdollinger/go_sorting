package main

import (
	"fmt"
	"log"

	"github.com/maxdolliger/timesort/data"
)

func main() {
	out, err := NewOutput("results.txt")
	defer out.Close()
	if err != nil {
		log.Fatal(err)
	}

	generator := data.NewDataGenerator()
	s := generator.NearSorted(9)

	fmt.Println(s)
}
