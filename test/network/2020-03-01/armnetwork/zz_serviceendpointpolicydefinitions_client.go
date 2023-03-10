//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// ServiceEndpointPolicyDefinitionsClient contains the methods for the ServiceEndpointPolicyDefinitions group.
// Don't use this type directly, use NewServiceEndpointPolicyDefinitionsClient() instead.
type ServiceEndpointPolicyDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServiceEndpointPolicyDefinitionsClient creates a new instance of ServiceEndpointPolicyDefinitionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServiceEndpointPolicyDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceEndpointPolicyDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ServiceEndpointPolicyDefinitionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceEndpointPolicyDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a service endpoint policy definition in the specified service endpoint policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the service endpoint policy.
//   - serviceEndpointPolicyDefinitionName - The name of the service endpoint policy definition name.
//   - serviceEndpointPolicyDefinitions - Parameters supplied to the create or update service endpoint policy operation.
//   - options - ServiceEndpointPolicyDefinitionsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServiceEndpointPolicyDefinitionsClient.BeginCreateOrUpdate
//     method.
func (client *ServiceEndpointPolicyDefinitionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition, options *ServiceEndpointPolicyDefinitionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServiceEndpointPolicyDefinitionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceEndpointPolicyDefinitionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServiceEndpointPolicyDefinitionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a service endpoint policy definition in the specified service endpoint policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *ServiceEndpointPolicyDefinitionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition, options *ServiceEndpointPolicyDefinitionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, serviceEndpointPolicyDefinitions, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServiceEndpointPolicyDefinitionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, serviceEndpointPolicyDefinitions ServiceEndpointPolicyDefinition, options *ServiceEndpointPolicyDefinitionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if serviceEndpointPolicyDefinitionName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, serviceEndpointPolicyDefinitions)
}

// BeginDelete - Deletes the specified ServiceEndpoint policy definitions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the Service Endpoint Policy.
//   - serviceEndpointPolicyDefinitionName - The name of the service endpoint policy definition.
//   - options - ServiceEndpointPolicyDefinitionsClientBeginDeleteOptions contains the optional parameters for the ServiceEndpointPolicyDefinitionsClient.BeginDelete
//     method.
func (client *ServiceEndpointPolicyDefinitionsClient) BeginDelete(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsClientBeginDeleteOptions) (*runtime.Poller[ServiceEndpointPolicyDefinitionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServiceEndpointPolicyDefinitionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ServiceEndpointPolicyDefinitionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified ServiceEndpoint policy definitions.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
func (client *ServiceEndpointPolicyDefinitionsClient) deleteOperation(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, options)
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
func (client *ServiceEndpointPolicyDefinitionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if serviceEndpointPolicyDefinitionName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified service endpoint policy definitions from service endpoint policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the service endpoint policy name.
//   - serviceEndpointPolicyDefinitionName - The name of the service endpoint policy definition name.
//   - options - ServiceEndpointPolicyDefinitionsClientGetOptions contains the optional parameters for the ServiceEndpointPolicyDefinitionsClient.Get
//     method.
func (client *ServiceEndpointPolicyDefinitionsClient) Get(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsClientGetOptions) (ServiceEndpointPolicyDefinitionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, serviceEndpointPolicyDefinitionName, options)
	if err != nil {
		return ServiceEndpointPolicyDefinitionsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServiceEndpointPolicyDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceEndpointPolicyDefinitionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServiceEndpointPolicyDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, serviceEndpointPolicyDefinitionName string, options *ServiceEndpointPolicyDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions/{serviceEndpointPolicyDefinitionName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if serviceEndpointPolicyDefinitionName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyDefinitionName}", url.PathEscape(serviceEndpointPolicyDefinitionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServiceEndpointPolicyDefinitionsClient) getHandleResponse(resp *http.Response) (ServiceEndpointPolicyDefinitionsClientGetResponse, error) {
	result := ServiceEndpointPolicyDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceEndpointPolicyDefinition); err != nil {
		return ServiceEndpointPolicyDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all service endpoint policy definitions in a service end point policy.
//
// Generated from API version 2020-03-01
//   - resourceGroupName - The name of the resource group.
//   - serviceEndpointPolicyName - The name of the service endpoint policy name.
//   - options - ServiceEndpointPolicyDefinitionsClientListByResourceGroupOptions contains the optional parameters for the ServiceEndpointPolicyDefinitionsClient.NewListByResourceGroupPager
//     method.
func (client *ServiceEndpointPolicyDefinitionsClient) NewListByResourceGroupPager(resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPolicyDefinitionsClientListByResourceGroupOptions) *runtime.Pager[ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse]{
		More: func(page ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse) (ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, serviceEndpointPolicyName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ServiceEndpointPolicyDefinitionsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, serviceEndpointPolicyName string, options *ServiceEndpointPolicyDefinitionsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/serviceEndpointPolicies/{serviceEndpointPolicyName}/serviceEndpointPolicyDefinitions"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceEndpointPolicyName == "" {
		return nil, errors.New("parameter serviceEndpointPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceEndpointPolicyName}", url.PathEscape(serviceEndpointPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ServiceEndpointPolicyDefinitionsClient) listByResourceGroupHandleResponse(resp *http.Response) (ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse, error) {
	result := ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceEndpointPolicyDefinitionListResult); err != nil {
		return ServiceEndpointPolicyDefinitionsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
