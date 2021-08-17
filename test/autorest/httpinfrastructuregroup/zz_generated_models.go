//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

type B struct {
	MyException
	TextStatusCode *string `json:"textStatusCode,omitempty"`
}

type C struct {
	HTTPCode *string `json:"httpCode,omitempty"`
}

type D struct {
	HTTPStatusCode *string `json:"httpStatusCode,omitempty"`
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

// HTTPClientFailureDelete400Options contains the optional parameters for the HTTPClientFailure.Delete400 method.
type HTTPClientFailureDelete400Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureDelete407Options contains the optional parameters for the HTTPClientFailure.Delete407 method.
type HTTPClientFailureDelete407Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureDelete417Options contains the optional parameters for the HTTPClientFailure.Delete417 method.
type HTTPClientFailureDelete417Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureGet400Options contains the optional parameters for the HTTPClientFailure.Get400 method.
type HTTPClientFailureGet400Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureGet402Options contains the optional parameters for the HTTPClientFailure.Get402 method.
type HTTPClientFailureGet402Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureGet403Options contains the optional parameters for the HTTPClientFailure.Get403 method.
type HTTPClientFailureGet403Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureGet411Options contains the optional parameters for the HTTPClientFailure.Get411 method.
type HTTPClientFailureGet411Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureGet412Options contains the optional parameters for the HTTPClientFailure.Get412 method.
type HTTPClientFailureGet412Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureGet416Options contains the optional parameters for the HTTPClientFailure.Get416 method.
type HTTPClientFailureGet416Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureHead400Options contains the optional parameters for the HTTPClientFailure.Head400 method.
type HTTPClientFailureHead400Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureHead401Options contains the optional parameters for the HTTPClientFailure.Head401 method.
type HTTPClientFailureHead401Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureHead410Options contains the optional parameters for the HTTPClientFailure.Head410 method.
type HTTPClientFailureHead410Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureHead429Options contains the optional parameters for the HTTPClientFailure.Head429 method.
type HTTPClientFailureHead429Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureOptions400Options contains the optional parameters for the HTTPClientFailure.Options400 method.
type HTTPClientFailureOptions400Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureOptions403Options contains the optional parameters for the HTTPClientFailure.Options403 method.
type HTTPClientFailureOptions403Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailureOptions412Options contains the optional parameters for the HTTPClientFailure.Options412 method.
type HTTPClientFailureOptions412Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePatch400Options contains the optional parameters for the HTTPClientFailure.Patch400 method.
type HTTPClientFailurePatch400Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePatch405Options contains the optional parameters for the HTTPClientFailure.Patch405 method.
type HTTPClientFailurePatch405Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePatch414Options contains the optional parameters for the HTTPClientFailure.Patch414 method.
type HTTPClientFailurePatch414Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePost400Options contains the optional parameters for the HTTPClientFailure.Post400 method.
type HTTPClientFailurePost400Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePost406Options contains the optional parameters for the HTTPClientFailure.Post406 method.
type HTTPClientFailurePost406Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePost415Options contains the optional parameters for the HTTPClientFailure.Post415 method.
type HTTPClientFailurePost415Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePut400Options contains the optional parameters for the HTTPClientFailure.Put400 method.
type HTTPClientFailurePut400Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePut404Options contains the optional parameters for the HTTPClientFailure.Put404 method.
type HTTPClientFailurePut404Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePut409Options contains the optional parameters for the HTTPClientFailure.Put409 method.
type HTTPClientFailurePut409Options struct {
	// placeholder for future optional parameters
}

// HTTPClientFailurePut413Options contains the optional parameters for the HTTPClientFailure.Put413 method.
type HTTPClientFailurePut413Options struct {
	// placeholder for future optional parameters
}

// HTTPFailureGetEmptyErrorOptions contains the optional parameters for the HTTPFailure.GetEmptyError method.
type HTTPFailureGetEmptyErrorOptions struct {
	// placeholder for future optional parameters
}

// HTTPFailureGetNoModelEmptyOptions contains the optional parameters for the HTTPFailure.GetNoModelEmpty method.
type HTTPFailureGetNoModelEmptyOptions struct {
	// placeholder for future optional parameters
}

// HTTPFailureGetNoModelErrorOptions contains the optional parameters for the HTTPFailure.GetNoModelError method.
type HTTPFailureGetNoModelErrorOptions struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsDelete307Options contains the optional parameters for the HTTPRedirects.Delete307 method.
type HTTPRedirectsDelete307Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsGet300Options contains the optional parameters for the HTTPRedirects.Get300 method.
type HTTPRedirectsGet300Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsGet301Options contains the optional parameters for the HTTPRedirects.Get301 method.
type HTTPRedirectsGet301Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsGet302Options contains the optional parameters for the HTTPRedirects.Get302 method.
type HTTPRedirectsGet302Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsGet307Options contains the optional parameters for the HTTPRedirects.Get307 method.
type HTTPRedirectsGet307Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsHead300Options contains the optional parameters for the HTTPRedirects.Head300 method.
type HTTPRedirectsHead300Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsHead301Options contains the optional parameters for the HTTPRedirects.Head301 method.
type HTTPRedirectsHead301Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsHead302Options contains the optional parameters for the HTTPRedirects.Head302 method.
type HTTPRedirectsHead302Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsHead307Options contains the optional parameters for the HTTPRedirects.Head307 method.
type HTTPRedirectsHead307Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsOptions307Options contains the optional parameters for the HTTPRedirects.Options307 method.
type HTTPRedirectsOptions307Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsPatch302Options contains the optional parameters for the HTTPRedirects.Patch302 method.
type HTTPRedirectsPatch302Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsPatch307Options contains the optional parameters for the HTTPRedirects.Patch307 method.
type HTTPRedirectsPatch307Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsPost303Options contains the optional parameters for the HTTPRedirects.Post303 method.
type HTTPRedirectsPost303Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsPost307Options contains the optional parameters for the HTTPRedirects.Post307 method.
type HTTPRedirectsPost307Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsPut301Options contains the optional parameters for the HTTPRedirects.Put301 method.
type HTTPRedirectsPut301Options struct {
	// placeholder for future optional parameters
}

// HTTPRedirectsPut307Options contains the optional parameters for the HTTPRedirects.Put307 method.
type HTTPRedirectsPut307Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryDelete503Options contains the optional parameters for the HTTPRetry.Delete503 method.
type HTTPRetryDelete503Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryGet502Options contains the optional parameters for the HTTPRetry.Get502 method.
type HTTPRetryGet502Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryHead408Options contains the optional parameters for the HTTPRetry.Head408 method.
type HTTPRetryHead408Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryOptions502Options contains the optional parameters for the HTTPRetry.Options502 method.
type HTTPRetryOptions502Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryPatch500Options contains the optional parameters for the HTTPRetry.Patch500 method.
type HTTPRetryPatch500Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryPatch504Options contains the optional parameters for the HTTPRetry.Patch504 method.
type HTTPRetryPatch504Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryPost503Options contains the optional parameters for the HTTPRetry.Post503 method.
type HTTPRetryPost503Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryPut500Options contains the optional parameters for the HTTPRetry.Put500 method.
type HTTPRetryPut500Options struct {
	// placeholder for future optional parameters
}

// HTTPRetryPut504Options contains the optional parameters for the HTTPRetry.Put504 method.
type HTTPRetryPut504Options struct {
	// placeholder for future optional parameters
}

// HTTPServerFailureDelete505Options contains the optional parameters for the HTTPServerFailure.Delete505 method.
type HTTPServerFailureDelete505Options struct {
	// placeholder for future optional parameters
}

// HTTPServerFailureGet501Options contains the optional parameters for the HTTPServerFailure.Get501 method.
type HTTPServerFailureGet501Options struct {
	// placeholder for future optional parameters
}

// HTTPServerFailureHead501Options contains the optional parameters for the HTTPServerFailure.Head501 method.
type HTTPServerFailureHead501Options struct {
	// placeholder for future optional parameters
}

// HTTPServerFailurePost505Options contains the optional parameters for the HTTPServerFailure.Post505 method.
type HTTPServerFailurePost505Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessDelete200Options contains the optional parameters for the HTTPSuccess.Delete200 method.
type HTTPSuccessDelete200Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessDelete202Options contains the optional parameters for the HTTPSuccess.Delete202 method.
type HTTPSuccessDelete202Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessDelete204Options contains the optional parameters for the HTTPSuccess.Delete204 method.
type HTTPSuccessDelete204Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessGet200Options contains the optional parameters for the HTTPSuccess.Get200 method.
type HTTPSuccessGet200Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessHead200Options contains the optional parameters for the HTTPSuccess.Head200 method.
type HTTPSuccessHead200Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessHead204Options contains the optional parameters for the HTTPSuccess.Head204 method.
type HTTPSuccessHead204Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessHead404Options contains the optional parameters for the HTTPSuccess.Head404 method.
type HTTPSuccessHead404Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessOptions200Options contains the optional parameters for the HTTPSuccess.Options200 method.
type HTTPSuccessOptions200Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPatch200Options contains the optional parameters for the HTTPSuccess.Patch200 method.
type HTTPSuccessPatch200Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPatch202Options contains the optional parameters for the HTTPSuccess.Patch202 method.
type HTTPSuccessPatch202Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPatch204Options contains the optional parameters for the HTTPSuccess.Patch204 method.
type HTTPSuccessPatch204Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPost200Options contains the optional parameters for the HTTPSuccess.Post200 method.
type HTTPSuccessPost200Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPost201Options contains the optional parameters for the HTTPSuccess.Post201 method.
type HTTPSuccessPost201Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPost202Options contains the optional parameters for the HTTPSuccess.Post202 method.
type HTTPSuccessPost202Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPost204Options contains the optional parameters for the HTTPSuccess.Post204 method.
type HTTPSuccessPost204Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPut200Options contains the optional parameters for the HTTPSuccess.Put200 method.
type HTTPSuccessPut200Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPut201Options contains the optional parameters for the HTTPSuccess.Put201 method.
type HTTPSuccessPut201Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPut202Options contains the optional parameters for the HTTPSuccess.Put202 method.
type HTTPSuccessPut202Options struct {
	// placeholder for future optional parameters
}

// HTTPSuccessPut204Options contains the optional parameters for the HTTPSuccess.Put204 method.
type HTTPSuccessPut204Options struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model201ModelDefaultError200ValidOptions contains the optional parameters for the MultipleResponses.Get200Model201ModelDefaultError200Valid
// method.
type MultipleResponsesGet200Model201ModelDefaultError200ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model201ModelDefaultError201ValidOptions contains the optional parameters for the MultipleResponses.Get200Model201ModelDefaultError201Valid
// method.
type MultipleResponsesGet200Model201ModelDefaultError201ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model201ModelDefaultError400ValidOptions contains the optional parameters for the MultipleResponses.Get200Model201ModelDefaultError400Valid
// method.
type MultipleResponsesGet200Model201ModelDefaultError400ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model204NoModelDefaultError200ValidOptions contains the optional parameters for the MultipleResponses.Get200Model204NoModelDefaultError200Valid
// method.
type MultipleResponsesGet200Model204NoModelDefaultError200ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model204NoModelDefaultError201InvalidOptions contains the optional parameters for the MultipleResponses.Get200Model204NoModelDefaultError201Invalid
// method.
type MultipleResponsesGet200Model204NoModelDefaultError201InvalidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model204NoModelDefaultError202NoneOptions contains the optional parameters for the MultipleResponses.Get200Model204NoModelDefaultError202None
// method.
type MultipleResponsesGet200Model204NoModelDefaultError202NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model204NoModelDefaultError204ValidOptions contains the optional parameters for the MultipleResponses.Get200Model204NoModelDefaultError204Valid
// method.
type MultipleResponsesGet200Model204NoModelDefaultError204ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200Model204NoModelDefaultError400ValidOptions contains the optional parameters for the MultipleResponses.Get200Model204NoModelDefaultError400Valid
// method.
type MultipleResponsesGet200Model204NoModelDefaultError400ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA200InvalidOptions contains the optional parameters for the MultipleResponses.Get200ModelA200Invalid method.
type MultipleResponsesGet200ModelA200InvalidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA200NoneOptions contains the optional parameters for the MultipleResponses.Get200ModelA200None method.
type MultipleResponsesGet200ModelA200NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA200ValidOptions contains the optional parameters for the MultipleResponses.Get200ModelA200Valid method.
type MultipleResponsesGet200ModelA200ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError200ValidOptions contains the optional parameters for the MultipleResponses.Get200ModelA201ModelC404ModelDDefaultError200Valid
// method.
type MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError200ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError201ValidOptions contains the optional parameters for the MultipleResponses.Get200ModelA201ModelC404ModelDDefaultError201Valid
// method.
type MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError201ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError400ValidOptions contains the optional parameters for the MultipleResponses.Get200ModelA201ModelC404ModelDDefaultError400Valid
// method.
type MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError400ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError404ValidOptions contains the optional parameters for the MultipleResponses.Get200ModelA201ModelC404ModelDDefaultError404Valid
// method.
type MultipleResponsesGet200ModelA201ModelC404ModelDDefaultError404ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA202ValidOptions contains the optional parameters for the MultipleResponses.Get200ModelA202Valid method.
type MultipleResponsesGet200ModelA202ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA400InvalidOptions contains the optional parameters for the MultipleResponses.Get200ModelA400Invalid method.
type MultipleResponsesGet200ModelA400InvalidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA400NoneOptions contains the optional parameters for the MultipleResponses.Get200ModelA400None method.
type MultipleResponsesGet200ModelA400NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet200ModelA400ValidOptions contains the optional parameters for the MultipleResponses.Get200ModelA400Valid method.
type MultipleResponsesGet200ModelA400ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet202None204NoneDefaultError202NoneOptions contains the optional parameters for the MultipleResponses.Get202None204NoneDefaultError202None
// method.
type MultipleResponsesGet202None204NoneDefaultError202NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet202None204NoneDefaultError204NoneOptions contains the optional parameters for the MultipleResponses.Get202None204NoneDefaultError204None
// method.
type MultipleResponsesGet202None204NoneDefaultError204NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet202None204NoneDefaultError400ValidOptions contains the optional parameters for the MultipleResponses.Get202None204NoneDefaultError400Valid
// method.
type MultipleResponsesGet202None204NoneDefaultError400ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet202None204NoneDefaultNone202InvalidOptions contains the optional parameters for the MultipleResponses.Get202None204NoneDefaultNone202Invalid
// method.
type MultipleResponsesGet202None204NoneDefaultNone202InvalidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet202None204NoneDefaultNone204NoneOptions contains the optional parameters for the MultipleResponses.Get202None204NoneDefaultNone204None
// method.
type MultipleResponsesGet202None204NoneDefaultNone204NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet202None204NoneDefaultNone400InvalidOptions contains the optional parameters for the MultipleResponses.Get202None204NoneDefaultNone400Invalid
// method.
type MultipleResponsesGet202None204NoneDefaultNone400InvalidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGet202None204NoneDefaultNone400NoneOptions contains the optional parameters for the MultipleResponses.Get202None204NoneDefaultNone400None
// method.
type MultipleResponsesGet202None204NoneDefaultNone400NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultModelA200NoneOptions contains the optional parameters for the MultipleResponses.GetDefaultModelA200None method.
type MultipleResponsesGetDefaultModelA200NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultModelA200ValidOptions contains the optional parameters for the MultipleResponses.GetDefaultModelA200Valid method.
type MultipleResponsesGetDefaultModelA200ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultModelA400NoneOptions contains the optional parameters for the MultipleResponses.GetDefaultModelA400None method.
type MultipleResponsesGetDefaultModelA400NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultModelA400ValidOptions contains the optional parameters for the MultipleResponses.GetDefaultModelA400Valid method.
type MultipleResponsesGetDefaultModelA400ValidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultNone200InvalidOptions contains the optional parameters for the MultipleResponses.GetDefaultNone200Invalid method.
type MultipleResponsesGetDefaultNone200InvalidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultNone200NoneOptions contains the optional parameters for the MultipleResponses.GetDefaultNone200None method.
type MultipleResponsesGetDefaultNone200NoneOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultNone400InvalidOptions contains the optional parameters for the MultipleResponses.GetDefaultNone400Invalid method.
type MultipleResponsesGetDefaultNone400InvalidOptions struct {
	// placeholder for future optional parameters
}

// MultipleResponsesGetDefaultNone400NoneOptions contains the optional parameters for the MultipleResponses.GetDefaultNone400None method.
type MultipleResponsesGetDefaultNone400NoneOptions struct {
	// placeholder for future optional parameters
}

// Implements the error and azcore.HTTPResponse interfaces.
type MyException struct {
	raw        string
	StatusCode *string `json:"statusCode,omitempty"`
}

// Error implements the error interface for type MyException.
// The contents of the error text are not contractual and subject to change.
func (e MyException) Error() string {
	return e.raw
}
