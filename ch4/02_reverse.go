package main

import "fmt"

func reverse(s []int) {
  for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
    s[i], s[j] = s[j], s[i]
  }
}


func reversePointer(s *List) string {
  return "s"
}

func main() {
  input := [...]int{1, 2, 3, 4, 5}
  fmt.Printf("%T\n", input)
  reverse(input[:])
  fmt.Println(input)

  fmt.Println(reversePointer(input))
}
