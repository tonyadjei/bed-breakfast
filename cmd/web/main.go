package main

import (
	"fmt"
	"net/http"

	"github.com/tonyadjei/bed-breakfast/pkg/handlers"
)

const portNumber = ":8080"




// main is the main application function (entry point)
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
