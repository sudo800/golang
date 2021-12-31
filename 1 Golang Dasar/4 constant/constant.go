package main

import "fmt"

func main(){
	//constant variabel  yang tidak biasa diubah lagi isinya

	const firstName string 	= "Dodi"
	const lastName 			= "Susanto"

	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		firstName string = "Dodi"
		 lastName 		 = "Susanto"
	)
	
	fmt.Println(firstName)
	fmt.Println(lastName)



}