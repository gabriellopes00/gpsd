package main

import (
	"fmt"

	validator "github.com/asaskevich/govalidator"
)

type Position struct {
	Latitude  float64 `json:"latitude" valid:"latitude"`
	Longitude float64 `json:"longitude" valid:"longitude"`
}

func CreatePosition(latitude, longitude float64) *Position {
	return &Position{Latitude: latitude, Longitude: longitude}
}

func (p *Position) Validate() error {
	_, err := validator.ValidateStruct(p)
	return err
}

func main() {
	position := CreatePosition(-24.00097794064844, -46.1797712784327)
	err := position.Validate()
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	validator.SetFieldsRequiredByDefault(true)
}
