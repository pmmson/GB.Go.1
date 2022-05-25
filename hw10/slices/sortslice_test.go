package slices_test

import (
	"testing"

	"GB.Go.1/hw10/slices"
)

func BenchmarkSortSl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slices.SortSl([]int{6, 1, 5, 2, 8, 0, 3, 8, 2, 5, 7, 4, 3, 3})
	}
}

func BenchmarkSortFromPkg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = slices.SortFromPkg([]int{6, 1, 5, 2, 8, 0, 3, 8, 2, 5, 7, 4, 3, 3})
	}
}
