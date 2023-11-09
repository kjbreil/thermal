package thermal

import "testing"

func TestTemperature_PressureSaturation(t1 *testing.T) {
	Precision = 2
	tests := []struct {
		name    string
		celsius float64
		want    float64
	}{
		{
			name:    "0",
			celsius: 0,
			want:    611.21,
		},
		{
			name:    "5",
			celsius: 5,
			want:    872.49,
		},
		{
			name:    "10",
			celsius: 10,
			want:    1228.0,
		},
		{
			name:    "15",
			celsius: 15,
			want:    1705.45,
		},
		{
			name:    "20",
			celsius: 20,
			want:    2338.8,
		},
		{
			name:    "25",
			celsius: 25,
			want:    3169.22,
		},
		{
			name:    "30",
			celsius: 30,
			want:    4246.03,
		},
		{
			name:    "35",
			celsius: 35,
			want:    5627.82,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Celsius(tt.celsius)
			if got := t.PressureSaturation(); got != tt.want {
				t1.Errorf("PressureSaturation() = %v, want %v", got, tt.want)
			}
		})
	}
}
