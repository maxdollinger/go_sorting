package data

import (
	"math"
	"math/rand"
)

// Create slices of randomized times

type SortableNumber int64

func (n SortableNumber) SortValue() int64 {
	return int64(n)
}

func Sorted(n int) []SortableNumber {
	s := make([]SortableNumber, n)

	for i := range s {
		s[i] = SortableNumber(i)
	}

	return s
}

func Reversed(n int) []SortableNumber {
	s := make([]SortableNumber, n)

	for i := range s {
		s[i] = SortableNumber(n - i)
	}

	return s
}

func SortedDoubled(n int) []SortableNumber {
	s := make([]SortableNumber, 0, n)

	s1 := Sorted(int(math.Floor(float64(n) / 2)))
	s = append(s, s1...)

	s2 := Sorted(int(math.Ceil(float64(n) / 2)))
	s = append(s, s2...)

	return s
}

func Random(n int) []SortableNumber {
	s := make([]SortableNumber, n)

	for i := range s {
		rnd := rand.Int63n(int64(n))
		s[i] = SortableNumber(rnd)
	}

	return s
}

func NearSorted(n int) []SortableNumber {
	s := Sorted(n)

	for i := 0; i < n; i++ {
		rndInt := rand.Intn(n-i) + i
		s[i], s[rndInt] = s[rndInt], s[i]
		i += rndInt
	}

	return s
}
