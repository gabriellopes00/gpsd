package main

import (
	d "gps-worker/domain"
	u "gps-worker/usecases"

	m "gps-worker/app/mock"
)

func main() {
	entrypoint := d.Position{Name: "My House", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	positions := m.Positions()

	u.CalcDistance(&entrypoint, positions)
}
