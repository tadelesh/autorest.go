//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package formdatagroup

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw     string
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

// FormdataUploadFileOptions contains the optional parameters for the Formdata.UploadFile method.
type FormdataUploadFileOptions struct {
	// placeholder for future optional parameters
}

// FormdataUploadFileViaBodyOptions contains the optional parameters for the Formdata.UploadFileViaBody method.
type FormdataUploadFileViaBodyOptions struct {
	// placeholder for future optional parameters
}

// FormdataUploadFilesOptions contains the optional parameters for the Formdata.UploadFiles method.
type FormdataUploadFilesOptions struct {
	// placeholder for future optional parameters
}

type Paths1MqqetpFormdataStreamUploadfilePostRequestbodyContentMultipartFormDataSchema struct {
	// REQUIRED; File to upload.
	FileContent *azcore.ReadSeekCloser `json:"fileContent,omitempty"`

	// REQUIRED; File name to upload. Name has to be spelled exactly as written here.
	FileName *string `json:"fileName,omitempty"`
}

type Paths1P3Stk3FormdataStreamUploadfilesPostRequestbodyContentMultipartFormDataSchema struct {
	// REQUIRED; Files to upload.
	FileContent []azcore.ReadSeekCloser `json:"fileContent,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type Paths1P3Stk3FormdataStreamUploadfilesPostRequestbodyContentMultipartFormDataSchema.
func (p Paths1P3Stk3FormdataStreamUploadfilesPostRequestbodyContentMultipartFormDataSchema) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "fileContent", p.FileContent)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
