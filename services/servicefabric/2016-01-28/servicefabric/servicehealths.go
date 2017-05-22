package servicefabric

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

// ServiceHealthsClient is the client for the ServiceHealths methods of the
// Servicefabric service.
type ServiceHealthsClient struct {
    ManagementClient
}
// NewServiceHealthsClient creates an instance of the ServiceHealthsClient
// client.
func NewServiceHealthsClient(timeout *int32) ServiceHealthsClient {
        return NewServiceHealthsClientWithBaseURI(DefaultBaseURI, timeout)
        }

// NewServiceHealthsClientWithBaseURI creates an instance of the
// ServiceHealthsClient client.
    func NewServiceHealthsClientWithBaseURI(baseURI string, timeout *int32) ServiceHealthsClient {
        return ServiceHealthsClient{ NewWithBaseURI(baseURI, timeout)}
    }

// Get get service healths
//
// serviceName is the name of the service
func (client ServiceHealthsClient) Get(serviceName string) (result ServiceHealth, err error) {
    req, err := client.GetPreparer(serviceName)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServiceHealthsClient", "Get", nil , "Failure preparing request")
        return
    }

    resp, err := client.GetSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServiceHealthsClient", "Get", resp, "Failure sending request")
        return
    }

    result, err = client.GetResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServiceHealthsClient", "Get", resp, "Failure responding to request")
    }

    return
}

// GetPreparer prepares the Get request.
func (client ServiceHealthsClient) GetPreparer(serviceName string) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "serviceName": serviceName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsGet(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Services/{serviceName}/$/GetHealth",pathParameters),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceHealthsClient) GetSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ServiceHealthsClient) GetResponder(resp *http.Response) (result ServiceHealth, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}

// Send send service healths
//
// serviceName is the name of the service serviceHealthReport is the report of
// the service health
func (client ServiceHealthsClient) Send(serviceName string, serviceHealthReport ServiceHealthReport) (result String, err error) {
    req, err := client.SendPreparer(serviceName, serviceHealthReport)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServiceHealthsClient", "Send", nil , "Failure preparing request")
        return
    }

    resp, err := client.SendSender(req)
    if err != nil {
        result.Response = autorest.Response{Response: resp}
        err = autorest.NewErrorWithError(err, "servicefabric.ServiceHealthsClient", "Send", resp, "Failure sending request")
        return
    }

    result, err = client.SendResponder(resp)
    if err != nil {
        err = autorest.NewErrorWithError(err, "servicefabric.ServiceHealthsClient", "Send", resp, "Failure responding to request")
    }

    return
}

// SendPreparer prepares the Send request.
func (client ServiceHealthsClient) SendPreparer(serviceName string, serviceHealthReport ServiceHealthReport) (*http.Request, error) {
    pathParameters := map[string]interface{} {
    "serviceName": serviceName,
    }

        const APIVersion = "1.0.0"
    queryParameters := map[string]interface{} {
    "api-version": APIVersion,
    }
    if client.Timeout != nil {
        queryParameters["timeout"] = autorest.Encode("query",*client.Timeout)
    }

    preparer := autorest.CreatePreparer(
                        autorest.AsJSON(),
                        autorest.AsPost(),
                        autorest.WithBaseURL(client.BaseURI),
                        autorest.WithPathParameters("/Services/{serviceName}/$/ReportHealth",pathParameters),
                        autorest.WithJSON(serviceHealthReport),
                        autorest.WithQueryParameters(queryParameters))
    return preparer.Prepare(&http.Request{})
}

// SendSender sends the Send request. The method will close the
// http.Response Body if it receives an error.
func (client ServiceHealthsClient) SendSender(req *http.Request) (*http.Response, error) {
    return autorest.SendWithSender(client, req)
}

// SendResponder handles the response to the Send request. The method always
// closes the http.Response Body.
func (client ServiceHealthsClient) SendResponder(resp *http.Response) (result String, err error) {
    err = autorest.Respond(
            resp,
            client.ByInspecting(),
            azure.WithErrorUnlessStatusCode(http.StatusOK),
            autorest.ByUnmarshallingJSON(&result.Value),
            autorest.ByClosing())
    result.Response = autorest.Response{Response: resp}
    return
}
