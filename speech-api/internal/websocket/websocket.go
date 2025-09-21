package websocket

import (
	"sync"

	ws "github.com/gorilla/websocket"
)

type (
	WebSocketPool interface {
		Add(*ws.Conn)
		Remove(*ws.Conn)
	}

	websocketPool struct {
		mutex sync.RWMutex
		conns map[*ws.Conn]bool
	}
)

func NewWebSocketPool() WebSocketPool {
	return &websocketPool{
		conns: make(map[*ws.Conn]bool),
	}
}

func (w *websocketPool) Add(conn *ws.Conn) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.conns[conn] = true
}

func (w *websocketPool) Remove(conn *ws.Conn) {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	delete((w.conns), conn)
}
