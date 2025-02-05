package jira

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"testing"
)

func TestIssueLinkTypeService_Create(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		payload            *LinkTypeScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name: "CreateIssueLinksTypeWhenThePayloadIsCorrect",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},
		{
			name: "CreateIssueLinksTypeWhenTheResponseBodyHasADifferentFormat",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name: "CreateIssueLinksTypeWhenTheContextIsNil",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name: "CreateIssueLinksTypeWhenTheStatusCodeIsIncorrect",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name: "CreateIssueLinksTypeWhenTheMethodIsIncorrect",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "CreateIssueLinksTypeWhenThePayloadIsNil",
			payload:            nil,
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "CreateIssueLinksTypeWhenThePayloadIsEmpty",
			payload:            nil,
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusCreated,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueLinkTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Create(testCase.context, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
				assert.Equal(t, gotResult.Name, testCase.payload.Name)
			}
		})

	}

}

func TestIssueLinkTypeService_Get(t *testing.T) {

	testCases := []struct {
		name                   string
		mockFile               string
		IssueLinkTypeServiceID string
		wantHTTPMethod         string
		endpoint               string
		context                context.Context
		wantHTTPHeaders        map[string]string
		wantHTTPCodeReturn     int
		wantErr                bool
	}{
		{
			name:                   "GetIssueLinksTypeWhenTheJSONIsCorrect",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                false,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheContextIsNil",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                nil,
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheTheResponseBodyHasADifferentFormat",
			mockFile:               "./mocks/invalid-json.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheJSONIsEmpty",
			mockFile:               "./mocks/empty_json.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheMockedIDIsDifferent",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10001",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheMockedIDIsEmpty",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheHTTPResponseCodeIsOK",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodGet,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                false,
		},
		{
			name:                   "GetIssueLinksTypeWhenTheHTTPResponseCodeIsNotValid",
			mockFile:               "./mocks/get_issue_link_type_id_10000.json",
			wantHTTPCodeReturn:     http.StatusOK,
			wantHTTPMethod:         http.MethodPut,
			IssueLinkTypeServiceID: "10000",
			endpoint:               "/rest/api/3/issueLinkType/10000",
			context:                context.Background(),
			wantHTTPHeaders:        map[string]string{"Accept": "application/json"},
			wantErr:                true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueLinkTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Get(testCase.context, testCase.IssueLinkTypeServiceID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
				assert.Equal(t, gotResult.ID, testCase.IssueLinkTypeServiceID)
			}
		})

	}

}

