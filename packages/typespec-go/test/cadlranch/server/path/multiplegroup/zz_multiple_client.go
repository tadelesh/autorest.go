// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package multiplegroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MultipleClient contains the methods for the Server.Path.Multiple namespace.
// Don't use this type directly, use a constructor function instead.
type MultipleClient struct {
	internal *azcore.Client
	endpoint string
}

//   - options - MultipleClientNoOperationParamsOptions contains the optional parameters for the MultipleClient.NoOperationParams
//     method.
func (client *MultipleClient) NoOperationParams(ctx context.Context, options *MultipleClientNoOperationParamsOptions) (MultipleClientNoOperationParamsResponse, error) {
	var err error
	req, err := client.noOperationParamsCreateRequest(ctx, options)
	if err != nil {
		return MultipleClientNoOperationParamsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleClientNoOperationParamsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultipleClientNoOperationParamsResponse{}, err
	}
	return MultipleClientNoOperationParamsResponse{}, nil
}

// noOperationParamsCreateRequest creates the NoOperationParams request.
func (client *MultipleClient) noOperationParamsCreateRequest(ctx context.Context, options *MultipleClientNoOperationParamsOptions) (*policy.Request, error) {
	host := "{endpoint}/server/path/multiple/{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	req, err := runtime.NewRequest(ctx, http.MethodGet, host)
	if err != nil {
		return nil, err
	}
	return req, nil
}

//   - options - MultipleClientWithOperationPathParamOptions contains the optional parameters for the MultipleClient.WithOperationPathParam
//     method.
func (client *MultipleClient) WithOperationPathParam(ctx context.Context, keyword string, options *MultipleClientWithOperationPathParamOptions) (MultipleClientWithOperationPathParamResponse, error) {
	var err error
	req, err := client.withOperationPathParamCreateRequest(ctx, keyword, options)
	if err != nil {
		return MultipleClientWithOperationPathParamResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MultipleClientWithOperationPathParamResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MultipleClientWithOperationPathParamResponse{}, err
	}
	return MultipleClientWithOperationPathParamResponse{}, nil
}

// withOperationPathParamCreateRequest creates the WithOperationPathParam request.
func (client *MultipleClient) withOperationPathParamCreateRequest(ctx context.Context, keyword string, options *MultipleClientWithOperationPathParamOptions) (*policy.Request, error) {
	host := "{endpoint}/server/path/multiple/{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	urlPath := "/{keyword}"
	if keyword == "" {
		return nil, errors.New("parameter keyword cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{keyword}", url.PathEscape(keyword))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
