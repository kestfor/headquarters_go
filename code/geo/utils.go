package geo

import "math"

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

const R = 6371000

const MAX_ERROR = 200
