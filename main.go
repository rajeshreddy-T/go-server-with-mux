// sample go http server with mux
package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"

	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	fmt.Println("Initialize the server on  port 8080")
	log.Fatal(http.ListenAndServe(":8084", r))

	fmt.Println("Server started on port 8080")
}
