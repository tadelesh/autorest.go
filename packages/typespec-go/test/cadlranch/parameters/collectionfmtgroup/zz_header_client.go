//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package collectionfmtgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// HeaderClient contains the methods for the Parameters.CollectionFormat group.
// Don't use this type directly, use a constructor function instead.
type HeaderClient struct {
	internal *azcore.Client
}

// - colors - Possible values for colors are [blue,red,green]
// - options - HeaderClientCSVOptions contains the optional parameters for the HeaderClient.CSV method.
func (client *HeaderClient) CSV(ctx context.Context, colors []string, options *HeaderClientCSVOptions) (HeaderClientCSVResponse, error) {
	var err error
	req, err := client.csvCreateRequest(ctx, colors, options)
	if err != nil {
		return HeaderClientCSVResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientCSVResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientCSVResponse{}, err
	}
	return HeaderClientCSVResponse{}, nil
}

// csvCreateRequest creates the CSV request.
func (client *HeaderClient) csvCreateRequest(ctx context.Context, colors []string, options *HeaderClientCSVOptions) (*policy.Request, error) {
	urlPath := "/parameters/collection-format/header/csv"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["colors"] = []string{strings.Join(colors, ",")}
	return req, nil
}
