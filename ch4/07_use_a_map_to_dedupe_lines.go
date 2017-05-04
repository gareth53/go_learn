package main
/*
uses a map to remember input strings
displays a 'DUPE' message if the text has already been seen and is
already a key in the map
*/

import (
  "os"
  "bufio"
  "fmt"
)

func main() {
  seen := make(map[string]bool)
  input := bufio.NewScanner(os.Stdin)
  for input.Scan() {
    line := input.Text()
    if !seen[line] {
      seen[line] = true
      fmt.Println(seen)
    } else {
      fmt.Println("DUPE")
    }
  }
  err := input.Err()
  if err != nil {
    fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
    os.Exit(1)
  }
}
