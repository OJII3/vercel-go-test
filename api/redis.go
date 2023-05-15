package handler

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

var client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379", // use default Addr
	Password: "",               // no password set
	DB:       0,                // use default DB
	TLSConfig: &tls.Config{
		MinVersion: tls.VersionTLS12,
	},
})

func Reis(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")

	// set a foo key on upstash value
	client.Set(ctx, "foo", "value-of-upstash-redis", 0)

	// get the foo key from upstash
	foo := client.Get(ctx, "foo")

	resp := make(map[string]string)
	resp["foo"] = foo.Val()
	resp["github"] = "https://ojii3.vercel.app"
	body, err := json.Marshal(resp)

	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		w.Write(body)
	}
}
