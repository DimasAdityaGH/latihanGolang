package main

import "fmt"

func main () {
	months := [...]string {"januari", "februari", "maret", "april", "mei", "juni", "juli", "agustus", "september", "oktober", "november", "desember"}
	fmt.Println(months)
	slice1 := months[4:7]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)

	slice2 := months[10:]

	slice3 := append(slice2, "januari lagi")
	fmt.Println(slice3)

	forCopy := make([]string, 2, 5)
	forCopy[0] = "dimas"
	forCopy[1] = "fifi"
	fmt.Println(len(forCopy))
	fmt.Println(cap(forCopy))
	fmt.Println(forCopy)

	thisCopy := make([]string, len(forCopy), cap(forCopy))
	copy(thisCopy, forCopy)
	fmt.Println(len(thisCopy))
	fmt.Println(cap(thisCopy))
	fmt.Println(thisCopy)
	 

}