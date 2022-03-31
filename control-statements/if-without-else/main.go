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
	}
	fmt.Println("End of the Program")
}

/* The first lines are identical to the program of the previous section. In short, we define two int variables that will be initialized with a random value between 0 and 109.

The boolean expression is agePaul > ageJohn. When this expression is evaluated to true, the statement fmt.Println("Paul is older than John") is executed. */

/* You can note that in the first execution, the boolean expression is evaluated to false (John is older than Paul, 32 < 64 is false). As a consequence, the statement fmt.Println("Paul is older than John") is not executed.

However, in the second execution, the expression is evaluated to true (72 < 10 is true). In that run Paul is older than John is printed to screen! */
