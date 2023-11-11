package main

import (
	"fmt"

	entities "github.com/andresousadotpt/teste/entities"
)

func main() {
	a := entities.User{
		Name: "vski", // it has to be Name because name in entitier file does not exists
	}
	fmt.Println(a)
}
