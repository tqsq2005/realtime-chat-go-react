package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,

		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func main() {
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func reader(conn *websocket.Conn) {
	for true {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Fatal(err)
			return
		}
	}
}

func serverWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Simple Server\nNow is:%v\n", time.Now().String())
	})

	http.HandleFunc("/ws", serverWs)
}
