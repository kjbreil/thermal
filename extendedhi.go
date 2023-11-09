package thermal

import (
	"fmt"
	"math"
	"reflect"
)

// ExtendedHeatIndex is an implementation of Lu and Romps equation
// Code is adaptation of https://romps.org/papers/pubdata/2020/heatindex/heatindex.py
// Still needs a lot of cleanup to refactor for Go and eventually output clothing information
func (t *Thermal) ExtendedHeatIndex() (*Temperature, error) {
	if t.RelativeHumidity == nil {
		return nil, fmt.Errorf("relative humidity needed for humidex")
	}
	var err error
	RH := *t.RelativeHumidity / float64(100)
	ev, err := findEqvar(t.Temperature.Kelvin, RH)
	if err != nil {
		return nil, err
	}
	var T float64
	switch ev.name {
	case "phi", "Rf", "Rs", "Rs*", "dTcdt":
		T, err = findTemperature(ev)
		if err != nil {
			return nil, err
		}
	}
	if t.Temperature.Kelvin == 0.0 {
		T = 0.0
	}
	return Kelvin(T), nil
}

const (
	ttrip   = 273.16
	ptrip   = 611.65
	e0v     = 2374000.0
	e0s     = 333700.0
	rgasa   = 287.04
	rgasv   = 461.0
	cva     = 719.0
	cvv     = 1418.0
	cvl     = 4119.0
	cvs     = 1861.0
	cpa     = cva + rgasa
	cpv     = cvv + rgasv
	r       = 124.0
	phiSalt = 0.9 // Effective relative humidity of saline solution
	tc      = 310.0
)

// Usage Constants
const (
	maxIter = 100
)

// Constants
var (
	MetobolicRate = 180.0 // Metabolic rate per skin area for walking

	AverageCoreMass = 83.6
	AverageHeight   = 1.69

	pc = phiSalt * pvstar(tc)
)

func pvstar(T float64) float64 {
	if T == 0 {
		return 0.0
	} else if T < ttrip {
		return ptrip * math.Pow(T/ttrip, (cpv-cvs)/rgasv) * math.Exp((e0v+e0s-(cvv-cvs)*ttrip)/rgasv*(1.0/ttrip-1.0/float64(T)))
	} else {
		return ptrip * math.Pow(T/ttrip, (cpv-cvl)/rgasv) * math.Exp((e0v-(cvv-cvl)*ttrip)/rgasv*(1.0/ttrip-1.0/float64(T)))
	}
}

func le(T float64) float64 {
	return e0v + (cvv-cvl)*(T-ttrip) + rgasv*T
}

func qv(Ta float64, Pa float64) float64 {
	p := 101300.0
	eta := 1.43e-06
	return eta * MetobolicRate * (cpa*(tc-Ta) + le(310.0)*rgasa/(p*rgasv)*(pc-Pa))
}

func zs(Rs float64) float64 {
	return func() float64 {
		if reflect.DeepEqual(Rs, 0.0387) {
			return 52.1
		}
		return 600000000.0 * math.Pow(Rs, 5)
	}()
}

// Resistance to heat transfer through the  boundary layer of air in contact with exposed skin Regions II and III
func ra(ts, ta float64) float64 {
	hc := 17.4
	phiRad := 0.85
	return heatTransfer(ts, ta, hc, phiRad)
}

// Resistance to heat transfer through the boundary layer of air in contact with clothing Regions II and III
func raBar(ts, ta float64) float64 {
	hc := 11.6
	phiRad := 0.79
	return heatTransfer(ts, ta, hc, phiRad)
}

// Resistance to heat transfer through the boundary layer of air in contact with the exposed skin Regions IV and V
func raUn(ts, ta float64) float64 {
	hc := 12.3
	phiRad := 0.8
	return heatTransfer(ts, ta, hc, phiRad)
}

