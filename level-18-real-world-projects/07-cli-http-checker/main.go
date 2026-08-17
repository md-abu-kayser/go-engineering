package main

import (
	"fmt"
	"net/http"
)

func handlerCliHttpChecker(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Cli Http Checker")
}

func main() {
	h := http.HandlerFunc(handlerCliHttpChecker)
	fmt.Printf("handler=%T topic=Cli Http Checker\n", h)
}
