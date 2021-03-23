package main

import (
	"fmt"
	"modules/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	error := checkmail.ValidateFormat("123")
	fmt.Println(error)
}
