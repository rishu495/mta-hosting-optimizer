package util

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func populateMap(hash map[string][]HostNameIPStatus) {
	ip := [6]string{"127.0.0.1", "127.0.0.2", "127.0.0.3", "127.0.0.4", "127.0.0.5", "127.0.0.6"}
	hostName := [6]string{"mta-prod-1", "mta-prod-1", "mta-prod-2", "mta-prod-2", "mta-prod-2", "mta-prod-3"}
	active := [6]bool{true, false, true, true, false, false}

	for idx := 0; idx < len(ip); idx++ {

		hash[hostName[idx]] = append(hash[hostName[idx]], HostNameIPStatus{
			IP:     ip[idx],
			Status: active[idx],
		})
	}
}

func parseIntFromString(thresholdEnv string) (int, error) {
	// string to int
	i, err := strconv.Atoi(thresholdEnv)
	if err != nil {
		return 0, err
	}

	return i, err
}
