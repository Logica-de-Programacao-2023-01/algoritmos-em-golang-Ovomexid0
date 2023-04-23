//Faça um algoritmo que leia um número inteiro entre 1 e 7 e mostre o dia da semana correspondente (1 = domingo, 2 = segunda-feira, etc.).

package main

import "fmt"

func main() {

	var x int

	fmt.Print("escolha um número de 1 a 7 correspondente aos dias da semana: ")
	fmt.Scan(&x)

	if x == 1 {
		println("1 = domingo")
	} else if x == 2 {
		println("2 = segunda-feira")
	} else if x == 3 {
		println("3 = terça-feira")
	} else if x == 4 {
		println("4 = quarta-feira")
	} else if x == 5 {
		println("5 = quinta-feira")
	} else if x == 6 {
		println("6 = sexta-feira")
	} else if x == 7 {
		println("7 = sábado")
	} else {
		println("desculpe essa não é uma opção válida")
	}

}
