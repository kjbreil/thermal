package thermal

import (
	"fmt"
	"math"
)

// Net calculates the Normal Effective Temperature from Missenard (1993) equation.
func (t *Thermal) Net() (*Temperature, error) {
	if t.RelativeHumidity == nil {
		return nil, fmt.Errorf("relative humidity needed for humidex")
	}
	v := 0.0
	frac := 1.0 / (1.76 + 1.4*math.Pow(v, 0.75))
	n := round(37 - (37-t.Temperature.Celsius)/(0.68-0.0014**t.RelativeHumidity+frac) - 0.29*t.Temperature.Celsius*(1-0.01**t.RelativeHumidity))
	return Celsius(n), nil
}
