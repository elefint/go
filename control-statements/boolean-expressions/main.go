package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    var rooms, roomsOccupied int = 100, rand.Intn(100)

    fmt.Println("rooms :", rooms, " roomsOccupied :", roomsOccupied)

    // Example 1: is there more rooms than roomsOccupied
    fmt.Println("boolean expression : 'rooms > roomsOccupied' is equal to :")
    fmt.Println(rooms > roomsOccupied) //*\label{condExpBool1}

    // Example 2: is there more roomsOccupied than rooms
    fmt.Println("boolean expression : 'roomsOccupied > rooms' is equal to :")
    fmt.Println(roomsOccupied > rooms) //*\label{condExpBool2}

    // Example 3: is roomsOccupied equal to rooms
    fmt.Println("boolean expression : 'roomsOccupied == rooms' is equal to :")
    fmt.Println(roomsOccupied == rooms) //*\label{condExpBool3}
}

/* In this program, we begin by calling the function Seed from the package math/randed. This function will “seed” The pseudo-random generator. When seeded, the generator will output different numbers for each run.

Then we define two int variables rooms, roomsOccupied that are initialized with respectively 100 and the value returned by rand.Intn(100) (a random int between 0 and 100).

On the line [condExpBool1] we create a boolean expression : rooms > roomsOccupied. When the program is executed, it will check if rooms is greater than roomsOccupied. If it’s the case, the value of rooms > roomsOccupied will be true. If not, the value will be false.

On the line [condExpBool2] we create a boolean expression : roomsOccupied > rooms. When the program is executed, it will check if roomsOccupied is greater than rooms. If it’s the case, the value of roomsOccupied > rooms will be true. If not, the value will be false.

In this program, we defined a third boolean expression [condExpBool3] : roomsOccupied == rooms. This expression will be true if the two variable values are equal, if not, it will be equal to false. */