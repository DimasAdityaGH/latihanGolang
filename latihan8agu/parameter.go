 package main 

 import "fmt"

func sayHelloTo (firstName string, midName string, lastName string) {
	fmt.Println("hello", firstName, midName, lastName)
}

func main() {
	sayHelloTo("esa", "dimas", "aditya")
}