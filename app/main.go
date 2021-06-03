package main

import (
	m "gps-worker/app/mock"
	d "gps-worker/domain"
	"log"

	c "gps-worker/usecases/calc"
	"gps-worker/usecases/mail"
	p "gps-worker/usecases/path"
	v "gps-worker/usecases/validation"
)

func main() {
	entrypoint := d.Position{Name: "Victim", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	positions := m.Positions()

	paths := make(chan d.Position)

	if err := v.ValidateCoordinates(positions); err != nil {
		log.Fatalln(err)
	}

	ordered := c.Sort(c.GetDistances(&entrypoint, c.GetRadius(&entrypoint, positions)))
	go p.GetPositionPath(entrypoint, ordered, paths)
	mail.SendHelperMail(entrypoint, paths)
}
