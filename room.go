package main

import (
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
	"sync"
)

const (
	socketBuffSize    = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:  socketBuffSize,
	WriteBufferSize: messageBufferSize,
}

/*
 * Rooms collection
 */
type rooms struct {
	lock  sync.RWMutex
	rooms map[string]*room
}

func newRooms() *rooms {
	return &rooms{rooms: make(map[string]*room)}
}

func (self *rooms) get(name string) *room {
	self.lock.RLock();
	room, ok := self.rooms[name]
	if !ok {
		self.lock.RUnlock()
		self.lock.Lock()
		room = newRoom(name)
		self.rooms[name] = room
		self.lock.Unlock()
		go room.run()
	}

	return room
}

/*
 * Single room
 */
type room struct {
	name    string
	msg     chan *clientMsg
	join    chan *client
	leave   chan *client
	clients map[*client]bool
}

func newRoom(name string) *room {
	return &room{
		name:    name,
		msg:     make(chan *clientMsg),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (self *room) run() {
	log.Printf("Running room %s\n", self.name)
	for {
		select {
		case client := <-self.join:
			self.clients[client] = true
			log.Println("New client joined")

		case client := <-self.leave:
			delete(self.clients, client)
			close(client.send)
			log.Println("Client left")

		case msg := <-self.msg:
			response := self.processMessage(msg)
			for client := range self.clients {
				select {
				case client.send <- response:
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

func (self *room) processMessage(clientMsg *clientMsg) []byte {
	msg := Msg{}

	if err := json.Unmarshal(clientMsg.msg, &msg); err != nil {
		switch msg.Cmd {
		case "vote":
			log.Printf(" -- received vote: %s, %s\n", msg.Vote, clientMsg.msg)
			// save the vote, send the MsgSendVote message
			clientMsg.client.vote = msg.Vote
			sendVote, err := json.Marshal(&Msg{Cmd: "vote", User: clientMsg.client.name})
			if err == nil {
				return sendVote
			}
			log.Printf(" -- error encoding message: %s\n", err)

		case "storyDesc":
			sendStoryDesc, err := json.Marshal(&Msg{Cmd: "storyDesc", StoryDesc: msg.StoryDesc})
			if err == nil {
				return sendStoryDesc
			}
			log.Printf(" -- error encoding message: %s\n", err)

		default:
			log.Printf(" -- unknown command: %s\n", msg.Cmd)
		}

	} else {
		log.Printf(" -- error decoding message: %s\n", err)
	}

	return clientMsg.msg
}