func heatTransfer(ts, ta, hc, phiRad float64) float64 {
	sigma := 5.67e-08
	epsilon := 0.97
	hr := epsilon * phiRad * sigma * (math.Pow(ts, 2) + math.Pow(ta, 2)) * (ts + ta)
	return 1.0 / (hc + hr)
}

type equalVariable struct {
	name  string
	phi   float64
	rf    float64
	rs    float64
	dtcdt float64
}

func findEqvar(Ta float64, RH float64) (equalVariable, error) {
	var Rf float64
	var eqvarName string

	Pa := RH * pvstar(Ta)
	Rs := 0.0387
	phi := 0.84
	dTcdt := 0.0
	tol := 1e-08

	// Variables used only in here
	zaUn := 60.6 / 12.3
	za := 60.6 / 17.4    // Resistance to water transfer through the boundary layer of air in contact with exposed skin
	zaBar := 60.6 / 11.6 // Resistance to water transfer through the boundary layer of air in contact with clothing
	areaOfSkin := 0.202 * math.Pow(AverageCoreMass, 0.425) * math.Pow(AverageHeight, 0.725)
	specificCoreHeatCapacity := 3492.0

	heatCapacityOfCore := AverageCoreMass * specificCoreHeatCapacity / areaOfSkin

	m := (pc - Pa) / (zs(Rs) + za)
	mBar := (pc - Pa) / (zs(Rs) + zaBar)
	Ts, err := solve(func(Ts float64) (float64, error) {
		return (Ts-Ta)/ra(Ts, Ta) + (pc-Pa)/(zs(Rs)+za) - (tc-Ts)/Rs, nil
	},
		max(0, min(tc, Ta)-Rs*math.Abs(m)),
		max(tc, Ta)+Rs*math.Abs(m),
		tol,
		maxIter,
	)
	if err != nil {
		return equalVariable{}, err
	}

	Tf, err := solve(func(Tf float64) (float64, error) {
		return (Tf-Ta)/raBar(Tf, Ta) + (pc-Pa)/(zs(Rs)+zaBar) - (tc-Tf)/Rs, nil
	},
		max(0., min(tc, Ta)-Rs*math.Abs(mBar)),
		max(tc, Ta)+Rs*math.Abs(mBar),
		tol,
		maxIter,
	)
	if err != nil {
		return equalVariable{}, err
	}

	flux1 := MetobolicRate - qv(Ta, Pa) - (1.0-phi)*(tc-Ts)/Rs
	flux2 := MetobolicRate - qv(Ta, Pa) - (1.0-phi)*(tc-Ts)/Rs - phi*(tc-Tf)/Rs

	switch {
	case flux1 <= 0.0:
		phi = 1.0 - (MetobolicRate-qv(Ta, Pa))*Rs/(tc-Ts)
		Rf = math.Inf(1)
		return equalVariable{
			name:  "phi",
			phi:   phi,
			rf:    Rf,
			rs:    Rs,
			dtcdt: dTcdt,
		}, nil
	case flux2 <= 0.0:
		TsBar := tc - (MetobolicRate-qv(Ta, Pa))*Rs/phi + (1.0/phi-1.0)*(tc-Ts)
		Tf, err = solve(func(Tf float64) (float64, error) {
			return (Tf-Ta)/raBar(Tf, Ta) + (pc-Pa)*(Tf-Ta)/((zs(Rs)+zaBar)*(Tf-Ta)+r*raBar(Tf, Ta)*(TsBar-Tf)) - (tc-TsBar)/Rs, nil
		}, Ta, TsBar, tol, maxIter)
		if err != nil {
			return equalVariable{}, err
		}
		Rf = raBar(Tf, Ta) * (TsBar - Tf) / (Tf - Ta)
		return equalVariable{
			name:  "Rf",
			phi:   phi,
			rf:    Rf,
			rs:    Rs,
			dtcdt: dTcdt,
		}, nil
	default:
		Rf = 0.0
		flux3 := MetobolicRate - qv(Ta, Pa) - (tc-Ta)/raUn(tc, Ta) - (phiSalt*pvstar(tc)-Pa)/zaUn
		if flux3 < 0.0 {
			Ts, err = solve(func(Ts float64) (float64, error) {
				return (Ts-Ta)/raUn(Ts, Ta) + (pc-Pa)/(zs((tc-Ts)/(MetobolicRate-qv(Ta, Pa)))+zaUn) - (MetobolicRate - qv(Ta, Pa)), nil
			}, 0.0, tc, tol, maxIter)
			if err != nil {
				return equalVariable{}, err
			}
			Rs = (tc - Ts) / (MetobolicRate - qv(Ta, Pa))
			eqvarName = "Rs"
			Ps := pc - (pc-Pa)*zs(Rs)/(zs(Rs)+zaUn)
			if Ps > phiSalt*pvstar(Ts) {
				Ts, err = solve(func(Ts float64) (float64, error) {
					return (Ts-Ta)/raUn(Ts, Ta) + (phiSalt*pvstar(Ts)-Pa)/zaUn - (MetobolicRate - qv(Ta, Pa)), nil
				}, 0.0, tc, tol, maxIter)
				if err != nil {
					return equalVariable{}, err
				}
				Rs = (tc - Ts) / (MetobolicRate - qv(Ta, Pa))
				eqvarName = "Rs*"
			}
		} else {
			Rs = 0.0
			eqvarName = "dTcdt"
			dTcdt = 1.0 / heatCapacityOfCore * flux3
		}
		return equalVariable{
			name:  eqvarName,
			phi:   phi,
			rf:    Rf,
			rs:    Rs,
			dtcdt: dTcdt,
		}, nil
	}
}

