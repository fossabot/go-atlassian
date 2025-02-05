package jira

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type VoteService struct{ client *Client }

type IssueVoteScheme struct {
	Self     string        `json:"self,omitempty"`
	Votes    int           `json:"votes,omitempty"`
	HasVoted bool          `json:"hasVoted,omitempty"`
	Voters   []*UserScheme `json:"voters,omitempty"`
}

// Returns details about the votes on an issue.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/vote#get-votes
func (v *VoteService) Gets(ctx context.Context, issueKeyOrID string) (result *IssueVoteScheme, response *Response, err error) {

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = v.client.Do(request)
	if err != nil {
		return
	}

	result = new(IssueVoteScheme)
	if err = json.Unmarshal(response.BodyAsBytes, &result); err != nil {
		return
	}

	return
}

// Adds the user's vote to an issue. This is the equivalent of the user clicking Vote on an issue in Jira.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/vote#add-vote
func (v *VoteService) Add(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = v.client.Do(request)
	if err != nil {
		return
	}

	return
}

// Deletes a user's vote from an issue. This is the equivalent of the user clicking Unvote on an issue in Jira.
// Docs: https://docs.go-atlassian.io/jira-software-cloud/issues/vote#delete-vote
func (v *VoteService) Delete(ctx context.Context, issueKeyOrID string) (response *Response, err error) {

	if len(issueKeyOrID) == 0 {
		return nil, fmt.Errorf("error, please provide a issueKeyOrID value")
	}

	var endpoint = fmt.Sprintf("rest/api/3/issue/%v/votes", issueKeyOrID)

	request, err := v.client.newRequest(ctx, http.MethodDelete, endpoint, nil)
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err = v.client.Do(request)
	if err != nil {
		return
	}

	return
}
