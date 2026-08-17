package main

import (
	"fmt"
	"net/http"
)

func handlerTemplateEscaping(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Template Escaping")
}

func main() {
	h := http.HandlerFunc(handlerTemplateEscaping)
	fmt.Printf("handler=%T topic=Template Escaping\n", h)
}
