package main

import (
	"encoding/json"
	"gps-worker/domain"
	"log"

	"github.com/gorilla/websocket"
)

type User struct {
	Name     string
	Role     string
	Position domain.Position
	Conn     *websocket.Conn
	Chan     *Channel
}

func (u *User) Read() {
	for {
		if _, message, err := u.Conn.ReadMessage(); err != nil {
			log.Println("Error on read message:", err.Error())
			break
		} else {
			u.Chan.messages <- NewMessage(string(message), u.Name)
		}
	}

	u.Chan.leave <- u
}

func (u *User) Write(message *Message) {
	b, _ := json.Marshal(message)

	if err := u.Conn.WriteMessage(websocket.TextMessage, b); err != nil {
		log.Println("Error on write message:", err.Error())
	}
}
