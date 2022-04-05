package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed((time.Now().UTC().UnixNano()))
	var ageJohn int = rand.Intn(110)
	fmt.Println("John is", ageJohn, "years old")

	for i := 0; i < ageJohn; i++ {
		fmt.Println("interation N", i)
	}
}

/* We define the variable ageJohn and we initialize it to a random value between 0 and 109. After printing its value with the statement fmt.Println("John is", ageJohn, "years old") we add a for statement. This for statement is composed of the following elements : i := 0 this is the init statement. It will be executed at the beginning of the for loop

We create and initialize a variable i

This is our counter. (Usually, counters are named: i,j or k)
i < ageJohn is the condition.

This is a boolean expression, which means that it will be evaluated to true or false at runtime

This boolean expression will be evaluated at each loop.

If it returns false, the loop stops.

i++ is the post statement.

It will be executed after each loop.

i++ is a shorthand for i = i +1

We increment the value of i; in other words, we add 1 to i.
fmt.Println("interation N", i) will be executed at each loop. The program will print “interation nX” ageJohn times. If ageJohn is 3, then we will print on the screen :

interation n 0

iteration n 1

iteration n 2

We printed it three times. The value of i is 0, then 1, and then 2. */

/* The first value of i is 0 and not 1; remember that our init statement is i := 0.

If the init statement is i := 1 The first value of i is 1. The first sentence printed on screen would be interation N° 1

This is often deconcerting for beginners. In the real world, when we start an enumeration with the value one and not 0!

This is (often) not the case in computer science. Indexes usually start with the value 0.

For loops with range clauses are used intensively. */
