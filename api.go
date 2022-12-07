package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Read from .env
var CLIENT_ID = "CLIENT ID"
var APP_SECRET = "APP_SECRET"

var base = "https://api-m.sandbox.paypal.com"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
}

func createOrder() {
	// This will be generated from the generateAccessToken() func
	// accessToken := "AT"

	// body := []byte("{
	// 	intent: "CAPTURE",
	// 	purchase_units: [
	// 	  {
	// 		amount: {
	// 		  currency_code: "USD",
	// 		  value: "100.00",
	// 		}
	// 	  }
	// 	]
	//   }")

	// req, err := http.NewRequest("POST", base+"/v2/checkout/orders", bytes.NewBuffer(body))
	// req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer "+auth)
}

func generateAccessToken() {
	auth := base64.StdEncoding.EncodeToString([]byte(CLIENT_ID + ":" + APP_SECRET))

	body := []byte("grant_type=client_credentials")

	req, err := http.NewRequest("POST", base+"/v1/oauth2/token", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	fmt.Println(resp.StatusCode)
	if resp.StatusCode == http.StatusOK {

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		bodyString := string(bodyBytes)
		fmt.Println(bodyString)

		// TODO return jsonObject.access_token
	}
}

func handleRequest() {
	// http.HandleFunc("/api/orderse", createOrder)
	generateAccessToken()
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequest()
}
