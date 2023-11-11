package benchmarks_test

import (
	"testing"

	"github.com/maxdolliger/timesort/data"
	"github.com/maxdolliger/timesort/sorting"
)

func BenchmarkSort1k(b *testing.B) {
	n := 1000

	// b.Run("fn=standart", stdSort(n))
	b.Run("fn=radix", radixSort(n))

}

func BenchmarkSort10k(b *testing.B) {
	n := 10_000

	// b.Run("fn=standart", stdSort(n))
	b.Run("fn=radix", radixSort(n))

}

func BenchmarkSort100k(b *testing.B) {
	n := 100_000

	// b.Run("fn=standart", stdSort(n))
	b.Run("fn=radix", radixSort(n))
}

func stdSort(n int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			s := data.Random(n)
			b.StartTimer()
			sorting.StandartSort(s)
		}
	}
}

func radixSort(n int) func(*testing.B) {
	return func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			b.StopTimer()
			s := data.Random(n)
			b.StartTimer()
			sorting.RadixSort(s)
		}
	}
}
