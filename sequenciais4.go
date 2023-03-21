package main

import "fmt"

func main() {

	var (
		x int64
		y int64
		z int64
	)
	var (
		a int64
		b int64
		c int64
	)
	a = 2
	b = 3
	c = 5

	fmt.Print("qual o valor de x?")
	fmt.Scan(&x)
	fmt.Print("qual o valor de y?")
	fmt.Scan(&y)
	fmt.Print("qual o valor de z?")
	fmt.Scan(&z)

	resultado := (x*2+y*3*z+5)/a + b + c
	println("o resultado da média ponderada é", resultado)
	fmt.Scan(&resultado)
}
