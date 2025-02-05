package main

import (
	"context"
	"encoding/json"
	"github.com/ctreminiom/go-atlassian/jira"
	"log"
	"os"
)

func main() {

	var (
		host  = os.Getenv("HOST")
		mail  = os.Getenv("MAIL")
		token = os.Getenv("TOKEN")
	)

	atlassian, err := jira.New(nil, host)
	if err != nil {
		return
	}

	atlassian.Auth.SetBasicAuth(mail, token)
	atlassian.Auth.SetUserAgent("curl/7.54.0")

	var (
		issueKey = "DESK-11"
		expand   = []string{"serviceDesk", "requestType"}
	)

	request, response, err := atlassian.ServiceManagement.Request.Get(context.Background(), issueKey, expand)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	dataAsJson, err := json.MarshalIndent(request, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(dataAsJson))

}
