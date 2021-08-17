//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package extenumsgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// PetClient contains the methods for the Pet group.
// Don't use this type directly, use NewPetClient() instead.
type PetClient struct {
	con *Connection
}

// NewPetClient creates a new instance of PetClient with the specified values.
func NewPetClient(con *Connection) *PetClient {
	return &PetClient{con: con}
}

// AddPet - add pet
// If the operation fails it returns a generic error.
func (client *PetClient) AddPet(ctx context.Context, options *PetAddPetOptions) (PetAddPetResponse, error) {
	req, err := client.addPetCreateRequest(ctx, options)
	if err != nil {
		return PetAddPetResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PetAddPetResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PetAddPetResponse{}, client.addPetHandleError(resp)
	}
	return client.addPetHandleResponse(resp)
}

// addPetCreateRequest creates the AddPet request.
func (client *PetClient) addPetCreateRequest(ctx context.Context, options *PetAddPetOptions) (*azcore.Request, error) {
	urlPath := "/extensibleenums/pet/addPet"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil && options.PetParam != nil {
		return req, req.MarshalAsJSON(*options.PetParam)
	}
	return req, nil
}

// addPetHandleResponse handles the AddPet response.
func (client *PetClient) addPetHandleResponse(resp *azcore.Response) (PetAddPetResponse, error) {
	result := PetAddPetResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Pet); err != nil {
		return PetAddPetResponse{}, err
	}
	return result, nil
}

// addPetHandleError handles the AddPet error response.
func (client *PetClient) addPetHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetByPetID - get pet by id
// If the operation fails it returns a generic error.
func (client *PetClient) GetByPetID(ctx context.Context, petID string, options *PetGetByPetIDOptions) (PetGetByPetIDResponse, error) {
	req, err := client.getByPetIDCreateRequest(ctx, petID, options)
	if err != nil {
		return PetGetByPetIDResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return PetGetByPetIDResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return PetGetByPetIDResponse{}, client.getByPetIDHandleError(resp)
	}
	return client.getByPetIDHandleResponse(resp)
}

// getByPetIDCreateRequest creates the GetByPetID request.
func (client *PetClient) getByPetIDCreateRequest(ctx context.Context, petID string, options *PetGetByPetIDOptions) (*azcore.Request, error) {
	urlPath := "/extensibleenums/pet/{petId}"
	if petID == "" {
		return nil, errors.New("parameter petID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{petId}", url.PathEscape(petID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getByPetIDHandleResponse handles the GetByPetID response.
func (client *PetClient) getByPetIDHandleResponse(resp *azcore.Response) (PetGetByPetIDResponse, error) {
	result := PetGetByPetIDResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Pet); err != nil {
		return PetGetByPetIDResponse{}, err
	}
	return result, nil
}

// getByPetIDHandleError handles the GetByPetID error response.
func (client *PetClient) getByPetIDHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
