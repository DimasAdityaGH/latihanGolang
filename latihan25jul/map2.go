package main

import "fmt"

func main () {
	Pelanggan := map[string]string {
		"nama": "dimas",
		"id": "092",
	}
	fmt.Println(Pelanggan)
	fmt.Println(Pelanggan["nama"])
	fmt.Println(Pelanggan["id"])

	var penjual map[string]string = make(map[string]string)
	penjual["nama"] = "dwi"
	penjual["id"] = "012"
	fmt.Println(penjual)
	fmt.Println(penjual["nama"])
	fmt.Println(penjual["id"])

	developer := make(map[string]string)
	developer["web"] = "john"
	developer["mobile"] = "jeny"
	fmt.Println(developer)
	fmt.Println(developer["web"])
	fmt.Println(len(developer))

	developer["web"] = "denis"
	fmt.Println(developer["web"])
}