package main

import (
	"fmt"
	"math"
	"os"
)

func main() {

	defer func() {
		err:= recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, !): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			panic("Второе число не должно быть равно 0! На 0 делить нельзя!")
			// os.Exit(1)
		}
		res = a / b
	case "!":
		res = 1
		for i := 1.0; i <= a; i++ {
			res = res * i
		}
	case "^":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	if op == "!" {
		fmt.Printf("Результат выполнения операции %.2f%s = %.2f\n", a, op, res)
	} else {
		fmt.Printf("Результат выполнения операции %.2f%s%.2f = %.2f\n", a, op, b, res)
	}
}
