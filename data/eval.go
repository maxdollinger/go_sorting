package data

import (
	"fmt"
	"math"
	"runtime"
	"sort"
	"time"
)

type Evaluation struct {
	Size         int
	Distribution string
	Method       string
	MemoryRaw    []uint64
	ExectimeRaw  []time.Duration
}

const (
	GB = 1000000000
	MB = 1000000
	KB = 1000
)

func NewEvaluation(size int, dist, method string) *Evaluation {
	return &Evaluation{Size: size, Distribution: dist, Method: method}
}

func (e *Evaluation) AddMemorySnapshot() {
	memSnap := runtime.MemStats{}
	runtime.ReadMemStats(&memSnap)

	e.MemoryRaw = append(e.MemoryRaw, memSnap.HeapAlloc)
}

func (e *Evaluation) AddExecTime(t time.Duration) {
	e.ExectimeRaw = append(e.ExectimeRaw, t)
}

func (e *Evaluation) MemoryMeanStr() string {
	mean := float64(e.MemoryMean())

	if mean >= GB {
		return fmt.Sprintf("%.2f GB", mean/GB)
	} else if mean >= MB {
		return fmt.Sprintf("%.2f MB", mean/MB)
	} else if mean >= KB {
		return fmt.Sprintf("%.2f KB", mean/KB)
	} else {
		return fmt.Sprintf("%.2f B", mean)
	}
}

func (e *Evaluation) MemoryMean() uint64 {
	mean := uint64(0)
	for i := range e.MemoryRaw {
		mean += e.MemoryRaw[i]
	}

	return (mean / uint64(len(e.MemoryRaw)))
}

func (e *Evaluation) ExectimeMean() time.Duration {
	mean := time.Duration(0)
	for i := range e.ExectimeRaw {
		mean += e.ExectimeRaw[i]
	}

	return mean / time.Duration(len(e.ExectimeRaw))
}

func (e *Evaluation) ExectimeMedian() float64 {
	s := make([]time.Duration, len(e.ExectimeRaw))
	for i := range e.ExectimeRaw {
		s[i] = e.ExectimeRaw[i]
	}
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	l := len(s)
	if l%2 != 0 {
		return float64(s[(l/2)-1])
	} else {
		return (float64(s[(l/2)-1]) + float64(s[l/2])) / 2
	}
}

func (e *Evaluation) ExectimeFastest() float64 {
	min := time.Hour
	for i := range e.ExectimeRaw {
		if min > e.ExectimeRaw[i] {
			min = e.ExectimeRaw[i]
		}
	}

	return float64(min)
}

func (e *Evaluation) ExectimeSlowest() float64 {
	max := time.Nanosecond
	for i := range e.ExectimeRaw {
		if max < e.ExectimeRaw[i] {
			max = e.ExectimeRaw[i]
		}
	}

	return float64(max)
}

func (e *Evaluation) ExectimeP(p float64) float64 {
	l := len(e.ExectimeRaw)
	s := make([]time.Duration, len(e.ExectimeRaw))
	copy(s, e.ExectimeRaw)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	n := float64(l)
	np := n * p

	if math.Mod(np, 1.0) == 0 {
		i := int(math.Floor(np) + 1)
		return float64(s[i])
	} else {
		d := math.Mod(np, 1.0)
		i := int(math.Floor(np))
		i1 := int(math.Ceil(np))

		xi := float64(s[i])
		xi1 := float64(s[i1])

		return xi + d*(xi1-xi)
	}
}

func (e *Evaluation) ExectimeSD() time.Duration {
	mean := e.ExectimeMean()
	devSqSum := float64(0)
	for _, v := range e.ExectimeRaw {
		devSqSum += math.Pow(float64(v-mean), 2)
	}

	vari := devSqSum / float64(len(e.ExectimeRaw))

	return time.Duration(math.Sqrt(vari))
}
