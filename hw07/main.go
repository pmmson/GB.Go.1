package main

import (
	"flag"
	"fmt"
	"io"
	"math/rand"
	"time"
)

// Необходимо объявить свой тип, обернув в него тип []byte - (слайс байтов)
type Slb []byte

// реализовать на нем такие методы, чтобы он удовлетворял интферйесам io.Writer
func (s *Slb) Write(x []byte) (n int, err error) {
	*s = append(*s, x...)
	return
}

var k int // счетчик уже считанных байт
// реализовать на нем такие методы, чтобы он удовлетворял интферйесам io.Reader
func (s *Slb) Read(x []byte) (n int, err error) {
	// 	n = copy(x, *s) // при копировании x станет длины *s, но если больше 512 - строки не совпадают !!

	for i, v := range *s {
		if i < k {
			continue // пропускаем все уже считанные эл-ты
		}
		x[i-k] = v // записываем в x остатки
		n++
		if n > len(x)-1 { // если n превысит длину x выходим
			break
		}
	}
	k += n // увеличиваем k счетчик
	if len(*s) == k {
		err = io.EOF // выходим если считаны все данные из s
	}
	return
}

func main() {

	cntS := flag.Int("n", 5, "length of string for generator")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	var str = randStr(*cntS)
	fmt.Println("Сгенерирована строка:", str)

	// С помощью io.WriteString() записать в переменную вашего типа произвольную строку
	var z Slb
	io.WriteString(&z, str)
	fmt.Printf("Записано в наш тип: %v\n", z)

	// С помощью io.ReadAll() считать вашу строку обратно (вообще говоря, он возвращает слайс байт, но его легко привести к виду строки)
	var x []byte
	x, _ = io.ReadAll(&z)

	fmt.Printf("Считано из нашего типа: %s\n", string(x))
	fmt.Printf("Введеные строки совпадают: %t\n", string(x) == str)

}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
