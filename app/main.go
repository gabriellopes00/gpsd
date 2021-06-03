package main

import (
	"gps-worker/app/mock"
	d "gps-worker/domain"
	"log"

	c "gps-worker/usecases/calc"
	"gps-worker/usecases/mail"
	p "gps-worker/usecases/path"
	v "gps-worker/usecases/validation"
)

func main() {
	entrypoint := &d.Position{Name: "Victim", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	positions := mock.Positions()

	if err := v.ValidateCoordinates(positions); err != nil {
		log.Fatalln(err)
	}

	paths := make(chan *d.Position)

	inRadius := c.GetRadius(entrypoint, positions)
	c.GetDistances(entrypoint, inRadius)

	ordered := c.Sort(inRadius)
	go p.GetPositionPath(entrypoint, ordered, paths)
	mail.SendHelperMail(paths)
}
