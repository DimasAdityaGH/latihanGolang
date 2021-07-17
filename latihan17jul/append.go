package main

import "fmt"

func main () {
	bulan := [...]string {"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	slice1 := bulan[10:]
	fmt.Println(slice1)

	slice2 := append(slice1, "new")
	fmt.Println(slice2)

	slice2[1] = "new append"
	fmt.Println(slice2)
	fmt.Println(slice1)
	

}