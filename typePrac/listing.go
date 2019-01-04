package main

import (
	"fmt"
	"golang_prc/entities"
)


func main(){
	a := entities.Admin{
		Rights: 5,
	}

	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)
}