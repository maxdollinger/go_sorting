package benchmarks_test

import (
	"testing"
	"time"
)

var timeSlice = []time.Time{time.Now(), time.Now().Add(time.Hour)}

func lessComp(a, b int) bool {
	return timeSlice[a].Compare(timeSlice[b]) < 0
}

func lessBefore(a, b int) bool {
	return timeSlice[a].Before(timeSlice[b])
}

func lessSub(a, b int) bool {
	return timeSlice[a].Sub(timeSlice[b]) < 0
}

func lessUnixComp(a, b int) bool {
	return timeSlice[a].Unix() < timeSlice[b].Unix()
}

func BenchmarkLessComp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lessComp(1, 0)
	}
}

func BenchmarkLessBefore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lessBefore(1, 0)
	}
}

func BenchmarkLessSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lessSub(1, 0)
	}
}

func BenchmarkLessUnixComp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lessUnixComp(1, 0)
	}
}
