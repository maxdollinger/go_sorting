package data

import (
	"math"
	"math/rand"
	"time"
)

type SliceGenerator struct{}

func NewSliceGenerator() *SliceGenerator {
	return &SliceGenerator{}
}

func Sorted(n int) []time.Time {
	s := make([]time.Time, n)
	t := time.Now().UTC()

	for i := range s {
		s[i] = t.Add(time.Second * time.Duration(i))
	}

	return s
}

func Reversed(n int) []time.Time {
	s := make([]time.Time, n)
	t := time.Now().UTC()

	for i := range s {
		s[i] = t.Add(time.Second * time.Duration(n-i))
	}

	return s
}

func SortedDoubled(n int) []time.Time {
	s := make([]time.Time, 0, n)

	s1 := Sorted(int(math.Floor(float64(n) / 2)))
	s = append(s, s1...)

	s2 := Sorted(int(math.Ceil(float64(n) / 2)))
	s = append(s, s2...)

	return s
}

func Random(n int) []time.Time {
	s := make([]time.Time, n)
	t := time.Now().UTC().AddDate(-25, 0, 0)

	for i := range s {
		rnd := rand.Int63n(438300 * time.Hour.Nanoseconds())
		s[i] = t.Add(time.Duration(rnd))
	}

	return s
}

func NearSorted(n int) []time.Time {
	s := Sorted(n)

	for i := 0; i < n; i++ {
		rndInt := rand.Intn(n-i) + i
		s[i], s[rndInt] = s[rndInt], s[i]
		i += rndInt
	}

	return s
}
