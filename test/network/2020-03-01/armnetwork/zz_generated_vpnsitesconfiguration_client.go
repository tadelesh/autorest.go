// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// VPNSitesConfigurationClient contains the methods for the VPNSitesConfiguration group.
// Don't use this type directly, use NewVPNSitesConfigurationClient() instead.
type VPNSitesConfigurationClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVPNSitesConfigurationClient creates a new instance of VPNSitesConfigurationClient with the specified values.
func NewVPNSitesConfigurationClient(con *armcore.Connection, subscriptionID string) *VPNSitesConfigurationClient {
	return &VPNSitesConfigurationClient{con: con, subscriptionID: subscriptionID}
}

// BeginDownload - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNSitesConfigurationClient) BeginDownload(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationBeginDownloadOptions) (VPNSitesConfigurationDownloadPollerResponse, error) {
	resp, err := client.download(ctx, resourceGroupName, virtualWANName, request, options)
	if err != nil {
		return VPNSitesConfigurationDownloadPollerResponse{}, err
	}
	result := VPNSitesConfigurationDownloadPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewLROPoller("VPNSitesConfigurationClient.Download", "location", resp, client.con.Pipeline(), client.downloadHandleError)
	if err != nil {
		return VPNSitesConfigurationDownloadPollerResponse{}, err
	}
	result.Poller = &VPNSitesConfigurationDownloadPoller{
		pt: pt,
	}
	return result, nil
}

// Download - Gives the sas-url to download the configurations for vpn-sites in a resource group.
// If the operation fails it returns the *CloudError error type.
func (client *VPNSitesConfigurationClient) download(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationBeginDownloadOptions) (*azcore.Response, error) {
	req, err := client.downloadCreateRequest(ctx, resourceGroupName, virtualWANName, request, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.downloadHandleError(resp)
	}
	return resp, nil
}

// downloadCreateRequest creates the Download request.
func (client *VPNSitesConfigurationClient) downloadCreateRequest(ctx context.Context, resourceGroupName string, virtualWANName string, request GetVPNSitesConfigurationRequest, options *VPNSitesConfigurationBeginDownloadOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualWans/{virtualWANName}/vpnConfiguration"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualWANName == "" {
		return nil, errors.New("parameter virtualWANName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualWANName}", url.PathEscape(virtualWANName))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(request)
}

// downloadHandleError handles the Download error response.
func (client *VPNSitesConfigurationClient) downloadHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	errType := CloudError{raw: string(body)}
	if err := resp.UnmarshalAsJSON(&errType); err != nil {
		return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
	}
	return azcore.NewResponseError(&errType, resp.Response)
}
