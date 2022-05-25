package numbers_test

import (
	"fmt"
	"testing"

	"GB.Go.1/hw10/numbers"
)

func ExampleSimpleN() {
	fmt.Println(numbers.SimpleN(10))
	// Output:
	// [1 2 3 5 7]
}

func TestSimpleN(t *testing.T) {
	numbs := []struct {
		name string
		n    int
		res  []int
	}{
		{"простые числа до 1", 1, []int{1}},
		{"простые числа до 30", 30, []int{1, 2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{"простые числа до 0", 0, []int{0}},
		{"простые числа до -5", -5, []int{}},
	}
	for _, x := range numbs {
		t.Run(x.name, func(t *testing.T) {
			ourres := numbers.SimpleN(x.n)
			if len(x.res) != len(ourres) {
				t.Errorf("FAIL: got %v, want %v", ourres, x.res)
			}
			for i, v := range ourres {
				if v != x.res[i] {
					t.Errorf("FAIL: got %v, want %v", ourres, x.res)
				}
			}
		})
	}
}
