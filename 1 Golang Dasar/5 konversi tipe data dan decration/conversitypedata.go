package main

import "fmt"

func main(){
	//constant variabel  yang tidak biasa diubah lagi isinya
	var nilai32 int32 = 32000
	var nilai64 int64 = int64(nilai32)

	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var name = "Dodi"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)

	//type declaration 
	//merubah tipe data 

	type NoKTP string
	type Married bool

	var ktpDodi NoKTP = "800800800"
	var marriedStatus = true
	fmt.Println(ktpDodi)
	fmt.Println(marriedStatus)

}