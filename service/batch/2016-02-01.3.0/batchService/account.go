package batchservice

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
	"github.com/Azure/go-autorest/autorest/date"
	"net/http"
)

// AccountClient is the a client for issuing REST requests to the Azure Batch
// service.
type AccountClient struct {
	ManagementClient
}

// NewAccountClient creates an instance of the AccountClient client.
func NewAccountClient() AccountClient {
	return NewAccountClientWithBaseURI(DefaultBaseURI)
}

// NewAccountClientWithBaseURI creates an instance of the AccountClient client.
func NewAccountClientWithBaseURI(baseURI string) AccountClient {
	return AccountClient{NewWithBaseURI(baseURI)}
}

// ListNodeAgentSkus lists all node agent SKUs supported by the Azure Batch
// service.
//
// filter is an OData $filter clause. maxResults is the maximum number of items
// to return in the response. timeout is the maximum time that the server can
// spend processing the request, in seconds. The default is 30 seconds.
// clientRequestID is the caller-generated request identity, in the form of a
// GUID with no decoration such as curly braces, e.g.
// 9C4D50EE-2D56-4CD3-8152-34347DC9F2B0. returnClientRequestID is whether the
// server should return the client-request-id identifier in the response.
// ocpDate is the time the request was issued. If not specified, this header
// will be automatically populated with the current system clock time.
func (client AccountClient) ListNodeAgentSkus(filter string, maxResults *int32, timeout *int32, clientRequestID string, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (result AccountListNodeAgentSkusResult, err error) {
	req, err := client.ListNodeAgentSkusPreparer(filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchservice.AccountClient", "ListNodeAgentSkus", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListNodeAgentSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchservice.AccountClient", "ListNodeAgentSkus", resp, "Failure sending request")
		return
	}

	result, err = client.ListNodeAgentSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchservice.AccountClient", "ListNodeAgentSkus", resp, "Failure responding to request")
	}

	return
}

// ListNodeAgentSkusPreparer prepares the ListNodeAgentSkus request.
func (client AccountClient) ListNodeAgentSkusPreparer(filter string, maxResults *int32, timeout *int32, clientRequestID string, returnClientRequestID *bool, ocpDate *date.TimeRFC1123) (*http.Request, error) {
	const APIVersion = "2016-02-01.3.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if maxResults != nil {
		queryParameters["maxresults"] = autorest.Encode("query", *maxResults)
	}
	if timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/nodeagentskus"),
		autorest.WithQueryParameters(queryParameters))
	if len(clientRequestID) > 0 {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("client-request-id", autorest.String(clientRequestID)))
	}
	if returnClientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("return-client-request-id", autorest.String(returnClientRequestID)))
	}
	if ocpDate != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("ocp-date", autorest.String(ocpDate)))
	}
	return preparer.Prepare(&http.Request{})
}

// ListNodeAgentSkusSender sends the ListNodeAgentSkus request. The method will close the
// http.Response Body if it receives an error.
func (client AccountClient) ListNodeAgentSkusSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListNodeAgentSkusResponder handles the response to the ListNodeAgentSkus request. The method always
// closes the http.Response Body.
func (client AccountClient) ListNodeAgentSkusResponder(resp *http.Response) (result AccountListNodeAgentSkusResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNodeAgentSkusNextResults retrieves the next set of results, if any.
func (client AccountClient) ListNodeAgentSkusNextResults(lastResults AccountListNodeAgentSkusResult) (result AccountListNodeAgentSkusResult, err error) {
	req, err := lastResults.AccountListNodeAgentSkusResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "batchservice.AccountClient", "ListNodeAgentSkus", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListNodeAgentSkusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "batchservice.AccountClient", "ListNodeAgentSkus", resp, "Failure sending next results request")
	}

	result, err = client.ListNodeAgentSkusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchservice.AccountClient", "ListNodeAgentSkus", resp, "Failure responding to next results request")
	}

	return
}

// ListNodeAgentSkusComplete gets all elements from the list without paging.
func (client AccountClient) ListNodeAgentSkusComplete(filter string, maxResults *int32, timeout *int32, clientRequestID string, returnClientRequestID *bool, ocpDate *date.TimeRFC1123, cancel <-chan struct{}) (<-chan NodeAgentSku, <-chan error) {
	resultChan := make(chan NodeAgentSku)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListNodeAgentSkus(filter, maxResults, timeout, clientRequestID, returnClientRequestID, ocpDate)
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
		for list.OdataNextLink != nil {
			list, err = client.ListNodeAgentSkusNextResults(list)
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
