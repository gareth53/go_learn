package main

import "fmt"

func remove_dupes(str string) string {
  output := str[0:1]
  prev := str[:1]
  var char string
  for i := 1; i<len(str); i++ {
    char = str[i:i+1]
    if char != prev {
      output += char
      prev = char
    }
  }
  return output
}

func main() {
  testcases := [...]string{"AA", "Aa", "BumbleBee", "Zzzzz", "Hello  Sailor!"}
  for _, str := range testcases {
    fmt.Println(remove_dupes(str))
  }
}
