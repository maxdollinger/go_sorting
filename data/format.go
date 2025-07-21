package data

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

type Formater struct {
	evals []*Evaluation
}

func NewFormater(evals []*Evaluation) *Formater {
	sort.Slice(evals, func(i, j int) bool {
		if evals[i].Size == evals[j].Size {
			return evals[i].Method < evals[j].Method
		} else {
			return evals[i].Size < evals[j].Size
		}
	})
	return &Formater{
		evals: evals,
	}
}

func (f *Formater) String() string {
	str := strings.Builder{}

	str.WriteString("method;dist;n;Q1;median;3Q;min;max;memory\n")

	for i := range f.evals {
		_, err := str.WriteString(f.csvEntry(f.evals[i]))
		if err != nil {
			log.Fatal(fmt.Errorf("Error writing to EvalFormatString: %v", err))
		}
	}

	return str.String()
}

func (f *Formater) csvEntry(eval *Evaluation) string {
	str := strings.Builder{}

	str.WriteString(eval.Method)
	str.WriteRune(';')
	str.WriteString(eval.Distribution)
	str.WriteRune(';')
	str.WriteString(numShort(eval.Size))
	str.WriteRune(';')
	str.WriteString(timeFormat(eval.ExectimeP(0.25)))
	str.WriteRune(';')
	str.WriteString(timeFormat(eval.ExectimeMedian()))
	str.WriteRune(';')
	str.WriteString(timeFormat(eval.ExectimeP(0.75)))
	str.WriteRune(';')
	str.WriteString(timeFormat(eval.ExectimeFastest()))
	str.WriteRune(';')
	str.WriteString(timeFormat(eval.ExectimeSlowest()))
	str.WriteRune(';')
	str.WriteString(eval.MemoryMeanStr())

	str.WriteRune('\n')
	return str.String()
}

func (f *Formater) RawData() string {
	str := strings.Builder{}

	for i := range f.evals {
		str.WriteString(f.evals[i].Method)
		str.WriteRune(',')
		str.WriteString(f.evals[i].Distribution)
		str.WriteRune(',')
		str.WriteString(fmt.Sprint(f.evals[i].Size))

		str.WriteRune(',')
		str.WriteRune('[')
		for j := range f.evals[i].ExectimeRaw {
			str.WriteString(fmt.Sprint(f.evals[i].ExectimeRaw[j].Nanoseconds()))
			str.WriteRune(',')
		}
		str.WriteRune(']')

		str.WriteRune(',')
		str.WriteString(strings.ReplaceAll(fmt.Sprint(f.evals[i].MemoryRaw), ",", " "))
		str.WriteRune('\n')
	}

	return str.String()
}

func numShort(n int) string {
	unit := ""
	if n >= 1000000 {
		n = n / 1000000
		unit = "m"
	} else if n >= 1000 {
		n = n / 1000
		unit = "k"
	}

	return fmt.Sprintf("%v%s", n, unit)
}

const (
	µs = 1000
	ms = 1000000
	s  = 1000000000
)

func timeFormat(t float64) string {
	unit := "ns"

	if t >= s {
		t = t / s
		unit = "s"
	} else if t >= ms {
		t = t / ms
		unit = "ms"
	} else if t >= µs {
		t = t / µs
		unit = "µs"
	}

	return fmt.Sprintf("%.2f%s", t, unit)
}
