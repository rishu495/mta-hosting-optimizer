package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishu495/mta-hosting-optimizer/usecase"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/inactiveIPAddressHostname", usecase.GetHostName).Methods("GET")
	fmt.Printf("Server started at port :: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
