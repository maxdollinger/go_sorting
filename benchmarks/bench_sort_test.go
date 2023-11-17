package benchmarks_test

import (
	"testing"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func BenchmarkSort10k(b *testing.B) {
	n := 10_000
	sortN(n, b)
}

func BenchmarkSort100k(b *testing.B) {
	n := 100_000
	sortN(n, b)
}

func sortN(n int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := data.Random(n)
		b.StartTimer()
		sorting.AmericanFlagSort(s)
	}
}
