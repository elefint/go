package main

import "fmt"

func main() {
	// maximum value of an int is 9223372036854775807
	// 9223372036854775808 (max + 1 ) overflows int
	const profit = 9223372036854775808
	var profit2 int64 = profit
	fmt.Println(profit2)
}

/* Let’s try to compile the program :

$ go build main.go
# command-line-arguments
./main.go:9:7: constant 9223372036854775808 overflows int64
We get a compilation error, and that’s fine. What we are trying to do is illegal.*/
