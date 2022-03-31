package main

import "fmt"

func main() {
	const occupancyLimit = 12

	var occupancyLimit4 string

	// Uncomment line below for error!
	// occupancyLimit4 = occupancyLimit

	fmt.Println(occupancyLimit4)
}

/* In this program, we define a constant occupancyLimit which has a value of 12 (an integer). We define a variable occupancyLimit4 with a string type. Then we try to assign to occupancyLimit4 the value of our constant.

We try to convert an integer to a string. Will this program compile? The answer is no! The compilation error is :

./main.go:10:19: cannot use occupancyLimit (type int) as type string in assignment
An untyped constant has a default type that is defined by the value assigned to it at compilation. In our example, occupancyLimit has a default type of int . An int cannot be assigned to a string variable.

Untyped constants default types are :

bool (for any boolean value)

rune (for any rune value)

int (for any integer value)

float64 (for any floating-point value)

complex128 (for any complex value)

string (for any string value) */
