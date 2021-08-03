package main 

import "fmt" 

func main () {
	var hewan = map[string]string {
		"nama": "kucing",
		"warna": "oren",
	}
	fmt.Println(hewan)

	manusia := make(map[string]string)
	manusia["nama"] = "afifah"
	manusia["umur"] = "16"
	manusia["sekolah"] = "smpn 1 karanggeneng"

	fmt.Println(manusia)

	//hapus
	delete(manusia, "sekolah")
	fmt.Println(manusia)


	//tambah 
	manusia["wajah"] = "cantik"
	fmt.Println(manusia)

	//ubah
	manusia["nama"] = "afifah masturina"
	fmt.Println(manusia)
}