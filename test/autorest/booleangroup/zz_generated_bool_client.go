//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// BoolClient contains the methods for the Bool group.
// Don't use this type directly, use NewBoolClient() instead.
type BoolClient struct {
	con *Connection
}

// NewBoolClient creates a new instance of BoolClient with the specified values.
func NewBoolClient(con *Connection) *BoolClient {
	return &BoolClient{con: con}
}

// GetFalse - Get false Boolean value
// If the operation fails it returns the *Error error type.
func (client *BoolClient) GetFalse(ctx context.Context, options *BoolGetFalseOptions) (BoolGetFalseResponse, error) {
	req, err := client.getFalseCreateRequest(ctx, options)
	if err != nil {
		return BoolGetFalseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BoolGetFalseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolGetFalseResponse{}, client.getFalseHandleError(resp)
	}
	return client.getFalseHandleResponse(resp)
}

// getFalseCreateRequest creates the GetFalse request.
func (client *BoolClient) getFalseCreateRequest(ctx context.Context, options *BoolGetFalseOptions) (*azcore.Request, error) {
	urlPath := "/bool/false"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getFalseHandleResponse handles the GetFalse response.
func (client *BoolClient) getFalseHandleResponse(resp *azcore.Response) (BoolGetFalseResponse, error) {
	result := BoolGetFalseResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return BoolGetFalseResponse{}, err
	}
	return result, nil
}

// getFalseHandleError handles the GetFalse error response.
func (client *BoolClient) getFalseHandleError(resp *azcore.Response) error {
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

// GetInvalid - Get invalid Boolean value
// If the operation fails it returns the *Error error type.
func (client *BoolClient) GetInvalid(ctx context.Context, options *BoolGetInvalidOptions) (BoolGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return BoolGetInvalidResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BoolGetInvalidResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolGetInvalidResponse{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *BoolClient) getInvalidCreateRequest(ctx context.Context, options *BoolGetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/bool/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *BoolClient) getInvalidHandleResponse(resp *azcore.Response) (BoolGetInvalidResponse, error) {
	result := BoolGetInvalidResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return BoolGetInvalidResponse{}, err
	}
	return result, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *BoolClient) getInvalidHandleError(resp *azcore.Response) error {
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

// GetNull - Get null Boolean value
// If the operation fails it returns the *Error error type.
func (client *BoolClient) GetNull(ctx context.Context, options *BoolGetNullOptions) (BoolGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return BoolGetNullResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BoolGetNullResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolGetNullResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *BoolClient) getNullCreateRequest(ctx context.Context, options *BoolGetNullOptions) (*azcore.Request, error) {
	urlPath := "/bool/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *BoolClient) getNullHandleResponse(resp *azcore.Response) (BoolGetNullResponse, error) {
	result := BoolGetNullResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return BoolGetNullResponse{}, err
	}
	return result, nil
}

// getNullHandleError handles the GetNull error response.
func (client *BoolClient) getNullHandleError(resp *azcore.Response) error {
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

// GetTrue - Get true Boolean value
// If the operation fails it returns the *Error error type.
func (client *BoolClient) GetTrue(ctx context.Context, options *BoolGetTrueOptions) (BoolGetTrueResponse, error) {
	req, err := client.getTrueCreateRequest(ctx, options)
	if err != nil {
		return BoolGetTrueResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BoolGetTrueResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolGetTrueResponse{}, client.getTrueHandleError(resp)
	}
	return client.getTrueHandleResponse(resp)
}

// getTrueCreateRequest creates the GetTrue request.
func (client *BoolClient) getTrueCreateRequest(ctx context.Context, options *BoolGetTrueOptions) (*azcore.Request, error) {
	urlPath := "/bool/true"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getTrueHandleResponse handles the GetTrue response.
func (client *BoolClient) getTrueHandleResponse(resp *azcore.Response) (BoolGetTrueResponse, error) {
	result := BoolGetTrueResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return BoolGetTrueResponse{}, err
	}
	return result, nil
}

// getTrueHandleError handles the GetTrue error response.
func (client *BoolClient) getTrueHandleError(resp *azcore.Response) error {
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

// PutFalse - Set Boolean value false
// If the operation fails it returns the *Error error type.
func (client *BoolClient) PutFalse(ctx context.Context, options *BoolPutFalseOptions) (BoolPutFalseResponse, error) {
	req, err := client.putFalseCreateRequest(ctx, options)
	if err != nil {
		return BoolPutFalseResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BoolPutFalseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolPutFalseResponse{}, client.putFalseHandleError(resp)
	}
	return BoolPutFalseResponse{RawResponse: resp.Response}, nil
}

// putFalseCreateRequest creates the PutFalse request.
func (client *BoolClient) putFalseCreateRequest(ctx context.Context, options *BoolPutFalseOptions) (*azcore.Request, error) {
	urlPath := "/bool/false"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(false)
}

// putFalseHandleError handles the PutFalse error response.
func (client *BoolClient) putFalseHandleError(resp *azcore.Response) error {
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

// PutTrue - Set Boolean value true
// If the operation fails it returns the *Error error type.
func (client *BoolClient) PutTrue(ctx context.Context, options *BoolPutTrueOptions) (BoolPutTrueResponse, error) {
	req, err := client.putTrueCreateRequest(ctx, options)
	if err != nil {
		return BoolPutTrueResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BoolPutTrueResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolPutTrueResponse{}, client.putTrueHandleError(resp)
	}
	return BoolPutTrueResponse{RawResponse: resp.Response}, nil
}

// putTrueCreateRequest creates the PutTrue request.
func (client *BoolClient) putTrueCreateRequest(ctx context.Context, options *BoolPutTrueOptions) (*azcore.Request, error) {
	urlPath := "/bool/true"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// putTrueHandleError handles the PutTrue error response.
func (client *BoolClient) putTrueHandleError(resp *azcore.Response) error {
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
