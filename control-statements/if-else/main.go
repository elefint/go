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
		fmt.Println("Paul is younger than John, or both have the same age")
	}
}

/* We begin by seeding our pseudo-random number generator. Then, we define two int variables : ageJohn and ageJohn. We assign the value rand.Intn(110) to those variables. It represents a pseudo-random number between 00 and 109109.

We print the value of those variables. Then the most interesting part is on the next line. We start an if/else structure :

The boolean expression is agePaul > ageJohn

The code that will be executed if this expression is true is fmt.Println("Paul is older than John")

If the expression is false, this code will be executed : fmt.Println("Paul is younger than John, or both have the same age") */

// Our program can have two different output depending on the value of agePaul and ageJohn. With an if-else statement, we can create two branches in the program
