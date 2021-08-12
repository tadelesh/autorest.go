// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type pipelineClient struct {
	con *connection
}

// BeginCreateOrUpdatePipeline - Creates or updates a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) BeginCreateOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineBeginCreateOrUpdatePipelineOptions) (PipelineCreateOrUpdatePipelinePollerResponse, error) {
	resp, err := client.createOrUpdatePipeline(ctx, pipelineName, pipeline, options)
	if err != nil {
		return PipelineCreateOrUpdatePipelinePollerResponse{}, err
	}
	result := PipelineCreateOrUpdatePipelinePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("pipelineClient.CreateOrUpdatePipeline", resp, client.con.Pipeline(), client.createOrUpdatePipelineHandleError)
	if err != nil {
		return PipelineCreateOrUpdatePipelinePollerResponse{}, err
	}
	result.Poller = &PipelineCreateOrUpdatePipelinePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdatePipeline - Creates or updates a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) createOrUpdatePipeline(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineBeginCreateOrUpdatePipelineOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdatePipelineCreateRequest(ctx, pipelineName, pipeline, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdatePipelineHandleError(resp)
	}
	return resp, nil
}

// createOrUpdatePipelineCreateRequest creates the CreateOrUpdatePipeline request.
func (client *pipelineClient) createOrUpdatePipelineCreateRequest(ctx context.Context, pipelineName string, pipeline PipelineResource, options *PipelineBeginCreateOrUpdatePipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Header.Set("If-Match", *options.IfMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(pipeline)
}

// createOrUpdatePipelineHandleError handles the CreateOrUpdatePipeline error response.
func (client *pipelineClient) createOrUpdatePipelineHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// CreatePipelineRun - Creates a run of a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) CreatePipelineRun(ctx context.Context, pipelineName string, options *PipelineCreatePipelineRunOptions) (PipelineCreatePipelineRunResponse, error) {
	req, err := client.createPipelineRunCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return PipelineCreatePipelineRunResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineCreatePipelineRunResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return PipelineCreatePipelineRunResponse{}, client.createPipelineRunHandleError(resp)
	}
	return client.createPipelineRunHandleResponse(resp)
}

// createPipelineRunCreateRequest creates the CreatePipelineRun request.
func (client *pipelineClient) createPipelineRunCreateRequest(ctx context.Context, pipelineName string, options *PipelineCreatePipelineRunOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}/createRun"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	if options != nil && options.ReferencePipelineRunID != nil {
		reqQP.Set("referencePipelineRunId", *options.ReferencePipelineRunID)
	}
	if options != nil && options.IsRecovery != nil {
		reqQP.Set("isRecovery", strconv.FormatBool(*options.IsRecovery))
	}
	if options != nil && options.StartActivityName != nil {
		reqQP.Set("startActivityName", *options.StartActivityName)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, req.MarshalAsJSON(options.Parameters)
	}
	return req, nil
}

// createPipelineRunHandleResponse handles the CreatePipelineRun response.
func (client *pipelineClient) createPipelineRunHandleResponse(resp *azcore.Response) (PipelineCreatePipelineRunResponse, error) {
	result := PipelineCreatePipelineRunResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.CreateRunResponse); err != nil {
		return PipelineCreatePipelineRunResponse{}, err
	}
	return result, nil
}

// createPipelineRunHandleError handles the CreatePipelineRun error response.
func (client *pipelineClient) createPipelineRunHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginDeletePipeline - Deletes a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) BeginDeletePipeline(ctx context.Context, pipelineName string, options *PipelineBeginDeletePipelineOptions) (PipelineDeletePipelinePollerResponse, error) {
	resp, err := client.deletePipeline(ctx, pipelineName, options)
	if err != nil {
		return PipelineDeletePipelinePollerResponse{}, err
	}
	result := PipelineDeletePipelinePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("pipelineClient.DeletePipeline", resp, client.con.Pipeline(), client.deletePipelineHandleError)
	if err != nil {
		return PipelineDeletePipelinePollerResponse{}, err
	}
	result.Poller = &PipelineDeletePipelinePoller{
		pt: pt,
	}
	return result, nil
}

