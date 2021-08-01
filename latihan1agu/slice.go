package main 

import "fmt" 

func main () {
	var months = [...]string {"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	fmt.Println(months)

	slice := months[4:7]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice1 := make([]string, 3, 6)
	slice1[0] = "vscode"
	slice1[1] = "sublime"
	slice1[2] = "notepad"
	fmt.Println(slice1)
	fmt.Println(cap(slice1))

	slice2 := append(slice1, "Atom")
	fmt.Println(slice2)

	slice3 := months[10:] 

	slice4 := append(slice3, "januari")
	fmt.Println(slice4)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "satu"
	newSlice[1] = "dua"
	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}