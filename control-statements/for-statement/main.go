package main

import "fmt"

func main() {
	const emailToSend = 3
	emailSent := 0

	for emailSent < emailToSend {
		fmt.Println("Sending email")
		emailSent++
	}
	fmt.Println("End of program")
}

/* We define a constant of type integer emailToSend. The value of the constant is set to 3.

Next, we define a variable emailSent wich is initialized with the value 0.

Then we create a for loop :

The condition of the for loop is emailSent < emailToSend

The runtime evaluates this condition.

If the evaluation result is true, then the statements

fmt.Println("sending email..")
and

emailSent++
will be executed.

When the condition is evaluated to false, the previous statements are not executed. */

/* Double-check your boolean expression to be sure that it will be false at some point in your program.

If you have a boolean expression that is always true, your program will run indefinitely.

You might have noticed, that syntactically the condition is optional. We will cover this later. */
