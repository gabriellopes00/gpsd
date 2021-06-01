package main

import (
	m "gps-worker/app/mock"
	d "gps-worker/domain"
	"log"

	"gps-worker/usecases/cache"
	c "gps-worker/usecases/calc"
	mail "gps-worker/usecases/mail"
	p "gps-worker/usecases/path"
	v "gps-worker/usecases/validation"
)

func main() {
	entrypoint := d.Position{Name: "Victim", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	positions := m.Positions()

	channel := make(chan d.Position)

	if err := v.ValidateCoordinates(positions); err != nil {
		log.Fatalln(err)
	}

	go cache.CachePositions(positions)
	go c.GetRadius(entrypoint, positions, channel)
	go c.GetDistances(&entrypoint, channel)
	go p.GetPositionPath(entrypoint, channel)
	mail.SendHelperMail(entrypoint, channel)
}
