package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type ScreenService struct {
	client *Client
	Tab    *ScreenTabService
	Scheme *ScreenSchemeService
}

type ScreenScheme struct {
	ID          int                            `json:"id,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
}

type ScreenFieldPageScheme struct {
	Self       string                 `json:"self,omitempty"`
	NextPage   string                 `json:"nextPage,omitempty"`
	MaxResults int                    `json:"maxResults,omitempty"`
	StartAt    int                    `json:"startAt,omitempty"`
	Total      int                    `json:"total,omitempty"`
	IsLast     bool                   `json:"isLast,omitempty"`
	Values     []*ScreenWithTabScheme `json:"values,omitempty"`
}

type ScreenWithTabScheme struct {
	ID          int                            `json:"id,omitempty"`
	Name        string                         `json:"name,omitempty"`
	Description string                         `json:"description,omitempty"`
	Scope       *TeamManagedProjectScopeScheme `json:"scope,omitempty"`
	Tab         *ScreenTabScheme               `json:"tab,omitempty"`
}

// Returns a paginated list of the screens a field is used in.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens-for-a-field
func (s *ScreenService) Fields(ctx context.Context, fieldID string, startAt, maxResults int) (result *ScreenFieldPageScheme, response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid fieldID value ")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/screens?%v", fieldID, params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenFieldPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ScreenSearchPageScheme struct {
	Self       string          `json:"self,omitempty"`
	MaxResults int             `json:"maxResults,omitempty"`
	StartAt    int             `json:"startAt,omitempty"`
	Total      int             `json:"total,omitempty"`
	IsLast     bool            `json:"isLast,omitempty"`
	Values     []*ScreenScheme `json:"values,omitempty"`
}

// Returns a paginated list of all screens or those specified by one or more screen IDs.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#get-screens
func (s *ScreenService) Gets(ctx context.Context, screenIDs []int, startAt, maxResults int) (result *ScreenSearchPageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, screenID := range screenIDs {
		params.Add("id", strconv.Itoa(screenID))
	}

	var endpoint = fmt.Sprintf("rest/api/3/screens?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenSearchPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Creates a screen with a default field tab.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#create-screen
func (s *ScreenService) Create(ctx context.Context, name, description string) (result *ScreenScheme, response *Response, err error) {

	if len(name) == 0 {
		return nil, nil, fmt.Errorf("error, please project a valid screen name value")
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = "rest/api/3/screens"

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Adds a field to the default tab of the default screen.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#add-field-to-default-screen
func (s *ScreenService) AddToDefault(ctx context.Context, fieldID string) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/screens/addToDefault/%v", fieldID)

	request, err := s.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Updates a screen. Only screens used in classic projects can be updated.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#update-screen
func (s *ScreenService) Update(ctx context.Context, screenID int, name, description string) (result *ScreenScheme, response *Response, err error) {

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v", screenID)

	request, err := s.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Deletes a screen.
// A screen cannot be deleted if it is used in a screen scheme,
// workflow, or workflow draft. Only screens used in classic projects can be deleted.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#delete-screen
func (s *ScreenService) Delete(ctx context.Context, screenID int) (response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v", screenID)

	request, err := s.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	return
}

type AvailableScreenFieldScheme struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Returns the fields that can be added to a tab on a screen.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens#get-available-screen-fields
func (s *ScreenService) Available(ctx context.Context, screenID int) (result *[]AvailableScreenFieldScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/screens/%v/availableFields", screenID)

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new([]AvailableScreenFieldScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}
