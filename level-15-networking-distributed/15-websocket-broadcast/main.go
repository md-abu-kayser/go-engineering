package main

import (
	"fmt"
	"net/http"
)

func handlerWebsocketBroadcast(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Websocket Broadcast")
}

func main() {
	h := http.HandlerFunc(handlerWebsocketBroadcast)
	fmt.Printf("handler=%T topic=Websocket Broadcast\n", h)
}
