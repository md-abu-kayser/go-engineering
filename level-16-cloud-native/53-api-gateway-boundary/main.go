package main

import (
	"fmt"
	"net/http"
)

func handlerApiGatewayBoundary(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Gateway Boundary")
}

func main() {
	h := http.HandlerFunc(handlerApiGatewayBoundary)
	fmt.Printf("handler=%T topic=Api Gateway Boundary\n", h)
}
