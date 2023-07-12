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

	scheme := "https://app.o95.info"
	url := fmt.Sprintf("%s%s", scheme, r.URL)

	fmt.Fprintln(w, url)
	pdf.Pdf(url)
}

// test
func main() {
	log.Println("Create router")
	router := httprouter.New()
	router.GET("/receipt", Pdf)

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
