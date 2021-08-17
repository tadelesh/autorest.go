//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// SkipURLEncodingClient contains the methods for the SkipURLEncoding group.
// Don't use this type directly, use NewSkipURLEncodingClient() instead.
type SkipURLEncodingClient struct {
	con *Connection
}

// NewSkipURLEncodingClient creates a new instance of SkipURLEncodingClient with the specified values.
func NewSkipURLEncodingClient(con *Connection) *SkipURLEncodingClient {
	return &SkipURLEncodingClient{con: con}
}

// GetMethodPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
// If the operation fails it returns the *Error error type.
func (client *SkipURLEncodingClient) GetMethodPathValid(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingGetMethodPathValidOptions) (SkipURLEncodingGetMethodPathValidResponse, error) {
	req, err := client.getMethodPathValidCreateRequest(ctx, unencodedPathParam, options)
	if err != nil {
		return SkipURLEncodingGetMethodPathValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SkipURLEncodingGetMethodPathValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SkipURLEncodingGetMethodPathValidResponse{}, client.getMethodPathValidHandleError(resp)
	}
	return SkipURLEncodingGetMethodPathValidResponse{RawResponse: resp.Response}, nil
}

// getMethodPathValidCreateRequest creates the GetMethodPathValid request.
func (client *SkipURLEncodingClient) getMethodPathValidCreateRequest(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingGetMethodPathValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/method/path/valid/{unencodedPathParam}"
	if unencodedPathParam == "" {
		return nil, errors.New("parameter unencodedPathParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{unencodedPathParam}", unencodedPathParam)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodPathValidHandleError handles the GetMethodPathValid error response.
func (client *SkipURLEncodingClient) getMethodPathValidHandleError(resp *azcore.Response) error {
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

// GetMethodQueryNull - Get method with unencoded query parameter with value null
// If the operation fails it returns the *Error error type.
func (client *SkipURLEncodingClient) GetMethodQueryNull(ctx context.Context, options *SkipURLEncodingGetMethodQueryNullOptions) (SkipURLEncodingGetMethodQueryNullResponse, error) {
	req, err := client.getMethodQueryNullCreateRequest(ctx, options)
	if err != nil {
		return SkipURLEncodingGetMethodQueryNullResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SkipURLEncodingGetMethodQueryNullResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SkipURLEncodingGetMethodQueryNullResponse{}, client.getMethodQueryNullHandleError(resp)
	}
	return SkipURLEncodingGetMethodQueryNullResponse{RawResponse: resp.Response}, nil
}

// getMethodQueryNullCreateRequest creates the GetMethodQueryNull request.
func (client *SkipURLEncodingClient) getMethodQueryNullCreateRequest(ctx context.Context, options *SkipURLEncodingGetMethodQueryNullOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/method/query/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	unencodedParams := []string{}
	if options != nil && options.Q1 != nil {
		unencodedParams = append(unencodedParams, "q1="+*options.Q1)
	}
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodQueryNullHandleError handles the GetMethodQueryNull error response.
func (client *SkipURLEncodingClient) getMethodQueryNullHandleError(resp *azcore.Response) error {
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

// GetMethodQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
// If the operation fails it returns the *Error error type.
func (client *SkipURLEncodingClient) GetMethodQueryValid(ctx context.Context, q1 string, options *SkipURLEncodingGetMethodQueryValidOptions) (SkipURLEncodingGetMethodQueryValidResponse, error) {
	req, err := client.getMethodQueryValidCreateRequest(ctx, q1, options)
	if err != nil {
		return SkipURLEncodingGetMethodQueryValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SkipURLEncodingGetMethodQueryValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SkipURLEncodingGetMethodQueryValidResponse{}, client.getMethodQueryValidHandleError(resp)
	}
	return SkipURLEncodingGetMethodQueryValidResponse{RawResponse: resp.Response}, nil
}

// getMethodQueryValidCreateRequest creates the GetMethodQueryValid request.
func (client *SkipURLEncodingClient) getMethodQueryValidCreateRequest(ctx context.Context, q1 string, options *SkipURLEncodingGetMethodQueryValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/method/query/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	unencodedParams := []string{}
	unencodedParams = append(unencodedParams, "q1="+q1)
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodQueryValidHandleError handles the GetMethodQueryValid error response.
func (client *SkipURLEncodingClient) getMethodQueryValidHandleError(resp *azcore.Response) error {
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

// GetPathQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
// If the operation fails it returns the *Error error type.
func (client *SkipURLEncodingClient) GetPathQueryValid(ctx context.Context, q1 string, options *SkipURLEncodingGetPathQueryValidOptions) (SkipURLEncodingGetPathQueryValidResponse, error) {
	req, err := client.getPathQueryValidCreateRequest(ctx, q1, options)
	if err != nil {
		return SkipURLEncodingGetPathQueryValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SkipURLEncodingGetPathQueryValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SkipURLEncodingGetPathQueryValidResponse{}, client.getPathQueryValidHandleError(resp)
	}
	return SkipURLEncodingGetPathQueryValidResponse{RawResponse: resp.Response}, nil
}

// getPathQueryValidCreateRequest creates the GetPathQueryValid request.
func (client *SkipURLEncodingClient) getPathQueryValidCreateRequest(ctx context.Context, q1 string, options *SkipURLEncodingGetPathQueryValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/path/query/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	unencodedParams := []string{}
	unencodedParams = append(unencodedParams, "q1="+q1)
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPathQueryValidHandleError handles the GetPathQueryValid error response.
func (client *SkipURLEncodingClient) getPathQueryValidHandleError(resp *azcore.Response) error {
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

// GetPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
// If the operation fails it returns the *Error error type.
func (client *SkipURLEncodingClient) GetPathValid(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingGetPathValidOptions) (SkipURLEncodingGetPathValidResponse, error) {
	req, err := client.getPathValidCreateRequest(ctx, unencodedPathParam, options)
	if err != nil {
		return SkipURLEncodingGetPathValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SkipURLEncodingGetPathValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SkipURLEncodingGetPathValidResponse{}, client.getPathValidHandleError(resp)
	}
	return SkipURLEncodingGetPathValidResponse{RawResponse: resp.Response}, nil
}

// getPathValidCreateRequest creates the GetPathValid request.
func (client *SkipURLEncodingClient) getPathValidCreateRequest(ctx context.Context, unencodedPathParam string, options *SkipURLEncodingGetPathValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/path/path/valid/{unencodedPathParam}"
	if unencodedPathParam == "" {
		return nil, errors.New("parameter unencodedPathParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{unencodedPathParam}", unencodedPathParam)
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPathValidHandleError handles the GetPathValid error response.
func (client *SkipURLEncodingClient) getPathValidHandleError(resp *azcore.Response) error {
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

// GetSwaggerPathValid - Get method with unencoded path parameter with value 'path1/path2/path3'
// If the operation fails it returns the *Error error type.
func (client *SkipURLEncodingClient) GetSwaggerPathValid(ctx context.Context, options *SkipURLEncodingGetSwaggerPathValidOptions) (SkipURLEncodingGetSwaggerPathValidResponse, error) {
	req, err := client.getSwaggerPathValidCreateRequest(ctx, options)
	if err != nil {
		return SkipURLEncodingGetSwaggerPathValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SkipURLEncodingGetSwaggerPathValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SkipURLEncodingGetSwaggerPathValidResponse{}, client.getSwaggerPathValidHandleError(resp)
	}
	return SkipURLEncodingGetSwaggerPathValidResponse{RawResponse: resp.Response}, nil
}

// getSwaggerPathValidCreateRequest creates the GetSwaggerPathValid request.
func (client *SkipURLEncodingClient) getSwaggerPathValidCreateRequest(ctx context.Context, options *SkipURLEncodingGetSwaggerPathValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/swagger/path/valid/{unencodedPathParam}"
	urlPath = strings.ReplaceAll(urlPath, "{unencodedPathParam}", "path1/path2/path3")
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSwaggerPathValidHandleError handles the GetSwaggerPathValid error response.
func (client *SkipURLEncodingClient) getSwaggerPathValidHandleError(resp *azcore.Response) error {
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

// GetSwaggerQueryValid - Get method with unencoded query parameter with value 'value1&q2=value2&q3=value3'
// If the operation fails it returns the *Error error type.
func (client *SkipURLEncodingClient) GetSwaggerQueryValid(ctx context.Context, options *SkipURLEncodingGetSwaggerQueryValidOptions) (SkipURLEncodingGetSwaggerQueryValidResponse, error) {
	req, err := client.getSwaggerQueryValidCreateRequest(ctx, options)
	if err != nil {
		return SkipURLEncodingGetSwaggerQueryValidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SkipURLEncodingGetSwaggerQueryValidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SkipURLEncodingGetSwaggerQueryValidResponse{}, client.getSwaggerQueryValidHandleError(resp)
	}
	return SkipURLEncodingGetSwaggerQueryValidResponse{RawResponse: resp.Response}, nil
}

// getSwaggerQueryValidCreateRequest creates the GetSwaggerQueryValid request.
func (client *SkipURLEncodingClient) getSwaggerQueryValidCreateRequest(ctx context.Context, options *SkipURLEncodingGetSwaggerQueryValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/skipUrlEncoding/swagger/query/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	unencodedParams := []string{}
	unencodedParams = append(unencodedParams, "q1="+"value1&q2=value2&q3=value3")
	req.URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSwaggerQueryValidHandleError handles the GetSwaggerQueryValid error response.
func (client *SkipURLEncodingClient) getSwaggerQueryValidHandleError(resp *azcore.Response) error {
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
