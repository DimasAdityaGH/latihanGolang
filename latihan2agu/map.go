package main 

import "fmt" 

func main () {
	var hewan = map[string]string {
		"nama": "kucing",
		"warna": "oren",
	}
	fmt.Println(hewan)

	//plus
	hewan["kaki"] = "4"
	fmt.Println(hewan)

	//delete
	delete(hewan, "warna")
	fmt.Println(hewan)

	//ubah
	hewan["nama"] = "singa"
	fmt.Println(hewan)

}