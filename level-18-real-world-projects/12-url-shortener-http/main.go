package main

import (
	"fmt"
	"net/http"
)

func handlerUrlShortenerHttp(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Url Shortener Http")
}

func main() {
	h := http.HandlerFunc(handlerUrlShortenerHttp)
	fmt.Printf("handler=%T topic=Url Shortener Http\n", h)
}
