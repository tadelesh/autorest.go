// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package srvdrivengroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ResiliencyServiceDrivenClient - Test that we can grow up a service spec and service deployment into a multi-versioned service
// with full client support.
// There are three concepts that should be clarified:
// 1. Client spec version: refers to the spec that the client is generated from. 'v1' is a client generated from old.tsp and
// 'v2' is a client generated from main.tsp.
// 2. Service deployment version: refers to a deployment version of the service. 'v1' represents the initial deployment of
// the service with a single api version. 'v2' represents the new deployment of a service with multiple api versions
// 3. Api version: The initial deployment of the service only supports api version 'v1'. The new deployment of the service
// supports api versions 'v1' and 'v2'.
// We test the following configurations from this service spec:
// - A client generated from the second service spec can call the second deployment of a service with api version v1
// - A client generated from the second service spec can call the second deployment of a service with api version v2
// Don't use this type directly, use a constructor function instead.
type ResiliencyServiceDrivenClient struct {
	internal                 *azcore.Client
	endpoint                 string
	serviceDeploymentVersion string
}

// AddOperation - Added operation
//   - options - ResiliencyServiceDrivenClientAddOperationOptions contains the optional parameters for the ResiliencyServiceDrivenClient.AddOperation
//     method.
func (client *ResiliencyServiceDrivenClient) AddOperation(ctx context.Context, options *ResiliencyServiceDrivenClientAddOperationOptions) (ResiliencyServiceDrivenClientAddOperationResponse, error) {
	var err error
	req, err := client.addOperationCreateRequest(ctx, options)
	if err != nil {
		return ResiliencyServiceDrivenClientAddOperationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResiliencyServiceDrivenClientAddOperationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResiliencyServiceDrivenClientAddOperationResponse{}, err
	}
	return ResiliencyServiceDrivenClientAddOperationResponse{}, nil
}

// addOperationCreateRequest creates the AddOperation request.
func (client *ResiliencyServiceDrivenClient) addOperationCreateRequest(ctx context.Context, options *ResiliencyServiceDrivenClientAddOperationOptions) (*policy.Request, error) {
	host := "{endpoint}/resiliency/service-driven/client:v2/service:{serviceDeploymentVersion}/api-version:{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{serviceDeploymentVersion}", client.serviceDeploymentVersion)
	urlPath := "/add-operation"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// FromNone - Test that grew up from accepting no parameters to an optional input parameter
//   - options - ResiliencyServiceDrivenClientFromNoneOptions contains the optional parameters for the ResiliencyServiceDrivenClient.FromNone
//     method.
func (client *ResiliencyServiceDrivenClient) FromNone(ctx context.Context, options *ResiliencyServiceDrivenClientFromNoneOptions) (ResiliencyServiceDrivenClientFromNoneResponse, error) {
	var err error
	req, err := client.fromNoneCreateRequest(ctx, options)
	if err != nil {
		return ResiliencyServiceDrivenClientFromNoneResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResiliencyServiceDrivenClientFromNoneResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResiliencyServiceDrivenClientFromNoneResponse{}, err
	}
	return ResiliencyServiceDrivenClientFromNoneResponse{}, nil
}

// fromNoneCreateRequest creates the FromNone request.
func (client *ResiliencyServiceDrivenClient) fromNoneCreateRequest(ctx context.Context, options *ResiliencyServiceDrivenClientFromNoneOptions) (*policy.Request, error) {
	host := "{endpoint}/resiliency/service-driven/client:v2/service:{serviceDeploymentVersion}/api-version:{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{serviceDeploymentVersion}", client.serviceDeploymentVersion)
	urlPath := "/add-optional-param/from-none"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.NewParameter != nil {
		reqQP.Set("new-parameter", *options.NewParameter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// FromOneOptional - Tests that we can grow up an operation from accepting one optional parameter to accepting two optional
// parameters.
//   - options - ResiliencyServiceDrivenClientFromOneOptionalOptions contains the optional parameters for the ResiliencyServiceDrivenClient.FromOneOptional
//     method.
func (client *ResiliencyServiceDrivenClient) FromOneOptional(ctx context.Context, options *ResiliencyServiceDrivenClientFromOneOptionalOptions) (ResiliencyServiceDrivenClientFromOneOptionalResponse, error) {
	var err error
	req, err := client.fromOneOptionalCreateRequest(ctx, options)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, err
	}
	return ResiliencyServiceDrivenClientFromOneOptionalResponse{}, nil
}

// fromOneOptionalCreateRequest creates the FromOneOptional request.
func (client *ResiliencyServiceDrivenClient) fromOneOptionalCreateRequest(ctx context.Context, options *ResiliencyServiceDrivenClientFromOneOptionalOptions) (*policy.Request, error) {
	host := "{endpoint}/resiliency/service-driven/client:v2/service:{serviceDeploymentVersion}/api-version:{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{serviceDeploymentVersion}", client.serviceDeploymentVersion)
	urlPath := "/add-optional-param/from-one-optional"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.NewParameter != nil {
		reqQP.Set("new-parameter", *options.NewParameter)
	}
	if options != nil && options.Parameter != nil {
		reqQP.Set("parameter", *options.Parameter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// FromOneRequired - Operation that grew up from accepting one required parameter to accepting a required parameter and an
// optional parameter.
//   - parameter - I am a required parameter
//   - options - ResiliencyServiceDrivenClientFromOneRequiredOptions contains the optional parameters for the ResiliencyServiceDrivenClient.FromOneRequired
//     method.
func (client *ResiliencyServiceDrivenClient) FromOneRequired(ctx context.Context, parameter string, options *ResiliencyServiceDrivenClientFromOneRequiredOptions) (ResiliencyServiceDrivenClientFromOneRequiredResponse, error) {
	var err error
	req, err := client.fromOneRequiredCreateRequest(ctx, parameter, options)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, err
	}
	return ResiliencyServiceDrivenClientFromOneRequiredResponse{}, nil
}

// fromOneRequiredCreateRequest creates the FromOneRequired request.
func (client *ResiliencyServiceDrivenClient) fromOneRequiredCreateRequest(ctx context.Context, parameter string, options *ResiliencyServiceDrivenClientFromOneRequiredOptions) (*policy.Request, error) {
	host := "{endpoint}/resiliency/service-driven/client:v2/service:{serviceDeploymentVersion}/api-version:{apiVersion}"
	host = strings.ReplaceAll(host, "{endpoint}", client.endpoint)
	host = strings.ReplaceAll(host, "{serviceDeploymentVersion}", client.serviceDeploymentVersion)
	urlPath := "/add-optional-param/from-one-required"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.NewParameter != nil {
		reqQP.Set("new-parameter", *options.NewParameter)
	}
	reqQP.Set("parameter", parameter)
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}
