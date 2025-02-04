// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package addlpropsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// IsUnknownDerivedClient contains the methods for the Type.Property.AdditionalProperties namespace.
// Don't use this type directly, use [AdditionalPropertiesClient.NewIsUnknownDerivedClient] instead.
type IsUnknownDerivedClient struct {
	internal *azcore.Client
}

// Get - Get call
//   - options - IsUnknownDerivedClientGetOptions contains the optional parameters for the IsUnknownDerivedClient.Get method.
func (client *IsUnknownDerivedClient) Get(ctx context.Context, options *IsUnknownDerivedClientGetOptions) (IsUnknownDerivedClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return IsUnknownDerivedClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsUnknownDerivedClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IsUnknownDerivedClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IsUnknownDerivedClient) getCreateRequest(ctx context.Context, options *IsUnknownDerivedClientGetOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isRecordUnknownDerived"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IsUnknownDerivedClient) getHandleResponse(resp *http.Response) (IsUnknownDerivedClientGetResponse, error) {
	result := IsUnknownDerivedClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IsUnknownAdditionalPropertiesDerived); err != nil {
		return IsUnknownDerivedClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put operation
//   - body - body
//   - options - IsUnknownDerivedClientPutOptions contains the optional parameters for the IsUnknownDerivedClient.Put method.
func (client *IsUnknownDerivedClient) Put(ctx context.Context, body IsUnknownAdditionalPropertiesDerived, options *IsUnknownDerivedClientPutOptions) (IsUnknownDerivedClientPutResponse, error) {
	var err error
	req, err := client.putCreateRequest(ctx, body, options)
	if err != nil {
		return IsUnknownDerivedClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IsUnknownDerivedClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return IsUnknownDerivedClientPutResponse{}, err
	}
	return IsUnknownDerivedClientPutResponse{}, nil
}

// putCreateRequest creates the Put request.
func (client *IsUnknownDerivedClient) putCreateRequest(ctx context.Context, body IsUnknownAdditionalPropertiesDerived, options *IsUnknownDerivedClientPutOptions) (*policy.Request, error) {
	urlPath := "/type/property/additionalProperties/isRecordUnknownDerived"
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
