package main

import "fmt"

func main() {
  salam := "world"
  
  //if expression
  //else if expression
  //else expression	
  if salam == "world"{
    fmt.Println("hello world")
  }else if salam == "you" {
    fmt.Println("hello you")
  }else {
    fmt.Println("this false")
  }

  //short statement:

  //no short statement

//   var length = len(salam)
//   if length > 5 {
// 	  fmt.Println("nama terlalu panjang!")
//   } else {
// 	  fmt.Println("nama sudah benar")
//   }

  //use short statement 
  if length := len(salam); length > 5 {
	  fmt.Println("nama terlalu panjang")
  } else {
	  fmt.Println("nama sudah benar")
  }
}