package thermal

// ApparentTemperature is from the Australian Bureau of Meteorology
// wind velocity is taken into account but for indoor 0.0 fits well
// Closely matches Steadman 1994 and works well for indoor temperatures
// http://www.bom.gov.au/info/thermal_stress/
// If airspeed is not set assumes still air
func (t *Thermal) ApparentTemperature() (*Temperature, error) {
	var relativeAirSpeed float64
	if t.RelativeAirSpeed != nil {
		relativeAirSpeed = *t.RelativeAirSpeed
	}

	// dividing Pressure of water vapor by 100 since the at equation requires pVap to be in hPa
	pVap := t.PressureVapor() / 100

	var at float64
	if t.NetRadiationAbsorbed != nil {
		at = t.Temperature.Celsius + 0.348*pVap - 0.7*+0.7**t.NetRadiationAbsorbed/(relativeAirSpeed+10) - 4.25
	} else {
		at = t.Temperature.Celsius + 0.33*pVap - 0.7*relativeAirSpeed - 4.00
	}
	return Celsius(at), nil
}

// PressureVapor partial pressure of water vapor in moist air
// If humidity has not been set assumes a 50% humidity
func (t *Thermal) PressureVapor() float64 {
	if t.RelativeHumidity == nil || *t.RelativeHumidity == 0.0 {
		*t.RelativeHumidity = 50.0
	}
	return *t.RelativeHumidity / 100 * t.Temperature.PressureSaturation()
}
