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
	"time"
)

// DatetimeValueClient - Array of datetime values
// Don't use this type directly, use [ArrayClient.NewDatetimeValueClient] instead.
type DatetimeValueClient struct {
	internal *azcore.Client
}

// - options - DatetimeValueClientGetOptions contains the optional parameters for the DatetimeValueClient.Get method.
func (client *DatetimeValueClient) Get(ctx context.Context, options *DatetimeValueClientGetOptions) (DatetimeValueClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return DatetimeValueClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeValueClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeValueClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatetimeValueClient) getCreateRequest(ctx context.Context, options *DatetimeValueClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/array/datetime"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatetimeValueClient) getHandleResponse(resp *http.Response) (DatetimeValueClientGetResponse, error) {
	result := DatetimeValueClientGetResponse{}
	var aux []dateTimeRFC3339
	if err := runtime.UnmarshalAsJSON(resp, &aux); err != nil {
		return DatetimeValueClientGetResponse{}, err
	}
	cp := make([]time.Time, len(aux))
	for i := 0; i < len(aux); i++ {
		cp[i] = (time.Time)(aux[i])
	}
	result.Value = cp
	return result, nil
}

// - options - DatetimeValueClientPutOptions contains the optional parameters for the DatetimeValueClient.Put method.
func (client *DatetimeValueClient) Put(ctx context.Context, body []time.Time, options *DatetimeValueClientPutOptions) (DatetimeValueClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return DatetimeValueClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatetimeValueClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DatetimeValueClientPutResponse{}, err
	}
	return DatetimeValueClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *DatetimeValueClient) putCreateRequest(ctx context.Context, body []time.Time, options *DatetimeValueClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/array/datetime"
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
