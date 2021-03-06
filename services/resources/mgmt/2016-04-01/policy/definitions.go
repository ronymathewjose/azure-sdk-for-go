package policy

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DefinitionsClient is the to manage and control access to your resources, you can define customized policies and
// assign them at a scope.
type DefinitionsClient struct {
	BaseClient
}

// NewDefinitionsClient creates an instance of the DefinitionsClient client.
func NewDefinitionsClient(subscriptionID string) DefinitionsClient {
	return NewDefinitionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDefinitionsClientWithBaseURI creates an instance of the DefinitionsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDefinitionsClientWithBaseURI(baseURI string, subscriptionID string) DefinitionsClient {
	return DefinitionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a policy definition.
// Parameters:
// policyDefinitionName - the name of the policy definition to create.
// parameters - the policy definition properties.
func (client DefinitionsClient) CreateOrUpdate(ctx context.Context, policyDefinitionName string, parameters Definition) (result Definition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, policyDefinitionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DefinitionsClient) CreateOrUpdatePreparer(ctx context.Context, policyDefinitionName string, parameters Definition) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	parameters.ID = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) CreateOrUpdateResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes a policy definition.
// Parameters:
// policyDefinitionName - the name of the policy definition to delete.
func (client DefinitionsClient) Delete(ctx context.Context, policyDefinitionName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DefinitionsClient) DeletePreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the policy definition.
// Parameters:
// policyDefinitionName - the name of the policy definition to get.
func (client DefinitionsClient) Get(ctx context.Context, policyDefinitionName string) (result Definition, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, policyDefinitionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client DefinitionsClient) GetPreparer(ctx context.Context, policyDefinitionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"policyDefinitionName": autorest.Encode("path", policyDefinitionName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions/{policyDefinitionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) GetResponder(resp *http.Response) (result Definition, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all the policy definitions for a subscription.
// Parameters:
// filter - the filter to apply on the operation.
func (client DefinitionsClient) List(ctx context.Context, filter string) (result DefinitionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.List")
		defer func() {
			sc := -1
			if result.dlr.Response.Response != nil {
				sc = result.dlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure sending request")
		return
	}

	result.dlr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client DefinitionsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-04-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/policydefinitions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DefinitionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DefinitionsClient) ListResponder(resp *http.Response) (result DefinitionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DefinitionsClient) listNextResults(ctx context.Context, lastResults DefinitionListResult) (result DefinitionListResult, err error) {
	req, err := lastResults.definitionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policy.DefinitionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DefinitionsClient) ListComplete(ctx context.Context, filter string) (result DefinitionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DefinitionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}
