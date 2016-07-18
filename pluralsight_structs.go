package main

func main() {
  foo := newMyStruct()
  foo.myField = "bar"
  foo.printMessage()

  composed := superStruct{}
  composed.myField = "baa"
  composed.printMessage()
}

type myStruct struct {
  myField string
  myMap map[string]string
}

func (mp *myStruct) printMessage() {
  println(mp.myField)
}

func newMyStruct() *myStruct {
  result := myStruct{}
  result.myMap = map[string]string{}
  return &result
}

type superStruct struct {
  myStruct
}
