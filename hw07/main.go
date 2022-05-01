package main

import (
	"fmt"
	"io"
)

type Slb []byte // Необходимо объявить свой тип, обернув в него тип []byte - (слайс байтов)

// Затем, необходимо реализовать на нем такие методы, чтобы он удовлетворял интферйесам io.Writer
func (s *Slb) Write(x []byte) (n int, err error) {
	*s = make(Slb, len(x))
	n = copy(*s, x)
	return
}

// io.Reader
// func (s *Slb) Read(x []byte) (n int, err error) {
// 	sl := make([]byte, len(*s))
// 	x = append(x, sl...)
// 	n = copy(x, *s)
// 	return
// }

func (s *Slb) Read(x *[]byte) (n int, err error) {
	sl := make([]byte, len(*s))
	*x = append(*x, sl...)
	n = copy(*x, *s)
	return
}

func main() {

	var z Slb

	// С помощью io.WriteString записать в переменную вашего типа произвольную строку
	io.WriteString(&z, "QazWerTy=!")
	fmt.Println(z)

	var x []byte
	z.Read(&x)
	fmt.Println(string(x))

	// С помощью io.ReadAll( ) считать вашу строку обратно (вообще говоря, он возвращает слайс байт, но его легко привести к виду строки)
	// не нашел такой функции в пакете io

}
