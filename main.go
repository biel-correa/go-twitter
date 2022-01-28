package main

import (
	"fmt"
	"net/http"
	Home "example/go-twitter/pages/home"
)

func main() {
	GenerateRoutes()

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func GenerateRoutes() {
	http.HandleFunc("/", Home.Home)
}