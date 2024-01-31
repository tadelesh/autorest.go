//go:build go1.18
// +build go1.18

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

// NullableFloatValueClient contains the methods for the Type.Array group.
// Don't use this type directly, use a constructor function instead.
type NullableFloatValueClient struct {
	internal *azcore.Client
}

// - options - NullableFloatValueClientGetOptions contains the optional parameters for the NullableFloatValueClient.Get method.
func (client *NullableFloatValueClient) Get(ctx context.Context, options *NullableFloatValueClientGetOptions) (NullableFloatValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return NullableFloatValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableFloatValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NullableFloatValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NullableFloatValueClient) getCreateRequest(ctx context.Context, options *NullableFloatValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/array/nullable-float"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NullableFloatValueClient) getHandleResponse(resp *http.Response) (NullableFloatValueClientGetResponse, error) {
	result := NullableFloatValueClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return NullableFloatValueClientGetResponse{}, err
	}
	return result, nil
}

// - options - NullableFloatValueClientPutOptions contains the optional parameters for the NullableFloatValueClient.Put method.
func (client *NullableFloatValueClient) Put(ctx context.Context, body []*float32, options *NullableFloatValueClientPutOptions) (NullableFloatValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return NullableFloatValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NullableFloatValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return NullableFloatValueClientPutResponse{}, err
	}
	return NullableFloatValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *NullableFloatValueClient) putCreateRequest(ctx context.Context, body []*float32, options *NullableFloatValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/array/nullable-float"
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
