package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func serveDefault(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("No Port In Heroku : " + port)
	}
	return ":" + port
}

func main() {
	hub := H
	go hub.Run()
	http.HandleFunc("/", serveDefault)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(w, r)
	})
	//Listerning on port
	log.Fatal(http.ListenAndServe(getPort(), nil))
}
