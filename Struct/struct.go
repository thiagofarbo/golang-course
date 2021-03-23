package main

import "fmt"

type user struct {
	name string
	age  uint8
	address
}

type address struct {
	street  string
	zipCode string
}

func main() {

	fmt.Println("")
	var u user
	u.name = "Thiago"
	u.age = 32
	fmt.Println(u)

	addressExample := address{"Street", "06110111"}

	user2 := user{"Thiago", 32, addressExample}
	fmt.Println(user2)

	user3 := user{name: "Thiago Augusto"}
	fmt.Println(user3)

}
