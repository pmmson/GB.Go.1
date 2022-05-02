package main

import (
	"fmt"
	"time"
)

func main() {
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711 ...
	// 0  1  2  3  4  5  6  7   8   9   10  11  12   13   14   15   16   17    18    19    20    21     22    ...

	fmt.Print("Введите порядковый номер числа Фибоначчи: ")
	var n int
	fmt.Scanln(&n)

	// recurcive func
	start := time.Now()
	fb := fibbonacci(n)
	end := time.Now()
	totalTime := end.Sub(start)
	fmt.Println(fb, totalTime)

	// recurcive with map func
	start = time.Now()
	fb = fibbonacciMap(n)
	end = time.Now()
	totalTime = end.Sub(start)
	fmt.Println(fb, totalTime)
}

func fibbonacci(a int) int {
	if a <= 1 {
		return a
	}
	return fibbonacci(a-1) + fibbonacci(a-2)
}

var m = map[int]int{0: 0, 1: 1}

func fibbonacciMap(a int) int {
	nF, ok := m[a]
	if !ok {
		nF = fibbonacciMap(a-1) + fibbonacciMap(a-2)
		m[a] = nF
	}
	return nF
}
