package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var agePaul, ageJohn int = rand.Intn(110), rand.Intn(110)

	fmt.Println("John is", ageJohn, "years old")
	fmt.Println("Paul is", agePaul, "years old")

	if agePaul > ageJohn {
		fmt.Println("Paul is older than John")
	} else {
		// another if/else structure
		if agePaul == ageJohn {
			fmt.Println("Paul and John have the same age")
		} else {
			fmt.Println("Paul is younger than John")
		}
	}
}

/* Inside the else block, we created another if-else statement :

if agePaul == ageJohn {
fmt.Println("Paul and John have the same age")
} else {
fmt.Println("Paul is younger than John")
} */

// Even if this type of construct is legal, I do not recommend them. The reason is simple: it’s hard to read and to understand. You can easily avoid it by using other useful constructs. (like a switch, for example).
