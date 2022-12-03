package usecase

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/rishu495/mta-hosting-optimizer/model"
	"github.com/rishu495/mta-hosting-optimizer/util"
)

func GetHostName(w http.ResponseWriter, r *http.Request) {

	//Setting the content Type
	w.Header().Set("Content-Type", "application/json")

	var result []string = make([]string, 0)

	//Create New String to Struct Array Map
	ipMap := make(map[string][]model.HostNameIPStatus)

	//Populating Map
	util.PopulateMap(ipMap)

	// godotenv package
	thresholdEnv := util.GoDotEnvVariable("X")

	//Convert String into int
	threshold, err := util.ParseIntFromString(thresholdEnv)

	if err != nil {
		log.Println("Error converting string to Int")
		json.NewEncoder(w).Encode(model.GetHostNameResponse{
			Result: result,
			Status: "Error",
			Error:  err,
		})
		return
	}

	//Business Logic
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

	//Encoding result into JSON
	json.NewEncoder(w).Encode(model.GetHostNameResponse{
		Result: result,
		Status: "OK",
		Error:  nil,
	})
	return
}
