package main

import (
	"fmt"
	"net/http"
)

func handlerBackwardCompatibleApi(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Backward Compatible Api")
}

func main() {
	h := http.HandlerFunc(handlerBackwardCompatibleApi)
	fmt.Printf("handler=%T topic=Backward Compatible Api\n", h)
}
