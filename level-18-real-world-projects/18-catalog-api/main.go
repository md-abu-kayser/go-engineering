package main

import (
	"fmt"
	"net/http"
)

func handlerCatalogApi(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Catalog Api")
}

func main() {
	h := http.HandlerFunc(handlerCatalogApi)
	fmt.Printf("handler=%T topic=Catalog Api\n", h)
}
