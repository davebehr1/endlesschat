//-------------MOST OF THIS CODE IS MY OWN------------------------
//-------------alot of it is standard bootstrapping like setting up rest routes and a redis client---
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davebehr1/anonymous-chat/pkg"
	"github.com/davebehr1/anonymous-chat/pkg/websocket"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
)

func wsHandler(pool *websocket.Pool, w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	user := strings.TrimPrefix(r.URL.Path, "/anonChat/")

	fmt.Println("WebSocket Endpoint Hit", user)
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn:     conn,
		Pool:     pool,
		Username: user,
	}

	pool.Register <- client
	client.Read(redisClient)
}

type Response struct {
	Message string `json:"message"`
	Taken   bool   `json:"taken"`
}

func setupRoutes(redisClient *redis.Client) {
	pool := websocket.NewPool(redisClient)
	go pool.Start()

	http.HandleFunc("/username/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		name := strings.TrimPrefix(req.URL.Path, "/username/")

		fmt.Println(name)
		usernameTaken, err := redisClient.SIsMember("chat-users", name).Result()
		if err != nil {
			log.Fatalf("could not retrieve value from set")
		}
		if usernameTaken {

			response := Response{Taken: true, Message: "please enter another username"}
			json.NewEncoder(rw).Encode(response)
		} else {
			pool.RedisClient.SAdd("chat-users", name)
			response2 := Response{Taken: false, Message: "Welcome"}
			json.NewEncoder(rw).Encode(response2)

		}

	})

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("welcome to websocket server!"))
	})
	http.HandleFunc("/anonChat/", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(pool, w, r, redisClient)
	})
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start chat server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := pkg.GetConfig()

		redisClient := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		defer func() {
			redisClient.Close()
		}()

		_, err = redisClient.Ping().Result()

		if err != nil {
			log.Fatalf("failed to create redis client: %v", err)
		}

		redisClient.FlushAll()

		fmt.Println("Distributed Chat App v0.01")
		setupRoutes(redisClient)

		//http.ListenAndServeTLS(":5000", "https-server.crt", "https-server.key", nil)

		http.ListenAndServeTLS(":5000", "anonymous.com+5.pem", "anonymous.com+5-key.pem", nil)

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
