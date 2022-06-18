//-------------MOST OF THIS CODE IS MY OWN------------------------
//-------------alot of it is standard bootstrapping like setting up rest routes and a redis client---
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/davebehr1/endlesschat/pkg"
	"github.com/davebehr1/endlesschat/pkg/auth"
	"github.com/davebehr1/endlesschat/pkg/websocket"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

func wsHandler(pool *websocket.Pool, w http.ResponseWriter, r *http.Request, redisClient *redis.Client) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := "david"
	fmt.Println("WebSocket Endpoint Hit", user)

	// c, err := r.Cookie("token")
	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		// If the cookie is not set, return an unauthorized status
	// 		w.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	// For any other type of error, return a bad request status
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }
	// fmt.Println(c)

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

func setupRoutes(r *mux.Router, redisClient *redis.Client) {

	pool := websocket.NewPool(redisClient)
	go pool.Start()

	r.HandleFunc("/v1/username/{name}", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(req)

		name := vars["name"]

		fmt.Println(name)
		usernameTaken, err := redisClient.SIsMember("chat-users", name).Result()
		if err != nil {
			log.Fatalf("could not retrieve value from set")
		}
		if usernameTaken {

			response := auth.Response{Taken: true, Message: "please enter another username"}
			json.NewEncoder(rw).Encode(response)
		} else {
			pool.RedisClient.SAdd("chat-users", name)
			response2 := auth.Response{Taken: false, Message: "Welcome"}
			json.NewEncoder(rw).Encode(response2)

		}

	})

	r.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Write([]byte("welcome to websocket server!"))
	})

	r.HandleFunc("/v1/ws", func(w http.ResponseWriter, r *http.Request) { wsHandler(pool, w, r, redisClient) })

	r.HandleFunc("/v1/signin", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Println("yes boet")
		auth.Signin(w, r, redisClient, pool)
	}).Methods("POST")

	// http.HandleFunc("/welcome", auth.Welcome)
	// http.HandleFunc("/refresh", auth.Refresh)
}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start chat server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := pkg.GetConfig()

		logger := log.Default()
		logger.Printf("redis config: host %s, port %d, password %s", cfg.Redis.Host, cfg.Redis.Port, cfg.Redis.Password)

		redisClient := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
			Password: cfg.Redis.Password, // no password set
			DB:       0,                  // use default DB
		})

		defer func() {
			redisClient.Close()
		}()

		_, err = redisClient.Ping().Result()

		if err != nil {
			log.Fatalf("failed to create redis client: %v", err)
		}

		redisClient.FlushAll()

		r := mux.NewRouter()
		r.Use(mux.CORSMethodMiddleware(r))

		setupRoutes(r, redisClient)

		//http.ListenAndServeTLS(":5000", "https-server.crt", "https-server.key", nil)

		//http.ListenAndServeTLS(":5002", "anonymous.com+5.pem", "anonymous.com+5-key.pem", nil)
		fmt.Println("Distributed Chat App v0.01")
		http.ListenAndServe(":5002", r)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
