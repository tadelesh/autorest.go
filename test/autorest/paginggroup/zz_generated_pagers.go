//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package paginggroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"reflect"
)

// PagingFirstResponseEmptyPager provides operations for iterating over paged responses.
type PagingFirstResponseEmptyPager struct {
	client    *PagingClient
	current   PagingFirstResponseEmptyResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingFirstResponseEmptyResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingFirstResponseEmptyPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingFirstResponseEmptyPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResultValue.NextLink == nil || len(*p.current.ProductResultValue.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.firstResponseEmptyHandleError(resp)
		return false
	}
	result, err := p.client.firstResponseEmptyHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingFirstResponseEmptyResponse page.
func (p *PagingFirstResponseEmptyPager) PageResponse() PagingFirstResponseEmptyResponse {
	return p.current
}

// PagingGetMultiplePagesFailurePager provides operations for iterating over paged responses.
type PagingGetMultiplePagesFailurePager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesFailureResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesFailureResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesFailurePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesFailurePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesFailureHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesFailureHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesFailureResponse page.
func (p *PagingGetMultiplePagesFailurePager) PageResponse() PagingGetMultiplePagesFailureResponse {
	return p.current
}

// PagingGetMultiplePagesFailureURIPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesFailureURIPager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesFailureURIResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesFailureURIResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesFailureURIPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesFailureURIPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesFailureURIHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesFailureURIHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesFailureURIResponse page.
func (p *PagingGetMultiplePagesFailureURIPager) PageResponse() PagingGetMultiplePagesFailureURIResponse {
	return p.current
}

// PagingGetMultiplePagesFragmentNextLinkPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesFragmentNextLinkPager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesFragmentNextLinkResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesFragmentNextLinkResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesFragmentNextLinkPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesFragmentNextLinkPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ODataProductResult.ODataNextLink == nil || len(*p.current.ODataProductResult.ODataNextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesFragmentNextLinkHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesFragmentNextLinkHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesFragmentNextLinkResponse page.
func (p *PagingGetMultiplePagesFragmentNextLinkPager) PageResponse() PagingGetMultiplePagesFragmentNextLinkResponse {
	return p.current
}

// PagingGetMultiplePagesFragmentWithGroupingNextLinkPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesFragmentWithGroupingNextLinkPager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesFragmentWithGroupingNextLinkResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesFragmentWithGroupingNextLinkResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesFragmentWithGroupingNextLinkPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesFragmentWithGroupingNextLinkPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ODataProductResult.ODataNextLink == nil || len(*p.current.ODataProductResult.ODataNextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesFragmentWithGroupingNextLinkHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesFragmentWithGroupingNextLinkHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesFragmentWithGroupingNextLinkResponse page.
func (p *PagingGetMultiplePagesFragmentWithGroupingNextLinkPager) PageResponse() PagingGetMultiplePagesFragmentWithGroupingNextLinkResponse {
	return p.current
}

// PagingGetMultiplePagesLROPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesLROPager struct {
	client  *PagingClient
	current PagingGetMultiplePagesLROResponse
	err     error
	second  bool
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesLROPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesLROPager) NextPage(ctx context.Context) bool {
	if !p.second {
		p.second = true
		return true
	} else if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
	}
	req, err := azcore.NewRequest(ctx, http.MethodGet, *p.current.ProductResult.NextLink)
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		p.err = p.client.getMultiplePagesLROHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesLROHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesLROResponse page.
func (p *PagingGetMultiplePagesLROPager) PageResponse() PagingGetMultiplePagesLROResponse {
	return p.current
}

// PagingGetMultiplePagesPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesPager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesResponse page.
func (p *PagingGetMultiplePagesPager) PageResponse() PagingGetMultiplePagesResponse {
	return p.current
}

// PagingGetMultiplePagesRetryFirstPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesRetryFirstPager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesRetryFirstResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesRetryFirstResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesRetryFirstPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesRetryFirstPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesRetryFirstHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesRetryFirstHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesRetryFirstResponse page.
func (p *PagingGetMultiplePagesRetryFirstPager) PageResponse() PagingGetMultiplePagesRetryFirstResponse {
	return p.current
}

// PagingGetMultiplePagesRetrySecondPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesRetrySecondPager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesRetrySecondResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesRetrySecondResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesRetrySecondPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesRetrySecondPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesRetrySecondHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesRetrySecondHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesRetrySecondResponse page.
func (p *PagingGetMultiplePagesRetrySecondPager) PageResponse() PagingGetMultiplePagesRetrySecondResponse {
	return p.current
}

// PagingGetMultiplePagesWithOffsetPager provides operations for iterating over paged responses.
type PagingGetMultiplePagesWithOffsetPager struct {
	client    *PagingClient
	current   PagingGetMultiplePagesWithOffsetResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetMultiplePagesWithOffsetResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetMultiplePagesWithOffsetPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetMultiplePagesWithOffsetPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getMultiplePagesWithOffsetHandleError(resp)
		return false
	}
	result, err := p.client.getMultiplePagesWithOffsetHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetMultiplePagesWithOffsetResponse page.
func (p *PagingGetMultiplePagesWithOffsetPager) PageResponse() PagingGetMultiplePagesWithOffsetResponse {
	return p.current
}

