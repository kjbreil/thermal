package thermal

import "time"

type HighLow struct {
	High *Temperature
	Low  *Temperature
}

func NewHighLow() *HighLow {
	return &HighLow{
		High: MinTemperature(),
		Low:  MaxTemperature(),
	}
}

type TimeMap map[time.Time]*Temperature

func (tm TimeMap) HighLow() *HighLow {
	hl := NewHighLow()
	for _, t := range tm {
		if t.GreaterThan(hl.High) {
			hl.High = t
		}
		if t.LessThan(hl.Low) {
			hl.Low = t
		}
	}
	return hl
}
