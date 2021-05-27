package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	server := socketio.NewServer(nil)
	server.OnEvent("", "connection", func(so socketio.Conn) {
		so.Emit("req-position", "msg")
	})

	server.OnEvent("", "return-position", func(so socketio.Conn) {
		fmt.Println(so.Context())
	})

	http.Handle("/socket.io/", server)
	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
