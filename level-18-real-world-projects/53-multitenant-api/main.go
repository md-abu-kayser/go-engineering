package main

import (
	"fmt"
	"net/http"
)

func handlerMultitenantApi(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Multitenant Api")
}

func main() {
	h := http.HandlerFunc(handlerMultitenantApi)
	fmt.Printf("handler=%T topic=Multitenant Api\n", h)
}
