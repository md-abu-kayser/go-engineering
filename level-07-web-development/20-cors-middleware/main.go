// Lesson 20: Cors Middleware
//
// Goal: Exercise an HTTP handler in memory, including its status, headers,
// and JSON response, instead of binding a port for a teaching example.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

const lesson = "Cors Middleware"

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "hello, " + r.URL.Query().Get("name")})
}

func main() {
	request := httptest.NewRequest(http.MethodGet, "/greeting?name=Asha", nil)
	recorder := httptest.NewRecorder()
	greetingHandler(recorder, request)

	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("status: %d body: %s", recorder.Code, recorder.Body.String())
}