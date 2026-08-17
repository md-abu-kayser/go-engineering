package main

import (
	"fmt"
	"net/http"
)

func handlerApiVersioning(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Versioning")
}

func main() {
	h := http.HandlerFunc(handlerApiVersioning)
	fmt.Printf("handler=%T topic=Api Versioning\n", h)
}
