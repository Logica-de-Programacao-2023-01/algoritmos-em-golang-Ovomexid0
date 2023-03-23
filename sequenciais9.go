package main

import "fmt"

func main() {

	var valor float64

	fmt.Print("qual é o valor do produto?")
	fmt.Scan(&valor)

	desconto := (valor * 10) / 100
	fmt.Println("o desconto do protuto é", desconto)

}
