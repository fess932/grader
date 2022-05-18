package task_test

import "testing"

type bstruct struct {
	Name string
	Age  []int
}

const chanSize = 100
const sendSize = chanSize / 2

var (
	bbool = true
	bmap  = map[string]string{"name": "test"}
	bs    = bstruct{Name: "test"}
	bstr  = "test"
)

func Benchmark_bool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan bool, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- bbool
		}
	}
}

func Benchmark_boolPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan *bool, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- &bbool
		}
	}
}

func Benchmark_string(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan string, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- bstr
		}
	}
}

func Benchmark_stringPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan *string, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- &bstr
		}
	}
}

func Benchmark_struct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan bstruct, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- bs
		}
	}
}

func Benchmark_structPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan *bstruct, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- &bs
		}
	}
}

func Benchmark_map(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan map[string]string, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- bmap
		}
	}
}

func Benchmark_mapPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan *map[string]string, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- &bmap
		}
	}
}

func Benchmark_interface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(chan interface{}, chanSize)

		for j := 0; j < sendSize; j++ {
			a <- bmap
		}
	}
}
