package main

import (
	"encoding/json"
	"gps-worker/app/mock"
	d "gps-worker/domain"
	"io/ioutil"
	"log"
	"net/http"

	c "gps-worker/usecases/calc"
	m "gps-worker/usecases/mail"
	p "gps-worker/usecases/path"
	s "gps-worker/usecases/sort"
	v "gps-worker/usecases/validation"
)

func main() {
	http.HandleFunc("/help", HandleHelpRequest)

	log.Fatalln(http.ListenAndServe(":8778", nil))
}

func HandleHelpRequest(wr http.ResponseWriter, r *http.Request) {
	var position d.Position
	body, _ := ioutil.ReadAll(r.Body)
	if err := json.Unmarshal(body, &position); err != nil {
		http.Error(wr, "invalid body request", http.StatusBadRequest)
	}

	go GetHelp(position.Latitude, position.Longitude)
	wr.Write([]byte("Help on the way"))
}

func GetHelp(lat, lng float64) {
	entrypoint := &d.Position{Name: "Victim", Latitude: lat, Longitude: lng}
	positions := mock.Positions()

	if err := v.ValidateCoordinates(positions); err != nil {
		log.Fatalln(err)
	}

	paths := make(chan *d.Position)

	inRadius := c.GetInRadius(entrypoint, positions)
	c.GetDistances(entrypoint, inRadius)
	ordered := s.Sort(inRadius)
	go p.GetPositionPath(entrypoint, s.Filter(ordered), paths)
	m.SendHelperMail(paths)
}
