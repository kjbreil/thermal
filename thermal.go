package thermal

import "math"

type Thermal struct {
	Standard string `json:"standard"`

	Temperature            *Temperature `json:"temperature,omitempty"`
	MeanRadiantTemperature *Temperature `json:"mean_radiant_temperature,omitempty"`
	RelativeAirSpeed       *float64     `json:"relative_air_speed,omitempty"`
	RelativeHumidity       *float64     `json:"relative_humidity,omitempty"`
	MetabolicRate          *float64     `json:"metabolic_rate,omitempty"`
	ClothingInsulation     *float64     `json:"clothing_insulation,omitempty"`
	ExternalWork           *float64     `json:"external_work,omitempty"`
	AdaptiveCoefficient    *float64     `json:"adaptive_coefficient,omitempty"`
	RunningMean            *Temperature `json:"running_mean,omitempty"`
	NetRadiationAbsorbed   *float64     `json:"net_radiation_absorbed,omitempty"`
}

// NewThermal creates a new Thermal object with the specified temperature.
func NewThermal(t *Temperature) *Thermal {
	defaultExternalWork := 0.0
	return &Thermal{
		Standard:               "ISO",
		Temperature:            t,
		MeanRadiantTemperature: t,
		ExternalWork:           &defaultExternalWork,
	}
}

// Humidity adds humidity in percentage, i.e. 60 rather than .60 to the thermal object
func (t *Thermal) Humidity(h float64) *Thermal {
	t.RelativeHumidity = &h
	return t
}

func (t *Thermal) MeanRadiant(h float64) *Thermal {
	t.RelativeHumidity = &h
	return t
}

func (t *Thermal) AirSpeed(v float64) *Thermal {
	vr := v
	if t.MetabolicRate != nil && *t.MetabolicRate > 1 {
		vr = v + 0.3*(*t.MetabolicRate-1)
	}

	t.RelativeAirSpeed = &vr
	return t
}

func (t *Thermal) V(v float64) *Thermal {
	t.RelativeAirSpeed = &v
	return t
}

func (t *Thermal) Met(m float64) *Thermal {
	t.MetabolicRate = &m
	return t
}
func (t *Thermal) Clo(c float64) *Thermal {
	t.ClothingInsulation = &c
	return t
}
func (t *Thermal) ACoefficient(ac float64) *Thermal {
	t.AdaptiveCoefficient = &ac
	return t
}

func (t *Thermal) RunningMeanOutdoor(rm *Temperature) *Thermal {
	t.RunningMean = rm
	return t
}

func (t *Thermal) OperativeTemperature(ot *Temperature) *Thermal {
	t.Temperature = ot
	t.MeanRadiantTemperature = ot
	return t
}

func (t *Thermal) RadiationAbsorbed(ra float64) *Thermal {
	t.NetRadiationAbsorbed = &ra
	return t
}

func (t *Thermal) DynamicClo() *Thermal {
	clo := *t.ClothingInsulation
	if t.Standard == "ASHRAE" {
		if *t.MetabolicRate > 1.2 {
			clo = *t.ClothingInsulation * (0.6 + 0.4 / *t.MetabolicRate)
		}
	} else {
		if *t.MetabolicRate > 1.0 {
			clo = *t.ClothingInsulation * (0.6 + 0.4 / *t.MetabolicRate)
		}
	}

	clo = math.Round(clo*1000) / 1000

	t.ClothingInsulation = &clo
	return t
}

func (t *Thermal) Wme(work float64) *Thermal {
	t.ExternalWork = &work
	return t
}
