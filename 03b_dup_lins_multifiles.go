package main

import (
  "os"
  "fmt"
  "bufio"
)

func main() {
  count := make(map[string]int)
  occurrence := make(map[string]string)
  if len(os.Args) < 2 {
    fmt.Println("ERROR! You must specify files to be analysed")
  }
  for _, filepath := range os.Args[1:] {
      file, _ := os.Open(filepath)
      input := bufio.NewScanner(file)
      for input.Scan() {
        count[input.Text()]++
        occurrence[input.Text()] += " " + filepath
      }
  }
  for line, n := range count {
    if n > 1 {
      fmt.Printf("%d\t%s (%s)\n", n, line, occurrence[line])
    }
  }
}
