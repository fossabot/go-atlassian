package main

import (
	"context"
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

	/*
		We can add different share permissions, for example:

		---- Project ID only
		payload := jira.PermissionFilterBodyScheme{
				Type:      "project",
				ProjectID: "10000",
			}

		---- Project ID and role ID
		payload := jira.PermissionFilterBodyScheme{
				Type:          "project",
				ProjectID:     "10000",
				ProjectRoleID: "222222",
			}

		==== Group Name
		payload := jira.PermissionFilterBodyScheme{
				Type:          "group",
				GroupName: "jira-users",
			}
	*/

	payload := jira.PermissionFilterBodyScheme{
		Type:      "project",
		ProjectID: "10000",
	}

	permissions, response, err := atlassian.Filter.Share.Add(context.Background(), 1001, &payload)
	if err != nil {
		if response != nil {
			log.Println("Response HTTP Response", string(response.BodyAsBytes))
			log.Println("Response HTTP Code", response.StatusCode)
			log.Println("HTTP Endpoint Used", response.Endpoint)
		}
		return
	}

	log.Println("Response HTTP Code", response.StatusCode)
	log.Println("HTTP Endpoint Used", response.Endpoint)

	for index, permission := range *permissions {
		log.Println(index, permission.ID, permission.Type)
	}
}
