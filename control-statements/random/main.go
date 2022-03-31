package main

import (
	"fmt"
	"math/rand"
)

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)
	var roomsAvailable int
	var rooms, roomsOccupied int = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupied
	fmt.Println(roomsAvailable, "rooms available")
}

/* The result is strange. I’m sure you expected another value than 19! It seems that the random number is fixed. (you can try to execute it 3,4, 5 times.
You will still get the same value).To work properly, the random library needs to be seeded.
The seed is a number that is used to initialize a pseudo-random generator. */
