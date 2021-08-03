package main 

import "fmt" 

func main () {
	var nilaiujian = 78
	var nilaitugas = 80

	var nilaikKmujian = nilaiujian > 75
	var nilaikkmtugas = nilaitugas >= 80

	var hasil = nilaikKmujian && nilaikkmtugas
	fmt.Println(hasil)
}