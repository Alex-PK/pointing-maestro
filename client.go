package main

import (
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan []byte
	room   *room
	name   string
	vote   string
}

type clientMsg struct {
	client *client
	msg    []byte
}

func (self *client) read() {
	for {
		if _, msg, err := self.socket.ReadMessage(); err == nil {
			self.room.msg <- &clientMsg{client: self, msg: msg}
		} else {
			break
		}
	}

	self.socket.Close()
}

func (self *client) write() {
	for msg := range self.send {
		if err := self.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}

	self.socket.Close()
}
