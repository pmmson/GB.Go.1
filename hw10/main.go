package main

import (
	"fmt"
	"time"

	"GB.Go.1/hw10/fibn"
	"GB.Go.1/hw10/numbers"
	"GB.Go.1/hw10/slices"
)

func main() {
	fmt.Println("=== Простые числа!! ===")
	var n int
	fmt.Print("Вывести все простые числа от 0 до ")
	fmt.Scanln(&n)
	fmt.Println(numbers.SimpleN(n))

	fmt.Println("=== Сортировка массива вставками по убыванию!! ===")
	var m int
	fmt.Print("Введите кол-во элементов массива для сортировки: ")
	fmt.Scan(&m)
	fmt.Println("Введите э-ты массива (один элемент на строку):")
	var x int
	var v []int
	for i := 0; i < m; i++ {
		fmt.Scanln(&x)
		v = append(v, x)
	}
	fmt.Println("Массив до сортировки:", v)
	u := slices.SortSl(v)
	fmt.Println("после сортировки вставками:", u)
	u = slices.SortFromPkg(v)
	fmt.Println("после сортировки пакетом sort:", u)

	fmt.Println("=== Фибоначчии!! ===")
	fmt.Print("Введите порядковый номер числа Фибоначчи: ")
	var fn int
	fmt.Scanln(&fn)

	// recurcive func
	start := time.Now()
	fb := fibn.Fibbonacci(fn)
	end := time.Now()
	totalTime := end.Sub(start)
	fmt.Println(fb, totalTime)

	// recurcive with map func
	start = time.Now()
	fb = fibn.FibbonacciMap(fn)
	end = time.Now()
	totalTime = end.Sub(start)
	fmt.Println(fb, totalTime)

	// recurcive with map func old
	start = time.Now()
	fb = fibn.FibbonacciMapOld(fn)
	end = time.Now()
	totalTime = end.Sub(start)
	fmt.Println(fb, totalTime)
}
