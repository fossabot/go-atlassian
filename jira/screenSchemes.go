package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"net/url"
	"strconv"
)

type ScreenSchemeService struct{ client *Client }

type ScreenSchemePageScheme struct {
	Self       string                `json:"self,omitempty"`
	NextPage   string                `json:"nextPage,omitempty"`
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*ScreenSchemeScheme `json:"values,omitempty"`
}

type ScreenSchemeScheme struct {
	ID                     int                        `json:"id,omitempty"`
	Name                   string                     `json:"name,omitempty"`
	Description            string                     `json:"description,omitempty"`
	Screens                *ScreenTypesScheme         `json:"screens,omitempty"`
	IssueTypeScreenSchemes *IssueTypeSchemePageScheme `json:"issueTypeScreenSchemes,omitempty"`
}

type ScreenTypesScheme struct {
	Create  int `json:"create,omitempty"`
	Default int `json:"default" validate:"required"`
	View    int `json:"view" validate:"required"`
	Edit    int `json:"edit" validate:"required"`
}

// Returns a paginated list of screen schemes.
// Only screen schemes used in classic projects are returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#get-screen-schemes
func (s *ScreenSchemeService) Gets(ctx context.Context, screenSchemeIDs []int, startAt, maxResults int) (result *ScreenSchemePageScheme, response *Response, err error) {

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, schemeScheme := range screenSchemeIDs {
		params.Add("id", strconv.Itoa(schemeScheme))
	}

	var endpoint = fmt.Sprintf("rest/api/3/screenscheme?%v", params.Encode())

	request, err := s.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = s.client.Do(request)
	if err != nil {
		return
	}

	result = new(ScreenSchemePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type ScreenSchemePayloadScheme struct {
	Screens     *ScreenTypesScheme `json:"screens"`
	Name        string             `json:"name" validate:"required"`
	Description string             `json:"description,omitempty"`
}

// Creates a screen scheme.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#create-screen-scheme
func (s *ScreenSchemeService) Create(ctx context.Context, payload *ScreenSchemePayloadScheme) (result *ScreenSchemeScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, please provide a valid ScreenSchemePayloadScheme pointer")
	}

	validate := validator.New()
	if err = validate.Struct(payload); err != nil {
		return
	}

	var endpoint = "rest/api/3/screenscheme"

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

	result = new(ScreenSchemeScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Updates a screen scheme. Only screen schemes used in classic projects can be updated.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#update-screen-scheme
func (s *ScreenSchemeService) Update(ctx context.Context, screenSchemeID string, payload *ScreenSchemePayloadScheme) (response *Response, err error) {

	if len(screenSchemeID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid screenSchemeID value")
	}

	if payload == nil {
		return nil, fmt.Errorf("error, please provide a valid ScreenSchemeScheme pointer")
	}

	var endpoint = fmt.Sprintf("rest/api/3/screenscheme/%v", screenSchemeID)

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

	return
}

// Deletes a screen scheme.
// A screen scheme cannot be deleted if it is used in an issue type screen scheme.
// Only screens schemes used in classic projects can be deleted.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/screens/schemes#delete-screen-scheme
func (s *ScreenSchemeService) Delete(ctx context.Context, screenSchemeID string) (response *Response, err error) {

	if len(screenSchemeID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid screenSchemeID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/screenscheme/%v", screenSchemeID)

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
