package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Read from .env
var CLIENT_ID = ""
var APP_SECRET = ""

var base = "https://api-m.sandbox.paypal.com"

type Token struct {
	Scope        string `json:"scope"`
	Access_token string `json:"access_token"`
	Token_type   string `json:"token_type"`
	App_id       string `json:"app_id"`
	Expires_in   string `json:"expires_in"`
	Nonce        string `json:"nonce"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
}

// Work in Progress
func createOrder() {
	//This will be generated from the generateAccessToken() func
	accessToken := generateAccessToken()

	//fmt.Println(accessToken) => it prints the access token

	body := []byte(`{intent: "CAPTURE",purchase_units: [{amount: {currency_code: "USD",value: "100.00"}}]}`)

	req, err := http.NewRequest("POST", base+"/v2/checkout/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	fmt.Println(req)
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

		// // TODO return jsonObject.access_token
		// // Unmarshal the JSON response into a map
		// var result map[string]interface{}
		// err = json.Unmarshal(bodyBytes, &result)
		fmt.Println(bodyString)
	}
}

func generateAccessToken() string {
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

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	// Unmarshal the JSON response into a map
	var jsonMap map[string]interface{}
	json.Unmarshal(bodyBytes, &jsonMap)
	value := fmt.Sprint(jsonMap["access_token"])

	return value
}

func handleRequest() {
	//http.HandleFunc("/api/orders", createOrder)
	createOrder()
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequest()
}
