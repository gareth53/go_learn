package main

import "fmt"

func remove_dupes(str []string) []string {
  output := make([]string, len(str))
  output[0] = str[0]
  var test string
  offset := 0
  for i, val := range str {
    if val != test {
      output[i-offset] = val
      test = val
    } else {
      offset += 1
    }
  }
  return output
}

func main() {
  testcase := [...]string{"AA", "AA", "BB", "CC", "CC", "CC"}
  deduped := remove_dupes(testcase[:])
  fmt.Println(deduped)
}
