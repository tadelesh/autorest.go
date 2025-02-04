// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

// CreateResponse contains the response from method Client.Create.
type CreateResponse struct {
	// The response model for the Alias Create API for the case when the alias was successfully created.
	AliasesCreateResponse

	// AccessControlExposeHeaders contains the information returned from the Access-Control-Expose-Headers header response.
	AccessControlExposeHeaders *string
}

// GetScriptResponse contains the response from method Client.GetScript.
type GetScriptResponse struct {
	Value *string
}

// ListLROResponse contains the response from method Client.BeginListLRO.
type ListLROResponse struct {
	PagesOfThings
}

// ListResponseEnvelope contains the response from method Client.NewListPager.
type ListResponseEnvelope struct {
	// The response model for the List API. Returns a list of all the previously created aliases.
	ListResponse
}

// ListWithSharedNextOneResponse contains the response from method Client.NewListWithSharedNextOnePager.
type ListWithSharedNextOneResponse struct {
	// The response model for the List API. Returns a list of all the previously created aliases.
	ListResponse
}

// ListWithSharedNextTwoResponse contains the response from method Client.NewListWithSharedNextTwoPager.
type ListWithSharedNextTwoResponse struct {
	// The response model for the List API. Returns a list of all the previously created aliases.
	ListResponse
}

// PolicyAssignmentResponse contains the response from method Client.PolicyAssignment.
type PolicyAssignmentResponse struct {
	PolicyAssignmentProperties
}
