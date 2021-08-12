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

// LocalNetworkGatewaysClient contains the methods for the LocalNetworkGateways group.
// Don't use this type directly, use NewLocalNetworkGatewaysClient() instead.
type LocalNetworkGatewaysClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewLocalNetworkGatewaysClient creates a new instance of LocalNetworkGatewaysClient with the specified values.
func NewLocalNetworkGatewaysClient(con *armcore.Connection, subscriptionID string) *LocalNetworkGatewaysClient {
	return &LocalNetworkGatewaysClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a local network gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *LocalNetworkGatewaysClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway, options *LocalNetworkGatewaysBeginCreateOrUpdateOptions) (LocalNetworkGatewaysCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, localNetworkGatewayName, parameters, options)
	if err != nil {
		return LocalNetworkGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result := LocalNetworkGatewaysCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LocalNetworkGatewaysClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return LocalNetworkGatewaysCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &LocalNetworkGatewaysCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a local network gateway in the specified resource group.
// If the operation fails it returns the *CloudError error type.
func (client *LocalNetworkGatewaysClient) createOrUpdate(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway, options *LocalNetworkGatewaysBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, parameters, options)
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
func (client *LocalNetworkGatewaysClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters LocalNetworkGateway, options *LocalNetworkGatewaysBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localNetworkGatewayName == "" {
		return nil, errors.New("parameter localNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *LocalNetworkGatewaysClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes the specified local network gateway.
// If the operation fails it returns the *CloudError error type.
func (client *LocalNetworkGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysBeginDeleteOptions) (LocalNetworkGatewaysDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, localNetworkGatewayName, options)
	if err != nil {
		return LocalNetworkGatewaysDeletePollerResponse{}, err
	}
	result := LocalNetworkGatewaysDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("LocalNetworkGatewaysClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return LocalNetworkGatewaysDeletePollerResponse{}, err
	}
	result.Poller = &LocalNetworkGatewaysDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified local network gateway.
// If the operation fails it returns the *CloudError error type.
func (client *LocalNetworkGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, options)
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
func (client *LocalNetworkGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localNetworkGatewayName == "" {
		return nil, errors.New("parameter localNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *LocalNetworkGatewaysClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Gets the specified local network gateway in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *LocalNetworkGatewaysClient) Get(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysGetOptions) (LocalNetworkGatewaysGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, options)
	if err != nil {
		return LocalNetworkGatewaysGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LocalNetworkGatewaysGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LocalNetworkGatewaysGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *LocalNetworkGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, options *LocalNetworkGatewaysGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localNetworkGatewayName == "" {
		return nil, errors.New("parameter localNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
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

// getHandleResponse handles the Get response.
func (client *LocalNetworkGatewaysClient) getHandleResponse(resp *azcore.Response) (LocalNetworkGatewaysGetResponse, error) {
	result := LocalNetworkGatewaysGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LocalNetworkGateway); err != nil {
		return LocalNetworkGatewaysGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *LocalNetworkGatewaysClient) getHandleError(resp *azcore.Response) error {
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

// List - Gets all the local network gateways in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *LocalNetworkGatewaysClient) List(resourceGroupName string, options *LocalNetworkGatewaysListOptions) *LocalNetworkGatewaysListPager {
	return &LocalNetworkGatewaysListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp LocalNetworkGatewaysListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.LocalNetworkGatewayListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *LocalNetworkGatewaysClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *LocalNetworkGatewaysListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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
func (client *LocalNetworkGatewaysClient) listHandleResponse(resp *azcore.Response) (LocalNetworkGatewaysListResponse, error) {
	result := LocalNetworkGatewaysListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LocalNetworkGatewayListResult); err != nil {
		return LocalNetworkGatewaysListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *LocalNetworkGatewaysClient) listHandleError(resp *azcore.Response) error {
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

// UpdateTags - Updates a local network gateway tags.
// If the operation fails it returns the *CloudError error type.
func (client *LocalNetworkGatewaysClient) UpdateTags(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject, options *LocalNetworkGatewaysUpdateTagsOptions) (LocalNetworkGatewaysUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, localNetworkGatewayName, parameters, options)
	if err != nil {
		return LocalNetworkGatewaysUpdateTagsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return LocalNetworkGatewaysUpdateTagsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LocalNetworkGatewaysUpdateTagsResponse{}, client.updateTagsHandleError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *LocalNetworkGatewaysClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, localNetworkGatewayName string, parameters TagsObject, options *LocalNetworkGatewaysUpdateTagsOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/localNetworkGateways/{localNetworkGatewayName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if localNetworkGatewayName == "" {
		return nil, errors.New("parameter localNetworkGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{localNetworkGatewayName}", url.PathEscape(localNetworkGatewayName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *LocalNetworkGatewaysClient) updateTagsHandleResponse(resp *azcore.Response) (LocalNetworkGatewaysUpdateTagsResponse, error) {
	result := LocalNetworkGatewaysUpdateTagsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.LocalNetworkGateway); err != nil {
		return LocalNetworkGatewaysUpdateTagsResponse{}, err
	}
	return result, nil
}

// updateTagsHandleError handles the UpdateTags error response.
func (client *LocalNetworkGatewaysClient) updateTagsHandleError(resp *azcore.Response) error {
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
