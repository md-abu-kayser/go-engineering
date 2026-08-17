package main

import (
	"fmt"
	"net/http"
)

func handlerWebhookHmac(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Webhook Hmac")
}

func main() {
	h := http.HandlerFunc(handlerWebhookHmac)
	fmt.Printf("handler=%T topic=Webhook Hmac\n", h)
}
