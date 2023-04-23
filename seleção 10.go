package main

import "fmt"

func main() {

	var idade int

	fmt.Print("qual é a sua idade? ")
	fmt.Scan(&idade)

	if idade <= 9 {
		println("você é mirin ")
	} else if idade >= 10 && idade <= 13 {
		println("você é infantil ")
	} else if idade >= 14 && idade <= 17 {
		println("você é juvenil ")
	} else if idade >= 18 {
		println("você é adulto ")
	}

}
