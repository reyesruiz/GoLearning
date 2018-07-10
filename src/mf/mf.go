package main

import (
	"fmt"
	// "lengthconv"
	"github.com/reyesruiz/GoLearning/src/lengthconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		m := lengthconv.Meters(l)
		f := lengthconv.Feet(l)
		fmt.Printf("%s = %s, %s = %s\n", m, lengthconv.FeetToMeter(m), f, lengthconv.MeterToFeet(f))
	}
}
