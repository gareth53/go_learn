package main

import (
	"fmt"
	"strconv"
)

/*
The obligatory Hello World example.
run this with
> go run 01_helloworld.go

or, compile it and run the binary:

>go build 01_helloworld.go
> 01_helloword

*/

func main() {

	a := []string{"31313", "-1", "0.5", ".2", "-0.5", "t", "2"}
	for _, i := range a {
		fmt.Println(i)
		_, err := strconv.ParseUint(i, 10, 64)
		if err != nil {
			fmt.Println(err)
		}
	}
}
/*
func isPositiveInteger(val string) (int64, error) {
	value, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return value, nil
}

/*
the package main is special. It defines a stabdalone executable program, rather than a library
the function main() is also special - it;s where the execution of the program begins
*/
