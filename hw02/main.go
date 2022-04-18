package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Расчет площади прямоугольника")
	var a, b int
	fmt.Printf("Введите сторону a: ")
	fmt.Scanln(&a)
	fmt.Printf("Введите сторону b: ")
	fmt.Scanln(&b)
	fmt.Printf("Площадь прямоугольника %dx%d равна: %d\n", a, b, a*b)

	var s float64
	fmt.Printf("Расчет диаметра и длины окружности, площадью: ")
	fmt.Scanln(&s)
	var r = math.Sqrt(s / math.Pi)
	fmt.Printf("Диаметр: %f, длина: %f\n", 2*r, 2*math.Pi*r)

	var n, h, d, u int
	fmt.Printf("Введите трехзначное число: ")
	fmt.Scanln(&n)
	if n >= 100 && n < 1000 {
		u = n % 10
		d = (n / 10) % 10
		h = (n / 100) % 10
		fmt.Printf("В числе %d: сотен: %d, десятков: %d, единиц: %d\n", n, h, d, u)
	} else {
		fmt.Printf("Число %d не трехзначное\n", n)
	}
}
