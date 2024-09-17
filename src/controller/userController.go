package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type UserRequest struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}


func GetUser(w http.ResponseWriter, r *http.Request) {
	var requestData UserRequest

    // Read and unmarshal the request body
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&requestData)

    // Log the request data
    fmt.Printf("Request Data: %+v\n", requestData.Age)
	fmt.Fprintf(w, `{"user": 3}`)
}
