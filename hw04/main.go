package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	fmt.Print("Введите кол-во элементов массива для сортировки: ")
	fmt.Scanln(&n)
	// генерируем последовательность для сортировки
	var v []int
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < n; i++ {
		v = append(v, rand.Intn(100))
	}
	fmt.Println(v)

	for i := 0; i < len(v); i++ {
		key := v[i] // просматриваем эл-т исходного массива и его индекс
		j := i
		for j > 0 && v[j-1] > key { // получается нет проверки второго условия цикла когда j = 0 ?
			v[j] = v[j-1] // если предыдущий эл-т больше просматриваемого - меняем его и его позицию
			j = j - 1
		}
		v[j] = key
	}
	fmt.Println(v)
}
