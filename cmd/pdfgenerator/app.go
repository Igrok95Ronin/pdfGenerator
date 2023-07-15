package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"pdgGenerator/internal/pdf"
	"time"
)

func Pdf(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	queryParam := r.URL.Query().Get("url")

	fmt.Println(queryParam)

	pdf.Pdf(queryParam, w)

}

// test
func main() {
	log.Println("Create router")
	router := httprouter.New()
	router.GET("/pdf", Pdf)

	start(router)
}

func start(router *httprouter.Router) {
	const (
		port = ":1234"
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
