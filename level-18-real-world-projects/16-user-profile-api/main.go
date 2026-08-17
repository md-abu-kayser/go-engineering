package main

import (
	"fmt"
	"net/http"
)

func handlerUserProfileApi(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "User Profile Api")
}

func main() {
	h := http.HandlerFunc(handlerUserProfileApi)
	fmt.Printf("handler=%T topic=User Profile Api\n", h)
}
