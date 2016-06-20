package main

import (
	"io"
	"os"

	"github.com/docker/docker/builder/dockerfile/parser"
	"github.com/k0kubun/pp"
)

func main() {
	err := Main(os.Args)
	if err != nil {
		panic(err)
	}

}

func Main(args []string) error {
	for _, path := range args[1:] {
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		defer f.Close()

		err = Show(f)
		if err != nil {
			return err
		}
	}
	return nil
}

func Show(r io.Reader) error {
	n, err := parser.Parse(r)
	if err != nil {
		return err
	}

	pp.Println(n)
	return nil
}
