// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CodeVersionsClient contains the methods for the CodeVersions group.
// Don't use this type directly, use NewCodeVersionsClient() instead.
type CodeVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCodeVersionsClient creates a new instance of CodeVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCodeVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CodeVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CodeVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - body - Version entity to create or update.
//   - options - CodeVersionsClientCreateOrUpdateOptions contains the optional parameters for the CodeVersionsClient.CreateOrUpdate
//     method.
func (client *CodeVersionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body CodeVersionData, options *CodeVersionsClientCreateOrUpdateOptions) (CodeVersionsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, name, version, body, options)
	if err != nil {
		return CodeVersionsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CodeVersionsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return CodeVersionsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *CodeVersionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, body CodeVersionData, options *CodeVersionsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/codes/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *CodeVersionsClient) createOrUpdateHandleResponse(resp *http.Response) (CodeVersionsClientCreateOrUpdateResponse, error) {
	result := CodeVersionsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CodeVersionData); err != nil {
		return CodeVersionsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - options - CodeVersionsClientDeleteOptions contains the optional parameters for the CodeVersionsClient.Delete method.
func (client *CodeVersionsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *CodeVersionsClientDeleteOptions) (CodeVersionsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
	if err != nil {
		return CodeVersionsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CodeVersionsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return CodeVersionsClientDeleteResponse{}, err
	}
	return CodeVersionsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CodeVersionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *CodeVersionsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/codes/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get version.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - version - Version identifier. This is case-sensitive.
//   - options - CodeVersionsClientGetOptions contains the optional parameters for the CodeVersionsClient.Get method.
func (client *CodeVersionsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *CodeVersionsClientGetOptions) (CodeVersionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, name, version, options)
	if err != nil {
		return CodeVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CodeVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CodeVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CodeVersionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, version string, options *CodeVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/codes/{name}/versions/{version}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if version == "" {
		return nil, errors.New("parameter version cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{version}", url.PathEscape(version))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CodeVersionsClient) getHandleResponse(resp *http.Response) (CodeVersionsClientGetResponse, error) {
	result := CodeVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CodeVersionData); err != nil {
		return CodeVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List versions.
//
// Generated from API version 2022-02-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - name - Container name. This is case-sensitive.
//   - options - CodeVersionsClientListOptions contains the optional parameters for the CodeVersionsClient.NewListPager method.
func (client *CodeVersionsClient) NewListPager(resourceGroupName string, workspaceName string, name string, options *CodeVersionsClientListOptions) *runtime.Pager[CodeVersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CodeVersionsClientListResponse]{
		More: func(page CodeVersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CodeVersionsClientListResponse) (CodeVersionsClientListResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, name, options)
			}, nil)
			if err != nil {
				return CodeVersionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CodeVersionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, name string, options *CodeVersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/codes/{name}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.OrderBy != nil {
		reqQP.Set("$orderBy", *options.OrderBy)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", *options.Skip)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2022-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CodeVersionsClient) listHandleResponse(resp *http.Response) (CodeVersionsClientListResponse, error) {
	result := CodeVersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CodeVersionResourceArmPaginatedResult); err != nil {
		return CodeVersionsClientListResponse{}, err
	}
	return result, nil
}
