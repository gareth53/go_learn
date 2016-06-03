// echos cmd arguments
package main

import (
    "fmt"
    "os"
)

/*
Exercise 1.2
Modify the echo program to print the index and value of each of its arguments,
one per line.
*/


func main() {
    for index, arg := range os.Args[1:] {
        fmt.Println(index, arg)
    }
}
