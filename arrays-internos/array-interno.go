package main

import (
	"fmt"
)

func main() {

	slice := make([]float32, 3, 4) //Make cria um array de um arrayInterno.
	fmt.Println(slice)

	slice = append(slice, 6)
	slice = append(slice, 7)
	fmt.Println(slice)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}
