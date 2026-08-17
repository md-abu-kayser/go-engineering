package main

import (
	"fmt"
	"net/http"
)

func handlerWebsocketHeartbeat(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Websocket Heartbeat")
}

func main() {
	h := http.HandlerFunc(handlerWebsocketHeartbeat)
	fmt.Printf("handler=%T topic=Websocket Heartbeat\n", h)
}
