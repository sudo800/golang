package main

import "fmt"

func main(){
	var (
			a = 10
			b = 10
			c = a + b
		)

	fmt.Println(c)
	
	c += 5 //c = c+5
	fmt.Println(c)

	c++ //c=c+1
	fmt.Println(c)

	//operasi perbandingan

	var (
		name1 = "Dodi"
		name2 = "Dodi S"
	)

	var result bool = name1 == name2

	fmt.Println(result) //hasil true atau false

	var (
		value1 = 100
		value2 = 800
	)

	fmt.Println(value1 < value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)


	//operasi boolean
	//&& dan || atau  ! kebalikan

	var (
		nilaiAkhir = 90
		absensi = 80
	)

	var lulusNIlaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNIlaiAkhir && lulusAbsensi

	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
	
}