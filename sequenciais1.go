package main

import "fmt"

func main() {

	var x, y, z float64

	fmt.Print("qual é o valor de x?")
	fmt.Scan(&x)
	fmt.Print("qual é o valor de y?")
	fmt.Scan(&y)
	fmt.Print("qual é o valor de z?")
	fmt.Scan(&z)

	resultado := x + y + z
	fmt.Print("qual é a soma das 3 variavéis?", resultado)
}
