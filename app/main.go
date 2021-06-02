package main

import (
	m "gps-worker/app/mock"
	d "gps-worker/domain"
	"log"

	c "gps-worker/usecases/calc"
	mail "gps-worker/usecases/mail"
	p "gps-worker/usecases/path"
	v "gps-worker/usecases/validation"
)

func main() {
	entrypoint := d.Position{Name: "Victim", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	positions := m.Positions()

	radius := make(chan d.Position)
	distances := make(chan d.Position)
	paths := make(chan d.Position)

	if err := v.ValidateCoordinates(positions); err != nil {
		log.Fatalln(err)
	}

	// go cache.CachePositions(positions)
	go c.GetRadius(entrypoint, positions, radius)
	go c.GetDistances(&entrypoint, radius, distances)
	go p.GetPositionPath(entrypoint, distances, paths)
	mail.SendHelperMail(entrypoint, paths)
}
