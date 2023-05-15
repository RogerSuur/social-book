package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebsocketServer struct {
	upgrader websocket.Upgrader
	clients  map[*websocket.Conn]bool
}

func (w *WebsocketServer) WShandler(rw http.ResponseWriter, r *http.Request) {
	w.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := w.upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println("Cannot upgrade", err)
		rw.Write([]byte(err.Error()))
		return
	}
	defer conn.Close()
	log.Println("Successfully upgraded connection")
	go w.handleMessages(conn)
}

func (w *WebsocketServer) handleMessages(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Cannot read message")
			return
		}
		log.Printf("Message received: %s\n", msg)
	}
}

func New() *WebsocketServer {
	return &WebsocketServer{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		clients: make(map[*websocket.Conn]bool),
	}
}
