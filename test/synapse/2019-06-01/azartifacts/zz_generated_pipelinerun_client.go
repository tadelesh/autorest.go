//go:build go1.16
// +build go1.16

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

type pipelineRunClient struct {
	con *connection
}

// CancelPipelineRun - Cancel a pipeline run by its run ID.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineRunClient) CancelPipelineRun(ctx context.Context, runID string, options *PipelineRunCancelPipelineRunOptions) (PipelineRunCancelPipelineRunResponse, error) {
	req, err := client.cancelPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return PipelineRunCancelPipelineRunResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineRunCancelPipelineRunResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PipelineRunCancelPipelineRunResponse{}, client.cancelPipelineRunHandleError(resp)
	}
	return PipelineRunCancelPipelineRunResponse{RawResponse: resp.Response}, nil
}

// cancelPipelineRunCreateRequest creates the CancelPipelineRun request.
func (client *pipelineRunClient) cancelPipelineRunCreateRequest(ctx context.Context, runID string, options *PipelineRunCancelPipelineRunOptions) (*azcore.Request, error) {
	urlPath := "/pipelineruns/{runId}/cancel"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.IsRecursive != nil {
		reqQP.Set("isRecursive", strconv.FormatBool(*options.IsRecursive))
	}
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// cancelPipelineRunHandleError handles the CancelPipelineRun error response.
func (client *pipelineRunClient) cancelPipelineRunHandleError(resp *azcore.Response) error {
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

// GetPipelineRun - Get a pipeline run by its run ID.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineRunClient) GetPipelineRun(ctx context.Context, runID string, options *PipelineRunGetPipelineRunOptions) (PipelineRunGetPipelineRunResponse, error) {
	req, err := client.getPipelineRunCreateRequest(ctx, runID, options)
	if err != nil {
		return PipelineRunGetPipelineRunResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineRunGetPipelineRunResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PipelineRunGetPipelineRunResponse{}, client.getPipelineRunHandleError(resp)
	}
	return client.getPipelineRunHandleResponse(resp)
}

// getPipelineRunCreateRequest creates the GetPipelineRun request.
func (client *pipelineRunClient) getPipelineRunCreateRequest(ctx context.Context, runID string, options *PipelineRunGetPipelineRunOptions) (*azcore.Request, error) {
	urlPath := "/pipelineruns/{runId}"
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
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

// getPipelineRunHandleResponse handles the GetPipelineRun response.
func (client *pipelineRunClient) getPipelineRunHandleResponse(resp *azcore.Response) (PipelineRunGetPipelineRunResponse, error) {
	result := PipelineRunGetPipelineRunResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PipelineRun); err != nil {
		return PipelineRunGetPipelineRunResponse{}, err
	}
	return result, nil
}

// getPipelineRunHandleError handles the GetPipelineRun error response.
func (client *pipelineRunClient) getPipelineRunHandleError(resp *azcore.Response) error {
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

// QueryActivityRuns - Query activity runs based on input filter conditions.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineRunClient) QueryActivityRuns(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, options *PipelineRunQueryActivityRunsOptions) (PipelineRunQueryActivityRunsResponse, error) {
	req, err := client.queryActivityRunsCreateRequest(ctx, pipelineName, runID, filterParameters, options)
	if err != nil {
		return PipelineRunQueryActivityRunsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineRunQueryActivityRunsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PipelineRunQueryActivityRunsResponse{}, client.queryActivityRunsHandleError(resp)
	}
	return client.queryActivityRunsHandleResponse(resp)
}

// queryActivityRunsCreateRequest creates the QueryActivityRuns request.
func (client *pipelineRunClient) queryActivityRunsCreateRequest(ctx context.Context, pipelineName string, runID string, filterParameters RunFilterParameters, options *PipelineRunQueryActivityRunsOptions) (*azcore.Request, error) {
	urlPath := "/pipelines/{pipelineName}/pipelineruns/{runId}/queryActivityruns"
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	if runID == "" {
		return nil, errors.New("parameter runID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(filterParameters)
}

// queryActivityRunsHandleResponse handles the QueryActivityRuns response.
func (client *pipelineRunClient) queryActivityRunsHandleResponse(resp *azcore.Response) (PipelineRunQueryActivityRunsResponse, error) {
	result := PipelineRunQueryActivityRunsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.ActivityRunsQueryResponse); err != nil {
		return PipelineRunQueryActivityRunsResponse{}, err
	}
	return result, nil
}

// queryActivityRunsHandleError handles the QueryActivityRuns error response.
func (client *pipelineRunClient) queryActivityRunsHandleError(resp *azcore.Response) error {
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

// QueryPipelineRunsByWorkspace - Query pipeline runs in the workspace based on input filter conditions.
// If the operation fails it returns the *CloudError error type.
func (client *pipelineRunClient) QueryPipelineRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters, options *PipelineRunQueryPipelineRunsByWorkspaceOptions) (PipelineRunQueryPipelineRunsByWorkspaceResponse, error) {
	req, err := client.queryPipelineRunsByWorkspaceCreateRequest(ctx, filterParameters, options)
	if err != nil {
		return PipelineRunQueryPipelineRunsByWorkspaceResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PipelineRunQueryPipelineRunsByWorkspaceResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PipelineRunQueryPipelineRunsByWorkspaceResponse{}, client.queryPipelineRunsByWorkspaceHandleError(resp)
	}
	return client.queryPipelineRunsByWorkspaceHandleResponse(resp)
}

// queryPipelineRunsByWorkspaceCreateRequest creates the QueryPipelineRunsByWorkspace request.
func (client *pipelineRunClient) queryPipelineRunsByWorkspaceCreateRequest(ctx context.Context, filterParameters RunFilterParameters, options *PipelineRunQueryPipelineRunsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/queryPipelineRuns"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(filterParameters)
}

// queryPipelineRunsByWorkspaceHandleResponse handles the QueryPipelineRunsByWorkspace response.
func (client *pipelineRunClient) queryPipelineRunsByWorkspaceHandleResponse(resp *azcore.Response) (PipelineRunQueryPipelineRunsByWorkspaceResponse, error) {
	result := PipelineRunQueryPipelineRunsByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.PipelineRunsQueryResponse); err != nil {
		return PipelineRunQueryPipelineRunsByWorkspaceResponse{}, err
	}
	return result, nil
}

// queryPipelineRunsByWorkspaceHandleError handles the QueryPipelineRunsByWorkspace error response.
func (client *pipelineRunClient) queryPipelineRunsByWorkspaceHandleError(resp *azcore.Response) error {
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
