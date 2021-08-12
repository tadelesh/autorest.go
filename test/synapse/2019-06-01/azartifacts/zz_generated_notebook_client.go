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
	"strings"
)

type notebookClient struct {
	con *connection
}

// BeginCreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) BeginCreateOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (NotebookCreateOrUpdateNotebookPollerResponse, error) {
	resp, err := client.createOrUpdateNotebook(ctx, notebookName, notebook, options)
	if err != nil {
		return NotebookCreateOrUpdateNotebookPollerResponse{}, err
	}
	result := NotebookCreateOrUpdateNotebookPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("notebookClient.CreateOrUpdateNotebook", resp, client.con.Pipeline(), client.createOrUpdateNotebookHandleError)
	if err != nil {
		return NotebookCreateOrUpdateNotebookPollerResponse{}, err
	}
	result.Poller = &NotebookCreateOrUpdateNotebookPoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdateNotebook - Creates or updates a Note Book.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) createOrUpdateNotebook(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateNotebookCreateRequest(ctx, notebookName, notebook, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.createOrUpdateNotebookHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateNotebookCreateRequest creates the CreateOrUpdateNotebook request.
func (client *notebookClient) createOrUpdateNotebookCreateRequest(ctx context.Context, notebookName string, notebook NotebookResource, options *NotebookBeginCreateOrUpdateNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
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
	return req, req.MarshalAsJSON(notebook)
}

// createOrUpdateNotebookHandleError handles the CreateOrUpdateNotebook error response.
func (client *notebookClient) createOrUpdateNotebookHandleError(resp *azcore.Response) error {
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

// BeginDeleteNotebook - Deletes a Note book.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) BeginDeleteNotebook(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (NotebookDeleteNotebookPollerResponse, error) {
	resp, err := client.deleteNotebook(ctx, notebookName, options)
	if err != nil {
		return NotebookDeleteNotebookPollerResponse{}, err
	}
	result := NotebookDeleteNotebookPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("notebookClient.DeleteNotebook", resp, client.con.Pipeline(), client.deleteNotebookHandleError)
	if err != nil {
		return NotebookDeleteNotebookPollerResponse{}, err
	}
	result.Poller = &NotebookDeleteNotebookPoller{
		pt: pt,
	}
	return result, nil
}

// DeleteNotebook - Deletes a Note book.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) deleteNotebook(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*azcore.Response, error) {
	req, err := client.deleteNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteNotebookHandleError(resp)
	}
	return resp, nil
}

// deleteNotebookCreateRequest creates the DeleteNotebook request.
func (client *notebookClient) deleteNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookBeginDeleteNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
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

// deleteNotebookHandleError handles the DeleteNotebook error response.
func (client *notebookClient) deleteNotebookHandleError(resp *azcore.Response) error {
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

// GetNotebook - Gets a Note Book.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) GetNotebook(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (NotebookGetNotebookResponse, error) {
	req, err := client.getNotebookCreateRequest(ctx, notebookName, options)
	if err != nil {
		return NotebookGetNotebookResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return NotebookGetNotebookResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNotModified) {
		return NotebookGetNotebookResponse{}, client.getNotebookHandleError(resp)
	}
	return client.getNotebookHandleResponse(resp)
}

// getNotebookCreateRequest creates the GetNotebook request.
func (client *notebookClient) getNotebookCreateRequest(ctx context.Context, notebookName string, options *NotebookGetNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
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

// getNotebookHandleResponse handles the GetNotebook response.
func (client *notebookClient) getNotebookHandleResponse(resp *azcore.Response) (NotebookGetNotebookResponse, error) {
	result := NotebookGetNotebookResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NotebookResource); err != nil {
		return NotebookGetNotebookResponse{}, err
	}
	return result, nil
}

// getNotebookHandleError handles the GetNotebook error response.
func (client *notebookClient) getNotebookHandleError(resp *azcore.Response) error {
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

// GetNotebookSummaryByWorkSpace - Lists a summary of Notebooks.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) GetNotebookSummaryByWorkSpace(options *NotebookGetNotebookSummaryByWorkSpaceOptions) *NotebookGetNotebookSummaryByWorkSpacePager {
	return &NotebookGetNotebookSummaryByWorkSpacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getNotebookSummaryByWorkSpaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp NotebookGetNotebookSummaryByWorkSpaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
	}
}

