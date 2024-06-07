package websocket

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Interface interface {
	ServeWS(ctx *gin.Context)
	RunWebsocket()
}

type Message struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type Hub struct {
	clients          map[*websocket.Conn]*ClientInfo
	RegisterClient   chan *websocket.Conn
	RemovalClient    chan *websocket.Conn
	BroadCastMessage chan Message
}

type ClientInfo struct {
	Name string
	Conn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}


func Init() Interface {
	return &Hub{
		clients:          make(map[*websocket.Conn]*ClientInfo),
		RegisterClient:   make(chan *websocket.Conn),
		RemovalClient:    make(chan *websocket.Conn),
		BroadCastMessage: make(chan Message),
	}
}

// ServeWS implements Interface.
func (h *Hub) ServeWS(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		h.RemovalClient <- conn
		_ = conn.Close()
	}()

	h.RegisterClient <- conn


	client := &ClientInfo{
		Name:  ctx.Request.URL.Query().Get("name"),
		Conn: conn,
	}
	h.clients[conn] = client

	for {
		var msg Message
		msg.Name = client.Name
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			break
		}

		h.BroadCastMessage <- msg
	}}


func (h *Hub) RunWebsocket() {
	for {
		select {
		case conn := <-h.RegisterClient:
			clientInfo := &ClientInfo{}
			h.clients[conn] = clientInfo
		case conn := <-h.RemovalClient:
			delete(h.clients, conn)
		case msg := <-h.BroadCastMessage:
			for conn := range h.clients {
				err := conn.WriteJSON(msg)
				if err != nil {
					log.Printf("error: %v", err)
					conn.Close()
					delete(h.clients, conn)
				}
			}
		}
	}
}

