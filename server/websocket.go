package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var upgrader = websocket.Upgrader{}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	clients[conn] = true
}

func BroadcastMetric(m Metric) {
	for c := range clients {
		c.WriteJSON(m)
	}
}
