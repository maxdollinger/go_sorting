package benchmarks_test

import (
	"testing"
	"time"
)

var timeSlice = []time.Time{time.Now().Add(time.Hour), time.Now()}

func BenchmarkLessFns(b *testing.B) {

	b.Run("fn=unix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lessUnix(0, 1)
		}
	})

	b.Run("fn=compare", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lessCompare(0, 1)
		}
	})

	b.Run("fn=before", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lessBefore(0, 1)
		}
	})

	b.Run("fn=sub", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			lessSub(0, 1)
		}
	})

}

func lessCompare(a, b int) bool {
	return timeSlice[a].Compare(timeSlice[b]) < 0
}

func lessBefore(a, b int) bool {
	return timeSlice[a].Before(timeSlice[b])
}

func lessSub(a, b int) bool {
	return timeSlice[a].Sub(timeSlice[b]) < 0
}

func lessUnix(a, b int) bool {
	return timeSlice[a].Unix() < timeSlice[b].Unix()
}
