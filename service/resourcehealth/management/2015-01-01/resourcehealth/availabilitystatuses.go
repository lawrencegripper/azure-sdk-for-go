package resourcehealth

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

// AvailabilityStatusesClient is the the Resource Health Client.
type AvailabilityStatusesClient struct {
	ManagementClient
}

// NewAvailabilityStatusesClient creates an instance of the
// AvailabilityStatusesClient client.
func NewAvailabilityStatusesClient(subscriptionID string, resourceType string) AvailabilityStatusesClient {
	return NewAvailabilityStatusesClientWithBaseURI(DefaultBaseURI, subscriptionID, resourceType)
}

// NewAvailabilityStatusesClientWithBaseURI creates an instance of the
// AvailabilityStatusesClient client.
func NewAvailabilityStatusesClientWithBaseURI(baseURI string, subscriptionID string, resourceType string) AvailabilityStatusesClient {
	return AvailabilityStatusesClient{NewWithBaseURI(baseURI, subscriptionID, resourceType)}
}

// GetByResource gets current availability status for a single resource
//
// resourceURI is the fully qualified ID of the resource, including the
// resource name and resource type. Currently the API support not nested and
// one nesting level resource types :
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// and
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
// filter is the filter to apply on the operation. For more information please
// see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand is setting $expand=recommendedactions in url query expands the
// recommendedactions in the response.
func (client AvailabilityStatusesClient) GetByResource(resourceURI string, filter string, expand string) (result AvailabilityStatus, err error) {
	req, err := client.GetByResourcePreparer(resourceURI, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "GetByResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "GetByResource", resp, "Failure sending request")
		return
	}

	result, err = client.GetByResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "GetByResource", resp, "Failure responding to request")
	}

	return
}

// GetByResourcePreparer prepares the GetByResource request.
func (client AvailabilityStatusesClient) GetByResourcePreparer(resourceURI string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2015-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses/current", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetByResourceSender sends the GetByResource request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) GetByResourceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetByResourceResponder handles the response to the GetByResource request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) GetByResourceResponder(resp *http.Response) (result AvailabilityStatus, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists the historical availability statuses for a single resource. Use
// the nextLink property in the response to get the next page of availability
// status
//
// resourceURI is the fully qualified ID of the resource, including the
// resource name and resource type. Currently the API support not nested and
// one nesting level resource types :
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// and
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
// filter is the filter to apply on the operation. For more information please
// see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand is setting $expand=recommendedactions in url query expands the
// recommendedactions in the response.
func (client AvailabilityStatusesClient) List(resourceURI string, filter string, expand string) (result AvailabilityStatusListResult, err error) {
	req, err := client.ListPreparer(resourceURI, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client AvailabilityStatusesClient) ListPreparer(resourceURI string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2015-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.ResourceHealth/availabilityStatuses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) ListResponder(resp *http.Response) (result AvailabilityStatusListResult, err error) {
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
func (client AvailabilityStatusesClient) ListNextResults(lastResults AvailabilityStatusListResult) (result AvailabilityStatusListResult, err error) {
	req, err := lastResults.AvailabilityStatusListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// ListComplete gets all elements from the list without paging.
func (client AvailabilityStatusesClient) ListComplete(resourceURI string, filter string, expand string, cancel <-chan struct{}) (<-chan AvailabilityStatus, <-chan error) {
	resultChan := make(chan AvailabilityStatus)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.List(resourceURI, filter, expand)
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

// ListByResourceGroup lists the current availability status for all the
// resources in the resource group. Use the nextLink property in the response
// to get the next page of availability statuses.
//
// resourceGroupName is the name of the resource group. filter is the filter to
// apply on the operation. For more information please see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand is setting $expand=recommendedactions in url query expands the
// recommendedactions in the response.
func (client AvailabilityStatusesClient) ListByResourceGroup(resourceGroupName string, filter string, expand string) (result AvailabilityStatusListResult, err error) {
	req, err := client.ListByResourceGroupPreparer(resourceGroupName, filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client AvailabilityStatusesClient) ListByResourceGroupPreparer(resourceGroupName string, filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ResourceHealth/availabilityStatuses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) ListByResourceGroupResponder(resp *http.Response) (result AvailabilityStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroupNextResults retrieves the next set of results, if any.
func (client AvailabilityStatusesClient) ListByResourceGroupNextResults(lastResults AvailabilityStatusListResult) (result AvailabilityStatusListResult, err error) {
	req, err := lastResults.AvailabilityStatusListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", resp, "Failure sending next results request")
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListByResourceGroup", resp, "Failure responding to next results request")
	}

	return
}

// ListByResourceGroupComplete gets all elements from the list without paging.
func (client AvailabilityStatusesClient) ListByResourceGroupComplete(resourceGroupName string, filter string, expand string, cancel <-chan struct{}) (<-chan AvailabilityStatus, <-chan error) {
	resultChan := make(chan AvailabilityStatus)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListByResourceGroup(resourceGroupName, filter, expand)
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
			list, err = client.ListByResourceGroupNextResults(list)
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

// ListBySubscriptionID lists the current availability status for all the
// resources in the subscription. Use the nextLink property in the response to
// get the next page of availability statuses.
//
// filter is the filter to apply on the operation. For more information please
// see
// https://docs.microsoft.com/en-us/rest/api/apimanagement/apis?redirectedfrom=MSDN
// expand is setting $expand=recommendedactions in url query expands the
// recommendedactions in the response.
func (client AvailabilityStatusesClient) ListBySubscriptionID(filter string, expand string) (result AvailabilityStatusListResult, err error) {
	req, err := client.ListBySubscriptionIDPreparer(filter, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", resp, "Failure responding to request")
	}

	return
}

// ListBySubscriptionIDPreparer prepares the ListBySubscriptionID request.
func (client AvailabilityStatusesClient) ListBySubscriptionIDPreparer(filter string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2015-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/availabilityStatuses", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListBySubscriptionIDSender sends the ListBySubscriptionID request. The method will close the
// http.Response Body if it receives an error.
func (client AvailabilityStatusesClient) ListBySubscriptionIDSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListBySubscriptionIDResponder handles the response to the ListBySubscriptionID request. The method always
// closes the http.Response Body.
func (client AvailabilityStatusesClient) ListBySubscriptionIDResponder(resp *http.Response) (result AvailabilityStatusListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionIDNextResults retrieves the next set of results, if any.
func (client AvailabilityStatusesClient) ListBySubscriptionIDNextResults(lastResults AvailabilityStatusListResult) (result AvailabilityStatusListResult, err error) {
	req, err := lastResults.AvailabilityStatusListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", resp, "Failure sending next results request")
	}

	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourcehealth.AvailabilityStatusesClient", "ListBySubscriptionID", resp, "Failure responding to next results request")
	}

	return
}

// ListBySubscriptionIDComplete gets all elements from the list without paging.
func (client AvailabilityStatusesClient) ListBySubscriptionIDComplete(filter string, expand string, cancel <-chan struct{}) (<-chan AvailabilityStatus, <-chan error) {
	resultChan := make(chan AvailabilityStatus)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListBySubscriptionID(filter, expand)
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
			list, err = client.ListBySubscriptionIDNextResults(list)
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
