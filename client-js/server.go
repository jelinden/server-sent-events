package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/r3labs/sse"
)

var server *sse.Server
var indexPage []byte

func main() {
	indexPage = readIndexPage()
	server = sse.New()
	server.AutoReplay = false
	server.CreateStream("messages")

	mux := http.NewServeMux()
	mux.HandleFunc("/events", sseHandler)
	mux.HandleFunc("/index.html", indexHandler)
	go dataSender(server)

	log.Println("starting server at port 8080")
	http.ListenAndServe(":8080", mux)
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("new client", r.Header.Get("X-Forwarded-For"))
	server.HTTPHandler(w, r)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write(indexPage)
}

func dataSender(server *sse.Server) {
	i := 0
	for {
		server.Publish("messages", &sse.Event{
			Data: []byte(strconv.Itoa(i)),
		})
		i++
		time.Sleep(100 * time.Millisecond)
	}
}

func readIndexPage() []byte {
	indexFile, err := ioutil.ReadFile("index.html")
	if err != nil {
		log.Fatalln("Reading index.html failed. Exiting.")

	}
	return indexFile
}
