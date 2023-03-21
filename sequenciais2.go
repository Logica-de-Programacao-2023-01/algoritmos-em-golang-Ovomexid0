package main

import "fmt"

func main() {

	var x float64

	fmt.Print("qual é o valor de x?")
	fmt.Scan(&x)

	dobro := x * 2
	triplo := x * 3
	quadruplo := x * 4

	fmt.Println("o dobro de x é", dobro)
	fmt.Println("o triplo de x é", triplo)
	fmt.Println("o quadruplo de x é", quadruplo)

}
