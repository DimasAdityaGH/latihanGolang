package main 

import "fmt" 

func sayHelloTo (firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}


func main () {
	// terkadang saat membuat function kita membutuhkan data dari luar, atau disebut parameter
	// kita bisa membuat function tanpa parameter seperti sebelumnya yang sudah kita buat
	//namun jika kita menambahkan parameter difunction, maka ketika memanggil function tersebut, kita wajib memasukan data ke parameternya
	firstName := "dimas"
	sayHelloTo(firstName, "aditya")
	sayHelloTo("dimas", "aditya")
}