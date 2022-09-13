package hw3bench

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

// запускаем перед основными функциями по разу чтобы файл остался в памяти в файловом кеше
// ioutil.Discard - это ioutil.Writer который никуда не пишет
func init() {
	SlowSearch(ioutil.Discard)
	FastSearch(ioutil.Discard)
}

// -----
// go test -v

func TestSearch(t *testing.T) {
	slowOut := new(bytes.Buffer)
	SlowSearch(slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	FastSearch(fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}

// -----
// go test -bench . -benchmem

func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowSearch(ioutil.Discard)
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch(ioutil.Discard)
	}
}

func BenchmarkFast1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch1(io.Discard)
	}
}

func BenchmarkFast2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch2(io.Discard)
	}
}

func BenchmarkFast3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastSearch3(io.Discard)
	}
}
