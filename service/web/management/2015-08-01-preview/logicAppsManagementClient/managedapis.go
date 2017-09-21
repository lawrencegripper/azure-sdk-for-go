package logicappsmanagementclient

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// ManagedApisClient is the client for the ManagedApis methods of the
// Logicappsmanagementclient service.
type ManagedApisClient struct {
	ManagementClient
}

// NewManagedApisClient creates an instance of the ManagedApisClient client.
func NewManagedApisClient(subscriptionID string) ManagedApisClient {
	return NewManagedApisClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewManagedApisClientWithBaseURI creates an instance of the ManagedApisClient
// client.
func NewManagedApisClientWithBaseURI(baseURI string, subscriptionID string) ManagedApisClient {
	return ManagedApisClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets a managed API.
//
// location is the location. APIName is the managed API name. export is flag
// showing whether to export API definition in format specified by Accept
// header.
func (client ManagedApisClient) Get(location string, APIName string, export *bool) (result APIEntity, err error) {
	req, err := client.GetPreparer(location, APIName, export)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ManagedApisClient) GetPreparer(location string, APIName string, export *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"apiName":        autorest.Encode("path", APIName),
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if export != nil {
		queryParameters["export"] = autorest.Encode("query", *export)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/managedApis/{apiName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedApisClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ManagedApisClient) GetResponder(resp *http.Response) (result APIEntity, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of managed APIs.
//
// location is the location.
func (client ManagedApisClient) List(location string) (result ApisCollection, err error) {
	req, err := client.ListPreparer(location)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ManagedApisClient) ListPreparer(location string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-08-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/managedApis", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ManagedApisClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ManagedApisClient) ListResponder(resp *http.Response) (result ApisCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client ManagedApisClient) ListNextResults(lastResults ApisCollection) (result ApisCollection, err error) {
	req, err := lastResults.ApisCollectionPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "logicappsmanagementclient.ManagedApisClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client ManagedApisClient) ListComplete(location string, cancel <-chan struct{}) (<-chan APIEntity, <-chan error) {
	resultChan := make(chan APIEntity)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(location)
		if err != nil {
			errChan <- err
			return
		}
		if list.Value != nil {
			for _, item := range *list.Value {
				select {
				case <-cancel:
					return
				case resultChan <- item:
					// Intentionally left blank
				}
			}
		}
		for list.NextLink != nil {
			list, err = client.ListNextResults(list)
			if err != nil {
				errChan <- err
				return
			}
			if list.Value != nil {
				for _, item := range *list.Value {
					select {
					case <-cancel:
						return
					case resultChan <- item:
						// Intentionally left blank
					}
				}
			}
		}
	}()
	return resultChan, errChan
}
