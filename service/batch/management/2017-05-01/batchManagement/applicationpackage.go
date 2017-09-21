package batchmanagement

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ApplicationPackageClient is the client for the ApplicationPackage methods of
// the Batchmanagement service.
type ApplicationPackageClient struct {
	ManagementClient
}

// NewApplicationPackageClient creates an instance of the
// ApplicationPackageClient client.
func NewApplicationPackageClient(subscriptionID string) ApplicationPackageClient {
	return NewApplicationPackageClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewApplicationPackageClientWithBaseURI creates an instance of the
// ApplicationPackageClient client.
func NewApplicationPackageClientWithBaseURI(baseURI string, subscriptionID string) ApplicationPackageClient {
	return ApplicationPackageClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Activate activates the specified application package.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is the
// ID of the application. version is the version of the application to
// activate. parameters is the parameters for the request.
func (client ApplicationPackageClient) Activate(resourceGroupName string, accountName string, applicationID string, version string, parameters ActivateApplicationPackageParameters) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Format", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batchmanagement.ApplicationPackageClient", "Activate")
	}

	req, err := client.ActivatePreparer(resourceGroupName, accountName, applicationID, version, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Activate", nil, "Failure preparing request")
		return
	}

	resp, err := client.ActivateSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Activate", resp, "Failure sending request")
		return
	}

	result, err = client.ActivateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Activate", resp, "Failure responding to request")
	}

	return
}

// ActivatePreparer prepares the Activate request.
func (client ApplicationPackageClient) ActivatePreparer(resourceGroupName string, accountName string, applicationID string, version string, parameters ActivateApplicationPackageParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	const APIVersion = "2017-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}/versions/{version}/activate", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ActivateSender sends the Activate request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) ActivateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ActivateResponder handles the response to the Activate request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) ActivateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Create creates an application package record.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is the
// ID of the application. version is the version of the application.
func (client ApplicationPackageClient) Create(resourceGroupName string, accountName string, applicationID string, version string) (result ApplicationPackage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batchmanagement.ApplicationPackageClient", "Create")
	}

	req, err := client.CreatePreparer(resourceGroupName, accountName, applicationID, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Create", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Create", resp, "Failure sending request")
		return
	}

	result, err = client.CreateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Create", resp, "Failure responding to request")
	}

	return
}

// CreatePreparer prepares the Create request.
func (client ApplicationPackageClient) CreatePreparer(resourceGroupName string, accountName string, applicationID string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	const APIVersion = "2017-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) CreateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) CreateResponder(resp *http.Response) (result ApplicationPackage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an application package record and its associated binary file.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is the
// ID of the application. version is the version of the application to delete.
func (client ApplicationPackageClient) Delete(resourceGroupName string, accountName string, applicationID string, version string) (result autorest.Response, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batchmanagement.ApplicationPackageClient", "Delete")
	}

	req, err := client.DeletePreparer(resourceGroupName, accountName, applicationID, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ApplicationPackageClient) DeletePreparer(resourceGroupName string, accountName string, applicationID string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	const APIVersion = "2017-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets information about the specified application package.
//
// resourceGroupName is the name of the resource group that contains the Batch
// account. accountName is the name of the Batch account. applicationID is the
// ID of the application. version is the version of the application.
func (client ApplicationPackageClient) Get(resourceGroupName string, accountName string, applicationID string, version string) (result ApplicationPackage, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}},
		{TargetValue: accountName,
			Constraints: []validation.Constraint{{Target: "accountName", Name: validation.MaxLength, Rule: 24, Chain: nil},
				{Target: "accountName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "accountName", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "batchmanagement.ApplicationPackageClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, accountName, applicationID, version)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "batchmanagement.ApplicationPackageClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ApplicationPackageClient) GetPreparer(resourceGroupName string, accountName string, applicationID string, version string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"accountName":       autorest.Encode("path", accountName),
		"applicationId":     autorest.Encode("path", applicationID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"version":           autorest.Encode("path", version),
	}

	const APIVersion = "2017-05-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Batch/batchAccounts/{accountName}/applications/{applicationId}/versions/{version}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ApplicationPackageClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ApplicationPackageClient) GetResponder(resp *http.Response) (result ApplicationPackage, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
