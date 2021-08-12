// +build go1.13

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

// SnapshotsClient contains the methods for the Snapshots group.
// Don't use this type directly, use NewSnapshotsClient() instead.
type SnapshotsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewSnapshotsClient creates a new instance of SnapshotsClient with the specified values.
func NewSnapshotsClient(con *armcore.Connection, subscriptionID string) *SnapshotsClient {
	return &SnapshotsClient{con: con, subscriptionID: subscriptionID}
}

// BeginCreateOrUpdate - Creates or updates a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsBeginCreateOrUpdateOptions) (SnapshotsCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return SnapshotsCreateOrUpdatePollerResponse{}, err
	}
	result := SnapshotsCreateOrUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SnapshotsClient.CreateOrUpdate", "", resp, client.con.Pipeline(), client.createOrUpdateHandleError)
	if err != nil {
		return SnapshotsCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &SnapshotsCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) createOrUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SnapshotsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, snapshot Snapshot, options *SnapshotsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(snapshot)
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *SnapshotsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginDelete - Deletes a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) BeginDelete(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsBeginDeleteOptions) (SnapshotsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return SnapshotsDeletePollerResponse{}, err
	}
	result := SnapshotsDeletePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SnapshotsClient.Delete", "", resp, client.con.Pipeline(), client.deleteHandleError)
	if err != nil {
		return SnapshotsDeletePollerResponse{}, err
	}
	result.Poller = &SnapshotsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) deleteOperation(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsBeginDeleteOptions) (*azcore.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, snapshotName, options)
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
func (client *SnapshotsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsBeginDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *SnapshotsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets information about a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) Get(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsGetOptions) (SnapshotsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return SnapshotsGetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SnapshotsGetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SnapshotsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SnapshotsClient) getCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SnapshotsClient) getHandleResponse(resp *azcore.Response) (SnapshotsGetResponse, error) {
	result := SnapshotsGetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Snapshot); err != nil {
		return SnapshotsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *SnapshotsClient) getHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginGrantAccess - Grants access to a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) BeginGrantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsBeginGrantAccessOptions) (SnapshotsGrantAccessPollerResponse, error) {
	resp, err := client.grantAccess(ctx, resourceGroupName, snapshotName, grantAccessData, options)
	if err != nil {
		return SnapshotsGrantAccessPollerResponse{}, err
	}
	result := SnapshotsGrantAccessPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SnapshotsClient.GrantAccess", "location", resp, client.con.Pipeline(), client.grantAccessHandleError)
	if err != nil {
		return SnapshotsGrantAccessPollerResponse{}, err
	}
	result.Poller = &SnapshotsGrantAccessPoller{
		pt: pt,
	}
	return result, nil
}

// GrantAccess - Grants access to a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) grantAccess(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsBeginGrantAccessOptions) (*azcore.Response, error) {
	req, err := client.grantAccessCreateRequest(ctx, resourceGroupName, snapshotName, grantAccessData, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.grantAccessHandleError(resp)
	}
	return resp, nil
}

// grantAccessCreateRequest creates the GrantAccess request.
func (client *SnapshotsClient) grantAccessCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, grantAccessData GrantAccessData, options *SnapshotsBeginGrantAccessOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/beginGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(grantAccessData)
}

// grantAccessHandleError handles the GrantAccess error response.
func (client *SnapshotsClient) grantAccessHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - Lists snapshots under a subscription.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) List(options *SnapshotsListOptions) *SnapshotsListPager {
	return &SnapshotsListPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp SnapshotsListResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SnapshotList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SnapshotsClient) listCreateRequest(ctx context.Context, options *SnapshotsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/snapshots"
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
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SnapshotsClient) listHandleResponse(resp *azcore.Response) (SnapshotsListResponse, error) {
	result := SnapshotsListResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SnapshotList); err != nil {
		return SnapshotsListResponse{}, err
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *SnapshotsClient) listHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - Lists snapshots under a resource group.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) ListByResourceGroup(resourceGroupName string, options *SnapshotsListByResourceGroupOptions) *SnapshotsListByResourceGroupPager {
	return &SnapshotsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp SnapshotsListByResourceGroupResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.SnapshotList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SnapshotsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SnapshotsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots"
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
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SnapshotsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (SnapshotsListByResourceGroupResponse, error) {
	result := SnapshotsListByResourceGroupResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.SnapshotList); err != nil {
		return SnapshotsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *SnapshotsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginRevokeAccess - Revokes access to a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) BeginRevokeAccess(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsBeginRevokeAccessOptions) (SnapshotsRevokeAccessPollerResponse, error) {
	resp, err := client.revokeAccess(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return SnapshotsRevokeAccessPollerResponse{}, err
	}
	result := SnapshotsRevokeAccessPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SnapshotsClient.RevokeAccess", "location", resp, client.con.Pipeline(), client.revokeAccessHandleError)
	if err != nil {
		return SnapshotsRevokeAccessPollerResponse{}, err
	}
	result.Poller = &SnapshotsRevokeAccessPoller{
		pt: pt,
	}
	return result, nil
}

// RevokeAccess - Revokes access to a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) revokeAccess(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsBeginRevokeAccessOptions) (*azcore.Response, error) {
	req, err := client.revokeAccessCreateRequest(ctx, resourceGroupName, snapshotName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.revokeAccessHandleError(resp)
	}
	return resp, nil
}

// revokeAccessCreateRequest creates the RevokeAccess request.
func (client *SnapshotsClient) revokeAccessCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, options *SnapshotsBeginRevokeAccessOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}/endGetAccess"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	return req, nil
}

// revokeAccessHandleError handles the RevokeAccess error response.
func (client *SnapshotsClient) revokeAccessHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginUpdate - Updates (patches) a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) BeginUpdate(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsBeginUpdateOptions) (SnapshotsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return SnapshotsUpdatePollerResponse{}, err
	}
	result := SnapshotsUpdatePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("SnapshotsClient.Update", "", resp, client.con.Pipeline(), client.updateHandleError)
	if err != nil {
		return SnapshotsUpdatePollerResponse{}, err
	}
	result.Poller = &SnapshotsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates (patches) a snapshot.
// If the operation fails it returns a generic error.
func (client *SnapshotsClient) update(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsBeginUpdateOptions) (*azcore.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, snapshotName, snapshot, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *SnapshotsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, snapshotName string, snapshot SnapshotUpdate, options *SnapshotsBeginUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/snapshots/{snapshotName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if snapshotName == "" {
		return nil, errors.New("parameter snapshotName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{snapshotName}", url.PathEscape(snapshotName))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-11-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(snapshot)
}

// updateHandleError handles the Update error response.
func (client *SnapshotsClient) updateHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
