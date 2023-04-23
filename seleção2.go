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

	if x < y && x < z {
		fmt.Println(x, "é o menor valor")
	}
	if y < x && y < z {
		fmt.Println(y, "é o menor valor")
	}
	if z < x && z < x {
		fmt.Println(z, "é o menor valor")
	}
}
