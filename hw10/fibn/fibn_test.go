package fibn_test

import (
	"testing"

	"GB.Go.1/hw10/fibn"
)

// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711 ...
// 0  1  2  3  4  5  6  7   8   9   10  11  12   13   14   15   16   17    18    19    20    21     22    ...

var testCases = []struct {
	name string
	n    int
	out  int
}{
	{"fibonacci numbers for 0", 0, 0},
	{"fibonacci numbers for 1", 1, 1},
	{"fibonacci numbers for 2", 2, 1},
	{"fibonacci numbers for 3", 3, 2},
	{"fibonacci numbers for 4", 4, 3},
	{"fibonacci numbers for 5", 5, 5},
	{"fibonacci numbers for 6", 6, 8},
	{"fibonacci numbers for 7", 7, 13},
	{"fibonacci numbers for 8", 8, 21},
	{"fibonacci numbers for 9", 9, 34},
	{"fibonacci numbers for 11", 11, 89},
	{"fibonacci numbers for 15", 16, 987},
	{"fibonacci numbers for 21", 22, 17711},
	{"fibonacci numbers for -3", -3, 0},
}

func TestFibbonacci(t *testing.T) {
	for _, x := range testCases {
		t.Run(x.name, func(t *testing.T) {
			res := fibn.Fibbonacci(x.n)
			if res != x.out {
				t.Errorf("FAIL: got %v, want %v", res, x.out)
			}
		})
	}
}

func TestFibbonacciMap(t *testing.T) {
	for _, x := range testCases {
		t.Run(x.name, func(t *testing.T) {
			res := fibn.FibbonacciMap(x.n)
			if res != x.out {
				t.Errorf("FAIL: got %v, want %v", res, x.out)
			}
		})
	}
}

func TestFibbonacciMapOld(t *testing.T) {
	for _, x := range testCases {
		t.Run(x.name, func(t *testing.T) {
			res := fibn.FibbonacciMapOld(x.n)
			if res != x.out {
				t.Errorf("FAIL: got %v, want %v", res, x.out)
			}
		})
	}
}

func BenchmarkFibbonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibn.Fibbonacci(30)
	}
}

func BenchmarkFibbonacciMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibn.FibbonacciMap(30)
	}
}

func BenchmarkFibbonacciMapOld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fibn.FibbonacciMapOld(30)
	}
}
