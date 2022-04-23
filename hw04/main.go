package main

import (
	"fmt"
)

func main() {
	fmt.Println("Сортировка массива вставками по убыванию!!")
	var n int
	fmt.Print("Введите кол-во элементов массива для сортировки: ")
	fmt.Scan(&n)
	fmt.Println("Введите э-ты массива (один элемент на строку):")
	var x int
	var v []int
	for i := 0; i < n; i++ {
		fmt.Scanln(&x)
		v = append(v, x)
	}
	fmt.Println("Массив до сортировки:", v)

	for i := 0; i < len(v); i++ {
		m := v[i] // запомним i эл-т
		j := i    // запомним его индекс
		for j > 0 && m > v[j-1] {
			v[j] = v[j-1]
			j = j - 1
		}
		v[j] = m
	}
	fmt.Println("    после сортировки:", v)
}
