package main

import "fmt"

func main() {
	roomNumber, floorNumber := 154, 3
	fmt.Println(roomNumber, floorNumber)
}

// Warning : short variable declaration cannot be used outside functions !
// will not compile
// package main

// vatRat := 20

// func main(){

// }

// Warning : you cannot use the value nil for a short variable declaration. The compiler cannot infer the type of your variable.
