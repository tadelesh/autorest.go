//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package durationgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// HeaderClient contains the methods for the Encode.Duration group.
// Don't use this type directly, use a constructor function instead.
type HeaderClient struct {
	internal *azcore.Client
}

// - options - HeaderClientDefaultOptions contains the optional parameters for the HeaderClient.Default method.
func (client *HeaderClient) Default(ctx context.Context, duration string, options *HeaderClientDefaultOptions) (HeaderClientDefaultResponse, error) {
	var err error
	req, err := client.defaultCreateRequest(ctx, duration, options)
	if err != nil {
		return HeaderClientDefaultResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientDefaultResponse{}, err
	}
	return HeaderClientDefaultResponse{}, nil
}

// defaultCreateRequest creates the Default request.
func (client *HeaderClient) defaultCreateRequest(ctx context.Context, duration string, options *HeaderClientDefaultOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/header/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["duration"] = []string{duration}
	return req, nil
}

// - options - HeaderClientFloatSecondsOptions contains the optional parameters for the HeaderClient.FloatSeconds method.
func (client *HeaderClient) FloatSeconds(ctx context.Context, duration float32, options *HeaderClientFloatSecondsOptions) (HeaderClientFloatSecondsResponse, error) {
	var err error
	req, err := client.floatSecondsCreateRequest(ctx, duration, options)
	if err != nil {
		return HeaderClientFloatSecondsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientFloatSecondsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientFloatSecondsResponse{}, err
	}
	return HeaderClientFloatSecondsResponse{}, nil
}

// floatSecondsCreateRequest creates the FloatSeconds request.
func (client *HeaderClient) floatSecondsCreateRequest(ctx context.Context, duration float32, options *HeaderClientFloatSecondsOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/header/float-seconds"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["duration"] = []string{strconv.FormatFloat(float64(duration), 'f', -1, 32)}
	return req, nil
}

// - options - HeaderClientISO8601Options contains the optional parameters for the HeaderClient.ISO8601 method.
func (client *HeaderClient) ISO8601(ctx context.Context, duration string, options *HeaderClientISO8601Options) (HeaderClientISO8601Response, error) {
	var err error
	req, err := client.iso8601CreateRequest(ctx, duration, options)
	if err != nil {
		return HeaderClientISO8601Response{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientISO8601Response{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientISO8601Response{}, err
	}
	return HeaderClientISO8601Response{}, nil
}

// iso8601CreateRequest creates the ISO8601 request.
func (client *HeaderClient) iso8601CreateRequest(ctx context.Context, duration string, options *HeaderClientISO8601Options) (*policy.Request, error) {
	urlPath := "/encode/duration/header/iso8601"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["duration"] = []string{duration}
	return req, nil
}

// - options - HeaderClientISO8601ArrayOptions contains the optional parameters for the HeaderClient.ISO8601Array method.
func (client *HeaderClient) ISO8601Array(ctx context.Context, duration []string, options *HeaderClientISO8601ArrayOptions) (HeaderClientISO8601ArrayResponse, error) {
	var err error
	req, err := client.iso8601ArrayCreateRequest(ctx, duration, options)
	if err != nil {
		return HeaderClientISO8601ArrayResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientISO8601ArrayResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientISO8601ArrayResponse{}, err
	}
	return HeaderClientISO8601ArrayResponse{}, nil
}

// iso8601ArrayCreateRequest creates the ISO8601Array request.
func (client *HeaderClient) iso8601ArrayCreateRequest(ctx context.Context, duration []string, options *HeaderClientISO8601ArrayOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/header/iso8601-array"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["duration"] = []string{strings.Join(duration, ",")}
	return req, nil
}

// - options - HeaderClientInt32SecondsOptions contains the optional parameters for the HeaderClient.Int32Seconds method.
func (client *HeaderClient) Int32Seconds(ctx context.Context, duration int32, options *HeaderClientInt32SecondsOptions) (HeaderClientInt32SecondsResponse, error) {
	var err error
	req, err := client.int32SecondsCreateRequest(ctx, duration, options)
	if err != nil {
		return HeaderClientInt32SecondsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HeaderClientInt32SecondsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HeaderClientInt32SecondsResponse{}, err
	}
	return HeaderClientInt32SecondsResponse{}, nil
}

// int32SecondsCreateRequest creates the Int32Seconds request.
func (client *HeaderClient) int32SecondsCreateRequest(ctx context.Context, duration int32, options *HeaderClientInt32SecondsOptions) (*policy.Request, error) {
	urlPath := "/encode/duration/header/int32-seconds"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["duration"] = []string{strconv.FormatInt(int64(duration), 10)}
	return req, nil
}
