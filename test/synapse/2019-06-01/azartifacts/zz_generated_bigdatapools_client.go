//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type bigDataPoolsClient struct {
	con *connection
}

// Get - Get Big Data Pool
// If the operation fails it returns the *ErrorContract error type.
func (client *bigDataPoolsClient) Get(ctx context.Context, bigDataPoolName string, options *BigDataPoolsGetOptions) (BigDataPoolsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, bigDataPoolName, options)
	if err != nil {
		return BigDataPoolsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BigDataPoolsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BigDataPoolsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *bigDataPoolsClient) getCreateRequest(ctx context.Context, bigDataPoolName string, options *BigDataPoolsGetOptions) (*azcore.Request, error) {
	urlPath := "/bigDataPools/{bigDataPoolName}"
	if bigDataPoolName == "" {
		return nil, errors.New("parameter bigDataPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bigDataPoolName}", url.PathEscape(bigDataPoolName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *bigDataPoolsClient) getHandleResponse(resp *azcore.Response) (BigDataPoolsGetResponse, error) {
	result := BigDataPoolsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.BigDataPoolResourceInfo); err != nil {
		return BigDataPoolsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *bigDataPoolsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorContract{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - List Big Data Pools
// If the operation fails it returns the *ErrorContract error type.
func (client *bigDataPoolsClient) List(ctx context.Context, options *BigDataPoolsListOptions) (BigDataPoolsListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return BigDataPoolsListResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return BigDataPoolsListResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BigDataPoolsListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *bigDataPoolsClient) listCreateRequest(ctx context.Context, options *BigDataPoolsListOptions) (*azcore.Request, error) {
	urlPath := "/bigDataPools"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *bigDataPoolsClient) listHandleResponse(resp *azcore.Response) (BigDataPoolsListResponse, error) {
	result := BigDataPoolsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.BigDataPoolResourceInfoListResult); err != nil {
		return BigDataPoolsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *bigDataPoolsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := ErrorContract{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
