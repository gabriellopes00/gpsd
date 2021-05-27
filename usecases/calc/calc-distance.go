package calc

import m "math"

func CalcDistance(lat1, lng1, lat2, lng2 float64) float64 {
	var earthRadiusKm float64 = 6371

	degLat := deg2rad(lat2 - lat1)
	degLng := deg2rad(lng2 - lng1)

	lat1 = deg2rad(lat1)
	lat2 = deg2rad(lat2)

	a := m.Pow(m.Sin(degLat/2), 2) + m.Pow(m.Sin(degLng/2), 2)*m.Cos(lat1)*m.Cos(lat2)
	var c = 2 * m.Atan2(m.Sqrt(a), m.Sqrt(1-a))

	result := (earthRadiusKm * c) * 1000
	return toFixed(result, 1)
}

func deg2rad(deg float64) float64 {
	return (deg * m.Pi) / 180
}

// Round float values to integer
func round(num float64) int {
	return int(num + m.Copysign(0.5, num))
}

// Truncate distance decimal cases
func toFixed(num float64, precision int) float64 {
	output := m.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
