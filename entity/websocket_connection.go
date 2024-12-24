package entity

import "github.com/gorilla/websocket"

type WebsocketConnection struct {
	ID   string
	Conn *websocket.Conn
}
