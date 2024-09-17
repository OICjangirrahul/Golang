package main

import (
	"golang/route"
	"fmt"
	"net/http"
)


func main() {
	app := route.SetupRoutes()

	
	err := http.ListenAndServe(":8081", app)

	if err != nil {
		fmt.Println("Server failed to start:", err)
	} else {
		fmt.Println("Server started successfully on :8081")
	}

}