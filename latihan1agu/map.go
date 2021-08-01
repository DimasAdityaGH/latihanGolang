package main 

import "fmt" 

func main () {
	var hewan  map[string]string = map[string]string {
		"nama": "kucing",
		"makanan": "ikan",
	}
	fmt.Println(hewan)

	manusia := make(map[string]string)
	manusia["nama"] = "dimas"
	manusia["umur"] = "17"
	manusia["kelas"] = "2 smk"

	fmt.Println(manusia)

	//ubah 
	manusia["kelas"] = "3smk"
	fmt.Println(manusia)

	//tambah
	manusia["tinggi"] = "165cm"
	fmt.Println(manusia)

	//hapus
	delete(manusia, "kelas")
	fmt.Println(manusia)

	var makanan = map[string]string {
	"nama": "pecel lele",
	"dari": "lamongan",
	}
	fmt.Println(makanan)
}