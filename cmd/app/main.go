package main

import (
	"fmt"
	"github.com/gabriellopes00/gpsd/domain"
	"log"

	"github.com/gabriellopes00/gpsd/usecases/calc"
	sort "github.com/gabriellopes00/gpsd/usecases/utils"
	"github.com/gabriellopes00/gpsd/usecases/validation"
)

func main() {
	entrypoint := &domain.Position{Name: "EntryPoint", Latitude: -23.16860649763682, Longitude: -46.86898723418065}
	positions := Positions()

	if err := validation.ValidateCoordinates(positions); err != nil {
		log.Fatalln(err)
	}

	inRadius := calc.GetInRadius(entrypoint, positions)
	calc.GetDistances(entrypoint, inRadius)
	ordered := sort.Sort(inRadius)

	for _, p := range ordered {
		fmt.Println(p)
	}
}
