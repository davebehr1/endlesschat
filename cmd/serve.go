package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/davebehr1/anonymous-chat/pkg"
	"github.com/davebehr1/anonymous-chat/pkg/websocket"
	"github.com/go-redis/redis"
	"github.com/spf13/cobra"
)

func webSocketHandler(pool *websocket.Pool, w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	user := strings.TrimPrefix(r.URL.Path, "/chat/")

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

func setupRoutes(redisClient *redis.Client) {
	pool := websocket.NewPool(redisClient)
	go pool.Start()

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("you are good to go!"))
	})
	http.HandleFunc("/chat/", func(w http.ResponseWriter, r *http.Request) {
		// surveyID := r.URL.Query()["username"]
		// fmt.Println(surveyID)
		webSocketHandler(pool, w, r, redisClient)
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
		defer redisClient.Close()

		_, err = redisClient.Ping().Result()

		if err != nil {
			log.Fatalf("failed to create redis client: %v", err)
		}

		fmt.Println("Distributed Chat App v0.01")
		setupRoutes(redisClient)

		//http.ListenAndServeTLS(":5000", "https-server.crt", "https-server.key", nil)

		http.ListenAndServeTLS(":5000", "anonymous.com+5.pem", "anonymous.com+5-key.pem", nil)

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
