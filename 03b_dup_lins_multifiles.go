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
    fmt.Println("ERROR! Yuo must specify files to be analysed")
  }
  for _, filepath := range os.Args[1:] {
      file, _ := os.Open(filepath)
      input := bufio.NewScanner(file)
      for input.Scan() {
        count[input.Text()]++
        occurrence[input.Text()] += " " + filepath
      }
  }
  /*
  input_file, err := os.Open(input_filepath)
  if err != nil {
    fmt.Printf("Error! %s\n", err)
  }
  input := bufio.NewScanner(input_file)
  for input.Scan() {
    count[input.Text()]++
  }
  */
  for line, n := range count {
    if n > 1 {
      fmt.Printf("%d\t%s (%s)\n", n, line, occurrence[line])
    }
  }
}
