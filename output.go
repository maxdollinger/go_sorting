package main

import (
	"fmt"
	"os"
)

type Output struct {
	file *os.File
}

func NewOutput(name string) (*Output, error) {
	f, err := os.OpenFile(name, (os.O_CREATE + os.O_APPEND + os.O_WRONLY), 644)
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

func (o *Output) WriteFile(str string) error {
	_, err := o.file.WriteString(str)
	return err
}

func (o *Output) WriteStdOut(str string) {
	fmt.Println(str)
}
