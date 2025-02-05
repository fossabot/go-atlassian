package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type FieldContextService struct {
	client *Client
	Option *FieldOptionContextService
}

type FieldContextOptionsScheme struct {
	IsAnyIssueType  bool
	IsGlobalContext bool
	ContextID       []int
}

type CustomFieldContextPageScheme struct {
	MaxResults int                   `json:"maxResults,omitempty"`
	StartAt    int                   `json:"startAt,omitempty"`
	Total      int                   `json:"total,omitempty"`
	IsLast     bool                  `json:"isLast,omitempty"`
	Values     []*FieldContextScheme `json:"values,omitempty"`
}

type FieldContextScheme struct {
	ID              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Description     string   `json:"description,omitempty"`
	IsGlobalContext bool     `json:"isGlobalContext,omitempty"`
	IsAnyIssueType  bool     `json:"isAnyIssueType,omitempty"`
	ProjectIds      []string `json:"projectIds,omitempty"`
	IssueTypeIds    []string `json:"issueTypeIds,omitempty"`
}

// Returns a paginated list of contexts for a custom field. Contexts can be returned as follows:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts
func (f *FieldContextService) Gets(ctx context.Context, fieldID string, opts *FieldContextOptionsScheme, startAt, maxResults int) (result *CustomFieldContextPageScheme, response *Response, err error) {

	if fieldID == "" {
		return nil, nil, fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	if opts != nil {

		if opts.IsAnyIssueType {
			params.Add("isAnyIssueType", "true")
		}

		if opts.IsGlobalContext {
			params.Add("isGlobalContext", "true")
		}

		for _, contextID := range opts.ContextID {
			params.Add("contextId", strconv.Itoa(contextID))
		}

	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context?%v", fieldID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomFieldContextPageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type FieldContextPayloadScheme struct {
	IssueTypeIDs []int  `json:"issueTypeIds,omitempty"`
	ProjectIDs   []int  `json:"projectIds,omitempty"`
	Name         string `json:"name,omitempty"`
	Description  string `json:"description,omitempty"`
}

// Creates a custom field context.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#create-custom-field-context
func (f *FieldContextService) Create(ctx context.Context, fieldID string, payload *FieldContextPayloadScheme) (result *FieldContextScheme, response *Response, err error) {

	if payload == nil {
		return nil, nil, fmt.Errorf("error, payload value is nil, please provide a valid FieldContextPayloadScheme pointer")
	}

	if fieldID == "" {
		return nil, nil, fmt.Errorf("error, fieldID value is nil, please provide a valid fieldID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context", fieldID)
	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(FieldContextScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return nil, response, fmt.Errorf("unable to marshall the response body, error: %v", err.Error())
	}

	return
}

type CustomFieldDefaultValuePageScheme struct {
	MaxResults int                              `json:"maxResults,omitempty"`
	StartAt    int                              `json:"startAt,omitempty"`
	Total      int                              `json:"total,omitempty"`
	IsLast     bool                             `json:"isLast,omitempty"`
	Values     []*CustomFieldDefaultValueScheme `json:"values,omitempty"`
}

type CustomFieldDefaultValueScheme struct {
	ContextID         string   `json:"contextId,omitempty"`
	OptionID          string   `json:"optionId,omitempty"`
	CascadingOptionID string   `json:"cascadingOptionId,omitempty"`
	OptionIDs         []string `json:"optionIds,omitempty"`
	Type              string   `json:"type,omitempty"`
}

// Returns a paginated list of defaults for a custom field.
// The results can be filtered by contextId, otherwise all values are returned.
// If no defaults are set for a context, nothing is returned.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#get-custom-field-contexts-default-values
func (f *FieldContextService) GetDefaultValues(ctx context.Context, fieldID string, contextIDs []int, startAt, maxResults int) (result *CustomFieldDefaultValuePageScheme, response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	params := url.Values{}
	params.Add("startAt", strconv.Itoa(startAt))
	params.Add("maxResults", strconv.Itoa(maxResults))

	for _, contextID := range contextIDs {
		params.Add("contextId", strconv.Itoa(contextID))
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/defaultValue?%s", fieldID, params.Encode())
	request, err := f.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	result = new(CustomFieldDefaultValuePageScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

type FieldContextDefaultPayloadScheme struct {
	DefaultValues []*CustomFieldDefaultValueScheme `json:"defaultValues,omitempty"`
}

// Sets default for contexts of a custom field.
// Default are defined using these objects:
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#set-custom-field-contexts-default-values
func (f *FieldContextService) SetDefaultValue(ctx context.Context, fieldID string, payload *FieldContextDefaultPayloadScheme) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	if payload == nil {
		return nil, fmt.Errorf("error, please provide a valid slice of DefaultValueScheme pointers")
	}

	if len(payload.DefaultValues) == 0 {
		return nil, fmt.Errorf("error, please provide a valid Custom Field Context default value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/defaultValue", fieldID)
	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Updates a custom field context
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#update-custom-field-context
func (f *FieldContextService) Update(ctx context.Context, fieldID string, contextID int, name, description string) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	if contextID == 0 {
		return nil, fmt.Errorf("error, please provide a valid contextID value")
	}

	payload := struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		Name:        name,
		Description: description,
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes a custom field context.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#delete-custom-field-context
func (f *FieldContextService) Delete(ctx context.Context, fieldID string, contextID int) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	if contextID == 0 {
		return nil, fmt.Errorf("error, please provide a valid contextID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v", fieldID, contextID)

	request, err := f.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

//Add issue types
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#add-issue-types-to-context
func (f *FieldContextService) AddIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	if len(issueTypesIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueTypesIDs value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/issuetype", fieldID, contextID)

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{IssueTypeIds: issueTypesIDs}

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Removes issue types from a custom field context.
// A custom field context without any issue types applies to all issue types.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-issue-types-from-context
func (f *FieldContextService) RemoveIssueTypes(ctx context.Context, fieldID string, contextID int, issueTypesIDs []string) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	if len(issueTypesIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid issueTypesIDs value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/issuetype/remove", fieldID, contextID)

	payload := struct {
		IssueTypeIds []string `json:"issueTypeIds"`
	}{IssueTypeIds: issueTypesIDs}

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Assigns a custom field context to projects.
// If any project in the request is assigned to any context of the custom field, the operation fails.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#assign-custom-field-context-to-projects
func (f *FieldContextService) Link(ctx context.Context, fieldID string, contextID int, projectIDs []string) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	if len(projectIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid projectIDs value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/project", fieldID, contextID)

	payload := struct {
		ProjectIds []string `json:"projectIds"`
	}{ProjectIds: projectIDs}

	request, err := f.client.newRequest(ctx, http.MethodPut, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Removes a custom field context from projects.
// A custom field context without any projects applies to all projects.
// Removing all projects from a custom field context would result in it applying to all projects.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/fields/context#remove-custom-field-context-from-projects
func (f *FieldContextService) UnLink(ctx context.Context, fieldID string, contextID int, projectIDs []string) (response *Response, err error) {

	if len(fieldID) == 0 {
		return nil, fmt.Errorf("error, please provide a valid fieldID value")
	}

	if len(projectIDs) == 0 {
		return nil, fmt.Errorf("error, please provide a valid projectIDs value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/field/%v/context/%v/project/remove", fieldID, contextID)

	payload := struct {
		ProjectIds []string `json:"projectIds"`
	}{ProjectIds: projectIDs}

	request, err := f.client.newRequest(ctx, http.MethodPost, endpoint, &payload)
	if err != nil {
		return
	}

	request.Header.Set("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")

	response, err = f.client.Do(request)
	if err != nil {
		return
	}

	return
}
