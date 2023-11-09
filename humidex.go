package thermal

import (
	"fmt"
	"math"
)

// Humidex Calculates the humidex (short for "humidity index"). It has been  developed by the Canadian Meteorological
// service. It was introduced in 1965 and then it was revised by Masterson and Richardson (1979). . It aims  to
// describe how hot, humid weather is felt by the average person. The Humidex differs from the heat index in being
// related to the dew point rather than relative humidity.
func (t *Thermal) Humidex() (*Temperature, error) {
	if t.RelativeHumidity == nil {
		return nil, fmt.Errorf("relative humidity needed for humidex")
	}

	hd := t.Temperature.Celsius + float64(5)/float64(9)*(6.112*math.Pow(10, 7.5*float64(t.Temperature.Celsius)/(237.7+float64(t.Temperature.Celsius)))**t.RelativeHumidity/float64(100)-10)

	return Celsius(hd), nil
}
