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
			fmt.Println("CLIENT:", client.Username)

			usernameTaken, err := pool.RedisClient.SIsMember(users, client.Username).Result()
			if err != nil {
				fmt.Println(err, "ERRRROR")
			}
			fmt.Println(usernameTaken)

			pool.Clients[client] = true
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{Type: 1, Body: fmt.Sprintf(`%s joined the chat`, client.Username)})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("Size of Connection Pool: ", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(Message{Type: 1, Body: fmt.Sprintf(`%s left the chat`, client.Username)})
			}
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
