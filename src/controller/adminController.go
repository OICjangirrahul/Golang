package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type AdminRequest struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}


func GetAdmin(w http.ResponseWriter, r *http.Request) {
	var requestData AdminRequest

    // Read and unmarshal the request body
    decoder := json.NewDecoder(r.Body)
    decoder.Decode(&requestData)

    // Log the request data
    fmt.Printf("Request Data: %+v\n", requestData.Age)
	fmt.Fprintf(w, `{"Admin": 1}`)
}

