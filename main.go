package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web-dockerize/server"
)

func main() {
	fmt.Println("Server running on http://localhost:8080 \nTo stop the server, press Ctrl + C.")
	// Set up routes
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", server.MainHandler)
	http.HandleFunc("/layout", server.ResultHandler)
	// Start server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
