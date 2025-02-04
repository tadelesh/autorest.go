// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package arraygroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// StringValueClient - Array of string values
// Don't use this type directly, use [ArrayClient.NewStringValueClient] instead.
type StringValueClient struct {
	internal *azcore.Client
}

// - options - StringValueClientGetOptions contains the optional parameters for the StringValueClient.Get method.
func (client *StringValueClient) Get(ctx context.Context, options *StringValueClientGetOptions) (StringValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return StringValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StringValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StringValueClient) getCreateRequest(ctx context.Context, options *StringValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/array/string"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StringValueClient) getHandleResponse(resp *http.Response) (StringValueClientGetResponse, error) {
	result := StringValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return StringValueClientGetResponse{}, err
	}
	return result, nil
}

// - options - StringValueClientPutOptions contains the optional parameters for the StringValueClient.Put method.
func (client *StringValueClient) Put(ctx context.Context, body []string, options *StringValueClientPutOptions) (StringValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return StringValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StringValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return StringValueClientPutResponse{}, err
	}
	return StringValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *StringValueClient) putCreateRequest(ctx context.Context, body []string, options *StringValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/array/string"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
