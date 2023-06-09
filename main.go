package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// main is the main entrance to the application
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)

}
