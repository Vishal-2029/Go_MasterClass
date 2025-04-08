package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make([]*websocket.Conn,0)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil{
		log.Println("Error Connection", err)
		return
	}
	 
	clients = append(clients, conn)

	fmt.Println("New Client connected" )

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		for _, client := range clients{
			if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
				fmt.Println("Error writing message:", err)
				break
			}

		}
	}
	for i, client := range clients {
		if client == conn {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleConnection)
	fmt.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}