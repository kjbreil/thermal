package thermal

import "math"

// BlazejczykHeatIndex is an equation based from Steadman's 1979 table From Blazejczyk et al. 2012a
func (t *Thermal) BlazejczykHeatIndex() (*Temperature, error) {

	relativeHumidity := 50.0
	if t.RelativeHumidity != nil {
		relativeHumidity = *t.RelativeHumidity
	}

	hi := -8.784695 + 1.61139411*t.Temperature.Celsius + 2.338549*relativeHumidity - 0.14611605*t.Temperature.Celsius*relativeHumidity
	hi += -1.2308094*math.Pow(10, -2)*math.Pow(t.Temperature.Celsius, 2) - 1.6424828*math.Pow(10, -2)*math.Pow(relativeHumidity, 2)
	hi += 2.211732*math.Pow(10, -3)*math.Pow(t.Temperature.Celsius, 2)*relativeHumidity + 7.2546*math.Pow(10, -4)*t.Temperature.Celsius*math.Pow(relativeHumidity, 2)
	hi += -3.582 * math.Pow(10, -6) * math.Pow(t.Temperature.Celsius, 2) * math.Pow(relativeHumidity, 2)

	return Celsius(hi), nil
}
