package main
/*
enhance comma so that it deals correctly with gloating point numbers and an optional sign
*/
import (
  "bytes"
  "strings"
  "strconv"
)

func main() {
  tests := [6]string{"123456", "-12345678", "+123456789", "12345.25", "-1234567890.2", "+1234567890.2345"}
  for _, test := range tests {
    println(comma(test))
  }
}

func comma(s string) string {
  var buf bytes.Buffer
  // strip (and store for later) any +/- prefix
  char1 := s[0:1]
  if char1 == "+" || char1 == "-" {
    s = s[1:]
  } else {
    char1 = ""
  }
  // store fractional value
  intVal, _ := strconv.ParseInt(s, 0, 64)
  floatVal, _ := strconv.ParseFloat(s, 64)
  frac := floatVal - float64(intVal)
  var fracStr string
  if frac > 0 {
    decpnt := strings.LastIndex(s, ".")
    fracStr = s[decpnt+1:]
    s = s[0:decpnt]
  }
  limit := len(s)
  buf.WriteString(char1)
  for i, v := range s {
    // how many figures from end of the number string?
    r := limit - i
    if r%3 == 0 && r < limit && r > 0 {
      buf.WriteRune(',')
    }
    buf.WriteRune(v)
  }
  if fracStr != "" {
    buf.WriteString(".")
    buf.WriteString(fracStr)
  }
  return buf.String()
}
