//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// UsageClient contains the methods for the Usage group.
// Don't use this type directly, use NewUsageClient() instead.
type UsageClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewUsageClient creates a new instance of UsageClient with the specified values.
func NewUsageClient(con *armcore.Connection, subscriptionID string) *UsageClient {
	return &UsageClient{con: con, subscriptionID: subscriptionID}
}

// List - Gets, for the specified location, the current compute resource usage information as well as the limits for compute resources under the subscription.
// If the operation fails it returns a generic error.
func (client *UsageClient) List(location string, options *UsageListOptions) *UsageListPager {
	return &UsageListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, location, options)
		},
		advancer: func(ctx context.Context, resp UsageListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListUsagesResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *UsageClient) listCreateRequest(ctx context.Context, location string, options *UsageListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/usages"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *UsageClient) listHandleResponse(resp *azcore.Response) (UsageListResponse, error) {
	result := UsageListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ListUsagesResult); err != nil {
		return UsageListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *UsageClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
