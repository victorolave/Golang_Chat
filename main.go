package main

import (
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

func main() {

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	//Sockets
	server.OnConnect("/", func(s socketio.Conn) error {
		log.Println("A new user is connected")

		//Other events

		return nil
	})

	//Http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))


}