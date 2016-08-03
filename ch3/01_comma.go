package main

import (
  "bytes"
)

func main() {
  tests := [10]string{"1", "12", "123", "1234","12345", "123456", "1234567", "12345678", "123456789", "1234567890"}
  for _, test := range tests {
    println(comma(test))
    println(" vs ")
    println(BufferComma(test))
    println("\n")

  }
}

func comma(s string) string {
  n := len(s)
  if n < 3 {
    return s
  }
  return comma(s[:n-3]) + "," + s[n-3:]
}

func BufferComma(s string) string {
  var buf bytes.Buffer
  limit := len(s)
  for i, v := range s {
    // how many figures from end of the number string?
    r := limit - i
    if r%3 == 0 && r < limit && r > 0 {
      buf.WriteRune(',')
    }
    buf.WriteRune(v)
  }
  return buf.String()
}
