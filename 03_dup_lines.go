package main

import (
  "os"
  "fmt"
  "bufio"
)

func main() {
  input_filepath := "03_input.txt"
  count := make(map[string]int)
  if len(os.Args) > 1 {
    input_filepath = os.Args[1]
  }
  input_file, err := os.Open(input_filepath)
  if err != nil {
    fmt.Printf("Error! %s\n", err)
  }
  input := bufio.NewScanner(input_file)
  for input.Scan() {
    count[input.Text()]++
  }
  for line, n := range count {
    if n > 1 {
      fmt.Printf("%d\t%s\n", n, line)
    }
  }
}
