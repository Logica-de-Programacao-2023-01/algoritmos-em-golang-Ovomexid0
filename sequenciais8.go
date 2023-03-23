package main

import "fmt"

func main() {
	var dias, valor float64

	fmt.Print("você trabalhou por quantos dias?")
	fmt.Scan(&dias)
	fmt.Print("qual é o valor da sua diária?")
	fmt.Scan(&valor)

	salario := dias * valor
	fmt.Println("o valor de seu salário é", salario)

}
