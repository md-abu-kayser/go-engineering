package main

import (
	"fmt"
	"net/http"
)

func handlerPackageApis(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Package Apis")
}

func main() {
	h := http.HandlerFunc(handlerPackageApis)
	fmt.Printf("handler=%T topic=Package Apis\n", h)
}
