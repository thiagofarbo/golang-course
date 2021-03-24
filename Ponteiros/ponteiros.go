package main

import "fmt"

func main() {

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)
	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO E UMA REFERENCIA DE MEMORIA

	var variavel3 int
	var ponteiro *int //* faz com que essa varivael seja um ponteiro

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro)

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) //Para pegarmos um valor de um ponteiro, precisamos add o * antes do nosso ponteiro.

}
