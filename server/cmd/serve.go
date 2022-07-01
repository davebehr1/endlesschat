package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davebehr1/endlesschat/pkg"
	"github.com/davebehr1/endlesschat/pkg/auth"
	"github.com/davebehr1/endlesschat/pkg/storage/postgress"
	"github.com/davebehr1/endlesschat/pkg/websocket"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
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

func setupRoutes(r *mux.Router, redisClient *redis.Client, logger *log.Logger) {

	pool := websocket.NewPool(redisClient)
	go pool.Start()

	r.HandleFunc("/username/{name}", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Content-Type", "application/json")

		vars := mux.Vars(req)

		logger.Printf("userame request: %s", vars)

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
		logger.Print(req.RequestURI)
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Write([]byte(fmt.Sprintf("welcome to websocket server!: %s", req.RequestURI)))
	})

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		logger.Print(r.RequestURI)

		wsHandler(pool, w, r, redisClient)
	})

	r.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
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

		var log = logrus.New()

		loglevel, err := logrus.ParseLevel(cfg.LogLevel)
		if err != nil {
			log.WithError(err).Error("Invalid loglevel")
			os.Exit(1)
		}
		log.SetLevel(loglevel)

		postgresDB := postgress.New(log, pkg.GetPostgresConnectionString())

		postgresDB.RunMigrations()

		r := mux.NewRouter()
		// r.Use(mux.CORSMethodMiddleware(r))
		c := cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
		})

		handler := c.Handler(r)

		setupRoutes(r, redisClient, logger)

		fmt.Println("Distributed Chat App v0.01")
		http.ListenAndServe(":5003", handler)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