// getNotebookSummaryByWorkSpaceCreateRequest creates the GetNotebookSummaryByWorkSpace request.
func (client *notebookClient) getNotebookSummaryByWorkSpaceCreateRequest(ctx context.Context, options *NotebookGetNotebookSummaryByWorkSpaceOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/summary"
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

// getNotebookSummaryByWorkSpaceHandleResponse handles the GetNotebookSummaryByWorkSpace response.
func (client *notebookClient) getNotebookSummaryByWorkSpaceHandleResponse(resp *azcore.Response) (NotebookGetNotebookSummaryByWorkSpaceResponse, error) {
	result := NotebookGetNotebookSummaryByWorkSpaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NotebookListResponse); err != nil {
		return NotebookGetNotebookSummaryByWorkSpaceResponse{}, err
	}
	return result, nil
}

// getNotebookSummaryByWorkSpaceHandleError handles the GetNotebookSummaryByWorkSpace error response.
func (client *notebookClient) getNotebookSummaryByWorkSpaceHandleError(resp *azcore.Response) error {
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

// GetNotebooksByWorkspace - Lists Notebooks.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) GetNotebooksByWorkspace(options *NotebookGetNotebooksByWorkspaceOptions) *NotebookGetNotebooksByWorkspacePager {
	return &NotebookGetNotebooksByWorkspacePager{
		client: client,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.getNotebooksByWorkspaceCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp NotebookGetNotebooksByWorkspaceResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.NotebookListResponse.NextLink)
		},
	}
}

// getNotebooksByWorkspaceCreateRequest creates the GetNotebooksByWorkspace request.
func (client *notebookClient) getNotebooksByWorkspaceCreateRequest(ctx context.Context, options *NotebookGetNotebooksByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/notebooks"
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

// getNotebooksByWorkspaceHandleResponse handles the GetNotebooksByWorkspace response.
func (client *notebookClient) getNotebooksByWorkspaceHandleResponse(resp *azcore.Response) (NotebookGetNotebooksByWorkspaceResponse, error) {
	result := NotebookGetNotebooksByWorkspaceResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.NotebookListResponse); err != nil {
		return NotebookGetNotebooksByWorkspaceResponse{}, err
	}
	return result, nil
}

// getNotebooksByWorkspaceHandleError handles the GetNotebooksByWorkspace error response.
func (client *notebookClient) getNotebooksByWorkspaceHandleError(resp *azcore.Response) error {
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

// BeginRenameNotebook - Renames a notebook.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) BeginRenameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (NotebookRenameNotebookPollerResponse, error) {
	resp, err := client.renameNotebook(ctx, notebookName, request, options)
	if err != nil {
		return NotebookRenameNotebookPollerResponse{}, err
	}
	result := NotebookRenameNotebookPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := azcore.NewLROPoller("notebookClient.RenameNotebook", resp, client.con.Pipeline(), client.renameNotebookHandleError)
	if err != nil {
		return NotebookRenameNotebookPollerResponse{}, err
	}
	result.Poller = &NotebookRenameNotebookPoller{
		pt: pt,
	}
	return result, nil
}

// RenameNotebook - Renames a notebook.
// If the operation fails it returns the *CloudError error type.
func (client *notebookClient) renameNotebook(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*azcore.Response, error) {
	req, err := client.renameNotebookCreateRequest(ctx, notebookName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.renameNotebookHandleError(resp)
	}
	return resp, nil
}

// renameNotebookCreateRequest creates the RenameNotebook request.
func (client *notebookClient) renameNotebookCreateRequest(ctx context.Context, notebookName string, request ArtifactRenameRequest, options *NotebookBeginRenameNotebookOptions) (*azcore.Request, error) {
	urlPath := "/notebooks/{notebookName}/rename"
	if notebookName == "" {
		return nil, errors.New("parameter notebookName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{notebookName}", url.PathEscape(notebookName))
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

// renameNotebookHandleError handles the RenameNotebook error response.
func (client *notebookClient) renameNotebookHandleError(resp *azcore.Response) error {
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
