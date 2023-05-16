package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// write response from public/index.html
		http.ServeFile(w, r, "public/index.html")
	})
}
