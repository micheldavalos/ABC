package main

import (
    "fmt"
    "sort"
)

func main() {
 ints := []int{0,0,0}
 var abc string
 
 fmt.Scanln(&ints[0], &ints[1], &ints[2])
 fmt.Scanln(&abc)
 sort.Ints(ints)
 
 for i, d := range abc {
     c := string(d)
//      fmt.Println(i, c)
  if c == "A" {
    fmt.Print(ints[0])
    } else if c == "B" {
      fmt.Print(ints[1])
    } else if c == "C" {
      fmt.Print(ints[2])
    }
  if i != 2 {
    fmt.Print(" ")
  }
 }
//  fmt.Println(ints)
//  fmt.Println(abc)    
}