// PagingGetNoItemNamePagesPager provides operations for iterating over paged responses.
type PagingGetNoItemNamePagesPager struct {
	client    *PagingClient
	current   PagingGetNoItemNamePagesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetNoItemNamePagesResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetNoItemNamePagesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetNoItemNamePagesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResultValue.NextLink == nil || len(*p.current.ProductResultValue.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getNoItemNamePagesHandleError(resp)
		return false
	}
	result, err := p.client.getNoItemNamePagesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetNoItemNamePagesResponse page.
func (p *PagingGetNoItemNamePagesPager) PageResponse() PagingGetNoItemNamePagesResponse {
	return p.current
}

// PagingGetODataMultiplePagesPager provides operations for iterating over paged responses.
type PagingGetODataMultiplePagesPager struct {
	client    *PagingClient
	current   PagingGetODataMultiplePagesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetODataMultiplePagesResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetODataMultiplePagesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetODataMultiplePagesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ODataProductResult.ODataNextLink == nil || len(*p.current.ODataProductResult.ODataNextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getODataMultiplePagesHandleError(resp)
		return false
	}
	result, err := p.client.getODataMultiplePagesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetODataMultiplePagesResponse page.
func (p *PagingGetODataMultiplePagesPager) PageResponse() PagingGetODataMultiplePagesResponse {
	return p.current
}

// PagingGetPagingModelWithItemNameWithXMSClientNamePager provides operations for iterating over paged responses.
type PagingGetPagingModelWithItemNameWithXMSClientNamePager struct {
	client    *PagingClient
	current   PagingGetPagingModelWithItemNameWithXMSClientNameResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetPagingModelWithItemNameWithXMSClientNameResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetPagingModelWithItemNameWithXMSClientNamePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetPagingModelWithItemNameWithXMSClientNamePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResultValueWithXMSClientName.NextLink == nil || len(*p.current.ProductResultValueWithXMSClientName.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getPagingModelWithItemNameWithXMSClientNameHandleError(resp)
		return false
	}
	result, err := p.client.getPagingModelWithItemNameWithXMSClientNameHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetPagingModelWithItemNameWithXMSClientNameResponse page.
func (p *PagingGetPagingModelWithItemNameWithXMSClientNamePager) PageResponse() PagingGetPagingModelWithItemNameWithXMSClientNameResponse {
	return p.current
}

// PagingGetSinglePagesFailurePager provides operations for iterating over paged responses.
type PagingGetSinglePagesFailurePager struct {
	client    *PagingClient
	current   PagingGetSinglePagesFailureResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetSinglePagesFailureResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetSinglePagesFailurePager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetSinglePagesFailurePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getSinglePagesFailureHandleError(resp)
		return false
	}
	result, err := p.client.getSinglePagesFailureHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetSinglePagesFailureResponse page.
func (p *PagingGetSinglePagesFailurePager) PageResponse() PagingGetSinglePagesFailureResponse {
	return p.current
}

// PagingGetSinglePagesPager provides operations for iterating over paged responses.
type PagingGetSinglePagesPager struct {
	client    *PagingClient
	current   PagingGetSinglePagesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetSinglePagesResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetSinglePagesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetSinglePagesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getSinglePagesHandleError(resp)
		return false
	}
	result, err := p.client.getSinglePagesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetSinglePagesResponse page.
func (p *PagingGetSinglePagesPager) PageResponse() PagingGetSinglePagesResponse {
	return p.current
}

// PagingGetWithQueryParamsPager provides operations for iterating over paged responses.
type PagingGetWithQueryParamsPager struct {
	client    *PagingClient
	current   PagingGetWithQueryParamsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingGetWithQueryParamsResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingGetWithQueryParamsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingGetWithQueryParamsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ProductResult.NextLink == nil || len(*p.current.ProductResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getWithQueryParamsHandleError(resp)
		return false
	}
	result, err := p.client.getWithQueryParamsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingGetWithQueryParamsResponse page.
func (p *PagingGetWithQueryParamsPager) PageResponse() PagingGetWithQueryParamsResponse {
	return p.current
}

// PagingNextFragmentPager provides operations for iterating over paged responses.
type PagingNextFragmentPager struct {
	client    *PagingClient
	current   PagingNextFragmentResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingNextFragmentResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingNextFragmentPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingNextFragmentPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ODataProductResult.ODataNextLink == nil || len(*p.current.ODataProductResult.ODataNextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.nextFragmentHandleError(resp)
		return false
	}
	result, err := p.client.nextFragmentHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingNextFragmentResponse page.
func (p *PagingNextFragmentPager) PageResponse() PagingNextFragmentResponse {
	return p.current
}

// PagingNextFragmentWithGroupingPager provides operations for iterating over paged responses.
type PagingNextFragmentWithGroupingPager struct {
	client    *PagingClient
	current   PagingNextFragmentWithGroupingResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, PagingNextFragmentWithGroupingResponse) (*azcore.Request, error)
}

// Err returns the last error encountered while paging.
func (p *PagingNextFragmentWithGroupingPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *PagingNextFragmentWithGroupingPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.ODataProductResult.ODataNextLink == nil || len(*p.current.ODataProductResult.ODataNextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.nextFragmentWithGroupingHandleError(resp)
		return false
	}
	result, err := p.client.nextFragmentWithGroupingHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current PagingNextFragmentWithGroupingResponse page.
func (p *PagingNextFragmentWithGroupingPager) PageResponse() PagingNextFragmentWithGroupingResponse {
	return p.current
}
