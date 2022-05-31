//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package mediatypesgroup

const host = "http://localhost:3000"

// ContentType - Content type for upload
type ContentType string

const (
	// ContentTypeApplicationPDF - Content Type 'application/pdf'
	ContentTypeApplicationPDF ContentType = "application/pdf"
	// ContentTypeImageJPEG - Content Type 'image/jpeg'
	ContentTypeImageJPEG ContentType = "image/jpeg"
	// ContentTypeImagePNG - Content Type 'image/png'
	ContentTypeImagePNG ContentType = "image/png"
	// ContentTypeImageTIFF - Content Type 'image/tiff'
	ContentTypeImageTIFF ContentType = "image/tiff"
)

// PossibleContentTypeValues returns the possible values for the ContentType const type.
func PossibleContentTypeValues() []ContentType {
	return []ContentType{
		ContentTypeApplicationPDF,
		ContentTypeImageJPEG,
		ContentTypeImagePNG,
		ContentTypeImageTIFF,
	}
}

// ContentType1 - Content type for upload
type ContentType1 string

const (
	// ContentType1ApplicationJSON - Content Type 'application/json'
	ContentType1ApplicationJSON ContentType1 = "application/json"
	// ContentType1ApplicationOctetStream - Content Type 'application/octet-stream'
	ContentType1ApplicationOctetStream ContentType1 = "application/octet-stream"
)

// PossibleContentType1Values returns the possible values for the ContentType1 const type.
func PossibleContentType1Values() []ContentType1 {
	return []ContentType1{
		ContentType1ApplicationJSON,
		ContentType1ApplicationOctetStream,
	}
}

// ContentType1AutoGenerated - Content type for upload
type ContentType1AutoGenerated string

const (
	// ContentType1AutoGeneratedApplicationJSON - Content Type 'application/json'
	ContentType1AutoGeneratedApplicationJSON ContentType1AutoGenerated = "application/json"
	// ContentType1AutoGeneratedApplicationOctetStream - Content Type 'application/octet-stream'
	ContentType1AutoGeneratedApplicationOctetStream ContentType1AutoGenerated = "application/octet-stream"
	// ContentType1AutoGeneratedTextPlain - Content Type 'text/plain'
	ContentType1AutoGeneratedTextPlain ContentType1AutoGenerated = "text/plain"
)

// PossibleContentType1AutoGeneratedValues returns the possible values for the ContentType1AutoGenerated const type.
func PossibleContentType1AutoGeneratedValues() []ContentType1AutoGenerated {
	return []ContentType1AutoGenerated{
		ContentType1AutoGeneratedApplicationJSON,
		ContentType1AutoGeneratedApplicationOctetStream,
		ContentType1AutoGeneratedTextPlain,
	}
}

// ContentType1AutoGenerated2 - Content type for upload
type ContentType1AutoGenerated2 string

const (
	// ContentType1AutoGenerated2ApplicationJSON - Content Type 'application/json'
	ContentType1AutoGenerated2ApplicationJSON ContentType1AutoGenerated2 = "application/json"
	// ContentType1AutoGenerated2TextPlain - Content Type 'text/plain'
	ContentType1AutoGenerated2TextPlain ContentType1AutoGenerated2 = "text/plain"
)

// PossibleContentType1AutoGenerated2Values returns the possible values for the ContentType1AutoGenerated2 const type.
func PossibleContentType1AutoGenerated2Values() []ContentType1AutoGenerated2 {
	return []ContentType1AutoGenerated2{
		ContentType1AutoGenerated2ApplicationJSON,
		ContentType1AutoGenerated2TextPlain,
	}
}
