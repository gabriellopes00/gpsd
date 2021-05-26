package calc

import m "math"

func CalcDistance(lat1, lng1, lat2, lng2 float64) float64 {
	earthRadius := 6371
	dLat := deg2rad(lat2 - lat1)
	dLng := deg2rad(lng2 - lng1)

	lat1 = deg2rad(lat1)
	lat2 = deg2rad(lat2)

	a := m.Sin(dLat/2)*m.Sin(dLat/2) + m.Sin(dLng/2)*m.Sin(dLng/2)*m.Cos(lat1)*m.Cos(lat2)
	var c = 2 * m.Atan2(m.Sqrt(a), m.Sqrt(1-a))

	result := float64(earthRadius) * float64(c)
	return toFixed((result * 1000), 1)
}

func deg2rad(deg float64) float64 {
	return (deg * m.Pi) / 180
}
