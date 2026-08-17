package main

import (
	"fmt"
	"net/http"
)

func handlerWebhookSignatures(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Webhook Signatures")
}

func main() {
	h := http.HandlerFunc(handlerWebhookSignatures)
	fmt.Printf("handler=%T topic=Webhook Signatures\n", h)
}
