package main 

import "fmt" 

func main () {
		//break and continue adalah kata kunci yang bisa digunakan dalan perulanagan
	//break digunakan untuk menghentikan seluruh perulangan 
	//continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung melanjutkan ke perulangan selanjutnya

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("perulangan ke",i)
	}
}