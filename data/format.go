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
	str.WriteString(fmt.Sprint(f.evals[i].MemoryMeanInKibiB()))
	str.WriteString("kibiByte")
	str.WriteRune(',')
	str.WriteString(fmt.Sprint(f.evals[i].ExectimeMean()))

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
		str.WriteString(strings.ReplaceAll(fmt.Sprint(f.evals[i].MemoryRaw), ",", " "))
		str.WriteRune(',')

		str.WriteRune('[')
		for j := range f.evals[i].ExectimeRaw {
			str.WriteString(fmt.Sprint(f.evals[i].ExectimeRaw[j].Nanoseconds()))
			str.WriteRune(' ')
		}
		str.WriteRune(']')
		str.WriteRune('\n')
	}

	return str.String()
}
