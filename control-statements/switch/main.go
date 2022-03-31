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

	// switch expression
	switch ageJohn {
	case 10:
		fmt.Println("John is 10 years old")
	case 20:
		fmt.Println("John is 20 years old")
	case 100:
		fmt.Println("John is 100 years old")
	default:
		fmt.Println("John has neither 10, 20 nor 100 years old")
	}

	// switch statement; expression
	switch ageSum := ageJohn + agePaul; ageSum {
	case 10:
		fmt.Println("ageJohn + agePaul = 10")
	case 20, 30, 40: //*\label{switchMulti}
		fmt.Println("ageJohn + agePaul = 20 or 30 or 40")
	case 2 * agePaul:
		fmt.Println("ageJohn + agePaul = 2 times agePaul")
	}

	// switch (without statement, without expression)
	switch {
	case agePaul > ageJohn:
		fmt.Println("agePaul > ageJohn")
	case agePaul == ageJohn:
		fmt.Println("agePaul == ageJohn")
	case agePaul < ageJohn:
		fmt.Println("agePaul < ageJohn")
	}
}

/* In this example, we omitted the definition of the main function (and also the package definition). The definition of agePaul and ageJohn was also omitted. You can find the complete code in the companion code repository

We have defined a switch statement with the expression ageJohn

The program will evaluate the value of ageJohn and compare it to the value of expressions in the cases :

10``20 and 100

Here our expressions are just values.
If the value of ageJohn is not equal to 10, 20, or 100, then the default case is executed

Tehn, we have defined a switch statement with :

A first statement (which is a variable assignation) ageSum := ageJohn + agePaul

Then our expression is ageSum.

The program will first execute the statement that defines ageSum then it will evaluate ageSum and compare it to the value of the expressions define in the cases :

10``20, 30, 40 and 2 * agePaul

The second case (on line [switchMulti]) is a list of expressions.

The third case 2 * agePaul is an expression that needs to be evaluated.

Then, we have defined a switch statement without any expression.

A switch without an expression is equivalent to compare cases to the value. true. */

// The switch case is intensively used into go programs. It’s easy to read, and the syntax is also easy to learn. This is a must-have!
