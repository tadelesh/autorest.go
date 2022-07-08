//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ProximityPlacementGroupsClient contains the methods for the ProximityPlacementGroups group.
// Don't use this type directly, use NewProximityPlacementGroupsClient() instead.
type ProximityPlacementGroupsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewProximityPlacementGroupsClient creates a new instance of ProximityPlacementGroupsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewProximityPlacementGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProximityPlacementGroupsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ProximityPlacementGroupsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a proximity placement group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-12-01
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// parameters - Parameters supplied to the Create Proximity Placement Group operation.
// options - ProximityPlacementGroupsClientCreateOrUpdateOptions contains the optional parameters for the ProximityPlacementGroupsClient.CreateOrUpdate
// method.
func (client *ProximityPlacementGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup, options *ProximityPlacementGroupsClientCreateOrUpdateOptions) (ProximityPlacementGroupsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, parameters, options)
	if err != nil {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProximityPlacementGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroup, options *ProximityPlacementGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProximityPlacementGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientCreateOrUpdateResponse, error) {
	result := ProximityPlacementGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroup); err != nil {
		return ProximityPlacementGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a proximity placement group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-12-01
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// options - ProximityPlacementGroupsClientDeleteOptions contains the optional parameters for the ProximityPlacementGroupsClient.Delete
// method.
func (client *ProximityPlacementGroupsClient) Delete(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientDeleteOptions) (ProximityPlacementGroupsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, options)
	if err != nil {
		return ProximityPlacementGroupsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProximityPlacementGroupsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProximityPlacementGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProximityPlacementGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Retrieves information about a proximity placement group .
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-12-01
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// options - ProximityPlacementGroupsClientGetOptions contains the optional parameters for the ProximityPlacementGroupsClient.Get
// method.
func (client *ProximityPlacementGroupsClient) Get(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientGetOptions) (ProximityPlacementGroupsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, options)
	if err != nil {
		return ProximityPlacementGroupsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProximityPlacementGroupsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProximityPlacementGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, options *ProximityPlacementGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.IncludeColocationStatus != nil {
		reqQP.Set("includeColocationStatus", *options.IncludeColocationStatus)
	}
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProximityPlacementGroupsClient) getHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientGetResponse, error) {
	result := ProximityPlacementGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroup); err != nil {
		return ProximityPlacementGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists all proximity placement groups in a resource group.
// Generated from API version 2019-12-01
// resourceGroupName - The name of the resource group.
// options - ProximityPlacementGroupsClientListByResourceGroupOptions contains the optional parameters for the ProximityPlacementGroupsClient.ListByResourceGroup
// method.
func (client *ProximityPlacementGroupsClient) NewListByResourceGroupPager(resourceGroupName string, options *ProximityPlacementGroupsClientListByResourceGroupOptions) *runtime.Pager[ProximityPlacementGroupsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProximityPlacementGroupsClientListByResourceGroupResponse]{
		More: func(page ProximityPlacementGroupsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProximityPlacementGroupsClientListByResourceGroupResponse) (ProximityPlacementGroupsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProximityPlacementGroupsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProximityPlacementGroupsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProximityPlacementGroupsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ProximityPlacementGroupsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ProximityPlacementGroupsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ProximityPlacementGroupsClient) listByResourceGroupHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientListByResourceGroupResponse, error) {
	result := ProximityPlacementGroupsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroupListResult); err != nil {
		return ProximityPlacementGroupsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists all proximity placement groups in a subscription.
// Generated from API version 2019-12-01
// options - ProximityPlacementGroupsClientListBySubscriptionOptions contains the optional parameters for the ProximityPlacementGroupsClient.ListBySubscription
// method.
func (client *ProximityPlacementGroupsClient) NewListBySubscriptionPager(options *ProximityPlacementGroupsClientListBySubscriptionOptions) *runtime.Pager[ProximityPlacementGroupsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProximityPlacementGroupsClientListBySubscriptionResponse]{
		More: func(page ProximityPlacementGroupsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProximityPlacementGroupsClientListBySubscriptionResponse) (ProximityPlacementGroupsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProximityPlacementGroupsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ProximityPlacementGroupsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProximityPlacementGroupsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ProximityPlacementGroupsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ProximityPlacementGroupsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/proximityPlacementGroups"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ProximityPlacementGroupsClient) listBySubscriptionHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientListBySubscriptionResponse, error) {
	result := ProximityPlacementGroupsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroupListResult); err != nil {
		return ProximityPlacementGroupsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update a proximity placement group.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-12-01
// resourceGroupName - The name of the resource group.
// proximityPlacementGroupName - The name of the proximity placement group.
// parameters - Parameters supplied to the Update Proximity Placement Group operation.
// options - ProximityPlacementGroupsClientUpdateOptions contains the optional parameters for the ProximityPlacementGroupsClient.Update
// method.
func (client *ProximityPlacementGroupsClient) Update(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate, options *ProximityPlacementGroupsClientUpdateOptions) (ProximityPlacementGroupsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, proximityPlacementGroupName, parameters, options)
	if err != nil {
		return ProximityPlacementGroupsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ProximityPlacementGroupsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProximityPlacementGroupsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ProximityPlacementGroupsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, proximityPlacementGroupName string, parameters ProximityPlacementGroupUpdate, options *ProximityPlacementGroupsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/proximityPlacementGroups/{proximityPlacementGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if proximityPlacementGroupName == "" {
		return nil, errors.New("parameter proximityPlacementGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{proximityPlacementGroupName}", url.PathEscape(proximityPlacementGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ProximityPlacementGroupsClient) updateHandleResponse(resp *http.Response) (ProximityPlacementGroupsClientUpdateResponse, error) {
	result := ProximityPlacementGroupsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProximityPlacementGroup); err != nil {
		return ProximityPlacementGroupsClientUpdateResponse{}, err
	}
	return result, nil
}
