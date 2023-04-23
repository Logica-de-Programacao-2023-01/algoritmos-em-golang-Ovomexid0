package main

import "fmt"

func main() {
	var x, y float64

	fmt.Print("qual é o valor de x?")
	fmt.Scan(&x)
	fmt.Print("qual é o valor de y?")
	fmt.Scan(&y)

	if x > y {

		fmt.Println("x é maior que y")
	} else {
		fmt.Println("y é maior que x")
	}

}
