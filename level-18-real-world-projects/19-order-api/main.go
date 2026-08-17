package main

import (
	"fmt"
	"net/http"
)

func handlerOrderApi(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Order Api")
}

func main() {
	h := http.HandlerFunc(handlerOrderApi)
	fmt.Printf("handler=%T topic=Order Api\n", h)
}
