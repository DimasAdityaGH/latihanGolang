package main

import "fmt"

func main () {
	var human = map[string]string {
		"nama": "jon",
		"umur": "23",
	}
	fmt.Println(human)
	fmt.Println(human["nama"])
	fmt.Println(human["umur"])

	var length = len(human["nama"]) 
	fmt.Println(length)

	animal := make(map[string]string)
	animal["nama"] = "kucing"
	animal["makanan"] = "ikan"
	fmt.Println(animal["nama"])
	fmt.Println(animal["makanan"])

	//menambah data dimap 
	animal["kaki"] = "lima"
	fmt.Println(animal)

	//mengubah data dimap
	animal["kaki"] = "empat"
	fmt.Println(animal)

	//menghapus data dimap
	delete(animal, "kaki")
	fmt.Println(animal)

}