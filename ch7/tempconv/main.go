package main

import "fmt"

type Celsius float64
type Fahrenheit float64

type celsiusFlag struct {
	Celsius
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64

	fmt.Sscanf(s, "%f%s", &value, &unit)

	switch unit {
	case "C", "ºC":
		f.Celsius = Celsius(value)
		return nil
	case "F", "ºF":
		f.Celsius = FToC(Fahrenheit(value))
		return nil

	}
	return fmt.Errorf("invalid temperature %q", s)
}
