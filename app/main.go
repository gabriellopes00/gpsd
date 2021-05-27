package main

import (
	"fmt"
	m "gps-worker/app/mock"
	d "gps-worker/domain"
	"net/http"
	"strconv"

	"gps-worker/usecases/calc"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/distances", calcDistance)
	e.Logger.Fatal(e.Start(":7699"))

	// entrypoint := d.Position{Name: "Entrypoint", Latitude: -23.16862852698391, Longitude: -46.868998411087226}
	// positions := m.Positions()
	// res := calc.CalcMultipleDistances(&entrypoint, positions)
	// for i1 := range res {
	// 	for i2 := range res {
	// 		if res[i1].Distance < res[i2].Distance {
	// 			res[i1], res[i2] = res[i2], res[i1]
	// 		}
	// 	}
	// }
	// for _, v := range res {

	// 	fmt.Println(v)
	// }
}

func calcDistance(c echo.Context) error {
	name := c.QueryParam("name")
	latitude, _ := strconv.ParseFloat(c.QueryParam("lat"), 64)
	longitude, _ := strconv.ParseFloat(c.QueryParam("lng"), 64)

	entrypoint := d.Position{Name: name, Latitude: latitude, Longitude: longitude}
	fmt.Println(entrypoint)
	positions := m.Positions()
	res := calc.CalcMultipleDistances(&entrypoint, positions)

	order := c.QueryParam("order")
	fmt.Println(order)

	if order == "asc" {
		for i1 := range res {
			for i2 := range res {
				if res[i1].Distance < res[i2].Distance {
					res[i1], res[i2] = res[i2], res[i1]
				}
			}
		}
	} else if order == "desc" {
		for i1 := range res {
			for i2 := range res {
				if res[i1].Distance > res[i2].Distance {
					res[i1], res[i2] = res[i2], res[i1]
				}
			}
		}
	}
	return c.JSON(http.StatusOK, res)
}
