
package main

import (
    "encoding/json"
    "fmt"
)
type Message struct {
    Name string   `json:"name" valid:"required"`
    Body string   `json:"body" valid:"required"`
    Time int64    `json:"-" valid:"required"`
}
func main() {
    m := Message{"Alice", "Hello", 1294706395881547000}
    out, _ := json.Marshal(m)
    fmt.Println(string(out))
}
