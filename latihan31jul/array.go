package main 

import "fmt" 

func main () {
	var buah [3]string
	buah[0] = "mangga"
	buah[1] = "pelem"
	buah[2] = "manggis"
	fmt.Println(buah)

	hewan := [...]string {
		"kucing",
		"tikus",
	}
	fmt.Println(hewan)
}