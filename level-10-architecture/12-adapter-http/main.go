package main

import (
	"fmt"
	"net/http"
)

func handlerAdapterHttp(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Adapter Http")
}

func main() {
	h := http.HandlerFunc(handlerAdapterHttp)
	fmt.Printf("handler=%T topic=Adapter Http\n", h)
}
