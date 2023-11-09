package thermal

import "testing"

func Test_fahrenheitToCelsius(t *testing.T) {
	tests := []struct {
		name string
		f    float64
		c    float64
	}{
		{
			name: "0",
			f:    0,
			c:    -17.78,
		},
		{
			name: "32",
			f:    32,
			c:    0,
		},
		{
			name: "50",
			f:    50,
			c:    10,
		},

		{
			name: "70",
			f:    70,
			c:    21.11,
		},
		{
			name: "100",
			f:    100,
			c:    37.78,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fahrenheitToCelsius(tt.f); got != tt.c {
				t.Errorf("fahrenheitToCelsius() = %v, want %v", got, tt.c)
			}
		})
	}
}

func Test_fahrenheitToKelvin(t *testing.T) {
	tests := []struct {
		name string
		f    float64
		k    float64
	}{
		{
			name: "0",
			f:    0,
			k:    255.37,
		},
		{
			name: "32",
			f:    32,
			k:    273.15,
		},
		{
			name: "50",
			f:    50,
			k:    283.15,
		},
		{
			name: "70",
			f:    70,
			k:    294.26,
		},
		{
			name: "100",
			f:    100,
			k:    310.93,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fahrenheitToKelvin(tt.f); got != tt.k {
				t.Errorf("fahrenheitToKelvin() = %v, want %v", got, tt.k)
			}
		})
	}
}

func Test_kelvinToFarenheit(t *testing.T) {
	tests := []struct {
		name string
		k    float64
		f    float64
	}{
		{
			name: "0",
			k:    255.372,
			f:    0,
		},
		{
			name: "32",
			k:    273.15,
			f:    32,
		},
		{
			name: "50",
			k:    283.15,
			f:    50,
		},
		{
			name: "70",
			k:    294.26,
			f:    70,
		},
		{
			name: "100",
			k:    310.928,
			f:    100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kelvinToFahrenheit(tt.k); got != tt.f {
				t.Errorf("kelvinToFahrenheit() = %v, want %v", got, tt.f)
			}
		})
	}
}

func Test_round(t *testing.T) {
	Precision = 3
	type args struct {
		input float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "0",
			args: args{input: 0},
			want: 0,
		},
		{
			name: "32.2",
			args: args{input: 32.2},
			want: 32.2,
		},
		{
			name: "32.32313",
			args: args{input: 32.32313},
			want: 32.323,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := round(tt.args.input); got != tt.want {
				t.Errorf("round() = %v, want %v", got, tt.want)
			}
		})
	}
}
