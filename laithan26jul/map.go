package main

import "fmt"

func main() {
	var brand = map[string]string {
		"nama": "dimas",
		"id": "1",
	}
	fmt.Println(brand)
	fmt.Println(brand["nama"])
	fmt.Println(brand["id"])
	fmt.Println(len(brand))

	murid := make(map[string]int) 
	murid["nomor"] = 12
	murid["id"] = 1
	fmt.Println(murid["nomor"])
	fmt.Println(murid["id"])

	delete[murid["nomor"]] 
}
