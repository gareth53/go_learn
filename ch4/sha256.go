package main
// write a function that counts the number of diff bits in a sha256 hash

import (
  "fmt"
  "crypto/sha256"
  )

func main(){
  c1 := sha256.Sum256([]byte("x"))
  c2 := sha256.Sum256([]byte("X"))
  fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1==c2, c1)
  matches := 0
  for i, _ := range c1 {
      if c1[i] == c2[i] {
        matches = matches + 1
        println(i, " ", c1[i])
      }
  }
  println(matches)
}
