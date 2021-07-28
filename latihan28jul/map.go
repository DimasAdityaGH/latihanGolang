package main

import "fmt"

func main() {
	buah := map[string]string {
		"nama": "pisang",
		"rasa": "manis",
	}
	fmt.Println(buah)
	fmt.Println(buah["nama"])
	fmt.Println(buah["rasa"])

	var hewan map[string]string = map[string]string {
		"nama": "kucing",
		"warna": "oren",
	}
	fmt.Println(hewan)
	fmt.Println(hewan["nama"])
	fmt.Println(hewan["warna"])

	benua := make(map[string]string)
	benua["nama"] = "asia"
	benua["luas"] = "212133 km"
	fmt.Println(benua)
	fmt.Println(benua["nama"])
	fmt.Println(benua["luas"])

	delete(benua, "luas")
	fmt.Println(benua)
	benua["luas"] = "44,58 km²"
	fmt.Println(benua)
	benua["luas"] = "44,58 juta km²"
	fmt.Println(benua)

}