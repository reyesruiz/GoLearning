package main

import (
	"fmt"
)

func main() {
	const freezingF, boolingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", boolingF, fToC(boolingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
