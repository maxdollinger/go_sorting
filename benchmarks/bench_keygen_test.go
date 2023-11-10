package benchmarks_test

import (
	"testing"
	"time"
)

var t time.Time = time.Now()

func BenchmarkGetKey(b *testing.B) {

	b.Run("m=Unix", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t.Unix()
		}
	})

	b.Run("m=Format", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t.Format("2006-01-02T15:03:04")
		}
	})

	b.Run("m=Year", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t.Year()
		}
	})

	b.Run("m=YearDay", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t.YearDay()
		}
	})

	b.Run("m=Hour", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t.YearDay()
		}
	})

	b.Run("m=Minute", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t.Minute()
		}
	})

	b.Run("m=Second", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			t.Second()
		}
	})

}
