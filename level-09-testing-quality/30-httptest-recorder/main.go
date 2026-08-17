package main

import (
	"fmt"
	"net/http"
)

func handlerHttptestRecorder(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Httptest Recorder")
}

func main() {
	h := http.HandlerFunc(handlerHttptestRecorder)
	fmt.Printf("handler=%T topic=Httptest Recorder\n", h)
}
