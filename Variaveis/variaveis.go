package main

import "fmt"

func main() {

	var variavel1 string = "Thiago"
	fmt.Println(variavel1)

	//Declaracao Implicita
	variavel2 := "Correa"

	fmt.Println(variavel2)
	fmt.Println(variavel1, variavel2)

	var (
		variavel3 string = "Teste3"
		variavel4 string = "Teste4"
	)

	fmt.Println(variavel3, variavel4)

	//Inferencia de tipos
	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	//Invertendo valor de variaveis.
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
