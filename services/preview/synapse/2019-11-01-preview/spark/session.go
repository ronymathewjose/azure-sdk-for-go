package spark

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
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SessionClient is the client for the Session methods of the Spark service.
type SessionClient struct {
	BaseClient
}

// NewSessionClient creates an instance of the SessionClient client.
func NewSessionClient(endpoint string, sparkPoolName string) SessionClient {
	return SessionClient{New(endpoint, sparkPoolName)}
}

// CancelSparkSession cancels a running spark session.
// Parameters:
// sessionID - identifier for the session.
func (client SessionClient) CancelSparkSession(ctx context.Context, sessionID int32) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.CancelSparkSession")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelSparkSessionPreparer(ctx, sessionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CancelSparkSession", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSparkSessionSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CancelSparkSession", resp, "Failure sending request")
		return
	}

	result, err = client.CancelSparkSessionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CancelSparkSession", resp, "Failure responding to request")
	}

	return
}

// CancelSparkSessionPreparer prepares the CancelSparkSession request.
func (client SessionClient) CancelSparkSessionPreparer(ctx context.Context, sessionID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	pathParameters := map[string]interface{}{
		"sessionId": autorest.Encode("path", sessionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPathParameters("/sessions/{sessionId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSparkSessionSender sends the CancelSparkSession request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) CancelSparkSessionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CancelSparkSessionResponder handles the response to the CancelSparkSession request. The method always
// closes the http.Response Body.
func (client SessionClient) CancelSparkSessionResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// CancelSparkStatement kill a statement within a session.
// Parameters:
// sessionID - identifier for the session.
// statementID - identifier for the statement.
func (client SessionClient) CancelSparkStatement(ctx context.Context, sessionID int32, statementID int32) (result StatementCancellationResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.CancelSparkStatement")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CancelSparkStatementPreparer(ctx, sessionID, statementID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CancelSparkStatement", nil, "Failure preparing request")
		return
	}

	resp, err := client.CancelSparkStatementSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CancelSparkStatement", resp, "Failure sending request")
		return
	}

	result, err = client.CancelSparkStatementResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CancelSparkStatement", resp, "Failure responding to request")
	}

	return
}

// CancelSparkStatementPreparer prepares the CancelSparkStatement request.
func (client SessionClient) CancelSparkStatementPreparer(ctx context.Context, sessionID int32, statementID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	pathParameters := map[string]interface{}{
		"sessionId":   autorest.Encode("path", sessionID),
		"statementId": autorest.Encode("path", statementID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPathParameters("/sessions/{sessionId}/statements/{statementId}/cancel", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CancelSparkStatementSender sends the CancelSparkStatement request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) CancelSparkStatementSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CancelSparkStatementResponder handles the response to the CancelSparkStatement request. The method always
// closes the http.Response Body.
func (client SessionClient) CancelSparkStatementResponder(resp *http.Response) (result StatementCancellationResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateSparkSession create new spark session.
// Parameters:
// sparkSessionOptions - livy compatible batch job request payload.
// detailed - optional query param specifying whether detailed response is returned beyond plain livy.
func (client SessionClient) CreateSparkSession(ctx context.Context, sparkSessionOptions SessionOptions, detailed *bool) (result Session, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.CreateSparkSession")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: sparkSessionOptions,
			Constraints: []validation.Constraint{{Target: "sparkSessionOptions.Name", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("spark.SessionClient", "CreateSparkSession", err.Error())
	}

	req, err := client.CreateSparkSessionPreparer(ctx, sparkSessionOptions, detailed)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CreateSparkSession", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSparkSessionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CreateSparkSession", resp, "Failure sending request")
		return
	}

	result, err = client.CreateSparkSessionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CreateSparkSession", resp, "Failure responding to request")
	}

	return
}

// CreateSparkSessionPreparer prepares the CreateSparkSession request.
func (client SessionClient) CreateSparkSessionPreparer(ctx context.Context, sparkSessionOptions SessionOptions, detailed *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	queryParameters := map[string]interface{}{}
	if detailed != nil {
		queryParameters["detailed"] = autorest.Encode("query", *detailed)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPath("/sessions"),
		autorest.WithJSON(sparkSessionOptions),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSparkSessionSender sends the CreateSparkSession request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) CreateSparkSessionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateSparkSessionResponder handles the response to the CreateSparkSession request. The method always
// closes the http.Response Body.
func (client SessionClient) CreateSparkSessionResponder(resp *http.Response) (result Session, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateSparkStatement create statement within a spark session.
// Parameters:
// sessionID - identifier for the session.
// sparkStatementOptions - livy compatible batch job request payload.
func (client SessionClient) CreateSparkStatement(ctx context.Context, sessionID int32, sparkStatementOptions StatementOptions) (result Statement, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.CreateSparkStatement")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateSparkStatementPreparer(ctx, sessionID, sparkStatementOptions)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CreateSparkStatement", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateSparkStatementSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CreateSparkStatement", resp, "Failure sending request")
		return
	}

	result, err = client.CreateSparkStatementResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "CreateSparkStatement", resp, "Failure responding to request")
	}

	return
}

// CreateSparkStatementPreparer prepares the CreateSparkStatement request.
func (client SessionClient) CreateSparkStatementPreparer(ctx context.Context, sessionID int32, sparkStatementOptions StatementOptions) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	pathParameters := map[string]interface{}{
		"sessionId": autorest.Encode("path", sessionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPathParameters("/sessions/{sessionId}/statements", pathParameters),
		autorest.WithJSON(sparkStatementOptions))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSparkStatementSender sends the CreateSparkStatement request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) CreateSparkStatementSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// CreateSparkStatementResponder handles the response to the CreateSparkStatement request. The method always
// closes the http.Response Body.
func (client SessionClient) CreateSparkStatementResponder(resp *http.Response) (result Statement, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSparkSession gets a single spark session.
// Parameters:
// sessionID - identifier for the session.
// detailed - optional query param specifying whether detailed response is returned beyond plain livy.
func (client SessionClient) GetSparkSession(ctx context.Context, sessionID int32, detailed *bool) (result Session, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.GetSparkSession")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSparkSessionPreparer(ctx, sessionID, detailed)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkSession", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSparkSessionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkSession", resp, "Failure sending request")
		return
	}

	result, err = client.GetSparkSessionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkSession", resp, "Failure responding to request")
	}

	return
}

