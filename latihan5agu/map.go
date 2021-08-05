package main 

import "fmt" 

func main () {
	var mahasiswa = map[string]string {
		"nama": "dimas",
		"jurusan": "teknik informatika",
	}

	fmt.Println(mahasiswa)
	fmt.Println("nama:", mahasiswa["nama"])
	fmt.Println("jurusan:",mahasiswa["jurusan"])

	bahasa := make(map[string]string)
	bahasa["nama"] = "javascript"
	bahasa["skill"] = "expert"
	bahasa["dibuat"] = "1996"

	fmt.Println(bahasa)

	//ubah
	bahasa["skill"] = "begginer"
	fmt.Println(bahasa)

	//tambah 
	bahasa["bootcamp"] = "no"
	fmt.Println(bahasa)

	//hapus
	delete(bahasa, "dibuat")
	fmt.Println(bahasa)


}