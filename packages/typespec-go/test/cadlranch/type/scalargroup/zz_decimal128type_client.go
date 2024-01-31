//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package scalargroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
)

// Decimal128TypeClient contains the methods for the Type.Scalar group.
// Don't use this type directly, use a constructor function instead.
type Decimal128TypeClient struct {
	internal *azcore.Client
}

//   - options - Decimal128TypeClientRequestBodyOptions contains the optional parameters for the Decimal128TypeClient.RequestBody
//     method.
func (client *Decimal128TypeClient) RequestBody(ctx context.Context, body float64, options *Decimal128TypeClientRequestBodyOptions) (Decimal128TypeClientRequestBodyResponse, error) {
	var err error
	req, err := client.requestBodyCreateRequest(ctx, body, options)
	if err != nil {
		return Decimal128TypeClientRequestBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Decimal128TypeClientRequestBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Decimal128TypeClientRequestBodyResponse{}, err
	}
	return Decimal128TypeClientRequestBodyResponse{}, nil
}

// requestBodyCreateRequest creates the RequestBody request.
func (client *Decimal128TypeClient) requestBodyCreateRequest(ctx context.Context, body float64, options *Decimal128TypeClientRequestBodyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal128/resquest_body"
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

//   - options - Decimal128TypeClientRequestParameterOptions contains the optional parameters for the Decimal128TypeClient.RequestParameter
//     method.
func (client *Decimal128TypeClient) RequestParameter(ctx context.Context, value float64, options *Decimal128TypeClientRequestParameterOptions) (Decimal128TypeClientRequestParameterResponse, error) {
	var err error
	req, err := client.requestParameterCreateRequest(ctx, value, options)
	if err != nil {
		return Decimal128TypeClientRequestParameterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Decimal128TypeClientRequestParameterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return Decimal128TypeClientRequestParameterResponse{}, err
	}
	return Decimal128TypeClientRequestParameterResponse{}, nil
}

// requestParameterCreateRequest creates the RequestParameter request.
func (client *Decimal128TypeClient) requestParameterCreateRequest(ctx context.Context, value float64, options *Decimal128TypeClientRequestParameterOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal128/request_parameter"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("value", strconv.FormatFloat(value, 'f', -1, 64))
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

//   - options - Decimal128TypeClientResponseBodyOptions contains the optional parameters for the Decimal128TypeClient.ResponseBody
//     method.
func (client *Decimal128TypeClient) ResponseBody(ctx context.Context, options *Decimal128TypeClientResponseBodyOptions) (Decimal128TypeClientResponseBodyResponse, error) {
	var err error
	req, err := client.responseBodyCreateRequest(ctx, options)
	if err != nil {
		return Decimal128TypeClientResponseBodyResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return Decimal128TypeClientResponseBodyResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return Decimal128TypeClientResponseBodyResponse{}, err
	}
	resp, err := client.responseBodyHandleResponse(httpResp)
	return resp, err
}

// responseBodyCreateRequest creates the ResponseBody request.
func (client *Decimal128TypeClient) responseBodyCreateRequest(ctx context.Context, options *Decimal128TypeClientResponseBodyOptions) (*policy.Request, error) {
	urlPath := "/type/scalar/decimal128/response_body"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// responseBodyHandleResponse handles the ResponseBody response.
func (client *Decimal128TypeClient) responseBodyHandleResponse(resp *http.Response) (Decimal128TypeClientResponseBodyResponse, error) {
	result := Decimal128TypeClientResponseBodyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return Decimal128TypeClientResponseBodyResponse{}, err
	}
	return result, nil
}
