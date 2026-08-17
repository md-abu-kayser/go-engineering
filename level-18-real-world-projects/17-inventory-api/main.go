package main

import (
	"fmt"
	"net/http"
)

func handlerInventoryApi(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Inventory Api")
}

func main() {
	h := http.HandlerFunc(handlerInventoryApi)
	fmt.Printf("handler=%T topic=Inventory Api\n", h)
}