// DeletePipeline - Deletes a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) deletePipeline(ctx context.Context, pipelineName string, options *PipelineBeginDeletePipelineOptions) (*azcore.Response, error) {
	req, err := client.deletePipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deletePipelineHandleError(resp)
	}
	return resp, nil
}

// deletePipelineCreateRequest creates the DeletePipeline request.
func (client *pipelineClient) deletePipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineBeginDeletePipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// deletePipelineHandleError handles the DeletePipeline error response.
func (client *pipelineClient) deletePipelineHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetPipeline - Gets a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) GetPipeline(ctx context.Context, pipelineName string, options *PipelineGetPipelineOptions) (PipelineGetPipelineResponse, error) {
	req, err := client.getPipelineCreateRequest(ctx, pipelineName, options)
	if err != nil {
		return PipelineGetPipelineResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineGetPipelineResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return PipelineGetPipelineResponse{}, client.getPipelineHandleError(resp)
	}
	return client.getPipelineHandleResponse(resp)
}

// getPipelineCreateRequest creates the GetPipeline request.
func (client *pipelineClient) getPipelineCreateRequest(ctx context.Context, pipelineName string, options *PipelineGetPipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPipelineHandleResponse handles the GetPipeline response.
func (client *pipelineClient) getPipelineHandleResponse(resp *azcore.Response) (PipelineGetPipelineResponse, error) {
	result := PipelineGetPipelineResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PipelineResource); err != nil {
		return PipelineGetPipelineResponse{}, err
	}
	return result, nil
}

// getPipelineHandleError handles the GetPipeline error response.
func (client *pipelineClient) getPipelineHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// GetPipelinesByWorkspace - Lists pipelines.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) GetPipelinesByWorkspace(options *PipelineGetPipelinesByWorkspaceOptions) *PipelineGetPipelinesByWorkspacePager {
	return &PipelineGetPipelinesByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getPipelinesByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp PipelineGetPipelinesByWorkspaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.PipelineListResponse.NextLink)
		},
	}
}

// getPipelinesByWorkspaceCreateRequest creates the GetPipelinesByWorkspace request.
func (client *pipelineClient) getPipelinesByWorkspaceCreateRequest(ctx context.Context, options *PipelineGetPipelinesByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/pipelines"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPipelinesByWorkspaceHandleResponse handles the GetPipelinesByWorkspace response.
func (client *pipelineClient) getPipelinesByWorkspaceHandleResponse(resp *azcore.Response) (PipelineGetPipelinesByWorkspaceResponse, error) {
	result := PipelineGetPipelinesByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PipelineListResponse); err != nil {
		return PipelineGetPipelinesByWorkspaceResponse{}, err
	}
	return result, nil
}

// getPipelinesByWorkspaceHandleError handles the GetPipelinesByWorkspace error response.
func (client *pipelineClient) getPipelinesByWorkspaceHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}

// BeginRenamePipeline - Renames a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) BeginRenamePipeline(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineBeginRenamePipelineOptions) (PipelineRenamePipelinePollerResponse, error) {
	resp, err := client.renamePipeline(ctx, pipelineName, request, options)
	if err != nil {
		return PipelineRenamePipelinePollerResponse{}, err
	}
	result := PipelineRenamePipelinePollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("pipelineClient.RenamePipeline", resp, client.con.Pipeline(), client.renamePipelineHandleError)
	if err != nil {
		return PipelineRenamePipelinePollerResponse{}, err
	}
	result.Poller = &PipelineRenamePipelinePoller{
		pt: pt,
	}
	return result, nil
}

// RenamePipeline - Renames a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineClient) renamePipeline(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineBeginRenamePipelineOptions) (*azcore.Response, error) {
	req, err := client.renamePipelineCreateRequest(ctx, pipelineName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renamePipelineHandleError(resp)
	}
	return resp, nil
}

// renamePipelineCreateRequest creates the RenamePipeline request.
func (client *pipelineClient) renamePipelineCreateRequest(ctx context.Context, pipelineName string, request ArtifactRenameRequest, options *PipelineBeginRenamePipelineOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}/rename"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// renamePipelineHandleError handles the RenamePipeline error response.
func (client *pipelineClient) renamePipelineHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType.InnerError); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
