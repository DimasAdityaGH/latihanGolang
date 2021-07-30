package main 

import "fmt" 

func main () {
	var nama [3]string 
	nama[0] = "dimas"
	nama[1] = "afifah"
	nama[2] = "aditya"
	fmt.Println(nama)

	var color = [...]string {"merah", "kuning", "hijau", "biru", "hitam", "putih"}
	fmt.Println(color)

	slice1 := color[3:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := color[4:]

	slice3 := append(slice2, "kuning")
	fmt.Println(slice3)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "satu"
	newSlice[1] = "dua"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))
}