func findTemperature(e equalVariable) (float64, error) {
	var T float64
	var err error

	// Constants for this function only
	tolT := 1e-08
	pa0 := 1600.0

	switch e.name {
	case "phi":
		T, err = solve(func(T float64) (float64, error) {
			ev, err := findEqvar(T, 1.0)
			return ev.phi - e.phi, err
		}, 0.0, 240.0, tolT, maxIter)
		if err != nil {
			return 0, err
		}
	case "Rf":
		T, err = solve(func(T float64) (float64, error) {
			ev, err := findEqvar(T, min(1., pa0/pvstar(T)))
			return ev.rf - e.rf, err
		}, 230.0, 300.0, tolT, maxIter)
		if err != nil {
			return 0, err
		}
	case "Rs", "Rs*":
		T, err = solve(func(T float64) (float64, error) {
			ev, err := findEqvar(T, pa0/pvstar(T))
			return ev.rs - e.rs, err
		}, 295.0, 350.0, tolT, maxIter)
		if err != nil {
			return 0, err
		}
	default:
		T, err = solve(func(T float64) (float64, error) {
			ev, err := findEqvar(T, pa0/pvstar(T))
			return ev.dtcdt - e.dtcdt, err
		}, 340.0, 1000.0, tolT, maxIter)
		if err != nil {
			return 0, err
		}
	}
	return T, nil
}

func solve(f func(Ts float64) (float64, error), a float64, b float64, tol float64, maxIter int) (float64, error) {
	var err error
	var fa, fb, fc float64
	fa, err = f(a)
	if err != nil {
		return 0, err
	}
	fb, err = f(b)
	if err != nil {
		return 0, err
	}
	if fa*fb > 0.0 {
		return 0, fmt.Errorf("wrong initial interval in the root solver")
	} else {
		for i := 0; i < maxIter; i++ {
			c := (a + b) / 2.0
			fc, err = f(c)
			if err != nil {
				return 0, err
			}
			if fb*fc > 0.0 {
				b = c
				fb = fc
			} else {
				a = c
				fa = fc
			}
			if math.Abs(a-b) < tol {
				return c, nil
			}
			if i == maxIter-1 {
				return 0, fmt.Errorf("reaching maximum iteration in the root solver")
			}
		}
	}
	return 0, fmt.Errorf("reached end")
}
