package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/yemmyharry/backend/pkg/websocket"
	"log"
	"net/http"
	"os"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	setupRoutes()
	fmt.Println("Chat App Started at" + port)

	http.ListenAndServe(":"+ port, nil)
	fmt.Println("Chat App Started at " + port)

}