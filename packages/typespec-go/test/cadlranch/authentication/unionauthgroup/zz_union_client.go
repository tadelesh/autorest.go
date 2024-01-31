//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package unionauthgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// UnionClient contains the methods for the Authentication.Union group.
// Don't use this type directly, use a constructor function instead.
type UnionClient struct {
	internal *azcore.Client
}

// ValidKey - Check whether client is authenticated
//   - options - UnionClientValidKeyOptions contains the optional parameters for the UnionClient.ValidKey method.
func (client *UnionClient) ValidKey(ctx context.Context, options *UnionClientValidKeyOptions) (UnionClientValidKeyResponse, error) {
	var err error
	req, err := client.validKeyCreateRequest(ctx, options)
	if err != nil {
		return UnionClientValidKeyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionClientValidKeyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UnionClientValidKeyResponse{}, err
	}
	return UnionClientValidKeyResponse{}, nil
}

// validKeyCreateRequest creates the ValidKey request.
func (client *UnionClient) validKeyCreateRequest(ctx context.Context, options *UnionClientValidKeyOptions) (*policy.Request, error) {
	urlPath := "/authentication/union/validkey"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ValidToken - Check whether client is authenticated
//   - options - UnionClientValidTokenOptions contains the optional parameters for the UnionClient.ValidToken method.
func (client *UnionClient) ValidToken(ctx context.Context, options *UnionClientValidTokenOptions) (UnionClientValidTokenResponse, error) {
	var err error
	req, err := client.validTokenCreateRequest(ctx, options)
	if err != nil {
		return UnionClientValidTokenResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return UnionClientValidTokenResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return UnionClientValidTokenResponse{}, err
	}
	return UnionClientValidTokenResponse{}, nil
}

// validTokenCreateRequest creates the ValidToken request.
func (client *UnionClient) validTokenCreateRequest(ctx context.Context, options *UnionClientValidTokenOptions) (*policy.Request, error) {
	urlPath := "/authentication/union/validtoken"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
