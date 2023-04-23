//Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo salário com um aumento de 10%
//se o salário for menor que R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.

package main

import "fmt"

func main() {

	var salario int

	fmt.Print("qual é o valor do seu salário? ")
	fmt.Scan(&salario)

	if salario < 1000 {
		resultado := salario*10/100 + salario
		println("seu novo salário é: ", resultado)
	} else {
		resultado2 := salario*5/100 + salario
		println("seu novo salário é: ", resultado2)
	}
}
