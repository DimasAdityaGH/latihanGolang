package main

import "fmt"

func main() {
	//tipe data map
	
	buku := map[string]string {
		"name": "dimas",
		"alamat": "tejo",
	}

	fmt.Println(buku)
	fmt.Println(buku["name"])
	fmt.Println(buku["alamat"])

	var murid map[string]string = make(map[string]string)   	
		murid["cerdas"] = "dimas"
		murid["berkembang"] = "dwi"
		murid["out"] = "arif"
		fmt.Println(murid)
		fmt.Println(len(murid))

		//menghapus data di map dengan key
		delete(murid, "out")

		fmt.Println(murid)
		fmt.Println(len(murid))

		Member := make(map[string]string)
		Member["satu"] = "dian"
		Member["dua"] = "deni"
		Member["tiga"] = "jon"
		
		//mengubah data dimap dengan map key
		Member["tiga"] = "lia"

		fmt.Println(Member)
		fmt.Println(Member["satu"])
		fmt.Println(Member["dua"])
		fmt.Println(Member["tiga"])
}