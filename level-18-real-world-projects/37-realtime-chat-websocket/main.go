package main

import (
	"fmt"
	"net/http"
)

func handlerRealtimeChatWebsocket(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Realtime Chat Websocket")
}

func main() {
	h := http.HandlerFunc(handlerRealtimeChatWebsocket)
	fmt.Printf("handler=%T topic=Realtime Chat Websocket\n", h)
}
