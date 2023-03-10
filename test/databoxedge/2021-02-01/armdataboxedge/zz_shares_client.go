//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdataboxedge

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SharesClient contains the methods for the Shares group.
// Don't use this type directly, use NewSharesClient() instead.
type SharesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSharesClient creates a new instance of SharesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSharesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharesClient, error) {
	cl, err := arm.NewClient(moduleName+".SharesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SharesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new share or updates an existing share on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The share name.
//   - resourceGroupName - The resource group name.
//   - share - The share properties.
//   - options - SharesClientBeginCreateOrUpdateOptions contains the optional parameters for the SharesClient.BeginCreateOrUpdate
//     method.
func (client *SharesClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, share Share, options *SharesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SharesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, name, resourceGroupName, share, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SharesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SharesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new share or updates an existing share on the device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *SharesClient) createOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, share Share, options *SharesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, name, resourceGroupName, share, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SharesClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, share Share, options *SharesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, share)
}

// BeginDelete - Deletes the share on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The share name.
//   - resourceGroupName - The resource group name.
//   - options - SharesClientBeginDeleteOptions contains the optional parameters for the SharesClient.BeginDelete method.
func (client *SharesClient) BeginDelete(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientBeginDeleteOptions) (*runtime.Poller[SharesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, name, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SharesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SharesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the share on the Data Box Edge/Data Box Gateway device.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *SharesClient) deleteOperation(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SharesClient) deleteCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a share by name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The share name.
//   - resourceGroupName - The resource group name.
//   - options - SharesClientGetOptions contains the optional parameters for the SharesClient.Get method.
func (client *SharesClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientGetOptions) (SharesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return SharesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharesClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharesClient) getHandleResponse(resp *http.Response) (SharesClientGetResponse, error) {
	result := SharesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Share); err != nil {
		return SharesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataBoxEdgeDevicePager - Lists all the shares in a Data Box Edge/Data Box Gateway device.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - SharesClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the SharesClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *SharesClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *SharesClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[SharesClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharesClientListByDataBoxEdgeDeviceResponse]{
		More: func(page SharesClientListByDataBoxEdgeDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharesClientListByDataBoxEdgeDeviceResponse) (SharesClientListByDataBoxEdgeDeviceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SharesClientListByDataBoxEdgeDeviceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SharesClientListByDataBoxEdgeDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SharesClientListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *SharesClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *SharesClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *SharesClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (SharesClientListByDataBoxEdgeDeviceResponse, error) {
	result := SharesClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ShareList); err != nil {
		return SharesClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Refreshes the share metadata with the data from the cloud.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The share name.
//   - resourceGroupName - The resource group name.
//   - options - SharesClientBeginRefreshOptions contains the optional parameters for the SharesClient.BeginRefresh method.
func (client *SharesClient) BeginRefresh(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientBeginRefreshOptions) (*runtime.Poller[SharesClientRefreshResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refresh(ctx, deviceName, name, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[SharesClientRefreshResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[SharesClientRefreshResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Refresh - Refreshes the share metadata with the data from the cloud.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *SharesClient) refresh(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientBeginRefreshOptions) (*http.Response, error) {
	req, err := client.refreshCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// refreshCreateRequest creates the Refresh request.
func (client *SharesClient) refreshCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *SharesClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/shares/{name}/refresh"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
