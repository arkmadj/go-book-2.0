package unitconv

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FtToM(ft Foot) Meter {
	return Meter(ft / 3.2808)
}

func MToFt(m Meter) Foot {
	return Foot(m * 3.2808)
}

func PToK(p Pound) Kilogram {
	return Kilogram(p / 2.2046)
}

func KToP(k Kilogram) Pound {
	return Pound(k * 2.2046)
}
