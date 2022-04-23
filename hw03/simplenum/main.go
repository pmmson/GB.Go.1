package main

import "fmt"

func main() {
	var n int
	fmt.Print("Вывести все простые числа от 0 до ")
	fmt.Scanln(&n)

	var cntdiv int
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i%j != 0 {
				continue
			}
			cntdiv++
		}
		if cntdiv > 2 {
			cntdiv = 0
			continue
		}
		fmt.Printf("%d ", i)
		cntdiv = 0
	}
	fmt.Print("\n")
}
