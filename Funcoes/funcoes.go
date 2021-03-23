package main

import (
	"fmt"
)

func sum(value1 int8, value2 int8) int8 {
	return value1 + value2
}

//As funcoes em Go podem ter mais de um retorno
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	subtraction := n1 - n2
	return sum, subtraction
}

func main() {
	result := sum(10, 10)
	fmt.Println(result)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	response := f("Texto da funcao")
	fmt.Println(response)

	// responseSum, responseSubtraction := calculosMatematicos(30, 25)
	// fmt.Println(responseSum, responseSubtraction)

	//Se nos nao quisermos o retrno de uma das variaves, apenas adicione um underline no retorno.
	responseSum, _ := calculosMatematicos(30, 25)
	fmt.Println(responseSum)

}
