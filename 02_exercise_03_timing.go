// echos cmd arguments
package main

import (
    "fmt"
    "os"
    "strings"
    "time"
)
/*
Measure the difference in running time between our potentially
inefficient versions and the one
that uses strings.Join
*/
func method1() {
    var s, sep string
    for i :=1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = ", "
    }
    fmt.Println(s)
}

func method2() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = ", "
    }
    fmt.Println(s)
}

func method3() {
    fmt.Println(strings.Join(os.Args[1:], ", "))
}

func main() {
    var start = time.Now()
    method1()
    fmt.Println(time.Since(start).Seconds())
    start = time.Now()
    method2()
    fmt.Println(time.Since(start).Seconds())
    start = time.Now()
    method3()
    fmt.Println(time.Since(start).Seconds())
}
