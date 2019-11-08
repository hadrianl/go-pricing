package bs

import "math"

const DX_TARGET = 0.00001

func cdf(x float64) float64 {
	return 0.5 * (1 + math.Erf(x/math.Sqrt2))
}

func pdf(x float64) float64 {
	return math.Exp(-math.Pow(x, 2)/2) / (math.Sqrt2 * math.SqrtPi)
}

func calcD1(s, l, r, t, v float64) float64 {
	return (math.Log(s/l) + (r+0.5*math.Pow(v, 2))*t) / (v * math.Sqrt(t))
}

func CalcPrice(s, l, r, t, v, cp float64) float64 {
	if v <= 0 {
		return math.Max(0, (s-l)*cp)
	}

	d1 := calcD1(s, l, r, t, v)
	d2 := d1 - v*math.Sqrt(t)
	price := cp * (s*cdf(cp*d1) - l*cdf(cp*d2)*math.Exp(-r*t))

	return price
}

func CalcDelta(s, l, r, t, v, cp float64) float64 {
	if v <= 0 {
		return 0
	}

	d1 := calcD1(s, l, r, t, v)
	delta := cp * cdf(cp*d1) * s * 0.01

	return delta
}

func CalcGamma(s, l, r, t, v, cp float64) float64 {
	if v <= 0 {
		return 0
	}

	d1 := calcD1(s, l, r, t, v)
	gamma := pdf(d1) / (s * v * math.Sqrt(t)) * math.Pow(s, 2) * 0.0001

	return gamma
}

func CalcTheta(s, l, r, t, v, cp float64) float64 {
	if v <= 0 {
		return 0
	}

	d1 := calcD1(s, l, r, t, v)
	d2 := d1 - v*math.Sqrt(t)

	theta := -0.5*s*pdf(d1)*v/math.Sqrt(t) - cp*r*l*math.Exp(-r*t)*cdf(cp*d2)
	theta = theta / 240

	return theta
}

func calcOriginalVega(s, l, r, t, v, cp float64) float64 {
	if v <= 0 {
		return 0
	}

	d1 := calcD1(s, l, r, t, v)
	vega := s * pdf(d1) * math.Sqrt(t)

	return vega
}

func CalcVega(s, l, r, t, v, cp float64) float64 {
	vega := calcOriginalVega(s, l, r, t, v, cp) / 100

	return vega
}

func CalcGreeks(s, l, r, t, v, cp float64) (price, delta, gamma, theta, vega float64) {
	price = CalcPrice(s, l, r, t, v, cp)
	delta = CalcDelta(s, l, r, t, v, cp)
	gamma = CalcGamma(s, l, r, t, v, cp)
	theta = CalcTheta(s, l, r, t, v, cp)
	vega = CalcVega(s, l, r, t, v, cp)
	return price, delta, gamma, theta, vega
}

func CalcImpVol(price, s, l, r, t, cp float64) float64 {
	if price <= 0 {
		return 0
	}

	var meet bool
	switch {
	case cp == 1 && (price > (s-l)*math.Exp(-r*t)):
		meet = true
	case cp == -1 && (price > l*math.Exp(-r*t)-s):
		meet = true
	default:
		meet = false
	}

	if !meet {
		return 0
	}

	v := 0.3
	var p, vega, dx float64

	for i := 0; i < 50; i++ {
		p = CalcPrice(s, l, r, t, v, cp)
		vega = calcOriginalVega(s, l, r, t, v, cp)

		if vega == 0 {
			break
		}

		dx = (price - p) / vega

		if math.Abs(dx) < DX_TARGET {
			break
		}

		v += dx
	}

	if v <= 0 {
		return 0
	}

	return v
}
