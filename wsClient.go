package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func wsConnectTo(path string) (*websocket.Conn, error) {
	/*interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)*/

	wsConn, _, err := websocket.DefaultDialer.Dial(path, nil)
	/*if err != nil {
		log.Fatal("dial: ", err)
	}*/
	//defer c.Close()

	//done := make(chan struct{})

	return wsConn, err

	/*go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()*/
	/*select {
	case <-done:
		return
	case <-interrupt:
		log.Println("interrupt")

		// Cleanly close the connection by sending a close message and then
		// waiting (with timeout) for the server to close the connection.
		err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}
		select {
		case <-done:
		case <-time.After(time.Second * 2):
		}
		return
	}*/
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
		fmt.Printf("hala recv: %+v\n", ps.ShowAllSubscribers())
	}
}
