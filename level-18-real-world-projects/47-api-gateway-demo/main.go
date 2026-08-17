package main

import (
	"fmt"
	"net/http"
)

func handlerApiGatewayDemo(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Gateway Demo")
}

func main() {
	h := http.HandlerFunc(handlerApiGatewayDemo)
	fmt.Printf("handler=%T topic=Api Gateway Demo\n", h)
}
