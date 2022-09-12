package hw3bench

import (
	"io"
	"os"

	"github.com/uladzislaudd/go-playground/tasks/hw/hw3_bench/pkg/fast"
)

var (
	data []byte
)

func init() {
	var err error
	data, err = os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	fast.FastSearch(out, data)
}

func FastSearch1(out io.Writer) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fast.FastSearch(out, data)
}

func FastSearch2(out io.Writer) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	fast.FastSearch(out, data)
}
