package main

import (
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	ID   string `json:"id"`
	Body string `json:"body"`
	From string `json:"sender"`
}

func NewMessage(body string, sender string) *Message {
	return &Message{
		ID:   uuid.NewV4().String(),
		Body: body,
		From: sender,
	}
}
