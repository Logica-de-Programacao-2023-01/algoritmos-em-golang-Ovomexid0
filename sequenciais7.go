package main

import "fmt"

func main() {
	var x int64
	fmt.Print("qual o valor de x?")
	fmt.Scan(&x)

	sucessor := x + 1
	fmt.Println("o sucessor de x é", sucessor)
	antecessor := x - 1
	fmt.Println("o antecessor de x é", antecessor)

}
