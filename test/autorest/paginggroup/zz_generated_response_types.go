//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paginggroup

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// PagingFirstResponseEmptyResponse contains the response from method Paging.FirstResponseEmpty.
type PagingFirstResponseEmptyResponse struct {
	PagingFirstResponseEmptyResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingFirstResponseEmptyResult contains the result from method Paging.FirstResponseEmpty.
type PagingFirstResponseEmptyResult struct {
	ProductResultValue
}

// PagingGetMultiplePagesFailureResponse contains the response from method Paging.GetMultiplePagesFailure.
type PagingGetMultiplePagesFailureResponse struct {
	PagingGetMultiplePagesFailureResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesFailureResult contains the result from method Paging.GetMultiplePagesFailure.
type PagingGetMultiplePagesFailureResult struct {
	ProductResult
}

// PagingGetMultiplePagesFailureURIResponse contains the response from method Paging.GetMultiplePagesFailureURI.
type PagingGetMultiplePagesFailureURIResponse struct {
	PagingGetMultiplePagesFailureURIResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesFailureURIResult contains the result from method Paging.GetMultiplePagesFailureURI.
type PagingGetMultiplePagesFailureURIResult struct {
	ProductResult
}

// PagingGetMultiplePagesFragmentNextLinkResponse contains the response from method Paging.GetMultiplePagesFragmentNextLink.
type PagingGetMultiplePagesFragmentNextLinkResponse struct {
	PagingGetMultiplePagesFragmentNextLinkResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesFragmentNextLinkResult contains the result from method Paging.GetMultiplePagesFragmentNextLink.
type PagingGetMultiplePagesFragmentNextLinkResult struct {
	ODataProductResult
}

// PagingGetMultiplePagesFragmentWithGroupingNextLinkResponse contains the response from method Paging.GetMultiplePagesFragmentWithGroupingNextLink.
type PagingGetMultiplePagesFragmentWithGroupingNextLinkResponse struct {
	PagingGetMultiplePagesFragmentWithGroupingNextLinkResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesFragmentWithGroupingNextLinkResult contains the result from method Paging.GetMultiplePagesFragmentWithGroupingNextLink.
type PagingGetMultiplePagesFragmentWithGroupingNextLinkResult struct {
	ODataProductResult
}

// PagingGetMultiplePagesLROPollerResponse contains the response from method Paging.GetMultiplePagesLRO.
type PagingGetMultiplePagesLROPollerResponse struct {
	// Poller contains an initialized poller.
	Poller *PagingGetMultiplePagesLROPoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
func (l PagingGetMultiplePagesLROPollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (*PagingGetMultiplePagesLROPager, error) {
	respType := &PagingGetMultiplePagesLROPager{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.current.ProductResult)
	if err != nil {
		return respType, err
	}
	respType.current.RawResponse = resp
	respType.client = l.Poller.client
	return respType, nil
}

// Resume rehydrates a PagingGetMultiplePagesLROPollerResponse from the provided client and resume token.
func (l *PagingGetMultiplePagesLROPollerResponse) Resume(ctx context.Context, client *PagingClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("PagingClient.GetMultiplePagesLRO", token, client.con.Pipeline(), client.getMultiplePagesLROHandleError)
	if err != nil {
		return err
	}
	poller := &PagingGetMultiplePagesLROPoller{
		pt:     pt,
		client: client,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// PagingGetMultiplePagesLROResponse contains the response from method Paging.GetMultiplePagesLRO.
type PagingGetMultiplePagesLROResponse struct {
	PagingGetMultiplePagesLROResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesLROResult contains the result from method Paging.GetMultiplePagesLRO.
type PagingGetMultiplePagesLROResult struct {
	ProductResult
}

// PagingGetMultiplePagesResponse contains the response from method Paging.GetMultiplePages.
type PagingGetMultiplePagesResponse struct {
	PagingGetMultiplePagesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesResult contains the result from method Paging.GetMultiplePages.
type PagingGetMultiplePagesResult struct {
	ProductResult
}

// PagingGetMultiplePagesRetryFirstResponse contains the response from method Paging.GetMultiplePagesRetryFirst.
type PagingGetMultiplePagesRetryFirstResponse struct {
	PagingGetMultiplePagesRetryFirstResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesRetryFirstResult contains the result from method Paging.GetMultiplePagesRetryFirst.
type PagingGetMultiplePagesRetryFirstResult struct {
	ProductResult
}

// PagingGetMultiplePagesRetrySecondResponse contains the response from method Paging.GetMultiplePagesRetrySecond.
type PagingGetMultiplePagesRetrySecondResponse struct {
	PagingGetMultiplePagesRetrySecondResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesRetrySecondResult contains the result from method Paging.GetMultiplePagesRetrySecond.
type PagingGetMultiplePagesRetrySecondResult struct {
	ProductResult
}

// PagingGetMultiplePagesWithOffsetResponse contains the response from method Paging.GetMultiplePagesWithOffset.
type PagingGetMultiplePagesWithOffsetResponse struct {
	PagingGetMultiplePagesWithOffsetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetMultiplePagesWithOffsetResult contains the result from method Paging.GetMultiplePagesWithOffset.
type PagingGetMultiplePagesWithOffsetResult struct {
	ProductResult
}

// PagingGetNoItemNamePagesResponse contains the response from method Paging.GetNoItemNamePages.
type PagingGetNoItemNamePagesResponse struct {
	PagingGetNoItemNamePagesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetNoItemNamePagesResult contains the result from method Paging.GetNoItemNamePages.
type PagingGetNoItemNamePagesResult struct {
	ProductResultValue
}

// PagingGetNullNextLinkNamePagesResponse contains the response from method Paging.GetNullNextLinkNamePages.
type PagingGetNullNextLinkNamePagesResponse struct {
	PagingGetNullNextLinkNamePagesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetNullNextLinkNamePagesResult contains the result from method Paging.GetNullNextLinkNamePages.
type PagingGetNullNextLinkNamePagesResult struct {
	ProductResult
}

// PagingGetODataMultiplePagesResponse contains the response from method Paging.GetODataMultiplePages.
type PagingGetODataMultiplePagesResponse struct {
	PagingGetODataMultiplePagesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetODataMultiplePagesResult contains the result from method Paging.GetODataMultiplePages.
type PagingGetODataMultiplePagesResult struct {
	ODataProductResult
}

// PagingGetPagingModelWithItemNameWithXMSClientNameResponse contains the response from method Paging.GetPagingModelWithItemNameWithXMSClientName.
type PagingGetPagingModelWithItemNameWithXMSClientNameResponse struct {
	PagingGetPagingModelWithItemNameWithXMSClientNameResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetPagingModelWithItemNameWithXMSClientNameResult contains the result from method Paging.GetPagingModelWithItemNameWithXMSClientName.
type PagingGetPagingModelWithItemNameWithXMSClientNameResult struct {
	ProductResultValueWithXMSClientName
}

// PagingGetSinglePagesFailureResponse contains the response from method Paging.GetSinglePagesFailure.
type PagingGetSinglePagesFailureResponse struct {
	PagingGetSinglePagesFailureResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetSinglePagesFailureResult contains the result from method Paging.GetSinglePagesFailure.
type PagingGetSinglePagesFailureResult struct {
	ProductResult
}

// PagingGetSinglePagesResponse contains the response from method Paging.GetSinglePages.
type PagingGetSinglePagesResponse struct {
	PagingGetSinglePagesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetSinglePagesResult contains the result from method Paging.GetSinglePages.
type PagingGetSinglePagesResult struct {
	ProductResult
}

// PagingGetWithQueryParamsResponse contains the response from method Paging.GetWithQueryParams.
type PagingGetWithQueryParamsResponse struct {
	PagingGetWithQueryParamsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingGetWithQueryParamsResult contains the result from method Paging.GetWithQueryParams.
type PagingGetWithQueryParamsResult struct {
	ProductResult
}

// PagingNextFragmentResponse contains the response from method Paging.NextFragment.
type PagingNextFragmentResponse struct {
	PagingNextFragmentResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingNextFragmentResult contains the result from method Paging.NextFragment.
type PagingNextFragmentResult struct {
	ODataProductResult
}

// PagingNextFragmentWithGroupingResponse contains the response from method Paging.NextFragmentWithGrouping.
type PagingNextFragmentWithGroupingResponse struct {
	PagingNextFragmentWithGroupingResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingNextFragmentWithGroupingResult contains the result from method Paging.NextFragmentWithGrouping.
type PagingNextFragmentWithGroupingResult struct {
	ODataProductResult
}

// PagingNextOperationWithQueryParamsResponse contains the response from method Paging.NextOperationWithQueryParams.
type PagingNextOperationWithQueryParamsResponse struct {
	PagingNextOperationWithQueryParamsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PagingNextOperationWithQueryParamsResult contains the result from method Paging.NextOperationWithQueryParams.
type PagingNextOperationWithQueryParamsResult struct {
	ProductResult
}
