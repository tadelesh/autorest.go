//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package mediatypesgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strings"
)

// MediaTypesClient contains the methods for the MediaTypesClient group.
// Don't use this type directly, use NewMediaTypesClient() instead.
type MediaTypesClient struct {
	con *Connection
}

// NewMediaTypesClient creates a new instance of MediaTypesClient with the specified values.
func NewMediaTypesClient(con *Connection) *MediaTypesClient {
	return &MediaTypesClient{con: con}
}

// AnalyzeBody - Analyze body, that could be different media types.
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) AnalyzeBody(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyOptions) (MediaTypesClientAnalyzeBodyResponse, error) {
	req, err := client.analyzeBodyCreateRequest(ctx, contentType, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MediaTypesClientAnalyzeBodyResponse{}, client.analyzeBodyHandleError(resp)
	}
	return client.analyzeBodyHandleResponse(resp)
}

// analyzeBodyCreateRequest creates the AnalyzeBody request.
func (client *MediaTypesClient) analyzeBodyCreateRequest(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyOptions) (*azcore.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Content-Type", string(contentType))
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, req.SetBody(options.Input, string(contentType))
	}
	return req, nil
}

// analyzeBodyHandleResponse handles the AnalyzeBody response.
func (client *MediaTypesClient) analyzeBodyHandleResponse(resp *azcore.Response) (MediaTypesClientAnalyzeBodyResponse, error) {
	result := MediaTypesClientAnalyzeBodyResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	return result, nil
}

// analyzeBodyHandleError handles the AnalyzeBody error response.
func (client *MediaTypesClient) analyzeBodyHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// AnalyzeBodyWithSourcePath - Analyze body, that could be different media types.
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) AnalyzeBodyWithSourcePath(ctx context.Context, options *MediaTypesClientAnalyzeBodyWithSourcePathOptions) (MediaTypesClientAnalyzeBodyWithSourcePathResponse, error) {
	req, err := client.analyzeBodyWithSourcePathCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, client.analyzeBodyWithSourcePathHandleError(resp)
	}
	return client.analyzeBodyWithSourcePathHandleResponse(resp)
}

// analyzeBodyWithSourcePathCreateRequest creates the AnalyzeBodyWithSourcePath request.
func (client *MediaTypesClient) analyzeBodyWithSourcePathCreateRequest(ctx context.Context, options *MediaTypesClientAnalyzeBodyWithSourcePathOptions) (*azcore.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		return req, req.MarshalAsJSON(*options.Input)
	}
	return req, nil
}

// analyzeBodyWithSourcePathHandleResponse handles the AnalyzeBodyWithSourcePath response.
func (client *MediaTypesClient) analyzeBodyWithSourcePathHandleResponse(resp *azcore.Response) (MediaTypesClientAnalyzeBodyWithSourcePathResponse, error) {
	result := MediaTypesClientAnalyzeBodyWithSourcePathResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyWithSourcePathResponse{}, err
	}
	return result, nil
}

// analyzeBodyWithSourcePathHandleError handles the AnalyzeBodyWithSourcePath error response.
func (client *MediaTypesClient) analyzeBodyWithSourcePathHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ContentTypeWithEncoding - Pass in contentType 'text/plain; encoding=UTF-8' to pass test. Value for input does not matter
// If the operation fails it returns a generic error.
func (client *MediaTypesClient) ContentTypeWithEncoding(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingOptions) (MediaTypesClientContentTypeWithEncodingResponse, error) {
	req, err := client.contentTypeWithEncodingCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MediaTypesClientContentTypeWithEncodingResponse{}, client.contentTypeWithEncodingHandleError(resp)
	}
	return client.contentTypeWithEncodingHandleResponse(resp)
}

// contentTypeWithEncodingCreateRequest creates the ContentTypeWithEncoding request.
func (client *MediaTypesClient) contentTypeWithEncodingCreateRequest(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingOptions) (*azcore.Request, error) {
	urlPath := "/mediatypes/contentTypeWithEncoding"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.Input != nil {
		body := azcore.NopCloser(strings.NewReader(*options.Input))
		return req, req.SetBody(body, "text/plain; encoding=UTF-8")
	}
	return req, nil
}

// contentTypeWithEncodingHandleResponse handles the ContentTypeWithEncoding response.
func (client *MediaTypesClient) contentTypeWithEncodingHandleResponse(resp *azcore.Response) (MediaTypesClientContentTypeWithEncodingResponse, error) {
	result := MediaTypesClientContentTypeWithEncodingResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Value); err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	return result, nil
}

// contentTypeWithEncodingHandleError handles the ContentTypeWithEncoding error response.
func (client *MediaTypesClient) contentTypeWithEncodingHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
