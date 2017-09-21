package trafficmanager

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

// EndpointsClient is the client for the Endpoints methods of the
// Trafficmanager service.
type EndpointsClient struct {
	ManagementClient
}

// NewEndpointsClient creates an instance of the EndpointsClient client.
func NewEndpointsClient(subscriptionID string) EndpointsClient {
	return NewEndpointsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEndpointsClientWithBaseURI creates an instance of the EndpointsClient
// client.
func NewEndpointsClientWithBaseURI(baseURI string, subscriptionID string) EndpointsClient {
	return EndpointsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a Traffic Manager endpoint.
//
// resourceGroupName is the name of the resource group containing the Traffic
// Manager endpoint to be created or updated. profileName is the name of the
// Traffic Manager profile. endpointType is the type of the Traffic Manager
// endpoint to be created or updated. endpointName is the name of the Traffic
// Manager endpoint to be created or updated. parameters is the Traffic Manager
// endpoint parameters supplied to the CreateOrUpdate operation.
func (client EndpointsClient) CreateOrUpdate(resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (result Endpoint, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, profileName, endpointType, endpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client EndpointsClient) CreateOrUpdatePreparer(resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client EndpointsClient) CreateOrUpdateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a Traffic Manager endpoint.
//
// resourceGroupName is the name of the resource group containing the Traffic
// Manager endpoint to be deleted. profileName is the name of the Traffic
// Manager profile. endpointType is the type of the Traffic Manager endpoint to
// be deleted. endpointName is the name of the Traffic Manager endpoint to be
// deleted.
func (client EndpointsClient) Delete(resourceGroupName string, profileName string, endpointType string, endpointName string) (result DeleteOperationResult, err error) {
	req, err := client.DeletePreparer(resourceGroupName, profileName, endpointType, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client EndpointsClient) DeletePreparer(resourceGroupName string, profileName string, endpointType string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client EndpointsClient) DeleteResponder(resp *http.Response) (result DeleteOperationResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets a Traffic Manager endpoint.
//
// resourceGroupName is the name of the resource group containing the Traffic
// Manager endpoint. profileName is the name of the Traffic Manager profile.
// endpointType is the type of the Traffic Manager endpoint. endpointName is
// the name of the Traffic Manager endpoint.
func (client EndpointsClient) Get(resourceGroupName string, profileName string, endpointType string, endpointName string) (result Endpoint, err error) {
	req, err := client.GetPreparer(resourceGroupName, profileName, endpointType, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client EndpointsClient) GetPreparer(resourceGroupName string, profileName string, endpointType string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client EndpointsClient) GetResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update update a Traffic Manager endpoint.
//
// resourceGroupName is the name of the resource group containing the Traffic
// Manager endpoint to be updated. profileName is the name of the Traffic
// Manager profile. endpointType is the type of the Traffic Manager endpoint to
// be updated. endpointName is the name of the Traffic Manager endpoint to be
// updated. parameters is the Traffic Manager endpoint parameters supplied to
// the Update operation.
func (client EndpointsClient) Update(resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (result Endpoint, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, profileName, endpointType, endpointName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "trafficmanager.EndpointsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client EndpointsClient) UpdatePreparer(resourceGroupName string, profileName string, endpointType string, endpointName string, parameters Endpoint) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"endpointType":      autorest.Encode("path", endpointType),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-03-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/{endpointType}/{endpointName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client EndpointsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client EndpointsClient) UpdateResponder(resp *http.Response) (result Endpoint, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
