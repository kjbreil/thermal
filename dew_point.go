package thermal

import (
	"fmt"
	"math"
)

// DewPoint calculates the dew point from temperature and relative humidity
// http://wahiduddin.net/calc/density_algorithms.htm
func (t *Thermal) DewPoint() (float64, error) {
	if t.RelativeHumidity == nil {
		return 0, fmt.Errorf("relative humidity needed for dew point")
	}
	a0 := 373.15 / (273.15 + t.Temperature.Celsius)

	sum := -7.90298 * (a0 - 1)
	sum += 5.02808 * math.Log10(a0)
	sum += -1.3816e-7 * (math.Pow(10, 11.344*(1-1/a0)) - 1)
	sum += 8.1328e-3 * (math.Pow(10, -3.49149*(a0-1)) - 1)
	sum += math.Log10(1013.246)
	vp := math.Pow(10, sum-3) * *t.RelativeHumidity
	td := math.Log(vp / 0.61078)
	td = (241.88 * td) / (17.558 - td)

	return td, nil
}

func (t *Thermal) DewPointF() (float64, error) {
	d, err := t.DewPoint()
	return celsiusToFahrenheit(d), err
}
