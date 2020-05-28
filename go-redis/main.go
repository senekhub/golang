package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Universe.")
}
func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	client.Set("nuevo2", "nuevo2", 0)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
