package data

import (
	"math"
	"math/rand"
	"time"
)

type DataGenerator struct{}

func NewDataGenerator() *DataGenerator {
	return &DataGenerator{}
}

func (dg *DataGenerator) Sorted(n int) []time.Time {
	s := make([]time.Time, n)
	t := time.Now().UTC()

	for i := range s {
		s[i] = t.Add(time.Second * time.Duration(i))
	}

	return s
}

func (dg *DataGenerator) Reversed(n int) []time.Time {
	s := make([]time.Time, n)
	t := time.Now().UTC()

	for i := range s {
		s[i] = t.Add(time.Second * time.Duration(n-i))
	}

	return s
}

func (dg *DataGenerator) SortedDoubled(n int) []time.Time {
	s := make([]time.Time, 0, n)

	s1 := dg.Sorted(int(math.Floor(float64(n) / 2)))
	s = append(s, s1...)

	s2 := dg.Sorted(int(math.Ceil(float64(n) / 2)))
	s = append(s, s2...)

	return s
}

func (dg *DataGenerator) Random(n int) []time.Time {
	s := make([]time.Time, n)
	t := time.Now().UTC().AddDate(-25, 0, 0)

	for i := range s {
		rnd := rand.Int63n(438300 * time.Hour.Nanoseconds())
		s[i] = t.Add(time.Duration(rnd))
	}

	return s
}

func (dg *DataGenerator) NearSorted(n int) []time.Time {
	s := dg.Sorted(n)

	for i := 0; i < n; i++ {
		rndInt := rand.Intn(n-i-1) + i
		s[i], s[rndInt] = s[rndInt], s[i]
		i += rndInt
	}

	return s
}
