package main

import (
	"binanceWebsocket/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}
	currentChan := ps.Subscribe("global")
	for {
		dataInterface := <-currentChan

		data := dataInterface.([]byte)

		var depth models.Depth
		json.Unmarshal(data, &depth)

		depthResponse := sumDepth(depth)

		bytes, err := json.Marshal(depthResponse)
		if err != nil {
			ps.Unsubscribe("global", currentChan)
			close(currentChan)
			break
		}

		err = conn.WriteMessage(1, bytes)
		if err != nil {
			ps.Unsubscribe("global", currentChan)
			close(currentChan)
			break
		}
	}
}
