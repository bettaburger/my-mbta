package controllers 

import (
	"encoding/json"
	//"fmt"
	"log"
	"os"
	"io"
	"net/http"
	//"github.com/rs/cors"
	"github/bettaburger/my-mbta/model"
)

type Datas struct {
	Data [] model.Alert `json:"data"`
}

// Read request the MBTA api, parse the json response into Data struct
// Display the alerts to port 3000
func AlertHandler (writer http.ResponseWriter, req *http.Request) {
	
	// Allow CORS to the origin
	writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	urlStr := "https://api-v3.mbta.com/alerts"
	var client http.Client
	response, err := client.Get(urlStr)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Close the response body 
	defer response.Body.Close()

	// Request is sucessful (200OK)
	if response.StatusCode == http.StatusOK {
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		// Parse the json response into data
		var data Datas
		err = json.Unmarshal(responseBody, &data)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		
		// Display the alerts in json
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(data)
	}
}







// Request the mbta API and parse the response and return it. 
/*func AlertHandler(url string) *Data {
	response, err := http.Get(url)
	// If the call to http fails, returns the http error to the caller. 
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Close the response body 
	defer response.Body.Close()
	// Read the response
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// Then parse the JSON response into the Data.
	var data Data
	json.Unmarshal(responseBody, &data)
	
	return &data
}

func displayAlert(data *Data) {
	for _, alert := range data.Data {
		fmt.Printf("ID: %s\n", alert.ID)
		fmt.Printf("Alert Header: %s\n", alert.Attributes.Header)
		fmt.Printf("Severity level: %d\n", alert.Attributes.Severity)
		fmt.Printf("Cause: %s\n", alert.Attributes.Cause)
		fmt.Printf("Service Effect: %s\n", alert.Attributes.ServiceEffect)
		fmt.Printf("Created: %s\n", alert.Attributes.CreatedAt)
		fmt.Println()
	}
}*/


