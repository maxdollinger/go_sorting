package main

import (
	"fmt"
	"os"

	"github.com/maxdolliger/timesort/data"
)

type Output struct {
	file *os.File
}

func NewOutput(name string) (*Output, error) {
	f, err := os.OpenFile(name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 6644)
	if err != nil {
		return nil, err
	}
	o := &Output{
		file: f,
	}

	return o, nil
}

func (o *Output) Close() {
	o.file.Close()
}

func (o *Output) Write(form *data.Formater) error {
	fmt.Print(form)
	_, err := o.file.WriteString(form.RawData())
	return err
}
