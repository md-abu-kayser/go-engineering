package main

import (
	"fmt"
	"net/http"
)

func handlerWebhookService(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Webhook Service")
}

func main() {
	h := http.HandlerFunc(handlerWebhookService)
	fmt.Printf("handler=%T topic=Webhook Service\n", h)
}
