package main

import "fmt"

func main(){
	//variabel tempat untuk menyimpan data
	var name string

	name = "Neo Archie"
	fmt.Println(name)

	name = "Archie Tech"
	fmt.Println(name)

	var myname = "Dodi S"
	fmt.Println(myname)

	fullname := "Dodi Susanto"
	fmt.Println(fullname)

	var age int8 = 30
	fmt.Println(age)

	var (
		firstName = "Dodi"
		lastName = "Susanto"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}