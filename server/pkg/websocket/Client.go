//-------------MOST OF THIS CODE IS MY OWN------------------------
// ----- some inspiration was taken from https://medium.com/@nqbao/writing-a-chat-server-in-go-3b61ccc2a8ed--------
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
	From string `json:"from"`
	To   string `json:"to"`
}

type webSocketMessage struct {
	Message string `json:"message"`
	To      string `json:"to"`
}

func (c *Client) Read(redisClient *redis.Client) {

	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var wsMessage webSocketMessage
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		err = json.Unmarshal(p, &wsMessage)
		if err != nil {
			log.Println(err)
			return
		}
		message := Message{Type: messageType, Body: wsMessage.Message, From: c.Username, To: wsMessage.To}
		textMessage, err := json.Marshal(message)
		if err != nil {
			log.Println(err)
			return
		}
		err = redisClient.Publish(channel, textMessage).Err()
		if err != nil {
			log.Println("something went wrong publishing to channel", err)
		}
		fmt.Printf("Message Received: %+v", message)
	}
}
