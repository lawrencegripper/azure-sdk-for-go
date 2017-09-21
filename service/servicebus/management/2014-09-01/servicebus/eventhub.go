package servicebus

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

// EventHubClient is the azure Service Bus client
type EventHubClient struct {
	ManagementClient
}

// NewEventHubClient creates an instance of the EventHubClient client.
func NewEventHubClient(subscriptionID string) EventHubClient {
	return NewEventHubClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewEventHubClientWithBaseURI creates an instance of the EventHubClient
// client.
func NewEventHubClientWithBaseURI(baseURI string, subscriptionID string) EventHubClient {
	return EventHubClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetAuthorizationRule returns the specified authorization rule.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. namespaceName is the namespace name eventhubName is the
// eventhub name. authorizationRuleName is the authorizationrule name.
func (client EventHubClient) GetAuthorizationRule(resourceGroupName string, namespaceName string, eventhubName string, authorizationRuleName string) (result SharedAccessAuthorizationRuleResource, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}},
		{TargetValue: authorizationRuleName,
			Constraints: []validation.Constraint{{Target: "authorizationRuleName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "authorizationRuleName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.EventHubClient", "GetAuthorizationRule")
	}

	req, err := client.GetAuthorizationRulePreparer(resourceGroupName, namespaceName, eventhubName, authorizationRuleName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubClient", "GetAuthorizationRule", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetAuthorizationRuleSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.EventHubClient", "GetAuthorizationRule", resp, "Failure sending request")
		return
	}

	result, err = client.GetAuthorizationRuleResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubClient", "GetAuthorizationRule", resp, "Failure responding to request")
	}

	return
}

// GetAuthorizationRulePreparer prepares the GetAuthorizationRule request.
func (client EventHubClient) GetAuthorizationRulePreparer(resourceGroupName string, namespaceName string, eventhubName string, authorizationRuleName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"authorizationRuleName": autorest.Encode("path", authorizationRuleName),
		"eventhubName":          autorest.Encode("path", eventhubName),
		"namespaceName":         autorest.Encode("path", namespaceName),
		"resourceGroupName":     autorest.Encode("path", resourceGroupName),
		"subscriptionId":        autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/eventhubs/{eventhubName}/authorizationRules/{authorizationRuleName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetAuthorizationRuleSender sends the GetAuthorizationRule request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubClient) GetAuthorizationRuleSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetAuthorizationRuleResponder handles the response to the GetAuthorizationRule request. The method always
// closes the http.Response Body.
func (client EventHubClient) GetAuthorizationRuleResponder(resp *http.Response) (result SharedAccessAuthorizationRuleResource, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAuthorizationRules gets authorization rules for a eventhub.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. namespaceName is the namespace name eventhubName is the
// eventhub name.
func (client EventHubClient) ListAuthorizationRules(resourceGroupName string, namespaceName string, eventhubName string) (result SharedAccessAuthorizationRuleListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: namespaceName,
			Constraints: []validation.Constraint{{Target: "namespaceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "namespaceName", Name: validation.MinLength, Rule: 6, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "servicebus.EventHubClient", "ListAuthorizationRules")
	}

	req, err := client.ListAuthorizationRulesPreparer(resourceGroupName, namespaceName, eventhubName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubClient", "ListAuthorizationRules", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAuthorizationRulesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicebus.EventHubClient", "ListAuthorizationRules", resp, "Failure sending request")
		return
	}

	result, err = client.ListAuthorizationRulesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubClient", "ListAuthorizationRules", resp, "Failure responding to request")
	}

	return
}

// ListAuthorizationRulesPreparer prepares the ListAuthorizationRules request.
func (client EventHubClient) ListAuthorizationRulesPreparer(resourceGroupName string, namespaceName string, eventhubName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"eventhubName":      autorest.Encode("path", eventhubName),
		"namespaceName":     autorest.Encode("path", namespaceName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2014-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/eventhubs/{eventhubName}/authorizationRules", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListAuthorizationRulesSender sends the ListAuthorizationRules request. The method will close the
// http.Response Body if it receives an error.
func (client EventHubClient) ListAuthorizationRulesSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAuthorizationRulesResponder handles the response to the ListAuthorizationRules request. The method always
// closes the http.Response Body.
func (client EventHubClient) ListAuthorizationRulesResponder(resp *http.Response) (result SharedAccessAuthorizationRuleListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAuthorizationRulesNextResults retrieves the next set of results, if any.
func (client EventHubClient) ListAuthorizationRulesNextResults(lastResults SharedAccessAuthorizationRuleListResult) (result SharedAccessAuthorizationRuleListResult, err error) {
	req, err := lastResults.SharedAccessAuthorizationRuleListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicebus.EventHubClient", "ListAuthorizationRules", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAuthorizationRulesSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicebus.EventHubClient", "ListAuthorizationRules", resp, "Failure sending next results request")
	}

	result, err = client.ListAuthorizationRulesResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicebus.EventHubClient", "ListAuthorizationRules", resp, "Failure responding to next results request")
	}

	return
}

// ListAuthorizationRulesComplete gets all elements from the list without paging.
func (client EventHubClient) ListAuthorizationRulesComplete(resourceGroupName string, namespaceName string, eventhubName string, cancel <-chan struct{}) (<-chan SharedAccessAuthorizationRuleResource, <-chan error) {
	resultChan := make(chan SharedAccessAuthorizationRuleResource)
	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(resultChan)
			close(errChan)
		}()
		list, err := client.ListAuthorizationRules(resourceGroupName, namespaceName, eventhubName)
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
			list, err = client.ListAuthorizationRulesNextResults(list)
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
