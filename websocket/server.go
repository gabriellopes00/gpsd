package main

import (
	"encoding/json"
	"fmt"
	"gps-worker/app/mock"
	"gps-worker/domain"
	"gps-worker/usecases/calc"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
)

var up = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true }, // allow everywhere access
}

func Handler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	name := keys.Get("username")
	role := keys.Get("role")
	lat, _ := strconv.ParseFloat(keys.Get("lat"), 64)
	lng, _ := strconv.ParseFloat(keys.Get("lng"), 64)

	user := &User{
		Name:     name,
		Role:     role,
		Position: domain.Position{Latitude: lat, Longitude: lng},
	}

	fmt.Println(user)

	// ws, err := up.Upgrade(w, r, nil)
	// if err != nil {
	// 	log.Println(err)
	// }

}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		victim := domain.Position{}
		json.Unmarshal(p, &victim)

		positions := mock.Positions()
		res := calc.GetDistances(&victim, positions)

		for i1 := range res {
			for i2 := range res {
				if res[i1].Distance < res[i2].Distance {
					res[i1], res[i2] = res[i2], res[i1]
				}
			}
		}

		for i := 0; i < 5; i++ {
			fmt.Println(res[i])
		}

		if err := conn.WriteMessage(messageType, []byte("Help on the way")); err != nil {
			log.Println(err)
			return
		}

	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", Handler)
}

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
