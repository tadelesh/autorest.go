//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package bytegroup

// ByteGetEmptyOptions contains the optional parameters for the Byte.GetEmpty method.
type ByteGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// ByteGetInvalidOptions contains the optional parameters for the Byte.GetInvalid method.
type ByteGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// ByteGetNonASCIIOptions contains the optional parameters for the Byte.GetNonASCII method.
type ByteGetNonASCIIOptions struct {
	// placeholder for future optional parameters
}

// ByteGetNullOptions contains the optional parameters for the Byte.GetNull method.
type ByteGetNullOptions struct {
	// placeholder for future optional parameters
}

// BytePutNonASCIIOptions contains the optional parameters for the Byte.PutNonASCII method.
type BytePutNonASCIIOptions struct {
	// placeholder for future optional parameters
}

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
