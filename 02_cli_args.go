// echos cmd arguments
package main

import (
    "fmt"
    "os"
    "strings"
)

/*
to run:
> go run 02_cli_args.go foo bar fubar fugazi
*/

func method1() {
    var s, sep string
    for i :=1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = ", "
    }
    fmt.Println(s)
}

/*
note the type specification of s and sep
and that strings have a default value

the := is a short variable declaration
the type is inferred from the initial value
*/

func method2() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = ", "
    }
    fmt.Println(s)
}
/*
in method 2 we use the short var declaration to decalre two var values
we're using range too, which is simialr to Python's enumerate()
note: _ is a special variable name (a blank identifier) that cannot be used
if we had used a var name like 'temp' we would encounter a compiler error
the Go compiler HATE unused variables :)
*/

func method3() {
    fmt.Println(strings.Join(os.Args[1:], ", "))
}
/*
method3 func is the proper way of doing things
but had we jumped straight to this technique
we'd have not learned all that stuff about var declaration or blank identifiers or the use of range
*/

func main() {
    method1()
    method2()
    method3()
}
