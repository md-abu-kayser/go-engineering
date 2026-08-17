package main

import (
	"fmt"
	"net/http"
)

func handlerWebsocketConcept(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Websocket Concept")
}

func main() {
	h := http.HandlerFunc(handlerWebsocketConcept)
	fmt.Printf("handler=%T topic=Websocket Concept\n", h)
}
