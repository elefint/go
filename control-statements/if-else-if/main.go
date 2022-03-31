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
	} else if agePaul == ageJohn {
		fmt.Println("Paul and John have the same age")
	} else {
		fmt.Println("Paul is younger than John")
	}
}

/* Like previously, we define and initialize (to a random value) two int variables (agePaul and ageJohn). Then, the if / else if / else begins. The first boolean expression is agePaul > ageJohn. The second Boolean expression is agePaul == ageJohn.

If the first boolean condition is evaluated to true, the text "Paul is older than John" will be printed on the screen, then the program exits. The second boolean expression will not be evaluated.

When the first boolean expression is evaluated to false, the second boolean expression is evaluated. If the second one is true, then we print "Paul and John have the same age" on the screen.

When this second expression is false, it means that agePaul > ageJohn and agePaul == ageJohn are false. It means that agePaul < ageJohn. If it’s not clear for you, let’s express it with words :

Paul is not older than John

Paul and John have not the same age

In conclusion, Paul is younger than John. */

// As you noticed, the if / else if / else statement is verbose. I find it also complex to understand. A switch-case statement can be easier to understand.
