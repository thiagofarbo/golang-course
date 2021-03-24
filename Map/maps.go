package main

import (
	"fmt"
)

func main() {

	//[tipos das chaves] {tipos dos valores}
	user := map[string]string{
		"name":     "Thiago",
		"lastname": "Correa",
	}

	//Formas do print
	fmt.Println(user["name"])
	fmt.Println(user)

	address := map[string]map[string]string{
		"street": {
			"zipCode": "01234000",
			"name":    "Luis Scott",
		},
		"city": {
			"name":  "Barueri",
			"state": "SP",
		},
	}
	fmt.Println(address)
	fmt.Println("Deletando => -----------------")
	//deletando uma chave
	delete(address, "city")
	fmt.Println(address)

	fmt.Println("Adicionando => -----------------")

	//adicionando uma chave
	address["neighbor"] = map[string]string{
		"name": "Jardim Iracema",
	}
	address["city"] = map[string]string{
		"name":  "Barueri",
		"state": "SP",
	}
	fmt.Println(address)
}
