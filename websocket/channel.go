package main

type Channel struct {
	users    map[string]*User
	messages chan *Message
	join     chan *User
	leave    chan *User
}
