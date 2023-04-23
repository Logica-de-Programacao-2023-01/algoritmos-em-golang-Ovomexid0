//Faça um algoritmo que leia dois números inteiros e mostre o resultado da multiplicação entre eles,
//se ambos forem positivos; ou a soma, se pelo menos um deles for negativo.

package main

import "fmt"

func main() {

	var x, y int

	fmt.Print("qual é o valor de x: ")
	fmt.Scan(&x)
	fmt.Print("qual é o valor de y: ")
	fmt.Scan(&y)

	if x >= 0 && y >= 0 {
		resultado := x * y
		println("o resultado da multiplicação entre x e y é: ", resultado)
	}
	if x <= 0 && y <= 0 {
		resultado2 := x + y
		println("a soma entre x e y é: ", resultado2)
	} else {
		resultado2 := x + y
		println("a soma entre x e y é: ", resultado2)
	}

}
