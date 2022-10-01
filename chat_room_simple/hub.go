package main

type Hub struct {
	clients    map[*Client]bool //维护所有Client
	register   chan *Client     // Client注册请求通过管道来接收
	unregister chan *Client     // Client注销请求通过管道来接收
	broadcast  chan []byte      // 需要广播的消息
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte), // 同步管道，确保hub这里消息不会堆积。
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true //注册client
		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok { // 防止重复注销
				delete(hub.clients, client) // 注销client
				close(client.send)          // hub从此以后不再向该client广播消息le
			}
		case msg := <-hub.broadcast:
			// 消息广播到所有client
			for client := range hub.clients {
				client.send <- msg
			}
		}
	}
}
