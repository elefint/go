package main

import "fmt"

func main() {
	hotelName := "The Gopher Hotel"
	fmt.Println("Hotel " + hotelName)
	var roomsAvailable uint8
	var rooms, roomsOccupied uint8 = 100, 10
	roomsAvailable = rooms - roomsOccupied //*\label{cond1diff}
	fmt.Println(roomsAvailable)
}

/* In the previous listing, we added the declaration of 3 variables (all with the uint8 type)

roomsAvailable : the number of rooms that can be rent

rooms : the total number of rooms in the hotel

roomsOccupied : the number of rooms currently occupied

Note that roomsAvailable was declared and not initialized whereas rooms and roomsOccupied were declared and initialized at the same time.

Then we assign to roomsAvailable the value of rooms - roomsOccupied on line [cond1diff]. */
