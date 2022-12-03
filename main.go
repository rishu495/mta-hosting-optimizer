package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getHostName(w http.ResponseWriter, r *http.Request) {

	//Setting the content Type
	w.Header().Set("Content-Type", "application/json")

	var result []string = make([]string, 0)

	//Create New String to Struct Array Map
	ipMap := make(map[string][]HostNameIPStatus)

	//Populating Map
	populateMap(ipMap)

	// godotenv package
	thresholdEnv := goDotEnvVariable("X")
	threshold, err := parseIntFromString(thresholdEnv)

	if err != nil {
		log.Println("Error converting string to Int")
		json.NewEncoder(w).Encode(GetHostNameResponse{
			Result: result,
			Status: "Error",
			Error:  err,
		})
		return
	}

	for key, val := range ipMap {
		count := 0
		for _, ipStatusHash := range val {
			if ipStatusHash.Status {
				count++
			}
		}
		if count <= threshold {
			result = append(result, key)
		}
	}
	json.NewEncoder(w).Encode(GetHostNameResponse{
		Result: result,
		Status: "OK",
		Error:  nil,
	})
	return
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/inactiveIPAddressHostname", getHostName).Methods("GET")
	fmt.Printf("Server started at port :: 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
