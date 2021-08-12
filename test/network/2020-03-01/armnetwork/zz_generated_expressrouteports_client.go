// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// ExpressRoutePortsClient contains the methods for the ExpressRoutePorts group.
// Don't use this type directly, use NewExpressRoutePortsClient() instead.
type ExpressRoutePortsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewExpressRoutePortsClient creates a new instance of ExpressRoutePortsClient with the specified values.
func NewExpressRoutePortsClient(con *armcore.Connection, subscriptionID string) *ExpressRoutePortsClient {
	return &ExpressRoutePortsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates the specified ExpressRoutePort resource.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsBeginCreateOrUpdateOptions) (ExpressRoutePortsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, expressRoutePortName, parameters, options)
	if err != nil {
		return ExpressRoutePortsCreateOrUpdatePollerResponse{}, err
	}
	result := ExpressRoutePortsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ExpressRoutePortsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return ExpressRoutePortsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ExpressRoutePortsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates the specified ExpressRoutePort resource.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) createOrUpdate(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, expressRoutePortName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ExpressRoutePortsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters ExpressRoutePort, options *ExpressRoutePortsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *ExpressRoutePortsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDelete - Deletes the specified ExpressRoutePort resource.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) BeginDelete(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsBeginDeleteOptions) (ExpressRoutePortsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, expressRoutePortName, options)
	if err != nil {
		return ExpressRoutePortsDeletePollerResponse{}, err
	}
	result := ExpressRoutePortsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("ExpressRoutePortsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return ExpressRoutePortsDeletePollerResponse{}, err
	}
	result.Poller = &ExpressRoutePortsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified ExpressRoutePort resource.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) deleteOperation(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ExpressRoutePortsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *ExpressRoutePortsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// Get - Retrieves the requested ExpressRoutePort resource.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) Get(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsGetOptions) (ExpressRoutePortsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, expressRoutePortName, options)
	if err != nil {
		return ExpressRoutePortsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRoutePortsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRoutePortsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRoutePortsClient) getCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, options *ExpressRoutePortsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRoutePortsClient) getHandleResponse(resp *azcore.Response) (ExpressRoutePortsGetResponse, error) {
	result := ExpressRoutePortsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRoutePort); err != nil {
		return ExpressRoutePortsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *ExpressRoutePortsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// List - List all the ExpressRoutePort resources in the specified subscription.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) List(options *ExpressRoutePortsListOptions) *ExpressRoutePortsListPager {
	return &ExpressRoutePortsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp ExpressRoutePortsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *ExpressRoutePortsClient) listCreateRequest(ctx context.Context, options *ExpressRoutePortsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePorts"
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
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRoutePortsClient) listHandleResponse(resp *azcore.Response) (ExpressRoutePortsListResponse, error) {
	result := ExpressRoutePortsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRoutePortListResult); err != nil {
		return ExpressRoutePortsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *ExpressRoutePortsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// ListByResourceGroup - List all the ExpressRoutePort resources in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) ListByResourceGroup(resourceGroupName string, options *ExpressRoutePortsListByResourceGroupOptions) *ExpressRoutePortsListByResourceGroupPager {
	return &ExpressRoutePortsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp ExpressRoutePortsListByResourceGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ExpressRoutePortListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ExpressRoutePortsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ExpressRoutePortsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ExpressRoutePortsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (ExpressRoutePortsListByResourceGroupResponse, error) {
	result := ExpressRoutePortsListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRoutePortListResult); err != nil {
		return ExpressRoutePortsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *ExpressRoutePortsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// UpdateTags - Update ExpressRoutePort tags.
// If the operation fails it returns the *CloudError error type.
func (client *ExpressRoutePortsClient) UpdateTags(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters TagsObject, options *ExpressRoutePortsUpdateTagsOptions) (ExpressRoutePortsUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, expressRoutePortName, parameters, options)
	if err != nil {
		return ExpressRoutePortsUpdateTagsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return ExpressRoutePortsUpdateTagsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ExpressRoutePortsUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ExpressRoutePortsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, expressRoutePortName string, parameters TagsObject, options *ExpressRoutePortsUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ExpressRoutePorts/{expressRoutePortName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if expressRoutePortName == "" {
		return nil, errors.New("parameter expressRoutePortName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{expressRoutePortName}", url.PathEscape(expressRoutePortName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *ExpressRoutePortsClient) updateTagsHandleResponse(resp *azcore.Response) (ExpressRoutePortsUpdateTagsResponse, error) {
	result := ExpressRoutePortsUpdateTagsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ExpressRoutePort); err != nil {
		return ExpressRoutePortsUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *ExpressRoutePortsClient) updateTagsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
