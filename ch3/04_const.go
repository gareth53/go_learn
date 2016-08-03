package main

import(
  "fmt"
  "math"
)

const(
  x = math.Pi
  y = 1024/16
  z = 1024^2
)

func main() {
  // print types - const values are untyped - or implied types
  fmt.Printf("%T\n", x)
  fmt.Printf("%T\n", y)
  fmt.Printf("%T\n", z)

  // now do some operations using typed variables
  var v1 int = 3
  fmt.Print(float32(v1)/float32(x), "\n")

  fmt.Print(v1/int(x), "\n")
}
