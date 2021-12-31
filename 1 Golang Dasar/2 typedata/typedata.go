package main

import "fmt"

func main(){
	//type data
	// int -128 sd 127
	// int16 -32768 sd 32767
	// int32 -2137483648 sd 2137483647 2M
	// int64 -9223372036854775808 sd 9223372036854775808
	// uint8 0 sd 255
	// uint16 0 sd 65535
	// uint32 0 sd 4294967295
	// uint64 0 sd 18336744073709551615
	// float32 1.18X10pangkat-38 sd 3.4X10pangkat38
	// float64 2.23X10pangkat-308 sd 1.80X10pangkat308
	// complex64 dan comprek128
	// ALIAS
	// byte = uint8
	// rune = int32
	// int = tergantung OS yang dipakai int32/64
	// uint =tergantung OS yang dipakai uint32/64
	fmt.Println("NUMBER")
	fmt.Println("Satu = ",1)
	fmt.Println("Dua = ",2)
	fmt.Println("Tiga Koma Lima = ",3.5)

	//BOOLEAN
	//bernilai benar atau salah kata kuncinya bool true / false
	fmt.Println("BOOLEAN")
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	//STRING
	fmt.Println("STRING")
	fmt.Println("Ini String ya, Karna STRING itu adalah kumpulan huruf/angka")
	fmt.Println(len("Neo"))
	fmt.Println("Neo Archie"[0]) //tampilnya Byte nya N
	fmt.Println("Neo Archie Tech"[1]) //tampilnya Byte nya e

}