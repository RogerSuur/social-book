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
	conn, err := w.upgrader.Upgrade(rw, r, nil)
	if err != nil {
		log.Println("Cannot upgrade:", err)
		rw.Write([]byte(err.Error()))
		return
	}
	defer conn.Close()
	log.Println("Successfully upgraded connection")
	w.clients[conn] = true

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading:", err)
			break
		}
		log.Printf("Received: %s", message)
		err = conn.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Error during message writing:", err)
			break
		}
	}
}

func New() *WebsocketServer {
	return &WebsocketServer{
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				// required for CORS
				// should return true if origin is allowed
				// origin := r.Header.Get("Origin")
				return true
			},
		},
		clients: make(map[*websocket.Conn]bool),
	}
}
