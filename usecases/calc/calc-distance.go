package calc

import m "math"

func deg2rad(deg float64) float64 {
	return (deg * m.Pi) / 180
}

func CalcDistance(lat1, lng1, lat2, lng2 float64) uint16 {
	const earthRadius = 6371

	degLat := deg2rad(lat2 - lat1)
	degLng := deg2rad(lng2 - lng1)

	lat1 = deg2rad(lat1)
	lat2 = deg2rad(lat2)

	angle := m.Pow(m.Sin(degLat/2), 2) + m.Pow(m.Sin(degLng/2), 2)*m.Cos(lat1)*m.Cos(lat2)
	tan := 2 * m.Atan2(m.Sqrt(angle), m.Sqrt(1-angle))

	distance := earthRadius * tan * 1000
	return uint16(m.Trunc(distance))
}
