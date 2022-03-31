package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)
	var roomsAvailable int
	rand.Seed(time.Now().UTC().UnixNano())
	var rooms, roomsOccupied int = 100, rand.Intn(100)
	roomsAvailable = rooms - roomsOccupied
	fmt.Println(roomsAvailable, "rooms available")
}

/* We seed our random number generator with time.Now().UTC().UnixNano() which is the number of elapsed milliseconds between the 1st January 1970 at midnight in UTC and now.

Note that this point in time is also called the UNIX Epoch. */
