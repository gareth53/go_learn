// echos cmd arguments
package main

import (
    "fmt"
    "os"
    "strings"
)

/*
Exercise 1.1: MOdify th echo program to also print os.Args[0], the 
name of the command that invoked it
*/
func main() {
    fmt.Println(os.Args[0])
    fmt.Println(strings.Join(os.Args[1:], ", "))
}
