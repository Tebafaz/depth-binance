package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func wsConnectTo(path string) (*websocket.Conn, error) {
	wsConn, _, err := websocket.DefaultDialer.Dial(path, nil)
	return wsConn, err
}

func wsListenAndServe(wsConn *websocket.Conn, done chan struct{}) {
	defer close(done)
	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		ps.Publish("global", message)
		//shows how many connections are there with channels
		fmt.Printf("hala recv: %+v\n", ps.ShowAllSubscribers())
	}
}
