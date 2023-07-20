package main

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"pdgGenerator/internal/handlers"
	"time"
)

// test
func main() {
	log.Println("Create router")
	router := httprouter.New()
	router.GET("/generate_pdf", handlers.Pdf)

	start(router)
}

func start(router *httprouter.Router) {
	const (
		port = ":8080"
	)
	log.Println("Start application")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server is listening port " + port)
	log.Fatal(server.Serve(listener))
}
