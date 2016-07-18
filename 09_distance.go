// CLI to convert CM to Feet & Inches. sample usage: go run 09_distange.go 48
package main

import (
  "os"
  "fmt"
  "strconv"
)

type Inch float64
type Centimetre float64

func(i Inch) String() string {
  inches := uint(i) % 12
  feet := (uint(i) - inches) / 12
  if feet > 0 {
    return fmt.Sprintf("%d'%d\"", feet, inches)
  }
  return fmt.Sprintf("%d\"", inches)
}

func(c Centimetre) String() string {
  cm := uint(c) % 100
  m := (uint(c) - cm) / 100.0
  if m > 0 {
    return fmt.Sprintf("%.2fm", c/100)
  }
  return fmt.Sprintf("%dcm", uint(c))
}

func CmToI(c Centimetre) Inch {
  return Inch(c/2.49)
}

func main() {
  num, err := strconv.Atoi(os.Args[1])
  if err != nil {
    fmt.Println(err)
    return
  }
  cm := Centimetre(num)
  fmt.Println(cm)
  fmt.Println(CmToI(cm))
}
