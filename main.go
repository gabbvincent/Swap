//Name: Vincent G
//Date: 4/7/2020
//Description Swap

package main

import "fmt"

func swap(a, b int){
  //Print a and b, swap their values, then print them again.
  fmt.Println(a, b)
  var hold = a
  a = b
  b = hold
  fmt.Println("SWAP!")
  fmt.Println(a, b)
}

func main() {
  //declare variables for a and b as integers
  var a int 
  var b int
  //ask user for 2 numbers, scan for a for the first, and for b for the second.
  fmt.Println("enter a number")
  fmt.Scanln(&a)
  fmt.Println("enter in another number")
  fmt.Scanln(&b)
  //call the function swap(a, b)
   swap(a, b)
}
