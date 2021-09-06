package main

import (
	"fmt"
	"github.com/bdinesh/sample-go-webapp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	registerHandlers()
	fmt.Println("Starting the sample web app on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}

type HandlerDetail struct {
	Pattern     string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func registerHandlers() {
	handlerDetails := []HandlerDetail{
		//{Pattern: "*", HandlerFunc: http.NotFound},
		{Pattern: "/", HandlerFunc: handlers.Home},
		{Pattern: "/about", HandlerFunc: handlers.About},
	}

	for _, h := range handlerDetails {
		http.HandleFunc(h.Pattern, h.HandlerFunc)
	}
}
