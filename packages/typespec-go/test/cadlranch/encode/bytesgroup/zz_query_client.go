//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package bytesgroup

import (
	"context"
	"encoding/base64"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// QueryClient contains the methods for the Encode.Bytes group.
// Don't use this type directly, use a constructor function instead.
type QueryClient struct {
	internal *azcore.Client
}

// - options - QueryClientBase64Options contains the optional parameters for the QueryClient.Base64 method.
func (client *QueryClient) Base64(ctx context.Context, value []byte, options *QueryClientBase64Options) (QueryClientBase64Response, error) {
	var err error
	req, err := client.base64CreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientBase64Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientBase64Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientBase64Response{}, err
	}
	return QueryClientBase64Response{}, nil
}

// base64CreateRequest creates the Base64 request.
func (client *QueryClient) base64CreateRequest(ctx context.Context, value []byte, options *QueryClientBase64Options) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/base64"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", base64.StdEncoding.EncodeToString(value))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - options - QueryClientBase64URLOptions contains the optional parameters for the QueryClient.Base64URL method.
func (client *QueryClient) Base64URL(ctx context.Context, value []byte, options *QueryClientBase64URLOptions) (QueryClientBase64URLResponse, error) {
	var err error
	req, err := client.base64URLCreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientBase64URLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientBase64URLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientBase64URLResponse{}, err
	}
	return QueryClientBase64URLResponse{}, nil
}

// base64URLCreateRequest creates the Base64URL request.
func (client *QueryClient) base64URLCreateRequest(ctx context.Context, value []byte, options *QueryClientBase64URLOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/base64url"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", base64.RawURLEncoding.EncodeToString(value))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - options - QueryClientBase64URLArrayOptions contains the optional parameters for the QueryClient.Base64URLArray method.
func (client *QueryClient) Base64URLArray(ctx context.Context, value [][]byte, options *QueryClientBase64URLArrayOptions) (QueryClientBase64URLArrayResponse, error) {
	var err error
	req, err := client.base64URLArrayCreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientBase64URLArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientBase64URLArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientBase64URLArrayResponse{}, err
	}
	return QueryClientBase64URLArrayResponse{}, nil
}

// base64URLArrayCreateRequest creates the Base64URLArray request.
func (client *QueryClient) base64URLArrayCreateRequest(ctx context.Context, value [][]byte, options *QueryClientBase64URLArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/base64url-array"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", strings.Join(func() []string {
		encodedValue := make([]string, len(value))
		for i := 0; i < len(value); i++ {
			encodedValue[i] = base64.RawURLEncoding.EncodeToString(value[i])
		}
		return encodedValue
	}(), ","))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// - options - QueryClientDefaultOptions contains the optional parameters for the QueryClient.Default method.
func (client *QueryClient) Default(ctx context.Context, value []byte, options *QueryClientDefaultOptions) (QueryClientDefaultResponse, error) {
	var err error
	req, err := client.defaultCreateRequest(ctx, value, options)
	if err != nil {
		return QueryClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return QueryClientDefaultResponse{}, err
	}
	return QueryClientDefaultResponse{}, nil
}

// defaultCreateRequest creates the Default request.
func (client *QueryClient) defaultCreateRequest(ctx context.Context, value []byte, options *QueryClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/bytes/query/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", base64.StdEncoding.EncodeToString(value))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
