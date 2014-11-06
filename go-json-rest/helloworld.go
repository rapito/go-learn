package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
)

type Message struct {
	Body string
}

func messageHandler(w rest.ResponseWriter, req *rest.Request) {

	w.WriteJson(&Message{
	Body: "Hello Wo2rld!", })
}

func main() {
	handler := rest.ResourceHandler{}
	err := handler.SetRoutes(
	&rest.Route{"GET", "/message", messageHandler},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":8080", &handler))
}
