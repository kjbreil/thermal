package thermal

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func (t *Thermal) PmvPpd() (float64, float64) {
	pmv := t.Pmv()
	ppd := 100.0 - 95.0*math.Exp(-0.03353*math.Pow(pmv, 4.0)-0.2179*math.Pow(pmv, 2.0))
	return pmv, ppd
}

func (t *Thermal) Apmv() float64 {
	pmv := t.Pmv()
	return math.Round(pmv/(1+*t.AdaptiveCoefficient*pmv)*100) / 100
}

func (t *Thermal) Pmv() float64 {

	tdb := t.Temperature.Celsius
	tr := t.MeanRadiantTemperature.Celsius
	vr := *t.RelativeAirSpeed
	rh := *t.RelativeHumidity
	met := *t.MetabolicRate
	clo := *t.ClothingInsulation
	wme := *t.ExternalWork

	var f_cl float64
	pa := rh * 10 * math.Exp(16.6536-4030.183/(tr+235))
	icl := 0.155 * clo
	m := met * 58.15
	w := wme * 58.15
	mw := m - w
	if icl <= 0.078 {
		f_cl = 1 + 1.29*icl
	} else {
		f_cl = 1.05 + 0.645*icl
	}
	hcf := 12.1 * math.Sqrt(vr)
	hc := hcf
	taa := tdb + 273
	tra := tr + 273
	t_cla := taa + (35.5-tdb)/(3.5*icl+0.1)
	p1 := icl * f_cl
	p2 := p1 * 3.96
	p3 := p1 * 100
	p4 := p1 * taa
	p5 := 308.7 - 0.028*mw + p2*math.Pow(float64(tra)/100.0, 4)
	xn := t_cla / 100
	xf := t_cla / 50
	eps := 0.00015
	n := 0
	for math.Abs(xn-xf) > eps {
		xf = (xf + xn) / 2
		hcn := 2.38 * math.Pow(math.Abs(100.0*xf-taa), 0.25)
		if hcf > hcn {
			hc = hcf
		} else {
			hc = hcn
		}
		xn = (p5 + p4*hc - p2*math.Pow(xf, 4)) / (100 + p3*hc)
		n += 1
		if n > 150 {
			panic(fmt.Errorf("StopIteration: %v", "Max iterations exceeded"))
		}
	}
	tcl := 100*xn - 273
	hl1 := 3.05 * 0.001 * float64(5733-6.99*mw-pa)
	hl2 := 0.0
	if mw > 58.15 {
		hl2 = 0.42 * (mw - 58.15)
	}
	hl3 := 1.7 * 1e-05 * m * (5867 - pa)
	hl4 := 0.0014 * m * (34 - tdb)
	hl5 := 3.96 * f_cl * (math.Pow(xn, 4) - math.Pow(float64(tra)/100.0, 4))
	hl6 := f_cl * hc * (tcl - tdb)
	ts := 0.303*math.Exp(-0.036*m) + 0.028
	_pmv := ts * (mw - hl1 - hl2 - hl3 - hl4 - hl5 - hl6)
	return round(_pmv)
}

func (t *Thermal) AdaptiveAshrae() (float64, float64, float64, float64, float64, bool, bool) {
	tdb := t.Temperature.Celsius
	tr := t.MeanRadiantTemperature.Celsius
	v := *t.RelativeAirSpeed

	tRunningMean := t.RunningMean.Celsius

	to := tO(tdb, tr, v, t.Standard)
	ce := 0.0

	if v >= 0.6 && to >= 25.0 {
		ce = 999
	}
	if v < 0.9 && ce == 999 {
		ce = 1.2
	}
	if v < 1.2 && ce == 999 {
		ce = 1.8
	}
	if ce == 999 {
		ce = 2.2
	}

	tCmf := 0.31*tRunningMean + 17.8

	tCmf = math.Round(tCmf)
	tmpCmf80Low := tCmf - 3.5
	tmpCmf90Low := tCmf - 2.5
	tmpCmf80Up := tCmf + 3.5 + ce
	tmpCmf90Up := tCmf + 2.5 + ce
	acceptability80 := false
	if tmpCmf80Low <= to && to <= tmpCmf80Up {
		acceptability80 = true
	}
	acceptability90 := false

	if tmpCmf90Low <= to && to <= tmpCmf90Up {
		acceptability90 = true
	}

	return tCmf, tmpCmf80Low, tmpCmf80Up, tmpCmf90Low, tmpCmf90Up, acceptability80, acceptability90

}

func (t *Thermal) AdaptiveAshraeF() (float64, float64, float64, float64, float64, bool, bool) {
	t_cmf, tmp_cmf_80_low, tmp_cmf_80_up, tmp_cmf_90_low, tmp_cmf_90_up, acceptability_80, acceptability_90 := t.AdaptiveAshrae()

	return celsiusToFahrenheit(t_cmf), celsiusToFahrenheit(tmp_cmf_80_low), celsiusToFahrenheit(tmp_cmf_80_up), celsiusToFahrenheit(tmp_cmf_90_low), celsiusToFahrenheit(tmp_cmf_90_up), acceptability_80, acceptability_90
}

func (t *Thermal) AdaptiveThermalHeatBalance() float64 {
	met := *t.MetabolicRate
	tRunningMean := t.RunningMean.Celsius

	metAdapted := met - 0.234*tRunningMean/58.2

	metSave := t.MetabolicRate
	cloSave := t.ClothingInsulation

	t = t.Met(met - 0.234*tRunningMean/58.2)
	t.Clo(math.Pow(10, -0.17168-0.000485*tRunningMean+0.08176*metAdapted-0.00527*tRunningMean*metAdapted))

	pmvRes := t.Pmv()

	t.MetabolicRate = metSave
	t.ClothingInsulation = cloSave
	ts := 0.303*math.Exp(-0.036*metAdapted*58.15) + 0.028
	lAdapted := pmvRes / ts
	return math.Round((1.484+0.0276*lAdapted-0.9602*metAdapted-0.0342*tRunningMean+0.0002264*lAdapted*tRunningMean+0.018696*metAdapted*tRunningMean-0.0002909*lAdapted*metAdapted*tRunningMean)*1000) / 1000
}

func (t *Thermal) AdaptiveThermalHeatBalanceRange(low, high float64) (float64, float64) {

	var goodTemps []float64

	tempSave := t.Temperature

	for i := 50.0; i < 90; i += 0.1 {
		i = math.Round(i*100) / 100
		t.OperativeTemperature(Fahrenheit(i))
		athb := t.AdaptiveThermalHeatBalance()

		if athb > low && athb < high {
			goodTemps = append(goodTemps, i)
		}
	}

	t.OperativeTemperature(tempSave)

	sort.Float64s(goodTemps)

	return goodTemps[0], goodTemps[len(goodTemps)-1]
}

func tO(tdb, tr, v float64, standard string) float64 {
	if strings.ToLower(standard) == "iso" {
		return (tdb*math.Sqrt(10*v) + tr) / (1 + math.Sqrt(10*v))
	} else if strings.ToLower(standard) == "ashrae" {
		a := 0.0
		if v < 0.6 {
			a = 0.6
		} else {
			a = 0.7
		}
		if v < 0.2 {
			a = 0.5
		}

		return (a * tdb) + (1-a)*tr
	}
	return 0
}

func clothing(clo, met float64) float64 {
	if met > 1 {
		return clo * (0.6 + 0.4/met)
	} else {
		return clo
	}
}
