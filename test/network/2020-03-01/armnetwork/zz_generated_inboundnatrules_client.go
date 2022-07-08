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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// InboundNatRulesClient contains the methods for the InboundNatRules group.
// Don't use this type directly, use NewInboundNatRulesClient() instead.
type InboundNatRulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewInboundNatRulesClient creates a new instance of InboundNatRulesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewInboundNatRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InboundNatRulesClient, error) {
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
	client := &InboundNatRulesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// inboundNatRuleName - The name of the inbound nat rule.
// inboundNatRuleParameters - Parameters supplied to the create or update inbound nat rule operation.
// options - InboundNatRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the InboundNatRulesClient.BeginCreateOrUpdate
// method.
func (client *InboundNatRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[InboundNatRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, inboundNatRuleParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[InboundNatRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InboundNatRulesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *InboundNatRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, inboundNatRuleParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *InboundNatRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, inboundNatRuleParameters InboundNatRule, options *InboundNatRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, inboundNatRuleParameters)
}

// BeginDelete - Deletes the specified load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// inboundNatRuleName - The name of the inbound nat rule.
// options - InboundNatRulesClientBeginDeleteOptions contains the optional parameters for the InboundNatRulesClient.BeginDelete
// method.
func (client *InboundNatRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientBeginDeleteOptions) (*runtime.Poller[InboundNatRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[InboundNatRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[InboundNatRulesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
func (client *InboundNatRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *InboundNatRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified load balancer inbound nat rule.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// inboundNatRuleName - The name of the inbound nat rule.
// options - InboundNatRulesClientGetOptions contains the optional parameters for the InboundNatRulesClient.Get method.
func (client *InboundNatRulesClient) Get(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientGetOptions) (InboundNatRulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, loadBalancerName, inboundNatRuleName, options)
	if err != nil {
		return InboundNatRulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return InboundNatRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return InboundNatRulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *InboundNatRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, inboundNatRuleName string, options *InboundNatRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules/{inboundNatRuleName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if inboundNatRuleName == "" {
		return nil, errors.New("parameter inboundNatRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{inboundNatRuleName}", url.PathEscape(inboundNatRuleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InboundNatRulesClient) getHandleResponse(resp *http.Response) (InboundNatRulesClientGetResponse, error) {
	result := InboundNatRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InboundNatRule); err != nil {
		return InboundNatRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all the inbound nat rules in a load balancer.
// Generated from API version 2020-03-01
// resourceGroupName - The name of the resource group.
// loadBalancerName - The name of the load balancer.
// options - InboundNatRulesClientListOptions contains the optional parameters for the InboundNatRulesClient.List method.
func (client *InboundNatRulesClient) NewListPager(resourceGroupName string, loadBalancerName string, options *InboundNatRulesClientListOptions) *runtime.Pager[InboundNatRulesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InboundNatRulesClientListResponse]{
		More: func(page InboundNatRulesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InboundNatRulesClientListResponse) (InboundNatRulesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InboundNatRulesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return InboundNatRulesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InboundNatRulesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *InboundNatRulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, options *InboundNatRulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/inboundNatRules"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InboundNatRulesClient) listHandleResponse(resp *http.Response) (InboundNatRulesClientListResponse, error) {
	result := InboundNatRulesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InboundNatRuleListResult); err != nil {
		return InboundNatRulesClientListResponse{}, err
	}
	return result, nil
}
