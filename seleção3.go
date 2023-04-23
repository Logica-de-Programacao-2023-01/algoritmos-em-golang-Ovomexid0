package main

import "fmt"

func main() {

	var x int

	fmt.Print("qual é o seu número?")
	fmt.Scan(&x)

	if x%2 == 0 {
		fmt.Println("o seu número é par")
	} else {
		fmt.Println("o seu número é impar")
	}

}
