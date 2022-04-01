package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const hotelName = "Gopher Paris Inn"
	const totalRooms = 4
	const firstRoomNumber = 110

	rand.Seed(time.Now().UTC().UnixNano())
	roomsOccupied := rand.Intn(totalRooms)
	roomsAvailable := totalRooms - roomsOccupied

	occupancyRate := (float64(roomsOccupied) / float64(totalRooms)) * 100
	var occupancyLevel string
	if occupancyRate > 70 {
		occupancyLevel = "High"
	} else if occupancyRate > 20 {
		occupancyLevel = "Medium"
	} else {
		occupancyLevel = "Low"
	}

	fmt.Println("Hotel Name:", hotelName)
	fmt.Println("Number of Rooms:", totalRooms)
	fmt.Println("Rooms Available:", roomsAvailable)
	fmt.Println("Occupancy level:", occupancyLevel)
	fmt.Printf("Occupany rate: %0.2f %%\n", occupancyRate)

	if roomsAvailable > 0 {
		fmt.Println("Rooms:")
		for i := 0; roomsAvailable > i; i++ {
			roomNumber := firstRoomNumber + i
			size := rand.Intn(6) + 1
			nights := rand.Intn(10) + 1

			fmt.Println(roomNumber, ":", size, "people /", nights, "nights")
		}
	} else {
		fmt.Println("No rooms available for tonight")
	}
}

// https://www.practical-go-lessons.com/chap-9-control-statements#switch-case - website
