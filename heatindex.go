package thermal

import (
	"fmt"
	"math"
)

// HeatIndex is the NOAA heat index based off Steadman's 1979 table
// Equations are from the source code for the US National Weather http://www.wpc.ncep.noaa.gov/html/heatindex.shtml
// More variations of equations to find Steadman's can be found at https://www.ncbi.nlm.nih.gov/pmc/articles/PMC3801457/#r39
// Currently uses the NWS Fahrenheit formula
func (t *Thermal) HeatIndex() (*Temperature, error) {
	heatIndex := t.Temperature.Fahrenheit
	if t.RelativeHumidity == nil {
		return nil, fmt.Errorf("RelativeHumidity is required")
	}
	// Under 40 humidity has no effect on the Heat Index
	if heatIndex <= 40.0 {
		return Fahrenheit(heatIndex), nil
	}

	heatIndex = 61.0 + ((t.Temperature.Fahrenheit - 68.0) * 1.2) + (*t.RelativeHumidity * 0.094)

	heatIndex = 0.5 * (t.Temperature.Fahrenheit + heatIndex)

	if heatIndex > 79 {
		heatIndex = -42.379 + 2.04901523*t.Temperature.Fahrenheit
		heatIndex += 10.14333127 * *t.RelativeHumidity
		heatIndex -= 0.22475541 * t.Temperature.Fahrenheit * *t.RelativeHumidity
		heatIndex -= 0.00683783 * math.Pow(t.Temperature.Fahrenheit, 2)
		heatIndex -= 0.05481717 * math.Pow(*t.RelativeHumidity, 2)
		heatIndex += 0.00122874 * math.Pow(t.Temperature.Fahrenheit, 2) * *t.RelativeHumidity
		heatIndex += 0.00085282 * t.Temperature.Fahrenheit * math.Pow(*t.RelativeHumidity, 2)
		heatIndex -= 0.00000199 * math.Pow(t.Temperature.Fahrenheit, 2) * math.Pow(*t.RelativeHumidity, 2)

		if *t.RelativeHumidity < 13 && t.Temperature.Fahrenheit >= 80 && t.Temperature.Fahrenheit <= 112 {
			heatIndex -= ((13 - *t.RelativeHumidity) * 0.25) * math.Sqrt((17-math.Abs(t.Temperature.Fahrenheit-95))*0.05882)
		} else if *t.RelativeHumidity > 85 && t.Temperature.Fahrenheit >= 80 && t.Temperature.Fahrenheit <= 87 {
			heatIndex += ((*t.RelativeHumidity - 85) * 0.1) * ((87 - t.Temperature.Fahrenheit) * 0.2)
		}
	}

	return Fahrenheit(heatIndex), nil
}
