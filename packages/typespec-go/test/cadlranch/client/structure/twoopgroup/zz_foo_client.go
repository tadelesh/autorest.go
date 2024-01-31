//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package twoopgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FooClient contains the methods for the Client.Structure.Service group.
// Don't use this type directly, use a constructor function instead.
type FooClient struct {
	internal *azcore.Client
}

// - options - FooClientFourOptions contains the optional parameters for the FooClient.Four method.
func (client *FooClient) Four(ctx context.Context, options *FooClientFourOptions) (FooClientFourResponse, error) {
	var err error
	req, err := client.fourCreateRequest(ctx, options)
	if err != nil {
		return FooClientFourResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FooClientFourResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FooClientFourResponse{}, err
	}
	return FooClientFourResponse{}, nil
}

// fourCreateRequest creates the Four request.
func (client *FooClient) fourCreateRequest(ctx context.Context, options *FooClientFourOptions) (*policy.Request, error) {
	urlPath := "/four"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// - options - FooClientThreeOptions contains the optional parameters for the FooClient.Three method.
func (client *FooClient) Three(ctx context.Context, options *FooClientThreeOptions) (FooClientThreeResponse, error) {
	var err error
	req, err := client.threeCreateRequest(ctx, options)
	if err != nil {
		return FooClientThreeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FooClientThreeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FooClientThreeResponse{}, err
	}
	return FooClientThreeResponse{}, nil
}

// threeCreateRequest creates the Three request.
func (client *FooClient) threeCreateRequest(ctx context.Context, options *FooClientThreeOptions) (*policy.Request, error) {
	urlPath := "/three"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