// GetSparkSessionPreparer prepares the GetSparkSession request.
func (client SessionClient) GetSparkSessionPreparer(ctx context.Context, sessionID int32, detailed *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	pathParameters := map[string]interface{}{
		"sessionId": autorest.Encode("path", sessionID),
	}

	queryParameters := map[string]interface{}{}
	if detailed != nil {
		queryParameters["detailed"] = autorest.Encode("query", *detailed)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPathParameters("/sessions/{sessionId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSparkSessionSender sends the GetSparkSession request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) GetSparkSessionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSparkSessionResponder handles the response to the GetSparkSession request. The method always
// closes the http.Response Body.
func (client SessionClient) GetSparkSessionResponder(resp *http.Response) (result Session, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSparkSessions list all spark sessions which are running under a particular spark pool.
// Parameters:
// from - optional param specifying which index the list should begin from.
// size - optional param specifying the size of the returned list.
// By default it is 20 and that is the maximum.
// detailed - optional query param specifying whether detailed response is returned beyond plain livy.
func (client SessionClient) GetSparkSessions(ctx context.Context, from *int32, size *int32, detailed *bool) (result SessionCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.GetSparkSessions")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSparkSessionsPreparer(ctx, from, size, detailed)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkSessions", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSparkSessionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkSessions", resp, "Failure sending request")
		return
	}

	result, err = client.GetSparkSessionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkSessions", resp, "Failure responding to request")
	}

	return
}

// GetSparkSessionsPreparer prepares the GetSparkSessions request.
func (client SessionClient) GetSparkSessionsPreparer(ctx context.Context, from *int32, size *int32, detailed *bool) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	queryParameters := map[string]interface{}{}
	if from != nil {
		queryParameters["from"] = autorest.Encode("query", *from)
	}
	if size != nil {
		queryParameters["size"] = autorest.Encode("query", *size)
	}
	if detailed != nil {
		queryParameters["detailed"] = autorest.Encode("query", *detailed)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPath("/sessions"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSparkSessionsSender sends the GetSparkSessions request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) GetSparkSessionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSparkSessionsResponder handles the response to the GetSparkSessions request. The method always
// closes the http.Response Body.
func (client SessionClient) GetSparkSessionsResponder(resp *http.Response) (result SessionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSparkStatement gets a single statement within a spark session.
// Parameters:
// sessionID - identifier for the session.
// statementID - identifier for the statement.
func (client SessionClient) GetSparkStatement(ctx context.Context, sessionID int32, statementID int32) (result Statement, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.GetSparkStatement")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSparkStatementPreparer(ctx, sessionID, statementID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkStatement", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSparkStatementSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkStatement", resp, "Failure sending request")
		return
	}

	result, err = client.GetSparkStatementResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkStatement", resp, "Failure responding to request")
	}

	return
}

// GetSparkStatementPreparer prepares the GetSparkStatement request.
func (client SessionClient) GetSparkStatementPreparer(ctx context.Context, sessionID int32, statementID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	pathParameters := map[string]interface{}{
		"sessionId":   autorest.Encode("path", sessionID),
		"statementId": autorest.Encode("path", statementID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPathParameters("/sessions/{sessionId}/statements/{statementId}", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSparkStatementSender sends the GetSparkStatement request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) GetSparkStatementSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSparkStatementResponder handles the response to the GetSparkStatement request. The method always
// closes the http.Response Body.
func (client SessionClient) GetSparkStatementResponder(resp *http.Response) (result Statement, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSparkStatements gets a list of statements within a spark session.
// Parameters:
// sessionID - identifier for the session.
func (client SessionClient) GetSparkStatements(ctx context.Context, sessionID int32) (result StatementCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.GetSparkStatements")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSparkStatementsPreparer(ctx, sessionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkStatements", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSparkStatementsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkStatements", resp, "Failure sending request")
		return
	}

	result, err = client.GetSparkStatementsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "GetSparkStatements", resp, "Failure responding to request")
	}

	return
}

// GetSparkStatementsPreparer prepares the GetSparkStatements request.
func (client SessionClient) GetSparkStatementsPreparer(ctx context.Context, sessionID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	pathParameters := map[string]interface{}{
		"sessionId": autorest.Encode("path", sessionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPathParameters("/sessions/{sessionId}/statements", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSparkStatementsSender sends the GetSparkStatements request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) GetSparkStatementsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetSparkStatementsResponder handles the response to the GetSparkStatements request. The method always
// closes the http.Response Body.
func (client SessionClient) GetSparkStatementsResponder(resp *http.Response) (result StatementCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ResetSparkSessionTimeout sends a keep alive call to the current session to reset the session timeout.
// Parameters:
// sessionID - identifier for the session.
func (client SessionClient) ResetSparkSessionTimeout(ctx context.Context, sessionID int32) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SessionClient.ResetSparkSessionTimeout")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ResetSparkSessionTimeoutPreparer(ctx, sessionID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "ResetSparkSessionTimeout", nil, "Failure preparing request")
		return
	}

	resp, err := client.ResetSparkSessionTimeoutSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "ResetSparkSessionTimeout", resp, "Failure sending request")
		return
	}

	result, err = client.ResetSparkSessionTimeoutResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "spark.SessionClient", "ResetSparkSessionTimeout", resp, "Failure responding to request")
	}

	return
}

// ResetSparkSessionTimeoutPreparer prepares the ResetSparkSessionTimeout request.
func (client SessionClient) ResetSparkSessionTimeoutPreparer(ctx context.Context, sessionID int32) (*http.Request, error) {
	urlParameters := map[string]interface{}{
		"endpoint":       client.Endpoint,
		"livyApiVersion": client.LivyAPIVersion,
		"sparkPoolName":  client.SparkPoolName,
	}

	pathParameters := map[string]interface{}{
		"sessionId": autorest.Encode("path", sessionID),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPut(),
		autorest.WithCustomBaseURL("{endpoint}/livyApi/versions/{livyApiVersion}/sparkPools/{sparkPoolName}", urlParameters),
		autorest.WithPathParameters("/sessions/{sessionId}/reset-timeout", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ResetSparkSessionTimeoutSender sends the ResetSparkSessionTimeout request. The method will close the
// http.Response Body if it receives an error.
func (client SessionClient) ResetSparkSessionTimeoutSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ResetSparkSessionTimeoutResponder handles the response to the ResetSparkSessionTimeout request. The method always
// closes the http.Response Body.
func (client SessionClient) ResetSparkSessionTimeoutResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
