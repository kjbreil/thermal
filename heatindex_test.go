package thermal

import (
	"testing"
)

func TestThermal_ApparentTemperatureModels(t1 *testing.T) {
	Precision = 2
	type args struct {
		Celsius              float64
		Humidity             float64
		WindSpeed            float64
		NetRadiationAbsorbed float64
	}

	type want struct {
		ApparentTemperature float64
		HeatIndex           float64
		Blazejczy           float64
		Humidex             float64
		Net                 float64
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{

		{
			name: "30% Humidity",
			args: args{
				Celsius:              21,
				Humidity:             30,
				WindSpeed:            0,
				NetRadiationAbsorbed: 0,
			},
			want: want{
				ApparentTemperature: 19.46,
				HeatIndex:           19.94,
				Blazejczy:           24.5,
				Humidex:             19.58,
				Net:                 19.47,
			},
			wantErr: false,
		},
		{
			name: "40% Humidity",
			args: args{
				Celsius:              21,
				Humidity:             40,
				WindSpeed:            0,
				NetRadiationAbsorbed: 0,
			},
			want: want{
				ApparentTemperature: 20.28,
				HeatIndex:           20.2,
				Blazejczy:           25.01,
				Humidex:             20.96,
				Net:                 19.93,
			},
			wantErr: false,
		},
		{
			name: "50% Humidity",
			args: args{
				Celsius:              21,
				Humidity:             50,
				WindSpeed:            0,
				NetRadiationAbsorbed: 0,
			},
			want: want{
				ApparentTemperature: 21.1,
				HeatIndex:           20.46,
				Blazejczy:           24.98,
				Humidex:             22.34,
				Net:                 20.37,
			},
			wantErr: false,
		},
		{
			name: "60% Humidity",
			args: args{
				Celsius:              21,
				Humidity:             60,
				WindSpeed:            0,
				NetRadiationAbsorbed: 0,
			},
			want: want{
				ApparentTemperature: 21.93,
				HeatIndex:           20.72,
				Blazejczy:           24.38,
				Humidex:             23.72,
				Net:                 20.82,
			},
			wantErr: false,
		},

		{
			name: "70% Humidity",
			args: args{
				Celsius:              21,
				Humidity:             70,
				WindSpeed:            0,
				NetRadiationAbsorbed: 0,
			},
			want: want{
				ApparentTemperature: 22.75,
				HeatIndex:           20.98,
				Blazejczy:           23.24,
				Humidex:             25.1,
				Net:                 21.26,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewThermal(Celsius(tt.args.Celsius)).Humidity(tt.args.Humidity).AirSpeed(tt.args.WindSpeed)
			if tt.args.NetRadiationAbsorbed != 0 {
				t.RadiationAbsorbed(tt.args.NetRadiationAbsorbed)
			}

			got, err := t.ApparentTemperature()
			if (err != nil) != tt.wantErr {
				t1.Errorf("ApparentTemperature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.ApparentTemperature {
				t1.Errorf("ApparentTemperature() got = %v, want %v", got.Celsius, tt.want.ApparentTemperature)
			}

			got, err = t.HeatIndex()
			if (err != nil) != tt.wantErr {
				t1.Errorf("HeatIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.HeatIndex {
				t1.Errorf("HeatIndex() got = %v, want %v", got.Celsius, tt.want.HeatIndex)
			}

			got, err = t.BlazejczykHeatIndex()
			if (err != nil) != tt.wantErr {
				t1.Errorf("BlazejczykHeatIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.Blazejczy {
				t1.Errorf("BlazejczykHeatIndex() got = %v, want %v", got.Celsius, tt.want.Blazejczy)
			}

			got, err = t.Humidex()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Humidex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.Humidex {
				t1.Errorf("Humidex() got = %v, want %v", got.Celsius, tt.want.Humidex)
			}

			got, err = t.Net()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Net() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.Net {
				t1.Errorf("Net() got = %v, want %v", got.Celsius, tt.want.Net)
			}

		})
	}
}

func TestThermal_PyThermalComfortModels(t1 *testing.T) {
	Precision = 3
	type args struct {
		Celsius              float64
		Humidity             float64
		WindSpeed            float64
		NetRadiationAbsorbed float64
	}

	type want struct {
		ApparentTemperature float64
		Blazejczy           float64
		Humidex             float64
		Net                 float64
	}
	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{name: "Temp 60 20% Humidity", args: args{Celsius: 15.556, Humidity: 20, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 12.723, Blazejczy: 22.917, Humidex: 11.962, Net: 15.817}, wantErr: false},
		{name: "Temp 60 25% Humidity", args: args{Celsius: 15.556, Humidity: 25, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 13.014, Blazejczy: 24.569, Humidex: 12.453, Net: 15.941}, wantErr: false},
		{name: "Temp 60 30% Humidity", args: args{Celsius: 15.556, Humidity: 30, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 13.306, Blazejczy: 25.921, Humidex: 12.943, Net: 16.064}, wantErr: false},
		{name: "Temp 60 35% Humidity", args: args{Celsius: 15.556, Humidity: 35, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 13.597, Blazejczy: 26.973, Humidex: 13.433, Net: 16.186}, wantErr: false},
		{name: "Temp 60 40% Humidity", args: args{Celsius: 15.556, Humidity: 40, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 13.889, Blazejczy: 27.725, Humidex: 13.924, Net: 16.306}, wantErr: false},
		{name: "Temp 60 45% Humidity", args: args{Celsius: 15.556, Humidity: 45, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 14.181, Blazejczy: 28.176, Humidex: 14.414, Net: 16.425}, wantErr: false},
		{name: "Temp 60 50% Humidity", args: args{Celsius: 15.556, Humidity: 50, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 14.472, Blazejczy: 28.327, Humidex: 14.905, Net: 16.543}, wantErr: false},
		{name: "Temp 60 55% Humidity", args: args{Celsius: 15.556, Humidity: 55, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 14.764, Blazejczy: 28.177, Humidex: 15.395, Net: 16.66}, wantErr: false},
		{name: "Temp 60 60% Humidity", args: args{Celsius: 15.556, Humidity: 60, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 15.056, Blazejczy: 27.728, Humidex: 15.885, Net: 16.776}, wantErr: false},
		{name: "Temp 60 65% Humidity", args: args{Celsius: 15.556, Humidity: 65, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 15.347, Blazejczy: 26.978, Humidex: 16.376, Net: 16.89}, wantErr: false},
		{name: "Temp 60 70% Humidity", args: args{Celsius: 15.556, Humidity: 70, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 15.639, Blazejczy: 25.927, Humidex: 16.866, Net: 17.003}, wantErr: false},
		{name: "Temp 60 75% Humidity", args: args{Celsius: 15.556, Humidity: 75, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 15.931, Blazejczy: 24.577, Humidex: 17.357, Net: 17.114}, wantErr: false},
		{name: "Temp 60 80% Humidity", args: args{Celsius: 15.556, Humidity: 80, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 16.222, Blazejczy: 22.926, Humidex: 17.847, Net: 17.224}, wantErr: false},
		{name: "Temp 65 20% Humidity", args: args{Celsius: 18.333, Humidity: 20, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 15.724, Blazejczy: 22.952, Humidex: 15.116, Net: 17.448}, wantErr: false},
		{name: "Temp 65 25% Humidity", args: args{Celsius: 18.333, Humidity: 25, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 16.072, Blazejczy: 23.994, Humidex: 15.701, Net: 17.626}, wantErr: false},
		{name: "Temp 65 30% Humidity", args: args{Celsius: 18.333, Humidity: 30, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 16.42, Blazejczy: 24.819, Humidex: 16.285, Net: 17.802}, wantErr: false},
		{name: "Temp 65 35% Humidity", args: args{Celsius: 18.333, Humidity: 35, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 16.768, Blazejczy: 25.428, Humidex: 16.87, Net: 17.978}, wantErr: false},
		{name: "Temp 65 40% Humidity", args: args{Celsius: 18.333, Humidity: 40, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 17.115, Blazejczy: 25.821, Humidex: 17.455, Net: 18.152}, wantErr: false},
		{name: "Temp 65 45% Humidity", args: args{Celsius: 18.333, Humidity: 45, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 17.463, Blazejczy: 25.997, Humidex: 18.039, Net: 18.326}, wantErr: false},
		{name: "Temp 65 50% Humidity", args: args{Celsius: 18.333, Humidity: 50, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 17.811, Blazejczy: 25.956, Humidex: 18.624, Net: 18.498}, wantErr: false},
		{name: "Temp 65 55% Humidity", args: args{Celsius: 18.333, Humidity: 55, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 18.159, Blazejczy: 25.699, Humidex: 19.209, Net: 18.669}, wantErr: false},
		{name: "Temp 65 60% Humidity", args: args{Celsius: 18.333, Humidity: 60, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 18.507, Blazejczy: 25.226, Humidex: 19.793, Net: 18.839}, wantErr: false},
		{name: "Temp 65 65% Humidity", args: args{Celsius: 18.333, Humidity: 65, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 18.855, Blazejczy: 24.536, Humidex: 20.378, Net: 19.008}, wantErr: false},
		{name: "Temp 65 70% Humidity", args: args{Celsius: 18.333, Humidity: 70, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 19.202, Blazejczy: 23.63, Humidex: 20.963, Net: 19.175}, wantErr: false},
		{name: "Temp 65 75% Humidity", args: args{Celsius: 18.333, Humidity: 75, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 19.55, Blazejczy: 22.508, Humidex: 21.548, Net: 19.342}, wantErr: false},
		{name: "Temp 65 80% Humidity", args: args{Celsius: 18.333, Humidity: 80, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 19.898, Blazejczy: 21.169, Humidex: 22.132, Net: 19.507}, wantErr: false},
		{name: "Temp 70 20% Humidity", args: args{Celsius: 21.111, Humidity: 20, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 18.764, Blazejczy: 23.458, Humidex: 18.333, Net: 19.08}, wantErr: false},
		{name: "Temp 70 25% Humidity", args: args{Celsius: 21.111, Humidity: 25, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 19.177, Blazejczy: 24.047, Humidex: 19.028, Net: 19.311}, wantErr: false},
		{name: "Temp 70 30% Humidity", args: args{Celsius: 21.111, Humidity: 30, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 19.591, Blazejczy: 24.501, Humidex: 19.722, Net: 19.541}, wantErr: false},
		{name: "Temp 70 35% Humidity", args: args{Celsius: 21.111, Humidity: 35, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 20.004, Blazejczy: 24.819, Humidex: 20.417, Net: 19.771}, wantErr: false},
		{name: "Temp 70 40% Humidity", args: args{Celsius: 21.111, Humidity: 40, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 20.417, Blazejczy: 25.003, Humidex: 21.111, Net: 19.999}, wantErr: false},
		{name: "Temp 70 45% Humidity", args: args{Celsius: 21.111, Humidity: 45, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 20.83, Blazejczy: 25.051, Humidex: 21.806, Net: 20.226}, wantErr: false},
		{name: "Temp 70 50% Humidity", args: args{Celsius: 21.111, Humidity: 50, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 21.244, Blazejczy: 24.963, Humidex: 22.5, Net: 20.453}, wantErr: false},
		{name: "Temp 70 55% Humidity", args: args{Celsius: 21.111, Humidity: 55, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 21.657, Blazejczy: 24.741, Humidex: 23.195, Net: 20.678}, wantErr: false},
		{name: "Temp 70 60% Humidity", args: args{Celsius: 21.111, Humidity: 60, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 22.07, Blazejczy: 24.383, Humidex: 23.889, Net: 20.903}, wantErr: false},
		{name: "Temp 70 65% Humidity", args: args{Celsius: 21.111, Humidity: 65, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 22.484, Blazejczy: 23.889, Humidex: 24.583, Net: 21.126}, wantErr: false},
		{name: "Temp 70 70% Humidity", args: args{Celsius: 21.111, Humidity: 70, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 22.897, Blazejczy: 23.261, Humidex: 25.278, Net: 21.349}, wantErr: false},
		{name: "Temp 70 75% Humidity", args: args{Celsius: 21.111, Humidity: 75, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 23.31, Blazejczy: 22.497, Humidex: 25.972, Net: 21.571}, wantErr: false},
		{name: "Temp 70 80% Humidity", args: args{Celsius: 21.111, Humidity: 80, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 23.723, Blazejczy: 21.598, Humidex: 26.667, Net: 21.791}, wantErr: false},
		{name: "Temp 75 20% Humidity", args: args{Celsius: 23.889, Humidity: 20, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 21.846, Blazejczy: 24.434, Humidex: 21.621, Net: 20.713}, wantErr: false},
		{name: "Temp 75 25% Humidity", args: args{Celsius: 23.889, Humidity: 25, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 22.335, Blazejczy: 24.729, Humidex: 22.443, Net: 20.997}, wantErr: false},
		{name: "Temp 75 30% Humidity", args: args{Celsius: 23.889, Humidity: 30, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 22.825, Blazejczy: 24.967, Humidex: 23.265, Net: 21.281}, wantErr: false},
		{name: "Temp 75 35% Humidity", args: args{Celsius: 23.889, Humidity: 35, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 23.314, Blazejczy: 25.148, Humidex: 24.087, Net: 21.564}, wantErr: false},
		{name: "Temp 75 40% Humidity", args: args{Celsius: 23.889, Humidity: 40, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 23.803, Blazejczy: 25.272, Humidex: 24.908, Net: 21.846}, wantErr: false},
		{name: "Temp 75 45% Humidity", args: args{Celsius: 23.889, Humidity: 45, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 24.292, Blazejczy: 25.339, Humidex: 25.73, Net: 22.127}, wantErr: false},
		{name: "Temp 75 50% Humidity", args: args{Celsius: 23.889, Humidity: 50, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 24.782, Blazejczy: 25.349, Humidex: 26.552, Net: 22.408}, wantErr: false},
		{name: "Temp 75 55% Humidity", args: args{Celsius: 23.889, Humidity: 55, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 25.271, Blazejczy: 25.302, Humidex: 27.374, Net: 22.688}, wantErr: false},
		{name: "Temp 75 60% Humidity", args: args{Celsius: 23.889, Humidity: 60, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 25.76, Blazejczy: 25.198, Humidex: 28.196, Net: 22.967}, wantErr: false},
		{name: "Temp 75 65% Humidity", args: args{Celsius: 23.889, Humidity: 65, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 26.25, Blazejczy: 25.038, Humidex: 29.018, Net: 23.245}, wantErr: false},
		{name: "Temp 75 70% Humidity", args: args{Celsius: 23.889, Humidity: 70, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 26.739, Blazejczy: 24.82, Humidex: 29.84, Net: 23.523}, wantErr: false},
		{name: "Temp 75 75% Humidity", args: args{Celsius: 23.889, Humidity: 75, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 27.228, Blazejczy: 24.546, Humidex: 30.661, Net: 23.799}, wantErr: false},
		{name: "Temp 75 80% Humidity", args: args{Celsius: 23.889, Humidity: 80, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 27.717, Blazejczy: 24.214, Humidex: 31.483, Net: 24.075}, wantErr: false},
		{name: "Temp 80 20% Humidity", args: args{Celsius: 26.667, Humidity: 20, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 24.976, Blazejczy: 25.881, Humidex: 24.988, Net: 22.345}, wantErr: false},
		{name: "Temp 80 25% Humidity", args: args{Celsius: 26.667, Humidity: 25, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 25.553, Blazejczy: 26.04, Humidex: 25.957, Net: 22.683}, wantErr: false},
		{name: "Temp 80 30% Humidity", args: args{Celsius: 26.667, Humidity: 30, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 26.13, Blazejczy: 26.217, Humidex: 26.927, Net: 23.02}, wantErr: false},
		{name: "Temp 80 35% Humidity", args: args{Celsius: 26.667, Humidity: 35, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 26.707, Blazejczy: 26.413, Humidex: 27.896, Net: 23.357}, wantErr: false},
		{name: "Temp 80 40% Humidity", args: args{Celsius: 26.667, Humidity: 40, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 27.285, Blazejczy: 26.628, Humidex: 28.865, Net: 23.693}, wantErr: false},
		{name: "Temp 80 45% Humidity", args: args{Celsius: 26.667, Humidity: 45, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 27.862, Blazejczy: 26.861, Humidex: 29.834, Net: 24.028}, wantErr: false},
		{name: "Temp 80 50% Humidity", args: args{Celsius: 26.667, Humidity: 50, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 28.439, Blazejczy: 27.113, Humidex: 30.803, Net: 24.363}, wantErr: false},
		{name: "Temp 80 55% Humidity", args: args{Celsius: 26.667, Humidity: 55, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 29.016, Blazejczy: 27.384, Humidex: 31.773, Net: 24.697}, wantErr: false},
		{name: "Temp 80 60% Humidity", args: args{Celsius: 26.667, Humidity: 60, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 29.593, Blazejczy: 27.673, Humidex: 32.742, Net: 25.031}, wantErr: false},
		{name: "Temp 80 65% Humidity", args: args{Celsius: 26.667, Humidity: 65, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 30.171, Blazejczy: 27.981, Humidex: 33.711, Net: 25.364}, wantErr: false},
		{name: "Temp 80 70% Humidity", args: args{Celsius: 26.667, Humidity: 70, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 30.748, Blazejczy: 28.308, Humidex: 34.68, Net: 25.696}, wantErr: false},
		{name: "Temp 80 75% Humidity", args: args{Celsius: 26.667, Humidity: 75, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 31.325, Blazejczy: 28.653, Humidex: 35.649, Net: 26.028}, wantErr: false},
		{name: "Temp 80 80% Humidity", args: args{Celsius: 26.667, Humidity: 80, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 31.902, Blazejczy: 29.018, Humidex: 36.619, Net: 26.359}, wantErr: false},
		{name: "Temp 85 20% Humidity", args: args{Celsius: 29.444, Humidity: 20, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 28.158, Blazejczy: 27.798, Humidex: 28.444, Net: 23.976}, wantErr: false},
		{name: "Temp 85 25% Humidity", args: args{Celsius: 29.444, Humidity: 25, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 28.837, Blazejczy: 27.979, Humidex: 29.583, Net: 24.368}, wantErr: false},
		{name: "Temp 85 30% Humidity", args: args{Celsius: 29.444, Humidity: 30, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 29.515, Blazejczy: 28.251, Humidex: 30.722, Net: 24.758}, wantErr: false},
		{name: "Temp 85 35% Humidity", args: args{Celsius: 29.444, Humidity: 35, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 30.194, Blazejczy: 28.615, Humidex: 31.861, Net: 25.149}, wantErr: false},
		{name: "Temp 85 40% Humidity", args: args{Celsius: 29.444, Humidity: 40, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 30.872, Blazejczy: 29.07, Humidex: 33.0, Net: 25.539}, wantErr: false},
		{name: "Temp 85 45% Humidity", args: args{Celsius: 29.444, Humidity: 45, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 31.551, Blazejczy: 29.616, Humidex: 34.139, Net: 25.928}, wantErr: false},
		{name: "Temp 85 50% Humidity", args: args{Celsius: 29.444, Humidity: 50, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 32.229, Blazejczy: 30.255, Humidex: 35.278, Net: 26.317}, wantErr: false},
		{name: "Temp 85 55% Humidity", args: args{Celsius: 29.444, Humidity: 55, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 32.908, Blazejczy: 30.984, Humidex: 36.417, Net: 26.706}, wantErr: false},
		{name: "Temp 85 60% Humidity", args: args{Celsius: 29.444, Humidity: 60, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 33.587, Blazejczy: 31.805, Humidex: 37.556, Net: 27.094}, wantErr: false},
		{name: "Temp 85 65% Humidity", args: args{Celsius: 29.444, Humidity: 65, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 34.265, Blazejczy: 32.718, Humidex: 38.695, Net: 27.482}, wantErr: false},
		{name: "Temp 85 70% Humidity", args: args{Celsius: 29.444, Humidity: 70, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 34.944, Blazejczy: 33.722, Humidex: 39.834, Net: 27.869}, wantErr: false},
		{name: "Temp 85 75% Humidity", args: args{Celsius: 29.444, Humidity: 75, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 35.622, Blazejczy: 34.818, Humidex: 40.973, Net: 28.256}, wantErr: false},
		{name: "Temp 85 80% Humidity", args: args{Celsius: 29.444, Humidity: 80, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 36.301, Blazejczy: 36.005, Humidex: 42.112, Net: 28.642}, wantErr: false},
		{name: "Temp 90 20% Humidity", args: args{Celsius: 32.222, Humidity: 20, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 31.402, Blazejczy: 30.186, Humidex: 32.003, Net: 25.609}, wantErr: false},
		{name: "Temp 90 25% Humidity", args: args{Celsius: 32.222, Humidity: 25, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 32.197, Blazejczy: 30.547, Humidex: 33.337, Net: 26.053}, wantErr: false},
		{name: "Temp 90 30% Humidity", args: args{Celsius: 32.222, Humidity: 30, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 32.992, Blazejczy: 31.07, Humidex: 34.671, Net: 26.498}, wantErr: false},
		{name: "Temp 90 35% Humidity", args: args{Celsius: 32.222, Humidity: 35, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 33.787, Blazejczy: 31.754, Humidex: 36.005, Net: 26.942}, wantErr: false},
		{name: "Temp 90 40% Humidity", args: args{Celsius: 32.222, Humidity: 40, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 34.583, Blazejczy: 32.6, Humidex: 37.339, Net: 27.386}, wantErr: false},
		{name: "Temp 90 45% Humidity", args: args{Celsius: 32.222, Humidity: 45, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 35.378, Blazejczy: 33.607, Humidex: 38.674, Net: 27.829}, wantErr: false},
		{name: "Temp 90 50% Humidity", args: args{Celsius: 32.222, Humidity: 50, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 36.173, Blazejczy: 34.776, Humidex: 40.008, Net: 28.272}, wantErr: false},
		{name: "Temp 90 55% Humidity", args: args{Celsius: 32.222, Humidity: 55, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 36.968, Blazejczy: 36.106, Humidex: 41.342, Net: 28.715}, wantErr: false},
		{name: "Temp 90 60% Humidity", args: args{Celsius: 32.222, Humidity: 60, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 37.763, Blazejczy: 37.598, Humidex: 42.676, Net: 29.158}, wantErr: false},
		{name: "Temp 90 65% Humidity", args: args{Celsius: 32.222, Humidity: 65, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 38.558, Blazejczy: 39.252, Humidex: 44.01, Net: 29.6}, wantErr: false},
		{name: "Temp 90 70% Humidity", args: args{Celsius: 32.222, Humidity: 70, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 39.353, Blazejczy: 41.067, Humidex: 45.344, Net: 30.043}, wantErr: false},
		{name: "Temp 90 75% Humidity", args: args{Celsius: 32.222, Humidity: 75, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 40.148, Blazejczy: 43.044, Humidex: 46.678, Net: 30.484}, wantErr: false},
		{name: "Temp 90 80% Humidity", args: args{Celsius: 32.222, Humidity: 80, WindSpeed: 0, NetRadiationAbsorbed: 0}, want: want{ApparentTemperature: 40.943, Blazejczy: 45.182, Humidex: 48.012, Net: 30.926}, wantErr: false},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := NewThermal(Celsius(tt.args.Celsius)).Humidity(tt.args.Humidity).AirSpeed(tt.args.WindSpeed)
			if tt.args.NetRadiationAbsorbed != 0 {
				t.RadiationAbsorbed(tt.args.NetRadiationAbsorbed)
			}

			got, err := t.ApparentTemperature()
			if (err != nil) != tt.wantErr {
				t1.Errorf("ApparentTemperature() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.ApparentTemperature {
				t1.Errorf("ApparentTemperature() got = %v, want %v", got.Celsius, tt.want.ApparentTemperature)
			}

			got, err = t.BlazejczykHeatIndex()
			if (err != nil) != tt.wantErr {
				t1.Errorf("BlazejczykHeatIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.Blazejczy {
				t1.Errorf("BlazejczykHeatIndex() got = %v, want %v", got.Celsius, tt.want.Blazejczy)
			}

			got, err = t.Humidex()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Humidex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.Humidex {
				t1.Errorf("Humidex() got = %v, want %v", got.Celsius, tt.want.Humidex)
			}

			got, err = t.Net()
			if (err != nil) != tt.wantErr {
				t1.Errorf("Net() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Celsius != tt.want.Net {
				t1.Errorf("Net() got = %v, want %v", got.Celsius, tt.want.Net)
			}

		})
	}
}
