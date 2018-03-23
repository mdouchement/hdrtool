package hdrtool

import "math"

func scale(x float64, unit float64) int {
	return int(round(x, unit) / unit)
}

func unscale(x int, unit float64) float64 {
	return float64(x) * unit
}

func round(x, unit float64) float64 {
	if x > 0 {
		return math.Trunc(x/unit+0.5) * unit
	}
	return math.Trunc(x/unit-0.5) * unit
}
