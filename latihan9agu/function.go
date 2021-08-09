package main

import (
	"fmt"
	"strings"
)

func printMessage(message string, arr[]string) {
	var mMessage = strings.Join(arr, " ")
	fmt.Println(message, mMessage)
}

func main() {
	var names = []string {"dimas"}
	printMessage("hello", names)
}