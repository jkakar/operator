package network

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
// Code generated by Microsoft (R) AutoRest Code Generator 0.14.0.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
	"net/url"
)

// VirtualNetworksClient is the the Microsoft Azure Network management API
// provides a RESTful set of web services that interact with Microsoft Azure
// Networks service to manage your network resrources. The API has entities
// that capture the relationship between an end user and the Microsoft Azure
// Networks service.
type VirtualNetworksClient struct {
	ManagementClient
}

// NewVirtualNetworksClient creates an instance of the VirtualNetworksClient
// client.
func NewVirtualNetworksClient(subscriptionID string) VirtualNetworksClient {
	return NewVirtualNetworksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworksClientWithBaseURI creates an instance of the
// VirtualNetworksClient client.
func NewVirtualNetworksClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworksClient {
	return VirtualNetworksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate the Put VirtualNetwork operation creates/updates a virtual
// network in the specified resource group. This method may poll for
// completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. parameters is parameters supplied to the
// create/update Virtual Network operation
func (client VirtualNetworksClient) CreateOrUpdate(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.CreateOrUpdatePreparer(resourceGroupName, virtualNetworkName, parameters, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", nil, "Failure preparing request")
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", resp, "Failure sending request")
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworksClient) CreateOrUpdatePreparer(resourceGroupName string, virtualNetworkName string, parameters VirtualNetwork, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{Cancel: cancel},
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}"),
		autorest.WithJSON(parameters),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) CreateOrUpdateResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Delete the Delete VirtualNetwork operation deletes the specifed virtual
// network This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel
// polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network.
func (client VirtualNetworksClient) Delete(resourceGroupName string, virtualNetworkName string, cancel <-chan struct{}) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, virtualNetworkName, cancel)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", nil, "Failure preparing request")
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", resp, "Failure sending request")
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworksClient) DeletePreparer(resourceGroupName string, virtualNetworkName string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{Cancel: cancel},
		autorest.AsJSON(),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get the Get VirtualNetwork operation retrieves information about the
// specified virtual network.
//
// resourceGroupName is the name of the resource group. virtualNetworkName is
// the name of the virtual network. expand is expand references resources.
func (client VirtualNetworksClient) Get(resourceGroupName string, virtualNetworkName string, expand string) (result VirtualNetwork, err error) {
	req, err := client.GetPreparer(resourceGroupName, virtualNetworkName, expand)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", nil, "Failure preparing request")
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", resp, "Failure sending request")
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworksClient) GetPreparer(resourceGroupName string, virtualNetworkName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":  url.QueryEscape(resourceGroupName),
		"subscriptionId":     url.QueryEscape(client.SubscriptionID),
		"virtualNetworkName": url.QueryEscape(virtualNetworkName),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = expand
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks/{virtualNetworkName}"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) GetResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List the list VirtualNetwork returns all Virtual Networks in a resource
// group
//
// resourceGroupName is the name of the resource group.
func (client VirtualNetworksClient) List(resourceGroupName string) (result VirtualNetworkListResult, err error) {
	req, err := client.ListPreparer(resourceGroupName)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", nil, "Failure preparing request")
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure sending request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworksClient) ListPreparer(resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": url.QueryEscape(resourceGroupName),
		"subscriptionId":    url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualnetworks"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
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
func (client VirtualNetworksClient) ListNextResults(lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, err error) {
	req, err := lastResults.VirtualNetworkListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure sending next results request request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "List", resp, "Failure responding to next results request request")
	}

	return
}

// ListAll the list VirtualNetwork returns all Virtual Networks in a
// subscription
func (client VirtualNetworksClient) ListAll() (result VirtualNetworkListResult, err error) {
	req, err := client.ListAllPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", nil, "Failure preparing request")
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure sending request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure responding to request")
	}

	return
}

// ListAllPreparer prepares the ListAll request.
func (client VirtualNetworksClient) ListAllPreparer() (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": url.QueryEscape(client.SubscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/subscriptions/{subscriptionId}/providers/Microsoft.Network/virtualnetworks"),
		autorest.WithPathParameters(pathParameters),
		autorest.WithQueryParameters(queryParameters))
}

// ListAllSender sends the ListAll request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListAllSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListAllResponder handles the response to the ListAll request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListAllResponder(resp *http.Response) (result VirtualNetworkListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAllNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) ListAllNextResults(lastResults VirtualNetworkListResult) (result VirtualNetworkListResult, err error) {
	req, err := lastResults.VirtualNetworkListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", nil, "Failure preparing next results request request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListAllSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure sending next results request request")
	}

	result, err = client.ListAllResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "network.VirtualNetworksClient", "ListAll", resp, "Failure responding to next results request request")
	}

	return
}
