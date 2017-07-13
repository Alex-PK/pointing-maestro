package main

import (
	"github.com/gorilla/websocket"
	"log"
)

const (
	socketBuffSize = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:socketBuffSize,
	WriteBufferSize: messageBufferSize,
}

type room struct {
	name    string
	msg     chan []byte
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func newRoom(name string) *room {
	return &room {
		name:    name,
		msg:     make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (self *room) run() {
	log.Printf("Running room %s\n", self.name)
	for {
		select {
		case client := <- self.join:
			self.clients[client] = true
			log.Println("New client joined")

		case client := <- self.leave:
			delete(self.clients, client)
			close(client.send)
			log.Println("Client left")

		case msg := <- self.msg:
			for client := range self.clients {
				select {
				case client.send <- msg:
					// TODO
					log.Println(" -- sent to client")

				default:
					delete(self.clients, client)
					close(client.send)
					log.Println(" -- failed to send, cleaned up client")
				}
			}
		}
	}
}
