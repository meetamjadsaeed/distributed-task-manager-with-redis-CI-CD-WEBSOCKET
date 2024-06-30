package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Upgrade connection to a WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

var clients = make(map[*Client]bool)
var broadcast = make(chan []byte)

// Handle WebSocket connections
func HandleConnections(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	defer conn.Close()

	client := &Client{Conn: conn, Send: make(chan []byte)}
	clients[client] = true

	go handleMessages(client)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			delete(clients, client)
			break
		}
		broadcast <- message
	}
}

// Handle incoming messages
func handleMessages(client *Client) {
	for {
		message := <-client.Send
		err := client.Conn.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			client.Conn.Close()
			delete(clients, client)
			break
		}
	}
}

// Broadcast messages to all clients
func BroadcastMessages() {
	for {
		message := <-broadcast
		for client := range clients {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(clients, client)
			}
		}
	}
}
