//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPFailureClient contains the methods for the HTTPFailure group.
// Don't use this type directly, use NewHTTPFailureClient() instead.
type HTTPFailureClient struct {
	con *Connection
}

// NewHTTPFailureClient creates a new instance of HTTPFailureClient with the specified values.
func NewHTTPFailureClient(con *Connection) *HTTPFailureClient {
	return &HTTPFailureClient{con: con}
}

// GetEmptyError - Get empty error form server
// If the operation fails it returns the *Error error type.
func (client *HTTPFailureClient) GetEmptyError(ctx context.Context, options *HTTPFailureGetEmptyErrorOptions) (HTTPFailureGetEmptyErrorResponse, error) {
	req, err := client.getEmptyErrorCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureGetEmptyErrorResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPFailureGetEmptyErrorResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HTTPFailureGetEmptyErrorResponse{}, client.getEmptyErrorHandleError(resp)
	}
	return client.getEmptyErrorHandleResponse(resp)
}

// getEmptyErrorCreateRequest creates the GetEmptyError request.
func (client *HTTPFailureClient) getEmptyErrorCreateRequest(ctx context.Context, options *HTTPFailureGetEmptyErrorOptions) (*azcore.Request, error) {
	urlPath := "/http/failure/emptybody/error"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyErrorHandleResponse handles the GetEmptyError response.
func (client *HTTPFailureClient) getEmptyErrorHandleResponse(resp *azcore.Response) (HTTPFailureGetEmptyErrorResponse, error) {
	result := HTTPFailureGetEmptyErrorResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return HTTPFailureGetEmptyErrorResponse{}, err
	}
	return result, nil
}

// getEmptyErrorHandleError handles the GetEmptyError error response.
func (client *HTTPFailureClient) getEmptyErrorHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := Error{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetNoModelEmpty - Get empty response from server
// If the operation fails it returns a generic error.
func (client *HTTPFailureClient) GetNoModelEmpty(ctx context.Context, options *HTTPFailureGetNoModelEmptyOptions) (HTTPFailureGetNoModelEmptyResponse, error) {
	req, err := client.getNoModelEmptyCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureGetNoModelEmptyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPFailureGetNoModelEmptyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HTTPFailureGetNoModelEmptyResponse{}, client.getNoModelEmptyHandleError(resp)
	}
	return client.getNoModelEmptyHandleResponse(resp)
}

// getNoModelEmptyCreateRequest creates the GetNoModelEmpty request.
func (client *HTTPFailureClient) getNoModelEmptyCreateRequest(ctx context.Context, options *HTTPFailureGetNoModelEmptyOptions) (*azcore.Request, error) {
	urlPath := "/http/failure/nomodel/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNoModelEmptyHandleResponse handles the GetNoModelEmpty response.
func (client *HTTPFailureClient) getNoModelEmptyHandleResponse(resp *azcore.Response) (HTTPFailureGetNoModelEmptyResponse, error) {
	result := HTTPFailureGetNoModelEmptyResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return HTTPFailureGetNoModelEmptyResponse{}, err
	}
	return result, nil
}

// getNoModelEmptyHandleError handles the GetNoModelEmpty error response.
func (client *HTTPFailureClient) getNoModelEmptyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetNoModelError - Get empty error form server
// If the operation fails it returns a generic error.
func (client *HTTPFailureClient) GetNoModelError(ctx context.Context, options *HTTPFailureGetNoModelErrorOptions) (HTTPFailureGetNoModelErrorResponse, error) {
	req, err := client.getNoModelErrorCreateRequest(ctx, options)
	if err != nil {
		return HTTPFailureGetNoModelErrorResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return HTTPFailureGetNoModelErrorResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return HTTPFailureGetNoModelErrorResponse{}, client.getNoModelErrorHandleError(resp)
	}
	return client.getNoModelErrorHandleResponse(resp)
}

// getNoModelErrorCreateRequest creates the GetNoModelError request.
func (client *HTTPFailureClient) getNoModelErrorCreateRequest(ctx context.Context, options *HTTPFailureGetNoModelErrorOptions) (*azcore.Request, error) {
	urlPath := "/http/failure/nomodel/error"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNoModelErrorHandleResponse handles the GetNoModelError response.
func (client *HTTPFailureClient) getNoModelErrorHandleResponse(resp *azcore.Response) (HTTPFailureGetNoModelErrorResponse, error) {
	result := HTTPFailureGetNoModelErrorResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return HTTPFailureGetNoModelErrorResponse{}, err
	}
	return result, nil
}

// getNoModelErrorHandleError handles the GetNoModelError error response.
func (client *HTTPFailureClient) getNoModelErrorHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
