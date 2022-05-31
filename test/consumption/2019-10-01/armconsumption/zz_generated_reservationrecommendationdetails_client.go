//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ReservationRecommendationDetailsClient contains the methods for the ReservationRecommendationDetails group.
// Don't use this type directly, use NewReservationRecommendationDetailsClient() instead.
type ReservationRecommendationDetailsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewReservationRecommendationDetailsClient creates a new instance of ReservationRecommendationDetailsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewReservationRecommendationDetailsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ReservationRecommendationDetailsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ReservationRecommendationDetailsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// Get - Details of a reservation recommendation for what-if analysis of reserved instances.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2019-10-01
// billingScope - The scope associated with reservation recommendation details operations. This includes '/subscriptions/{subscriptionId}/'
// for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for resource group scope, /providers/Microsoft.Billing/billingAccounts/{billingAccountId}'
// for BillingAccount scope, and
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}' for billingProfile
// scope
// scope - Scope of the reservation.
// region - Used to select the region the recommendation should be generated for.
// term - Specify length of reservation recommendation term.
// lookBackPeriod - Filter the time period on which reservation recommendation results are based.
// product - Filter the products for which reservation recommendation results are generated. Examples: StandardDS1v2 (for
// VM), PremiumSSDManagedDisksP30 (for Managed Disks)
// options - ReservationRecommendationDetailsClientGetOptions contains the optional parameters for the ReservationRecommendationDetailsClient.Get
// method.
func (client *ReservationRecommendationDetailsClient) Get(ctx context.Context, billingScope string, scope Scope, region string, term Term, lookBackPeriod LookBackPeriod, product string, options *ReservationRecommendationDetailsClientGetOptions) (ReservationRecommendationDetailsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, billingScope, scope, region, term, lookBackPeriod, product, options)
	if err != nil {
		return ReservationRecommendationDetailsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReservationRecommendationDetailsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ReservationRecommendationDetailsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReservationRecommendationDetailsClient) getCreateRequest(ctx context.Context, billingScope string, scope Scope, region string, term Term, lookBackPeriod LookBackPeriod, product string, options *ReservationRecommendationDetailsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{billingScope}/providers/Microsoft.Consumption/reservationRecommendationDetails"
	urlPath = strings.ReplaceAll(urlPath, "{billingScope}", billingScope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	reqQP.Set("scope", string(scope))
	reqQP.Set("region", region)
	reqQP.Set("term", string(term))
	reqQP.Set("lookBackPeriod", string(lookBackPeriod))
	reqQP.Set("product", product)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReservationRecommendationDetailsClient) getHandleResponse(resp *http.Response) (ReservationRecommendationDetailsClientGetResponse, error) {
	result := ReservationRecommendationDetailsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ReservationRecommendationDetailsModel); err != nil {
		return ReservationRecommendationDetailsClientGetResponse{}, err
	}
	return result, nil
}
