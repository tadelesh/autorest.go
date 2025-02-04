// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package accessgroup

import "github.com/Azure/azure-sdk-for-go/sdk/azcore"

// AccessClient - Test for internal decorator.
// Don't use this type directly, use a constructor function instead.
type AccessClient struct {
	internal *azcore.Client
}

// NewInternalOperationClient creates a new instance of [InternalOperationClient].
func (client *AccessClient) NewInternalOperationClient() *InternalOperationClient {
	return &InternalOperationClient{
		internal: client.internal,
	}
}

// NewPublicOperationClient creates a new instance of [PublicOperationClient].
func (client *AccessClient) NewPublicOperationClient() *PublicOperationClient {
	return &PublicOperationClient{
		internal: client.internal,
	}
}

// NewRelativeModelInOperationClient creates a new instance of [RelativeModelInOperationClient].
func (client *AccessClient) NewRelativeModelInOperationClient() *RelativeModelInOperationClient {
	return &RelativeModelInOperationClient{
		internal: client.internal,
	}
}

// NewSharedModelInOperationClient creates a new instance of [SharedModelInOperationClient].
func (client *AccessClient) NewSharedModelInOperationClient() *SharedModelInOperationClient {
	return &SharedModelInOperationClient{
		internal: client.internal,
	}
}
