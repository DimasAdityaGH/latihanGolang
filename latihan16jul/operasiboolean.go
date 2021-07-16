package main

import "fmt"

func main() {
	var produk = "milk"
	var harga = 1000

	fmt.Println("produk:", produk)
	fmt.Println("harga:", harga)

	var uangPembeli = 10000
	var hargaProduk = 1000

	var total = hargaProduk - uangPembeli
	fmt.Println("total kembalian:", total)

	var dimas = 100
	var ani = 100

	var nilaiKKm = 90
	var hasilDimas = dimas > nilaiKKm
	var hasilani = ani > nilaiKKm

	var lulus = hasilDimas && hasilani

	fmt.Println(lulus)

	
}