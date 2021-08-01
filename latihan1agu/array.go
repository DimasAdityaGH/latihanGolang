package main 

import "fmt" 

func main () {
	var hari = [3]string {
		"senin",
		"selasa",
		"rabu",
	}
	fmt.Println(hari)

	var day [3]string
	day[0] = "senin"
	day[1] = "selasa"
	day[2] = "rabu"
	fmt.Println(day)
}