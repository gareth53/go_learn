package main

import "fmt"

/*
The obligatory Hello World example.
run this with
> go run 01_helloworld.go

or, compile it and run the binary:

>go build 01_helloworld.go
> 01_helloword

*/

func main() {
   fmt.Println("Hello Gareth, you little 壹 壱")
}

/*
the package main is special. It defines a stabdalone executable program, rather than a library
the function main() is also special - it;s where the execution of the program begins
*/
