package main

import "fmt"

func main() {
	// data := map[string]int {}
	// data["dimas"] = 10
	// fmt.Println(data)

	// var siswa = map[string]string {
	// 	"nama": "dimas",
	// 	"kelas": "3",
	// }
	// fmt.Println(siswa)

	// makanan := make(map[string]string)
	// makanan["nama"] = "jeruk"
	// makanan["warna"] = "oren"
	// fmt.Println(makanan)

	// for key, value := range makanan {
	// 	fmt.Println(key, "=", value)
	// }

	// //  Deteksi Keberadaan Item Dengan Key Tertentu
	// chicken := map[string]int {
	// 	"januari": 50,
	// 	"februari": 40,
	// }
	// var value, isExist = chicken["januari"]

	// if isExist {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exists")
	// }

	//kombinasi slice dan map
	var chickens = []map[string]string{
		{"name": "chicken blue",   "gender": "male"},
		{"name": "chicken red",    "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	
	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}