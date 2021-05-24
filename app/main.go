package main

import (
	"fmt"
	ctr "gps-worker/app/controllers"
	"gps-worker/usecases"
	"log"
)

func main() {
	controller := ctr.Controller{&usecases.Service{}}
	fmt.Println("called")
	position, err := controller.CreateService(-24.00097794064844, -46.1797712784327)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(position)
}
