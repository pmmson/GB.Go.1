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
func (s *Slb) Read(x []byte) (n int, err error) {
	n = copy(x, *s)
	err = io.EOF // добавил, так как при debug увидел, что выходом из бесконечного цикла в io.ReadAll является это условие
	return
}

func main() {

	// С помощью io.WriteString записать в переменную вашего типа произвольную строку
	var z Slb
	io.WriteString(&z, "QWertuIopaHDklzBG-+msjfnKZKNDKnskslmx")
	fmt.Println(z)

	// С помощью io.ReadAll( ) считать вашу строку обратно (вообще говоря, он возвращает слайс байт, но его легко привести к виду строки)
	var x []byte
	x, _ = io.ReadAll(&z)

	fmt.Println(string(x))

}
