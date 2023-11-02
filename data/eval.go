package data

import (
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

func (e *Evaluation) MemoryMeanInKibiB() uint64 {
	mean := uint64(0)
	for i := range e.MemoryRaw {
		mean += e.MemoryRaw[i]
	}

	return (mean / uint64(len(e.MemoryRaw))) / 1_024
}

func (e *Evaluation) ExectimeMean() time.Duration {
	mean := time.Duration(0)
	for i := range e.ExectimeRaw {
		mean += e.ExectimeRaw[i]
	}

	return mean / time.Duration(len(e.ExectimeRaw))
}

func (e *Evaluation) FirstExec() time.Duration {
	t := time.Duration(0)

	if len(e.ExectimeRaw) >= 0 {
		t = e.ExectimeRaw[0]
	}

	return t
}

func (e *Evaluation) ExectimeMedian() time.Duration {
	l := len(e.ExectimeRaw)
	if l%2 != 0 {
		return e.ExectimeRaw[(l+1)/2]
	} else {
		return (e.ExectimeRaw[l/2] + e.ExectimeRaw[l/2+1]) / 2
	}
}

func (e *Evaluation) ExectimeFastest() time.Duration {
	min := time.Hour
	for i := range e.ExectimeRaw {
		if min > e.ExectimeRaw[i] {
			min = e.ExectimeRaw[i]
		}
	}

	return min
}

func (e *Evaluation) ExectimeSlowest() time.Duration {
	max := time.Nanosecond
	for i := range e.ExectimeRaw {
		if max < e.ExectimeRaw[i] {
			max = e.ExectimeRaw[i]
		}
	}

	return max
}

func (e *Evaluation) ExectimeP(p int) time.Duration {
	l := len(e.ExectimeRaw)
	p = int(math.Floor(float64(l / 100 * p)))
	s := make([]time.Duration, l)
	for i := range e.ExectimeRaw {
		s[i] = e.ExectimeRaw[i]
	}

	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })

	return e.ExectimeRaw[p]
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
