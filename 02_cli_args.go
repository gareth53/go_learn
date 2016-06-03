// echos cmd arguments
package main

import (
    "fmt"
    "os"
)

/*
to run:
> go run 02_cli_args.go foo bar fubar fugazi
*/

func main() {
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
