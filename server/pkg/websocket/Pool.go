//-------------MOST OF THIS CODE IS MY OWN------------------------
// ----- some inspiration was taken from https://itnext.io/lets-learn-how-to-to-build-a-chat-application-with-redis-websocket-and-go-7995b5c7b5e5--------
package websocket

import (
	"encoding/json"
	"fmt"
	"log"

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

			fmt.Println()
			fmt.Printf("registered new user : %s\n\n", client.Username)

			pool.Clients[client] = true

			for poolClient, _ := range pool.Clients {
				fmt.Println(client)
				poolClient.Conn.WriteJSON(Message{Type: 1, Body: fmt.Sprintf(`%s joined the chat`, client.Username)})
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
			var newMessage Message
			err := json.Unmarshal([]byte(message.Payload), &newMessage)
			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println()
			fmt.Println("newMessage", newMessage)
			fmt.Println()
			fmt.Println("Sending message to all clients in Pool:", message)
			for client, _ := range pool.Clients {
				if client.Username == newMessage.To {
					if err := client.Conn.WriteJSON(message.Payload); err != nil {
						fmt.Println(err)
						return
					}
				}

			}
		}
	}
}
