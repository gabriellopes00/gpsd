package main

import (
	m "gps-worker/app/mock"
	d "gps-worker/domain"
	"net/http"
	"strconv"

	"gps-worker/usecases/calc"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/distances/:lat/:lng/:order", calcDistance)
	e.Logger.Fatal(e.Start(":7699"))
}

func calcDistance(c echo.Context) error {
	name := c.Param("lat")
	latitude, _ := strconv.ParseFloat(c.Param("lat"), 64)
	longitude, _ := strconv.ParseFloat(c.Param("lng"), 64)

	entrypoint := d.Position{Name: name, Latitude: latitude, Longitude: longitude}
	positions := m.Positions()
	res := calc.CalcMultipleDistances(&entrypoint, positions)

	order := c.Param("order")

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