func TestIssueLinkTypeService_Gets(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "GetIssueLinksTypesWhenTheJSONIsCorrect",
			mockFile:           "./mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            false,
		},

		{
			name:               "GetIssueLinksTypesWhenTheContextIsNil",
			mockFile:           "./mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            nil,
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},

		{
			name:               "GetIssueLinksTypesWhenTheStatusCodeIsIncorrect",
			mockFile:           "./mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheRequestMethodIsIncorrect",
			mockFile:           "./mocks/get_issue_link_types.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONIsEmpty",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsInvalid",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkTypeasasdadsads",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsEmpty",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheEndpointProvidedIsANumber",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "111",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPResponseCodeIsInvalid",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusBadGateway,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPResponseCodeIsDifferent",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: 499,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
		{
			name:               "GetIssueLinksTypesWhenTheJSONWhenTheHTTPMethodRequestIsNotGET",
			mockFile:           "./mocks/empty_json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPost,
			endpoint:           "/rest/api/3/issueLinkType",
			context:            context.Background(),
			wantHTTPHeaders:    map[string]string{"Accept": "application/json"},
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueLinkTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Gets(testCase.context)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueLinkTypeService_Delete(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueLinkTypeID    string
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:               "DeleteIssueLinkTypeWhenTheIDIsCorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            false,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheIssueLinkTypeIDIsNotProvided",
			issueLinkTypeID:    "",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheIDIsIncorrect",
			issueLinkTypeID:    "10002",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheEndpointIsIncorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLinkTypes/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheContextIsNil",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            nil,
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheRequestMethodIsIncorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusNoContent,
			wantErr:            true,
		},

		{
			name:               "DeleteIssueLinkTypeWhenTheStatusCodeIsIncorrect",
			issueLinkTypeID:    "10001",
			wantHTTPMethod:     http.MethodDelete,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueLinkTypeService{client: mockClient}

			gotResponse, err := i.Delete(testCase.context, testCase.issueLinkTypeID)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}

}

func TestIssueLinkTypeService_Update(t *testing.T) {

	testCases := []struct {
		name               string
		mockFile           string
		issueLinkTypeID    string
		payload            *LinkTypeScheme
		wantHTTPMethod     string
		endpoint           string
		context            context.Context
		wantHTTPHeaders    map[string]string
		wantHTTPCodeReturn int
		wantErr            bool
	}{
		{
			name:            "UpdateIssueLinksTypeWhenThePayloadIsCorrect",
			issueLinkTypeID: "10001",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            false,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheIssueLinkTypeIDIsNotProvided",
			issueLinkTypeID: "",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheIDInsIncorrect",
			issueLinkTypeID: "10000",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheContextIsNil",
			issueLinkTypeID: "10001",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            nil,
			wantErr:            true,
		},

		{
			name:               "UpdateIssueLinksTypeWhenThePayloadIsNil",
			issueLinkTypeID:    "10001",
			payload:            nil,
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheRequestMethodIsIncorrect",
			issueLinkTypeID: "10001",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodGet,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheStatusCodeIsIncorrect",
			issueLinkTypeID: "10001",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/create_issue_link_type_duplicate_case.json",
			wantHTTPCodeReturn: http.StatusBadRequest,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},

		{
			name:            "UpdateIssueLinksTypeWhenTheResponseBodyHasADifferentFormat",
			issueLinkTypeID: "10001",
			payload: &LinkTypeScheme{
				Inward:  "Duplicated by",
				Name:    "Duplicate",
				Outward: "Duplicates",
			},
			mockFile:           "./mocks/invalid-json.json",
			wantHTTPCodeReturn: http.StatusOK,
			wantHTTPMethod:     http.MethodPut,
			endpoint:           "/rest/api/3/issueLinkType/10001",
			context:            context.Background(),
			wantErr:            true,
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.name, func(t *testing.T) {

			//Init a new HTTP mock server
			mockOptions := mockServerOptions{
				Endpoint:           testCase.endpoint,
				MockFilePath:       testCase.mockFile,
				MethodAccepted:     testCase.wantHTTPMethod,
				Headers:            testCase.wantHTTPHeaders,
				ResponseCodeWanted: testCase.wantHTTPCodeReturn,
			}

			mockServer, err := startMockServer(&mockOptions)
			if err != nil {
				t.Fatal(err)
			}

			defer mockServer.Close()

			//Init the library instance
			mockClient, err := startMockClient(mockServer.URL)
			if err != nil {
				t.Fatal(err)
			}

			i := &IssueLinkTypeService{client: mockClient}

			gotResult, gotResponse, err := i.Update(testCase.context, testCase.issueLinkTypeID, testCase.payload)

			if testCase.wantErr {

				if err != nil {
					t.Logf("error returned: %v", err.Error())
				}

				assert.Error(t, err)
			} else {

				assert.NoError(t, err)
				assert.NotEqual(t, gotResponse, nil)
				assert.NotEqual(t, gotResult, nil)

				apiEndpoint, err := url.Parse(gotResponse.Endpoint)
				if err != nil {
					t.Fatal(err)
				}

				var endpointToAssert string

				if apiEndpoint.Query().Encode() != "" {
					endpointToAssert = fmt.Sprintf("%v?%v", apiEndpoint.Path, apiEndpoint.Query().Encode())
				} else {
					endpointToAssert = apiEndpoint.Path
				}

				t.Logf("HTTP Endpoint Wanted: %v, HTTP Endpoint Returned: %v", testCase.endpoint, endpointToAssert)
				assert.Equal(t, testCase.endpoint, endpointToAssert)
			}
		})

	}
}
