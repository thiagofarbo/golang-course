package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array1 [5]int

	array1[0] = 1
	array1[4] = 2

	fmt.Println(array1)

	array2 := [5]string{"Position 1", "Position 2", "Position 3", "Position 4", "Position 5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// slice nao tem um tamanho fixo.
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, -1)
	fmt.Println(slice)

	//Pegando o intervalo de um array
	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Position 1"
	fmt.Println(slice2)

}
