//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package azartifacts

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

type datasetClient struct {
	endpoint string
	pl       runtime.Pipeline
}

// newDatasetClient creates a new instance of datasetClient with the specified values.
// endpoint - The workspace development endpoint, for example https://myworkspace.dev.azuresynapse.net.
// pl - the pipeline used for sending requests and handling responses.
func newDatasetClient(endpoint string, pl runtime.Pipeline) *datasetClient {
	client := &datasetClient{
		endpoint: endpoint,
		pl:       pl,
	}
	return client
}

// BeginCreateOrUpdateDataset - Creates or updates a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// datasetName - The dataset name.
// dataset - Dataset resource definition.
// options - datasetClientBeginCreateOrUpdateDatasetOptions contains the optional parameters for the datasetClient.BeginCreateOrUpdateDataset
// method.
func (client *datasetClient) BeginCreateOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *datasetClientBeginCreateOrUpdateDatasetOptions) (*runtime.Poller[DatasetClientCreateOrUpdateDatasetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdateDataset(ctx, datasetName, dataset, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatasetClientCreateOrUpdateDatasetResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatasetClientCreateOrUpdateDatasetResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdateDataset - Creates or updates a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *datasetClient) createOrUpdateDataset(ctx context.Context, datasetName string, dataset DatasetResource, options *datasetClientBeginCreateOrUpdateDatasetOptions) (*http.Response, error) {
	req, err := client.createOrUpdateDatasetCreateRequest(ctx, datasetName, dataset, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateDatasetCreateRequest creates the CreateOrUpdateDataset request.
func (client *datasetClient) createOrUpdateDatasetCreateRequest(ctx context.Context, datasetName string, dataset DatasetResource, options *datasetClientBeginCreateOrUpdateDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, dataset)
}

// BeginDeleteDataset - Deletes a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// datasetName - The dataset name.
// options - datasetClientBeginDeleteDatasetOptions contains the optional parameters for the datasetClient.BeginDeleteDataset
// method.
func (client *datasetClient) BeginDeleteDataset(ctx context.Context, datasetName string, options *datasetClientBeginDeleteDatasetOptions) (*runtime.Poller[DatasetClientDeleteDatasetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteDataset(ctx, datasetName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatasetClientDeleteDatasetResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatasetClientDeleteDatasetResponse](options.ResumeToken, client.pl, nil)
	}
}

// DeleteDataset - Deletes a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *datasetClient) deleteDataset(ctx context.Context, datasetName string, options *datasetClientBeginDeleteDatasetOptions) (*http.Response, error) {
	req, err := client.deleteDatasetCreateRequest(ctx, datasetName, options)
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

// deleteDatasetCreateRequest creates the DeleteDataset request.
func (client *datasetClient) deleteDatasetCreateRequest(ctx context.Context, datasetName string, options *datasetClientBeginDeleteDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetDataset - Gets a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// datasetName - The dataset name.
// options - datasetClientGetDatasetOptions contains the optional parameters for the datasetClient.GetDataset method.
func (client *datasetClient) GetDataset(ctx context.Context, datasetName string, options *datasetClientGetDatasetOptions) (DatasetClientGetDatasetResponse, error) {
	req, err := client.getDatasetCreateRequest(ctx, datasetName, options)
	if err != nil {
		return DatasetClientGetDatasetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DatasetClientGetDatasetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return DatasetClientGetDatasetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDatasetHandleResponse(resp)
}

// getDatasetCreateRequest creates the GetDataset request.
func (client *datasetClient) getDatasetCreateRequest(ctx context.Context, datasetName string, options *datasetClientGetDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDatasetHandleResponse handles the GetDataset response.
func (client *datasetClient) getDatasetHandleResponse(resp *http.Response) (DatasetClientGetDatasetResponse, error) {
	result := DatasetClientGetDatasetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetResource); err != nil {
		return DatasetClientGetDatasetResponse{}, err
	}
	return result, nil
}

// NewGetDatasetsByWorkspacePager - Lists datasets.
// Generated from API version 2019-06-01-preview
// options - datasetClientGetDatasetsByWorkspaceOptions contains the optional parameters for the datasetClient.GetDatasetsByWorkspace
// method.
func (client *datasetClient) NewGetDatasetsByWorkspacePager(options *datasetClientGetDatasetsByWorkspaceOptions) *runtime.Pager[DatasetClientGetDatasetsByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatasetClientGetDatasetsByWorkspaceResponse]{
		More: func(page DatasetClientGetDatasetsByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatasetClientGetDatasetsByWorkspaceResponse) (DatasetClientGetDatasetsByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.getDatasetsByWorkspaceCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DatasetClientGetDatasetsByWorkspaceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return DatasetClientGetDatasetsByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatasetClientGetDatasetsByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.getDatasetsByWorkspaceHandleResponse(resp)
		},
	})
}

// getDatasetsByWorkspaceCreateRequest creates the GetDatasetsByWorkspace request.
func (client *datasetClient) getDatasetsByWorkspaceCreateRequest(ctx context.Context, options *datasetClientGetDatasetsByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/datasets"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDatasetsByWorkspaceHandleResponse handles the GetDatasetsByWorkspace response.
func (client *datasetClient) getDatasetsByWorkspaceHandleResponse(resp *http.Response) (DatasetClientGetDatasetsByWorkspaceResponse, error) {
	result := DatasetClientGetDatasetsByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatasetListResponse); err != nil {
		return DatasetClientGetDatasetsByWorkspaceResponse{}, err
	}
	return result, nil
}

// BeginRenameDataset - Renames a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
// datasetName - The dataset name.
// request - proposed new name.
// options - datasetClientBeginRenameDatasetOptions contains the optional parameters for the datasetClient.BeginRenameDataset
// method.
func (client *datasetClient) BeginRenameDataset(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *datasetClientBeginRenameDatasetOptions) (*runtime.Poller[DatasetClientRenameDatasetResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.renameDataset(ctx, datasetName, request, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatasetClientRenameDatasetResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatasetClientRenameDatasetResponse](options.ResumeToken, client.pl, nil)
	}
}

// RenameDataset - Renames a dataset.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-06-01-preview
func (client *datasetClient) renameDataset(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *datasetClientBeginRenameDatasetOptions) (*http.Response, error) {
	req, err := client.renameDatasetCreateRequest(ctx, datasetName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// renameDatasetCreateRequest creates the RenameDataset request.
func (client *datasetClient) renameDatasetCreateRequest(ctx context.Context, datasetName string, request ArtifactRenameRequest, options *datasetClientBeginRenameDatasetOptions) (*policy.Request, error) {
	urlPath := "/datasets/{datasetName}/rename"
	if datasetName == "" {
		return nil, errors.New("parameter datasetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{datasetName}", url.PathEscape(datasetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, request)
}
