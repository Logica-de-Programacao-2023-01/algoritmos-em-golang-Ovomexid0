package main

import "fmt"

func main() {

	var x, y, z int

	fmt.Print("qual é o valor de x?")
	fmt.Scan(&x)
	fmt.Print("qual é o valor de y?")
	fmt.Scan(&y)
	fmt.Print("qual é o valor de z?")
	fmt.Scan(&z)

	resultado := x + y + z
	fmt.Print("a soma das 3 Variavéis é", resultado)
}
