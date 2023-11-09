package thermal

import (
	"fmt"
	"testing"
)

func TestNewThermal(t *testing.T) {
	thermal := NewThermal(Fahrenheit(70)).
		Met(2).
		Humidity(60).
		AirSpeed(0.1).
		V(0.1).
		Clo(0.85).
		DynamicClo().
		ACoefficient(0.293).
		RunningMeanOutdoor(Fahrenheit(70))

	fmt.Println(thermal.AdaptiveThermalHeatBalanceRange(-0.15, 0.01))
	fmt.Println(thermal.AdaptiveThermalHeatBalance())
	fmt.Println(thermal.AdaptiveAshraeF())
	ehi, _ := thermal.ApparentTemperature()
	fmt.Println(ehi.Fahrenheit)
}

func TestThermal_Pmv(t1 *testing.T) {
	Precision = 3
	type args struct {
		Celsius                float64
		RelativeHumidity       float64
		MetabolicRate          float64
		RelativeAirSpeed       float64
		MeanRadiantTemperature float64
		ClothingInsulation     float64
		ExternalWork           float64
	}
	tests := []struct {
		name    string
		args    args
		wantPmv float64
	}{
		{name: "F: 65 Humidity 20 ", args: args{Celsius: 18.333, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.425},
		{name: "F: 66 Humidity 20 ", args: args{Celsius: 18.889, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.284},
		{name: "F: 67 Humidity 20 ", args: args{Celsius: 19.444, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.143},
		{name: "F: 68 Humidity 20 ", args: args{Celsius: 20.0, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.002},
		{name: "F: 69 Humidity 20 ", args: args{Celsius: 20.556, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.861},
		{name: "F: 70 Humidity 20 ", args: args{Celsius: 21.111, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.72},
		{name: "F: 71 Humidity 20 ", args: args{Celsius: 21.667, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.579},
		{name: "F: 72 Humidity 20 ", args: args{Celsius: 22.222, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.439},
		{name: "F: 73 Humidity 20 ", args: args{Celsius: 22.778, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.307},
		{name: "F: 74 Humidity 20 ", args: args{Celsius: 23.333, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.174},
		{name: "F: 75 Humidity 20 ", args: args{Celsius: 23.889, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.041},
		{name: "F: 76 Humidity 20 ", args: args{Celsius: 24.444, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.092},
		{name: "F: 77 Humidity 20 ", args: args{Celsius: 25.0, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.226},
		{name: "F: 78 Humidity 20 ", args: args{Celsius: 25.556, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.361},
		{name: "F: 79 Humidity 20 ", args: args{Celsius: 26.111, RelativeHumidity: 20, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.495},
		{name: "F: 65 Humidity 25 ", args: args{Celsius: 18.333, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.405},
		{name: "F: 66 Humidity 25 ", args: args{Celsius: 18.889, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.263},
		{name: "F: 67 Humidity 25 ", args: args{Celsius: 19.444, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.121},
		{name: "F: 68 Humidity 25 ", args: args{Celsius: 20.0, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.979},
		{name: "F: 69 Humidity 25 ", args: args{Celsius: 20.556, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.837},
		{name: "F: 70 Humidity 25 ", args: args{Celsius: 21.111, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.696},
		{name: "F: 71 Humidity 25 ", args: args{Celsius: 21.667, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.554},
		{name: "F: 72 Humidity 25 ", args: args{Celsius: 22.222, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.413},
		{name: "F: 73 Humidity 25 ", args: args{Celsius: 22.778, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.28},
		{name: "F: 74 Humidity 25 ", args: args{Celsius: 23.333, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.146},
		{name: "F: 75 Humidity 25 ", args: args{Celsius: 23.889, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.012},
		{name: "F: 76 Humidity 25 ", args: args{Celsius: 24.444, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.122},
		{name: "F: 77 Humidity 25 ", args: args{Celsius: 25.0, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.257},
		{name: "F: 78 Humidity 25 ", args: args{Celsius: 25.556, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.393},
		{name: "F: 79 Humidity 25 ", args: args{Celsius: 26.111, RelativeHumidity: 25, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.528},
		{name: "F: 65 Humidity 30 ", args: args{Celsius: 18.333, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.384},
		{name: "F: 66 Humidity 30 ", args: args{Celsius: 18.889, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.241},
		{name: "F: 67 Humidity 30 ", args: args{Celsius: 19.444, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.099},
		{name: "F: 68 Humidity 30 ", args: args{Celsius: 20.0, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.956},
		{name: "F: 69 Humidity 30 ", args: args{Celsius: 20.556, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.813},
		{name: "F: 70 Humidity 30 ", args: args{Celsius: 21.111, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.671},
		{name: "F: 71 Humidity 30 ", args: args{Celsius: 21.667, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.528},
		{name: "F: 72 Humidity 30 ", args: args{Celsius: 22.222, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.387},
		{name: "F: 73 Humidity 30 ", args: args{Celsius: 22.778, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.252},
		{name: "F: 74 Humidity 30 ", args: args{Celsius: 23.333, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.118},
		{name: "F: 75 Humidity 30 ", args: args{Celsius: 23.889, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.017},
		{name: "F: 76 Humidity 30 ", args: args{Celsius: 24.444, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.152},
		{name: "F: 77 Humidity 30 ", args: args{Celsius: 25.0, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.288},
		{name: "F: 78 Humidity 30 ", args: args{Celsius: 25.556, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.425},
		{name: "F: 79 Humidity 30 ", args: args{Celsius: 26.111, RelativeHumidity: 30, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.562},
		{name: "F: 65 Humidity 35 ", args: args{Celsius: 18.333, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.363},
		{name: "F: 66 Humidity 35 ", args: args{Celsius: 18.889, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.22},
		{name: "F: 67 Humidity 35 ", args: args{Celsius: 19.444, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.077},
		{name: "F: 68 Humidity 35 ", args: args{Celsius: 20.0, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.933},
		{name: "F: 69 Humidity 35 ", args: args{Celsius: 20.556, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.79},
		{name: "F: 70 Humidity 35 ", args: args{Celsius: 21.111, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.647},
		{name: "F: 71 Humidity 35 ", args: args{Celsius: 21.667, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.503},
		{name: "F: 72 Humidity 35 ", args: args{Celsius: 22.222, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.36},
		{name: "F: 73 Humidity 35 ", args: args{Celsius: 22.778, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.225},
		{name: "F: 74 Humidity 35 ", args: args{Celsius: 23.333, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.09},
		{name: "F: 75 Humidity 35 ", args: args{Celsius: 23.889, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.046},
		{name: "F: 76 Humidity 35 ", args: args{Celsius: 24.444, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.182},
		{name: "F: 77 Humidity 35 ", args: args{Celsius: 25.0, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.319},
		{name: "F: 78 Humidity 35 ", args: args{Celsius: 25.556, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.457},
		{name: "F: 79 Humidity 35 ", args: args{Celsius: 26.111, RelativeHumidity: 35, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.595},
		{name: "F: 65 Humidity 40 ", args: args{Celsius: 18.333, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.343},
		{name: "F: 66 Humidity 40 ", args: args{Celsius: 18.889, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.199},
		{name: "F: 67 Humidity 40 ", args: args{Celsius: 19.444, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.055},
		{name: "F: 68 Humidity 40 ", args: args{Celsius: 20.0, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.91},
		{name: "F: 69 Humidity 40 ", args: args{Celsius: 20.556, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.766},
		{name: "F: 70 Humidity 40 ", args: args{Celsius: 21.111, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.622},
		{name: "F: 71 Humidity 40 ", args: args{Celsius: 21.667, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.478},
		{name: "F: 72 Humidity 40 ", args: args{Celsius: 22.222, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.334},
		{name: "F: 73 Humidity 40 ", args: args{Celsius: 22.778, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.198},
		{name: "F: 74 Humidity 40 ", args: args{Celsius: 23.333, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.062},
		{name: "F: 75 Humidity 40 ", args: args{Celsius: 23.889, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.075},
		{name: "F: 76 Humidity 40 ", args: args{Celsius: 24.444, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.212},
		{name: "F: 77 Humidity 40 ", args: args{Celsius: 25.0, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.35},
		{name: "F: 78 Humidity 40 ", args: args{Celsius: 25.556, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.489},
		{name: "F: 79 Humidity 40 ", args: args{Celsius: 26.111, RelativeHumidity: 40, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.628},
		{name: "F: 65 Humidity 45 ", args: args{Celsius: 18.333, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.322},
		{name: "F: 66 Humidity 45 ", args: args{Celsius: 18.889, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.177},
		{name: "F: 67 Humidity 45 ", args: args{Celsius: 19.444, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.033},
		{name: "F: 68 Humidity 45 ", args: args{Celsius: 20.0, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.887},
		{name: "F: 69 Humidity 45 ", args: args{Celsius: 20.556, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.742},
		{name: "F: 70 Humidity 45 ", args: args{Celsius: 21.111, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.598},
		{name: "F: 71 Humidity 45 ", args: args{Celsius: 21.667, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.452},
		{name: "F: 72 Humidity 45 ", args: args{Celsius: 22.222, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.308},
		{name: "F: 73 Humidity 45 ", args: args{Celsius: 22.778, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.171},
		{name: "F: 74 Humidity 45 ", args: args{Celsius: 23.333, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.034},
		{name: "F: 75 Humidity 45 ", args: args{Celsius: 23.889, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.104},
		{name: "F: 76 Humidity 45 ", args: args{Celsius: 24.444, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.242},
		{name: "F: 77 Humidity 45 ", args: args{Celsius: 25.0, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.381},
		{name: "F: 78 Humidity 45 ", args: args{Celsius: 25.556, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.521},
		{name: "F: 79 Humidity 45 ", args: args{Celsius: 26.111, RelativeHumidity: 45, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.661},
		{name: "F: 65 Humidity 50 ", args: args{Celsius: 18.333, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.302},
		{name: "F: 66 Humidity 50 ", args: args{Celsius: 18.889, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.156},
		{name: "F: 67 Humidity 50 ", args: args{Celsius: 19.444, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.01},
		{name: "F: 68 Humidity 50 ", args: args{Celsius: 20.0, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.865},
		{name: "F: 69 Humidity 50 ", args: args{Celsius: 20.556, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.719},
		{name: "F: 70 Humidity 50 ", args: args{Celsius: 21.111, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.573},
		{name: "F: 71 Humidity 50 ", args: args{Celsius: 21.667, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.427},
		{name: "F: 72 Humidity 50 ", args: args{Celsius: 22.222, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.282},
		{name: "F: 73 Humidity 50 ", args: args{Celsius: 22.778, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.144},
		{name: "F: 74 Humidity 50 ", args: args{Celsius: 23.333, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.006},
		{name: "F: 75 Humidity 50 ", args: args{Celsius: 23.889, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.133},
		{name: "F: 76 Humidity 50 ", args: args{Celsius: 24.444, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.272},
		{name: "F: 77 Humidity 50 ", args: args{Celsius: 25.0, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.412},
		{name: "F: 78 Humidity 50 ", args: args{Celsius: 25.556, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.553},
		{name: "F: 79 Humidity 50 ", args: args{Celsius: 26.111, RelativeHumidity: 50, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.694},
		{name: "F: 65 Humidity 55 ", args: args{Celsius: 18.333, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.281},
		{name: "F: 66 Humidity 55 ", args: args{Celsius: 18.889, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.135},
		{name: "F: 67 Humidity 55 ", args: args{Celsius: 19.444, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.988},
		{name: "F: 68 Humidity 55 ", args: args{Celsius: 20.0, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.842},
		{name: "F: 69 Humidity 55 ", args: args{Celsius: 20.556, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.695},
		{name: "F: 70 Humidity 55 ", args: args{Celsius: 21.111, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.548},
		{name: "F: 71 Humidity 55 ", args: args{Celsius: 21.667, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.402},
		{name: "F: 72 Humidity 55 ", args: args{Celsius: 22.222, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.256},
		{name: "F: 73 Humidity 55 ", args: args{Celsius: 22.778, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.117},
		{name: "F: 74 Humidity 55 ", args: args{Celsius: 23.333, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.022},
		{name: "F: 75 Humidity 55 ", args: args{Celsius: 23.889, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.162},
		{name: "F: 76 Humidity 55 ", args: args{Celsius: 24.444, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.302},
		{name: "F: 77 Humidity 55 ", args: args{Celsius: 25.0, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.443},
		{name: "F: 78 Humidity 55 ", args: args{Celsius: 25.556, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.585},
		{name: "F: 79 Humidity 55 ", args: args{Celsius: 26.111, RelativeHumidity: 55, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.727},
		{name: "F: 65 Humidity 60 ", args: args{Celsius: 18.333, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.26},
		{name: "F: 66 Humidity 60 ", args: args{Celsius: 18.889, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.113},
		{name: "F: 67 Humidity 60 ", args: args{Celsius: 19.444, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.966},
		{name: "F: 68 Humidity 60 ", args: args{Celsius: 20.0, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.819},
		{name: "F: 69 Humidity 60 ", args: args{Celsius: 20.556, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.671},
		{name: "F: 70 Humidity 60 ", args: args{Celsius: 21.111, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.524},
		{name: "F: 71 Humidity 60 ", args: args{Celsius: 21.667, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.376},
		{name: "F: 72 Humidity 60 ", args: args{Celsius: 22.222, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.229},
		{name: "F: 73 Humidity 60 ", args: args{Celsius: 22.778, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.09},
		{name: "F: 74 Humidity 60 ", args: args{Celsius: 23.333, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.05},
		{name: "F: 75 Humidity 60 ", args: args{Celsius: 23.889, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.191},
		{name: "F: 76 Humidity 60 ", args: args{Celsius: 24.444, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.332},
		{name: "F: 77 Humidity 60 ", args: args{Celsius: 25.0, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.474},
		{name: "F: 78 Humidity 60 ", args: args{Celsius: 25.556, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.617},
		{name: "F: 79 Humidity 60 ", args: args{Celsius: 26.111, RelativeHumidity: 60, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.76},
		{name: "F: 65 Humidity 65 ", args: args{Celsius: 18.333, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.24},
		{name: "F: 66 Humidity 65 ", args: args{Celsius: 18.889, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.092},
		{name: "F: 67 Humidity 65 ", args: args{Celsius: 19.444, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.944},
		{name: "F: 68 Humidity 65 ", args: args{Celsius: 20.0, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.796},
		{name: "F: 69 Humidity 65 ", args: args{Celsius: 20.556, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.648},
		{name: "F: 70 Humidity 65 ", args: args{Celsius: 21.111, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.499},
		{name: "F: 71 Humidity 65 ", args: args{Celsius: 21.667, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.351},
		{name: "F: 72 Humidity 65 ", args: args{Celsius: 22.222, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.203},
		{name: "F: 73 Humidity 65 ", args: args{Celsius: 22.778, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.062},
		{name: "F: 74 Humidity 65 ", args: args{Celsius: 23.333, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.078},
		{name: "F: 75 Humidity 65 ", args: args{Celsius: 23.889, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.22},
		{name: "F: 76 Humidity 65 ", args: args{Celsius: 24.444, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.362},
		{name: "F: 77 Humidity 65 ", args: args{Celsius: 25.0, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.505},
		{name: "F: 78 Humidity 65 ", args: args{Celsius: 25.556, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.649},
		{name: "F: 79 Humidity 65 ", args: args{Celsius: 26.111, RelativeHumidity: 65, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.793},
		{name: "F: 65 Humidity 70 ", args: args{Celsius: 18.333, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.219},
		{name: "F: 66 Humidity 70 ", args: args{Celsius: 18.889, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.07},
		{name: "F: 67 Humidity 70 ", args: args{Celsius: 19.444, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.922},
		{name: "F: 68 Humidity 70 ", args: args{Celsius: 20.0, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.773},
		{name: "F: 69 Humidity 70 ", args: args{Celsius: 20.556, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.624},
		{name: "F: 70 Humidity 70 ", args: args{Celsius: 21.111, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.475},
		{name: "F: 71 Humidity 70 ", args: args{Celsius: 21.667, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.326},
		{name: "F: 72 Humidity 70 ", args: args{Celsius: 22.222, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.177},
		{name: "F: 73 Humidity 70 ", args: args{Celsius: 22.778, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.035},
		{name: "F: 74 Humidity 70 ", args: args{Celsius: 23.333, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.106},
		{name: "F: 75 Humidity 70 ", args: args{Celsius: 23.889, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.249},
		{name: "F: 76 Humidity 70 ", args: args{Celsius: 24.444, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.392},
		{name: "F: 77 Humidity 70 ", args: args{Celsius: 25.0, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.536},
		{name: "F: 78 Humidity 70 ", args: args{Celsius: 25.556, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.681},
		{name: "F: 79 Humidity 70 ", args: args{Celsius: 26.111, RelativeHumidity: 70, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.827},
		{name: "F: 65 Humidity 75 ", args: args{Celsius: 18.333, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.198},
		{name: "F: 66 Humidity 75 ", args: args{Celsius: 18.889, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.049},
		{name: "F: 67 Humidity 75 ", args: args{Celsius: 19.444, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.9},
		{name: "F: 68 Humidity 75 ", args: args{Celsius: 20.0, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.75},
		{name: "F: 69 Humidity 75 ", args: args{Celsius: 20.556, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.6},
		{name: "F: 70 Humidity 75 ", args: args{Celsius: 21.111, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.45},
		{name: "F: 71 Humidity 75 ", args: args{Celsius: 21.667, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.3},
		{name: "F: 72 Humidity 75 ", args: args{Celsius: 22.222, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.151},
		{name: "F: 73 Humidity 75 ", args: args{Celsius: 22.778, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.008},
		{name: "F: 74 Humidity 75 ", args: args{Celsius: 23.333, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.134},
		{name: "F: 75 Humidity 75 ", args: args{Celsius: 23.889, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.278},
		{name: "F: 76 Humidity 75 ", args: args{Celsius: 24.444, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.422},
		{name: "F: 77 Humidity 75 ", args: args{Celsius: 25.0, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.567},
		{name: "F: 78 Humidity 75 ", args: args{Celsius: 25.556, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.713},
		{name: "F: 79 Humidity 75 ", args: args{Celsius: 26.111, RelativeHumidity: 75, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.86},
		{name: "F: 65 Humidity 80 ", args: args{Celsius: 18.333, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.178},
		{name: "F: 66 Humidity 80 ", args: args{Celsius: 18.889, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 18.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -1.028},
		{name: "F: 67 Humidity 80 ", args: args{Celsius: 19.444, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 19.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.878},
		{name: "F: 68 Humidity 80 ", args: args{Celsius: 20.0, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.727},
		{name: "F: 69 Humidity 80 ", args: args{Celsius: 20.556, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 20.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.577},
		{name: "F: 70 Humidity 80 ", args: args{Celsius: 21.111, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.426},
		{name: "F: 71 Humidity 80 ", args: args{Celsius: 21.667, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 21.667, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.275},
		{name: "F: 72 Humidity 80 ", args: args{Celsius: 22.222, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.222, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: -0.124},
		{name: "F: 73 Humidity 80 ", args: args{Celsius: 22.778, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 22.778, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.019},
		{name: "F: 74 Humidity 80 ", args: args{Celsius: 23.333, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.333, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.163},
		{name: "F: 75 Humidity 80 ", args: args{Celsius: 23.889, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 23.889, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.307},
		{name: "F: 76 Humidity 80 ", args: args{Celsius: 24.444, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 24.444, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.452},
		{name: "F: 77 Humidity 80 ", args: args{Celsius: 25.0, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.0, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.598},
		{name: "F: 78 Humidity 80 ", args: args{Celsius: 25.556, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 25.556, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.745},
		{name: "F: 79 Humidity 80 ", args: args{Celsius: 26.111, RelativeHumidity: 80, MetabolicRate: 1.4, RelativeAirSpeed: 0.1, MeanRadiantTemperature: 26.111, ClothingInsulation: 0.5, ExternalWork: 0}, wantPmv: 0.893},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewThermal(Celsius(tt.args.Celsius)).
				MeanRadiant(tt.args.MeanRadiantTemperature).
				Humidity(tt.args.RelativeHumidity).
				Met(tt.args.MetabolicRate).
				V(tt.args.RelativeAirSpeed).
				Clo(tt.args.ClothingInsulation).
				Wme(tt.args.ExternalWork)

			got := t.Pmv()
			if got != tt.wantPmv {
				t1.Errorf("PmvPpd() got = %v, want %v", got, tt.wantPmv)
			}
		})
	}
}
