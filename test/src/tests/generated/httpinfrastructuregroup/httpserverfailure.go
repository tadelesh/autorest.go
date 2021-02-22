package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
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

// HTTPServerFailureClient is the test Infrastructure for AutoRest
type HTTPServerFailureClient struct {
	BaseClient
}

// NewHTTPServerFailureClient creates an instance of the HTTPServerFailureClient client.
func NewHTTPServerFailureClient() HTTPServerFailureClient {
	return NewHTTPServerFailureClientWithBaseURI(DefaultBaseURI)
}

// NewHTTPServerFailureClientWithBaseURI creates an instance of the HTTPServerFailureClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewHTTPServerFailureClientWithBaseURI(baseURI string) HTTPServerFailureClient {
	return HTTPServerFailureClient{NewWithBaseURI(baseURI)}
}

// Delete505 return 505 status code - should be represented in the client as an error
// Parameters:
// booleanValue - simple boolean value true
func (client HTTPServerFailureClient) Delete505(ctx context.Context, booleanValue *bool) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HTTPServerFailureClient.Delete505")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.Delete505Preparer(ctx, booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Delete505", nil, "Failure preparing request")
		return
	}

	resp, err := client.Delete505Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Delete505", resp, "Failure sending request")
		return
	}

	result, err = client.Delete505Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Delete505", resp, "Failure responding to request")
		return
	}

	return
}

// Delete505Preparer prepares the Delete505 request.
func (client HTTPServerFailureClient) Delete505Preparer(ctx context.Context, booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/505"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Delete505Sender sends the Delete505 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Delete505Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Delete505Responder handles the response to the Delete505 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Delete505Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get501 return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Get501(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HTTPServerFailureClient.Get501")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.Get501Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Get501", nil, "Failure preparing request")
		return
	}

	resp, err := client.Get501Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Get501", resp, "Failure sending request")
		return
	}

	result, err = client.Get501Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Get501", resp, "Failure responding to request")
		return
	}

	return
}

// Get501Preparer prepares the Get501 request.
func (client HTTPServerFailureClient) Get501Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/501"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Get501Sender sends the Get501 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Get501Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Get501Responder handles the response to the Get501 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Get501Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Head501 return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Head501(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HTTPServerFailureClient.Head501")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.Head501Preparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Head501", nil, "Failure preparing request")
		return
	}

	resp, err := client.Head501Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Head501", resp, "Failure sending request")
		return
	}

	result, err = client.Head501Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Head501", resp, "Failure responding to request")
		return
	}

	return
}

// Head501Preparer prepares the Head501 request.
func (client HTTPServerFailureClient) Head501Preparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsHead(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/501"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Head501Sender sends the Head501 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Head501Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Head501Responder handles the response to the Head501 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Head501Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Post505 return 505 status code - should be represented in the client as an error
// Parameters:
// booleanValue - simple boolean value true
func (client HTTPServerFailureClient) Post505(ctx context.Context, booleanValue *bool) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HTTPServerFailureClient.Post505")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.Post505Preparer(ctx, booleanValue)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Post505", nil, "Failure preparing request")
		return
	}

	resp, err := client.Post505Sender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Post505", resp, "Failure sending request")
		return
	}

	result, err = client.Post505Responder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPServerFailureClient", "Post505", resp, "Failure responding to request")
		return
	}

	return
}

// Post505Preparer prepares the Post505 request.
func (client HTTPServerFailureClient) Post505Preparer(ctx context.Context, booleanValue *bool) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/server/505"))
	if booleanValue != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(booleanValue))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Post505Sender sends the Post505 request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPServerFailureClient) Post505Sender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Post505Responder handles the response to the Post505 request. The method always
// closes the http.Response Body.
func (client HTTPServerFailureClient) Post505Responder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
