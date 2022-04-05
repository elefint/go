package main

// compute the price of a room based on its rate per night
// and the number of night
func computePrice(rate float32, nights int) float32 {
	return rate * float32(nights)
}
