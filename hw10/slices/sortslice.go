package slices

import "sort"

func SortSl(v []int) []int {

	for i := 0; i < len(v); i++ {
		m := v[i] // запомним i эл-т
		j := i    // запомним его индекс
		for j > 0 && m > v[j-1] {
			v[j] = v[j-1]
			j = j - 1
		}
		v[j] = m
	}
	return v
}

func SortFromPkg(v []int) []int {
	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})
	return v
}
