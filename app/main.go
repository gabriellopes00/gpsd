package main

import (
	m "gps-worker/app/mock"
	d "gps-worker/domain"

	"gps-worker/usecases/calc"
)

func main() {
	entrypoint := d.Position{Name: "Victim", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	positions := m.Positions()
	// res := calc.Sort(calc.GetDistances(&entrypoint, positions))

	calc.GetRadius(entrypoint, positions)

	// filtered, err := calc.Filter(&entrypoint, &res)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	panic(err)
	// }

	// for _, v := range *filtered {
	// 	fmt.Println(v)
	// }
}
