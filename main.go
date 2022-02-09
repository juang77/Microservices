package main

import (
	"Microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBay(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbay", gh)

	http.ListenAndServe(":9090", sm)
}
