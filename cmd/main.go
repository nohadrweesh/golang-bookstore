//routes --> Controll --> models
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nohadrweesh/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Printf("Starting Server at port 9010\n")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
