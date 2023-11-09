package thermal

import "math"

type Temperature struct {
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func Fahrenheit(f float64) *Temperature {
	return &Temperature{
		Fahrenheit: round(f),
		Celsius:    fahrenheitToCelsius(f),
		Kelvin:     fahrenheitToKelvin(f),
	}
}
func Celsius(c float64) *Temperature {
	return &Temperature{
		Fahrenheit: celsiusToFahrenheit(c),
		Celsius:    round(c),
		Kelvin:     celsiusToKelvin(c),
	}
}
func Kelvin(k float64) *Temperature {
	return &Temperature{
		Fahrenheit: kelvinToFahrenheit(k),
		Celsius:    kelvinToCelsius(k),
		Kelvin:     round(k),
	}
}

func MinTemperature() *Temperature {
	return Kelvin(0)
}

func MaxTemperature() *Temperature {
	return Kelvin(math.Pow(10, 32))
}

func (t *Temperature) LessThan(c *Temperature) bool {
	return t.Kelvin < c.Kelvin
}

func (t *Temperature) GreaterThan(c *Temperature) bool {
	return t.Kelvin > c.Kelvin
}

// PressureSaturation calculates vapour pressure of water at different temperatures
// Requires the air temperature only
func (t *Temperature) PressureSaturation() float64 {
	unroundedKelvin := t.Celsius + 273.15
	c1 := -5674.5359
	c2 := 6.3925247
	c3 := -0.9677843 * math.Pow(10, -2)
	c4 := 0.62215701 * math.Pow(10, -6)
	c5 := 0.20747825 * math.Pow(10, -8)
	c6 := -0.9484024 * math.Pow(10, -12)
	c7 := 4.1635019
	c8 := -5800.2206
	c9 := 1.391499321
	c10 := -0.048640239
	c11 := 0.41764768 * math.Pow(10, -4)
	c12 := -0.14452093 * math.Pow(10, -7)
	c13 := 6.5459673

	preExp := c8/unroundedKelvin + c9 + unroundedKelvin*(c10+unroundedKelvin*(c11+unroundedKelvin*c12)) + c13*math.Log(unroundedKelvin)
	pascals := math.Exp(preExp)
	if unroundedKelvin < 273.15 {
		pascals = math.Exp(c1/unroundedKelvin + c2 + unroundedKelvin*(c3+unroundedKelvin*(c4+t.Kelvin*(c5+c6*unroundedKelvin))) + c7*math.Log(unroundedKelvin))
	}

	return round(pascals)
}
