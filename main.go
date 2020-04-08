//Name: Vincent G
//Date: 4/7/2020
//Description Swap

package main

import "fmt"

func swap(a, b int){
  fmt.Println(a, b)
  var hold = a
  a = b
  b = hold
  fmt.Println("SWAP!")
  fmt.Println(a, b)
}

func main() {
  var a int 
  var b int
  fmt.Println("enter a number")
  fmt.Scanln(&a)
  fmt.Println("enter in another number")
  fmt.Scanln(&b)
   swap(a, b)
}