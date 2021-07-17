package main

import "fmt"

func main() {
	bulan := [...]string {"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	slice1 := bulan[3:8]
	fmt.Println(slice1)
	//len function
	fmt.Println(len(slice1))
	//cap function
	fmt.Println(cap(slice1))
	
}