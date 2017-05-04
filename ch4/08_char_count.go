package main

import (
  "bufio"
  "fmt"
  "os"
  "io"
  "unicode"
  "unicode/utf8"
)

func main() {
  counts := make(map[rune]int)
  var utflen [utf8.UTFMax + 1]int
  invalid := 0

  in := bufio.NewReader(os.Stdin)
  for {
    r, n, err := in.ReadRune()
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Fprintf(os.Stderr, "Error: %v\n", err)
    }
    if r == unicode.ReplacementChar && n == 1 {
      invalid++
      continue
    }
    counts[r]++
    utflen[n]++
  }
  // now print out the results
  fmt.Printf("Rune\tCount\n")
  for c, n := range counts {
    fmt.Printf("%q\t%d\n", c, n)
  }
  fmt.Printf("Length\tCount\n")
  for i, n := range utflen {
    if i > 0 {
      fmt.Printf("%q\t%d\n", i, n)
    }
  }
}
