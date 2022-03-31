package main

import "fmt"

func main() {
	const hotelName string = "Gopher Hotel"
	const longitude = 24.806078
	const latitude = -78.243027
	const occupancy int = 12

	fmt.Println(hotelName, longitude, latitude)
	fmt.Println(occupancy)
}

/* We begin with the definition of the typed constant hotelName and the definition of longitude and latitude (which are untyped). Then we define the variable occupancy (type: int) and set its value to 12. Note that another syntax is possible :

A short variable declaration
occupancy := 12
The type can be omitted in the standard declaration :
var occupancy = 12
We can make the declaration of the variable and the assignment in two steps
var occupancy int
occupancy = 12 */
