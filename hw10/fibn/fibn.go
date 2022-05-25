package fibn

import "fmt"

func Fibbonacci(a int) int {
	if a < 0 {
		fmt.Println("Данная функция только для положительных чисел")
		return 0
	}
	if a <= 1 {
		return a
	}
	return Fibbonacci(a-1) + Fibbonacci(a-2)
}

var m = map[int]int{0: 0, 1: 1}

func FibbonacciMap(a int) int {
	if a < 0 {
		fmt.Println("Данная функция только для положительных чисел")
		return 0
	}
	nF, ok := m[a]
	if !ok {
		nF = FibbonacciMap(a-1) + FibbonacciMap(a-2)
		m[a] = nF
	}
	return nF
}

var fib = map[int]int{}

func FibbonacciMapOld(a int) int {
	if a < 0 {
		fmt.Println("Данная функция только для положительных чисел")
		return 0
	}
	_, x := fib[a-1]
	_, y := fib[a-2]

	if a <= 1 && (!x || !y) {
		fib[a] = a
		return fib[a]
	}

	if !x || !y {
		v := FibbonacciMapOld(a-1) + FibbonacciMapOld(a-2)
		fib[a] = v
	} else {
		fib[a] = fib[a-1] + fib[a-2]
	}
	return fib[a]
}
