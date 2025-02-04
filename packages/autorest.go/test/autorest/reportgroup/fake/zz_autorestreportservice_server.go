// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	"generatortests/reportgroup"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
)

// AutoRestReportServiceServer is a fake server for instances of the reportgroup.AutoRestReportServiceClient type.
type AutoRestReportServiceServer struct {
	// GetOptionalReport is the fake for method AutoRestReportServiceClient.GetOptionalReport
	// HTTP status codes to indicate success: http.StatusOK
	GetOptionalReport func(ctx context.Context, options *reportgroup.AutoRestReportServiceClientGetOptionalReportOptions) (resp azfake.Responder[reportgroup.AutoRestReportServiceClientGetOptionalReportResponse], errResp azfake.ErrorResponder)

	// GetReport is the fake for method AutoRestReportServiceClient.GetReport
	// HTTP status codes to indicate success: http.StatusOK
	GetReport func(ctx context.Context, options *reportgroup.AutoRestReportServiceClientGetReportOptions) (resp azfake.Responder[reportgroup.AutoRestReportServiceClientGetReportResponse], errResp azfake.ErrorResponder)
}

// NewAutoRestReportServiceServerTransport creates a new instance of AutoRestReportServiceServerTransport with the provided implementation.
// The returned AutoRestReportServiceServerTransport instance is connected to an instance of reportgroup.AutoRestReportServiceClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAutoRestReportServiceServerTransport(srv *AutoRestReportServiceServer) *AutoRestReportServiceServerTransport {
	return &AutoRestReportServiceServerTransport{srv: srv}
}

// AutoRestReportServiceServerTransport connects instances of reportgroup.AutoRestReportServiceClient to instances of AutoRestReportServiceServer.
// Don't use this type directly, use NewAutoRestReportServiceServerTransport instead.
type AutoRestReportServiceServerTransport struct {
	srv *AutoRestReportServiceServer
}

// Do implements the policy.Transporter interface for AutoRestReportServiceServerTransport.
func (a *AutoRestReportServiceServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AutoRestReportServiceClient.GetOptionalReport":
		resp, err = a.dispatchGetOptionalReport(req)
	case "AutoRestReportServiceClient.GetReport":
		resp, err = a.dispatchGetReport(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AutoRestReportServiceServerTransport) dispatchGetOptionalReport(req *http.Request) (*http.Response, error) {
	if a.srv.GetOptionalReport == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetOptionalReport not implemented")}
	}
	qp := req.URL.Query()
	qualifierUnescaped, err := url.QueryUnescape(qp.Get("qualifier"))
	if err != nil {
		return nil, err
	}
	qualifierParam := getOptional(qualifierUnescaped)
	var options *reportgroup.AutoRestReportServiceClientGetOptionalReportOptions
	if qualifierParam != nil {
		options = &reportgroup.AutoRestReportServiceClientGetOptionalReportOptions{
			Qualifier: qualifierParam,
		}
	}
	respr, errRespr := a.srv.GetOptionalReport(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AutoRestReportServiceServerTransport) dispatchGetReport(req *http.Request) (*http.Response, error) {
	if a.srv.GetReport == nil {
		return nil, &nonRetriableError{errors.New("fake for method GetReport not implemented")}
	}
	qp := req.URL.Query()
	qualifierUnescaped, err := url.QueryUnescape(qp.Get("qualifier"))
	if err != nil {
		return nil, err
	}
	qualifierParam := getOptional(qualifierUnescaped)
	var options *reportgroup.AutoRestReportServiceClientGetReportOptions
	if qualifierParam != nil {
		options = &reportgroup.AutoRestReportServiceClientGetReportOptions{
			Qualifier: qualifierParam,
		}
	}
	respr, errRespr := a.srv.GetReport(req.Context(), options)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Value, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
