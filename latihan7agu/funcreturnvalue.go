package main 


 
import "fmt" 

func getHello(name string) string {
	if name == "" {
		return "hello bro"
	} else {
		return "hello " + name
	}
}

func main() {
	result := getHello("eko")
	fmt.Println(result)

	fmt.Println(getHello(""))
}

// function bisa mengembalikan data
	// untuk memberitahu bahwa function mengembalikan data, kita harus menuliskan tipe data kembalian dari function nya function tersebut
	//jika function tersebut kita deklarasikan dengan tipe data pengembalian, maka wajib didalam function nya kita harus mengembalikan data
	// untuk mengembalikan data dari function, kita bisa menggunkan kata kunci return diikut dengan datanya