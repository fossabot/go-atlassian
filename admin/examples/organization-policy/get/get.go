package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ctreminiom/go-atlassian/admin"
	"log"
	"os"
)

func main() {

	//ATLASSIAN_ADMIN_TOKEN
	var apiKey = os.Getenv("ATLASSIAN_ADMIN_TOKEN")

	cloudAdmin, err := admin.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	cloudAdmin.Auth.SetBearerToken(apiKey)
	cloudAdmin.Auth.SetUserAgent("curl/7.54.0")

	var (
		organizationID = "9a1jj823-jac8-123d-jj01-63315k059cb2"
		policyID       = "60f0f660-be3e-4d70-bd34-9c2858ec040f"
	)

	policy, response, err := cloudAdmin.Organization.Policy.Get(context.Background(), organizationID, policyID)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
		}
		log.Fatal(err)
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	policyAsJSONKeys, err := json.MarshalIndent(policy, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("MarshalIndent Struct keys output\n %s\n", string(policyAsJSONKeys))
}
