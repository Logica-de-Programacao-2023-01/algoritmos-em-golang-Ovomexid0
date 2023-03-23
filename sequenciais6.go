package main

import "fmt"

func main() {
	var salario uint
	fmt.Print("qual é o seu salário atual?")
	fmt.Scan(&salario)

	Novo := (salario*15)/100 + salario
	fmt.Println("seu novo salário é", Novo)

}
