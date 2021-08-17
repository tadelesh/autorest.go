//go:build go1.16
// +build go1.16

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

// VPNConnectionsClient contains the methods for the VPNConnections group.
// Don't use this type directly, use NewVPNConnectionsClient() instead.
type VPNConnectionsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNConnectionsClient creates a new instance of VPNConnectionsClient with the specified values.
func NewVPNConnectionsClient(con *armcore.Connection, subscriptionID string) *VPNConnectionsClient {
	return &VPNConnectionsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
// If the operation fails it returns the *CloudError error type.
func (client *VPNConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsBeginCreateOrUpdateOptions) (VPNConnectionsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
	if err != nil {
		return VPNConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result := VPNConnectionsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VPNConnectionsClient.CreateOrUpdate", "azure-async-operation", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return VPNConnectionsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VPNConnectionsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
// If the operation fails it returns the *CloudError error type.
func (client *VPNConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
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
func (client *VPNConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vpnConnectionParameters)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *VPNConnectionsClient) createOrUpdateHandleError(resp *azcore.Response) error {
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

// BeginDelete - Deletes a vpn connection.
// If the operation fails it returns the *CloudError error type.
func (client *VPNConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsBeginDeleteOptions) (VPNConnectionsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return VPNConnectionsDeletePollerResponse{}, err
	}
	result := VPNConnectionsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VPNConnectionsClient.Delete", "location", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return VPNConnectionsDeletePollerResponse{}, err
	}
	result.Poller = &VPNConnectionsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a vpn connection.
// If the operation fails it returns the *CloudError error type.
func (client *VPNConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
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
func (client *VPNConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *VPNConnectionsClient) deleteHandleError(resp *azcore.Response) error {
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

// Get - Retrieves the details of a vpn connection.
// If the operation fails it returns the *CloudError error type.
func (client *VPNConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsGetOptions) (VPNConnectionsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return VPNConnectionsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return VPNConnectionsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VPNConnectionsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *VPNConnectionsClient) getHandleResponse(resp *azcore.Response) (VPNConnectionsGetResponse, error) {
	result := VPNConnectionsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.VPNConnection); err != nil {
		return VPNConnectionsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *VPNConnectionsClient) getHandleError(resp *azcore.Response) error {
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

// ListByVPNGateway - Retrieves all vpn connections for a particular virtual wan vpn gateway.
// If the operation fails it returns the *CloudError error type.
func (client *VPNConnectionsClient) ListByVPNGateway(resourceGroupName string, gatewayName string, options *VPNConnectionsListByVPNGatewayOptions) *VPNConnectionsListByVPNGatewayPager {
	return &VPNConnectionsListByVPNGatewayPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByVPNGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
		},
		advancer: func(ctx context.Context, resp VPNConnectionsListByVPNGatewayResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ListVPNConnectionsResult.NextLink)
		},
	}
}

// listByVPNGatewayCreateRequest creates the ListByVPNGateway request.
func (client *VPNConnectionsClient) listByVPNGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNConnectionsListByVPNGatewayOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// listByVPNGatewayHandleResponse handles the ListByVPNGateway response.
func (client *VPNConnectionsClient) listByVPNGatewayHandleResponse(resp *azcore.Response) (VPNConnectionsListByVPNGatewayResponse, error) {
	result := VPNConnectionsListByVPNGatewayResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ListVPNConnectionsResult); err != nil {
		return VPNConnectionsListByVPNGatewayResponse{}, err
	}
	return result, nil
}

// listByVPNGatewayHandleError handles the ListByVPNGateway error response.
func (client *VPNConnectionsClient) listByVPNGatewayHandleError(resp *azcore.Response) error {
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
