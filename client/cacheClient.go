package client

import (
	"github.com/gorilla/websocket"
)

var Cache map[string]*websocket.Conn = map[string]*websocket.Conn{}
