package main 

import "fmt" 

func main () {
	var hari = [...]string {
		"senin",
		"selasa",
		"rabu",
		"kamis",
		"jumat",
		"sabtu",
		"minggu",
	}
	fmt.Println(hari)

	slice := hari[2:4]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice1 := hari[5:]

	slice2 := append(slice1, "seninlagi")
	fmt.Println(slice2)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "dimas"
	newSlice[1] = "afifah"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	inislice := []string {
		"example",
		"example",
	}
	fmt.Println(inislice)

	var iniarray [2]string
	iniarray[0] = "example"
	iniarray[1] = "example"
	fmt.Println(iniarray)

}