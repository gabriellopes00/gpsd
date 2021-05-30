package main

import (
	"fmt"
	m "gps-worker/app/mock"
	d "gps-worker/domain"
	"log"

	"gps-worker/usecases/calc"
)

func main() {
	entrypoint := d.Position{Name: "Victim", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	positions := m.Positions()
	res := calc.Sort(calc.GetDistances(&entrypoint, positions))

	const helpers uint8 = 10

	filtered, err := calc.Filter(res, helpers, uint8(0))
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	for _, v := range filtered {
		fmt.Println(v)
	}
}
