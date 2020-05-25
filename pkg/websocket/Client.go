package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID       string
	Conn     *websocket.Conn
	Pool     *Pool
	Username string
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
	User string `json:"user`
}

func (c *Client) Read(redisClient *redis.Client) {

	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: string(p), User: c.Username}
		textMessage, err := json.Marshal(message)
		err = redisClient.Publish(channel, textMessage).Err()
		if err != nil {
			log.Println("could not publish to channel", err)
		}

		// c.Pool.Broadcast <- message
		fmt.Printf("Message Received: %+v\n", message)
	}
}
