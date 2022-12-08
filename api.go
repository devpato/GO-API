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

type CreateOrder struct {
	Id     string  `json:"id"`
	Status string  `json:"status"`
	Links  []Links `json:"links"`
}

type Links struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage Endpoint Hit")
}

// Work in Progress
func createOrder(w http.ResponseWriter, r *http.Request) {
	accessToken := generateAccessToken()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Println(accessToken)

	body := []byte(`{
		"intent":"CAPTURE",
		"purchase_units":[
		   {
			  "amount":{
				 "currency_code":"USD",
				 "value":"100.00"
			  }
		   }
		]
	 }`)

	req, err := http.NewRequest("POST", base+"/v2/checkout/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var jsonMap CreateOrder

	error := json.NewDecoder(resp.Body).Decode(&jsonMap)

	if error != nil {
		log.Fatal(err)
	}

	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(jsonMap)

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
	access_token := fmt.Sprint(jsonMap["access_token"])

	return access_token
}

func handleRequest() {
	http.HandleFunc("/api/orders", createOrder)
	http.HandleFunc("api/orders/:orderID/capture", capturePayment)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func main() {
	handleRequest()
}
