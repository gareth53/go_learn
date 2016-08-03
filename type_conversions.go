package main

import (
  "fmt"
  "strconv"
)

const(
  float1 float32 = 1.1
  float2 float64 = 20000.1
  int1 int32 = 10
  int2 int64 = 9999999
  stringFloat string = "3.14"
  stringInt string = "100000"
  stringIntComplex string = "1,000,000"
)

// numbers to strings

func num_to_str(){
  fmt.Printf("float1 = %.1f\n", float1)
  fmt.Printf("float2 = %.1f\n", float2)
  println(strconv.FormatFloat(float2, 'f', 1, 32))
  fmt.Printf("int1 = %d\n", int1)
  fmt.Printf("int2 = %d\n", int2)
}

func str_to_int() {
  output1, _ := strconv.Atoi(stringInt)
  output2, _ := strconv.Atoi(stringIntComplex)
  fmt.Printf("output1, type=%T\n", output1)
  println(output1)
  fmt.Printf("output2, type=%T\n", output2)
  println(output2)
}

func main() {
  num_to_str()
  str_to_int()
  println("FINISH")
}
