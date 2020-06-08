//-------------MOST OF THIS CODE IS MY OWN------------------------
// ----- some inspiration was taken from https://itnext.io/lets-learn-how-to-to-build-a-chat-application-with-redis-websocket-and-go-7995b5c7b5e5--------
package websocket

import (
	"fmt"

	"github.com/go-redis/redis"
)

const channel = "chat"

type Pool struct {
	Register    chan *Client
	Unregister  chan *Client
	Clients     map[*Client]bool
	Broadcast   <-chan *redis.Message
	RedisClient *redis.Client
}

func NewPool(redisClient *redis.Client) *Pool {
	sub := redisClient.Subscribe(channel)
	messages := sub.Channel()
	return &Pool{
		Register:    make(chan *Client),
		Unregister:  make(chan *Client),
		Clients:     make(map[*Client]bool),
		Broadcast:   messages,
		RedisClient: redisClient,
	}
}

const users = "chat-users"

func (pool *Pool) Start() {

	for {
		select {
		case client := <-pool.Register:

			pool.Clients[client] = true

			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: fmt.Sprintf(`%s joined the chat`, client.Username)})
			}
			break
		case client := <-pool.Unregister:

			pool.RedisClient.SRem(users, client.Username)
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: fmt.Sprintf(`%s left the chat`, client.Username)})
			}
			delete(pool.Clients, client)
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool:", message)
			for client, _ := range pool.Clients {
				if err := client.Conn.WriteJSON(message.Payload); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
