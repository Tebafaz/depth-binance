package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var ps *Pubsub

func main() {

	wsConn, err := wsConnectTo("wss://stream.binance.com:9443/ws/btcusdt@depth20@100ms")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	ps = NewPubsub()

	go wsListenAndServe(wsConn, done)

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./public", true)))

	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c.Writer, c.Request)
	})

	r.Run("localhost:8080")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	select {
	case <-done:
		ps.Close()
		return
	case <-interrupt:
		log.Println("interrupt")

		// Cleanly close the connection by sending a close message and then
		// waiting (with timeout) for the server to close the connection.
		err := wsConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		if err != nil {
			log.Println("write close:", err)
			return
		}
		select {
		case <-done:
		case <-time.After(time.Second * 2):
		}
		ps.Close()
		return
	}

}
