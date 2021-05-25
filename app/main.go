package main

import (
	"fmt"
	ctr "gps-worker/app/controllers"
	"log"
)

func main() {
	controller := ctr.Controller{}
	position, err := controller.CreateService(-24.00097794064844, -46.1797712784327)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(position)
}
