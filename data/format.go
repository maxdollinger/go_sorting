package data

import (
	"fmt"
	"strings"
)

type Formater struct {
	evals []*Evaluation
}

func NewFormater(evals []*Evaluation) *Formater {
	return &Formater{
		evals: evals,
	}
}

func (f *Formater) String() string {
	str := strings.Builder{}

	for i := range f.evals {
		str.WriteString(f.eval2String(i))
	}

	return str.String()
}

func (f *Formater) eval2String(i int) string {
	str := strings.Builder{}

	str.WriteString(f.evals[i].Method)
	str.WriteRune(',')
	str.WriteString(f.evals[i].Distribution)
	str.WriteRune(',')
	str.WriteString(fmt.Sprint(f.evals[i].Size))
	str.WriteRune(',')
	str.WriteString("mean-")
	str.WriteString(fmt.Sprint(f.evals[i].ExectimeMean().Microseconds()))
	str.WriteString("µs")
	str.WriteRune(',')
	str.WriteString("sd-")
	str.WriteString(fmt.Sprint(f.evals[i].ExectimeSD().Microseconds()))
	str.WriteString("µs")
	str.WriteRune(',')
	str.WriteString("slow-")
	str.WriteString(fmt.Sprint(f.evals[i].ExectimeSlowest().Microseconds()))
	str.WriteString("µs")
	str.WriteRune(',')
	str.WriteString("fast-")
	str.WriteString(fmt.Sprint(f.evals[i].ExectimeFastest().Microseconds()))
	str.WriteString("µs")
	str.WriteRune(',')
	str.WriteString("p95-")
	str.WriteString(fmt.Sprint(f.evals[i].ExectimeP(5).Microseconds()))
	str.WriteString("µs")
	str.WriteRune(',')
	str.WriteString(fmt.Sprint(f.evals[i].MemoryMeanInKibiB()))
	str.WriteString("kB")

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
			str.WriteString(fmt.Sprint(f.evals[i].ExectimeRaw[j].Microseconds()))
			str.WriteRune(' ')
		}
		str.WriteRune(']')

		str.WriteRune(',')
		str.WriteString(strings.ReplaceAll(fmt.Sprint(f.evals[i].MemoryRaw), ",", " "))
		str.WriteRune('\n')
	}

	return str.String()
}
