//Faça um algoritmo que leia um número inteiro e mostre se ele é múltiplo de 3 e de 5 ao mesmo tempo.

package main

import "fmt"

func main() {

	var x int

	fmt.Print("qual é o valor de x ? ")
	fmt.Scan(&x)

	if x%5 == 0 {
		fmt.Println("x é divisível por 5")
	} else {
		fmt.Println("x não é divisivel por 5")
	}
	if x%3 == 0 {
		fmt.Println("x é divisível por 3 ")
	} else {
		fmt.Println("x não é divisivel por 3")
	}

}
