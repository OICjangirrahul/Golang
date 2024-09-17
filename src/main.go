package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UserRequest struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func getAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"user": 4}`)
}
func getAllUser2(w http.ResponseWriter, r *http.Request) {
	var requestData UserRequest

    // Read and unmarshal the request body
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&requestData)

    // Log the request data
    fmt.Printf("Request Data: %+v\n", requestData.Age)
	fmt.Fprintf(w, `{"user": 433}`)
}

func main() {
	app := http.NewServeMux()
	app.HandleFunc("GET  /api",  getAllUser)
	app.HandleFunc("POST  /api3",  getAllUser2)
	err := http.ListenAndServe(":8081", app)

	if err != nil {
		fmt.Println("Server failed to start:", err)
	} else {
		fmt.Println("Server started successfully on :8081")
	}

}