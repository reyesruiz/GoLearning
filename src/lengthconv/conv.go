package lengthconv

func FeetToMeter(feet Meters) Feet {
	return Feet(feet / 0.3048)
}

func MeterToFeet(meter Feet) Meters {
	return Meters(meter * 0.3048)
}
