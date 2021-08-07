package main 

import "fmt" 

	// function adalah sebuah blok kode yang sengaja dibuat program agar bisa digunakan berulang ulang
	// cara membuar function sangat sederhana, hanya dengan menggunakan kata kunci func lali diikuti dengan anam function nya dan blok kode isi functionya
	//setelah membuat function, kita bisa mengeksekusi function tersebut dengan memanggilanya  menggunakan kata kunci nama function nya diikuti tanda kurung buka, kurang tutup
	func sayHello() {
		fmt.Println("hello world!")
	}

	func main() {
		for i := 0; i < 10; i++ {
			sayHello()
		}
	}