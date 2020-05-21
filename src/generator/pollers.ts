/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { Session } from '@azure-tools/autorest-extension-base';
import { camelCase, pascalCase } from '@azure-tools/codegen';
import { CodeModel, SchemaResponse, Operation } from '@azure-tools/codemodel';
import { values } from '@azure-tools/linq';
import { PollerInfo, isSchemaResponse } from '../common/helpers';
import { contentPreamble, sortAscending } from './helpers';
import { ImportManager } from './imports';

// Creates the content in pollers.go
export async function generatePollers(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('net/http');
  imports.add('errors');
  imports.add('encoding/json');
  imports.add('net/url');
  imports.add('time');
  let bodyText = '';
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });
  for (const poller of values(pollers)) {
    const pollerInterface = poller.name;
    const pollerName = camelCase(poller.name);
    let responseType = 'HTTPResponse';
    // HTTP Pollers do not need to perform the final get request since they do not return a model
    let finalResponseDeclaration = 'FinalResponse() *http.Response';
    let finalResponse = `${finalResponseDeclaration} {
      return p.pt.latestResponse().Response;
    }`;
    let pollUntilDoneResponse = '(*http.Response, error)';
    let pollUntilDoneReturn = 'p.FinalResponse(), nil';
    let handleResponse = '';
    let rawResponse = ''; // used to access the raw response field on response envelopes
    const schemaResponse = <SchemaResponse>poller.op.responses![0];
    let unmarshalResponse = 'nil';
    if (isSchemaResponse(schemaResponse) && schemaResponse.schema.language.go!.responseType.value != undefined) {
      responseType = schemaResponse.schema.language.go!.responseType.name;
      pollUntilDoneResponse = `(*${responseType}, error)`;
      pollUntilDoneReturn = 'p.FinalResponse(ctx)';
      rawResponse = '.RawResponse';
      unmarshalResponse = `resp.UnmarshalAsJSON(&result.${schemaResponse.schema.language.go!.responseType.value})`;
      // for operations that do return a model add a final response method that handles the final get URL scenario
      finalResponseDeclaration = `FinalResponse(ctx context.Context) (*${responseType}, error)`;
      finalResponse = `FinalResponse(ctx context.Context) (*${responseType}, error) {
        if p.pt.finalGetURL() == "" {
          // we can end up in this situation if the async operation returns a 200
          // with no polling URLs.  in that case return the response which should
          // contain the JSON payload (only do this for successful terminal cases).
          if lr := p.pt.latestResponse(); lr != nil && p.pt.hasSucceeded() {
            result, err := p.handleResponse(lr)
            if err != nil {
              return nil, err
            }
            return result, nil
          }
          return nil, errors.New("missing URL for retrieving result")
        }
        u, err := url.Parse(p.pt.finalGetURL())
        if err != nil {
          return nil, err
        }
        req := azcore.NewRequest(http.MethodGet, *u)
        if err != nil {
          return nil, err
        }
        resp, err := p.pipeline.Do(ctx, req)
        if err != nil {
          return nil, err
        }
        return p.handleResponse(resp)
      }`;
      handleResponse = `
      func (p *${pollerName}) handleResponse(resp *azcore.Response) (*${responseType}, error) {
        result := ${responseType}{RawResponse: resp.Response}
        if (resp.HasStatusCode(http.StatusNoContent)) {
          return &result, nil
        }
        if !resp.HasStatusCode(pollingCodes[:]...) {
          return nil, p.pt.handleError(resp)
        }
        return &result, ${unmarshalResponse}
      }
      `;
    }
    bodyText += `// ${pollerInterface} provides polling facilities until the operation completes
type ${pollerInterface} interface {
	Done() bool
	Poll(ctx context.Context) (*http.Response, error)
	${finalResponseDeclaration}
	ResumeToken() (string, error)
}

type ${pollerName} struct {
	// the client for making the request
	pipeline azcore.Pipeline
	// polling tracker
  pt pollingTracker
}

// Done returns true if there was an error or polling has reached a terminal state
func (p *${pollerName}) Done() bool {
  return p.pt.hasTerminated()
}

// Poll will send poll the service endpoint and return an http.Response or error received from the service
func (p *${pollerName}) Poll(ctx context.Context) (*http.Response, error) {
  if lroPollDone(ctx, p.pipeline, p.pt) {
    return p.pt.latestResponse().Response, p.pt.pollingError()
  }
  return nil, p.pt.pollingError()
}

func (p *${pollerName}) ${finalResponse}

// ResumeToken generates the string token that can be used with the Resume${pollerInterface} method
// on the client to create a new poller from the data held in the current poller type
func (p *${pollerName}) ResumeToken() (string, error) {
  if p.pt.hasTerminated() {
    return "", errors.New("cannot create a ResumeToken from a poller in a terminal state")
  }
  js, err := json.Marshal(p.pt)
  if err != nil {
    return "", err
  }
  return string(js), nil
}

func (p *${pollerName}) pollUntilDone(ctx context.Context, frequency time.Duration) ${pollUntilDoneResponse} {
    for {
        resp, err := p.Poll(ctx)
        if err != nil {
            return nil, err
        }
        if p.Done() {
          break
        }
        if delay := azcore.RetryAfter(resp); delay > 0 {
            time.Sleep(delay)
        } else {
            time.Sleep(frequency)
        }
    }
    return ${pollUntilDoneReturn}
}
${handleResponse}
`;
  }
  text += imports.text();
  text += bodyText;
  return text;
}


// Creates the content in pollers_helper.go
export async function generatePollersHelper(session: Session<CodeModel>): Promise<string> {
  if (session.model.language.go!.pollerTypes === undefined) {
    return '';
  }
  let text = await contentPreamble(session);
  const pollers = <Array<PollerInfo>>session.model.language.go!.pollerTypes;
  pollers.sort((a: PollerInfo, b: PollerInfo) => { return sortAscending(a.name, b.name) });

  // add standard imports
  const imports = new ImportManager();
  imports.add('context');
  imports.add('encoding/json');
  imports.add('errors');
  imports.add('github.com/Azure/azure-sdk-for-go/sdk/azcore');
  imports.add('io/ioutil');
  imports.add('net/http');
  imports.add('net/url');
  imports.add('strings');
  imports.add('fmt');
  imports.add('errors');
  text += imports.text();
  // TODO separate this into manageable chunks of text by section of functionality
  text += `
const (
  headerAsyncOperation = "Azure-AsyncOperation"
	headerLocation       = "Location"
	)

const (
  operationInProgress string = "InProgress"
operationCanceled   string = "Canceled"
operationFailed     string = "Failed"
operationSucceeded  string = "Succeeded"
	)

var pollingCodes = [...]int{http.StatusNoContent, http.StatusAccepted, http.StatusCreated, http.StatusOK}

type pollingTracker interface {
  // these methods can differ per tracker

  // checks the response headers and status code to determine the polling mechanism
  updatePollingMethod() error

  // checks the response for tracker-specific error conditions
  checkForErrors() error

  // returns true if provisioning state should be checked
  provisioningStateApplicable() bool

  // methods common to all trackers

  // initializes a tracker's polling URL and method, called for each iteration.
  // these values can be overridden by each polling tracker as required.
  initPollingMethod() error

  // initializes the tracker's internal state, call this when the tracker is created
  initializeState() error

  // makes an HTTP request to check the status of the LRO
  pollForStatus(ctx context.Context, client azcore.Pipeline) error

  // updates internal tracker state, call this after each call to pollForStatus
  updatePollingState(provStateApl bool) error

  // returns the error response from the service, can be nil
  pollingError() error

  // returns the polling method being used
  pollingMethod() pollingMethodType

  // returns the state of the LRO as returned from the service
  pollingStatus() string

  // returns the URL used for polling status
  pollingURL() string

  // returns the URL used for the final GET to retrieve the resource
  finalGetURL() string

  // returns true if the LRO is in a terminal state
  hasTerminated() bool

  // returns true if the LRO is in a failed terminal state
  hasFailed() bool

  // returns true if the LRO is in a successful terminal state
  hasSucceeded() bool

  // returns the cached HTTP response after a call to pollForStatus(), can be nil
  latestResponse() *azcore.Response

  // converts an *azcore.Response to an error
  handleError(resp *azcore.Response) error
}

type methodErrorHandler func(resp *azcore.Response) error

type pollingTrackerBase struct {
  // resp is the last response, either from the submission of the LRO or from polling
  resp *azcore.Response

  // PollerType is the name of the type of poller that is created
  PollerType string \`json:"pollerType"\`

	// errorHandler is the method to invoke to unmarshall an error response
	errorHandler methodErrorHandler

		// method is the HTTP verb, this is needed for deserialization
		Method string \`json:"method"\`
	
		// rawBody is the raw JSON response body
		rawBody map[string]interface{}
	
		// denotes if polling is using async-operation or location header
		Pm pollingMethodType \`json:"pollingMethod"\`
	
		// the URL to poll for status
		URI string \`json:"pollingURI"\`
	
		// the state of the LRO as returned from the service
		State string \`json:"lroState"\`
	
		// the URL to GET for the final result
		FinalGetURI string \`json:"resultURI"\`
	
		// used to hold an error object returned from the service
		Err error \`json:"error,omitempty"\`
	}
	
	// done queries the service to see if the operation has completed.
	func lroPollDone(ctx context.Context, p azcore.Pipeline, pt pollingTracker) bool {
		if pt.hasTerminated() {
			return true
		}
		if err := pt.pollForStatus(ctx, p); err != nil {
			return true
		}
		if err := pt.checkForErrors(); err != nil {
			return true
		}
		if err := pt.updatePollingState(pt.provisioningStateApplicable()); err != nil {
			return true
		}
		if err := pt.initPollingMethod(); err != nil {
			return true
		}
		if err := pt.updatePollingMethod(); err != nil {
			return true
		}
		return pt.hasTerminated()
	}

	func (pt *pollingTrackerBase) initializeState() error {
		// determine the initial polling state based on response body and/or HTTP status
		// code.  this is applicable to the initial LRO response, not polling responses!
		pt.Method = pt.resp.Request.Method
		if err := pt.updateRawBody(); err != nil {
			pt.Err = err
			return err
		}
		switch pt.resp.StatusCode {
		case http.StatusOK:
			if ps := pt.getProvisioningState(); ps != nil {
				pt.State = *ps
				if pt.hasFailed() {
					pt.updateErrorFromResponse()
					return pt.pollingError()
				}
			} else {
				pt.State = operationSucceeded
			}
		case http.StatusCreated:
			if ps := pt.getProvisioningState(); ps != nil {
				pt.State = *ps
			} else {
				pt.State = operationInProgress
			}
		case http.StatusAccepted:
			pt.State = operationInProgress
		case http.StatusNoContent:
			pt.State = operationSucceeded
		default:
			pt.State = operationFailed
			pt.updateErrorFromResponse()
			return pt.pollingError()
		}
		return pt.initPollingMethod()
	}
	
	func (pt pollingTrackerBase) getProvisioningState() *string {
		if pt.rawBody != nil && pt.rawBody["properties"] != nil {
			p := pt.rawBody["properties"].(map[string]interface{})
			if ps := p["provisioningState"]; ps != nil {
				s := ps.(string)
				return &s
			}
		}
		return nil
	}
	
	func (pt *pollingTrackerBase) updateRawBody() error {
		pt.rawBody = map[string]interface{}{}
		if pt.resp.ContentLength != 0 {
			defer pt.resp.Body.Close()
			b, err := ioutil.ReadAll(pt.resp.Body)
			if err != nil {
				pt.Err = err
				return pt.Err
			}
			// observed in 204 responses over HTTP/2.0; the content length is -1 but body is empty
			if len(b) == 0 {
				return nil
			}
			if err = json.Unmarshal(b, &pt.rawBody); err != nil {
				pt.Err = err
				return pt.Err
			}
		}
		return nil
	}
	
	func (pt *pollingTrackerBase) pollForStatus(ctx context.Context, client azcore.Pipeline) error {
		u, err := url.Parse(pt.URI)
		if err != nil {
					pt.Err = err
			return err
		}
		req := azcore.NewRequest(http.MethodGet, *u)
		resp, err := client.Do(ctx, req)
		pt.resp = resp
		if err != nil {
			pt.Err = err
			return pt.Err
		}
		if pt.resp.HasStatusCode(pollingCodes[:]...) {
			// reset the service error on success case
			pt.Err = nil
			err = pt.updateRawBody()
		} else {
			// check response body for error content
			pt.updateErrorFromResponse()
			err = pt.pollingError()
		}
		return err
	}
	
	// attempts to unmarshal a ServiceError type from the response body.
	// if that fails then make a best attempt at creating something meaningful.
	// NOTE: this assumes that the async operation has failed.
	func (pt *pollingTrackerBase) updateErrorFromResponse() {
		pt.Err = pt.errorHandler(pt.resp)
	}
	
	func (pt *pollingTrackerBase) updatePollingState(provStateApl bool) error {
		if pt.Pm == pollingAsyncOperation && pt.rawBody["status"] != nil {
			pt.State = pt.rawBody["status"].(string)
		} else {
			if pt.resp.StatusCode == http.StatusAccepted {
				pt.State = operationInProgress
			} else if provStateApl {
				if ps := pt.getProvisioningState(); ps != nil {
					pt.State = *ps
				} else {
					pt.State = operationSucceeded
				}
			} else {
				pt.Err = fmt.Errorf("the response from the async operation has an invalid status code: %d", pt.resp.StatusCode)
				return pt.Err
			}
		}
		// if the operation has failed update the error state
		if pt.hasFailed() {
			pt.updateErrorFromResponse()
		}
		return nil
	}
	
	func (pt pollingTrackerBase) pollingError() error {
		return pt.Err
	}
	
	func (pt pollingTrackerBase) pollingMethod() pollingMethodType {
		return pt.Pm
	}
	
	func (pt pollingTrackerBase) pollingStatus() string {
		return pt.State
	}
	
	func (pt pollingTrackerBase) pollingURL() string {
		return pt.URI
	}
	
	func (pt pollingTrackerBase) finalGetURL() string {
		return pt.FinalGetURI
	}
	
	func (pt pollingTrackerBase) hasTerminated() bool {
		return strings.EqualFold(pt.State, operationCanceled) || strings.EqualFold(pt.State, operationFailed) || strings.EqualFold(pt.State, operationSucceeded)
	}
	
	func (pt pollingTrackerBase) hasFailed() bool {
		return strings.EqualFold(pt.State, operationCanceled) || strings.EqualFold(pt.State, operationFailed)
	}
	
	func (pt pollingTrackerBase) hasSucceeded() bool {
		return strings.EqualFold(pt.State, operationSucceeded)
	}
	
	func (pt pollingTrackerBase) latestResponse() *azcore.Response {
		return pt.resp
	}
	
	// error checking common to all trackers
	func (pt pollingTrackerBase) baseCheckForErrors() error {
		// for Azure-AsyncOperations the response body cannot be nil or empty
		if pt.Pm == pollingAsyncOperation {
			if pt.resp.Body == nil || pt.resp.ContentLength == 0 {
						pt.Err = errors.New("for Azure-AsyncOperation response body cannot be nil")
				return pt.Err
			}
			if pt.rawBody["status"] == nil {
						pt.Err = errors.New("missing status property in Azure-AsyncOperation response body")
				return pt.Err
			}
		}
		return nil
	}
	
	// default initialization of polling URL/method.  each verb tracker will update this as required.
	func (pt *pollingTrackerBase) initPollingMethod() error {
		if ao, err := getURLFromAsyncOpHeader(pt.resp); err != nil {
			pt.Err = err
			return err
		} else if ao != "" {
			pt.URI = ao
			pt.Pm = pollingAsyncOperation
			return nil
		}
		if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
					pt.Err = err
			return err
		} else if lh != "" {
			pt.URI = lh
			pt.Pm = pollingLocation
			return nil
		}
		// it's ok if we didn't find a polling header, this will be handled elsewhere
		return nil
	}
  
  func (pt *pollingTrackerBase) handleError(resp *azcore.Response) error {
    return pt.errorHandler(resp)
  }

	// DELETE
	
	type pollingTrackerDelete struct {
		pollingTrackerBase
	}

	func (pt *pollingTrackerDelete) updatePollingMethod() error {
		// for 201 the Location header is required
		if pt.resp.StatusCode == http.StatusCreated {
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
						pt.Err = err
				return err
			} else if lh == "" {
						pt.Err = errors.New("missing Location header in 201 response")
				return pt.Err
			} else {
				pt.URI = lh
			}
			pt.Pm = pollingLocation
			pt.FinalGetURI = pt.URI
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
						pt.Err = err
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			// if the Location header is invalid and we already have a polling URL
			// then we don't care if the Location header URL is malformed.
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil && pt.URI == "" {
						pt.Err = err
				return err
			} else if lh != "" {
				if ao == "" {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
				// when both headers are returned we use the value in the Location header for the final GET
				pt.FinalGetURI = lh
			}
			// make sure a polling URL was found
			if pt.URI == "" {
						pt.Err = errors.New("didn't get any suitable polling URLs in 202 response")
				return pt.Err
			}
		}
		return nil
	}
	
	func (pt pollingTrackerDelete) checkForErrors() error {
		return pt.baseCheckForErrors()
	}
	
	func (pt pollingTrackerDelete) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusNoContent
	}
	
	// PATCH
	
	type pollingTrackerPatch struct {
		pollingTrackerBase
	}

	func (pt *pollingTrackerPatch) updatePollingMethod() error {
		// by default we can use the original URL for polling and final GET
		if pt.URI == "" {
			pt.URI = pt.resp.Request.URL.String()
		}
		if pt.FinalGetURI == "" {
			pt.FinalGetURI = pt.resp.Request.URL.String()
		}
		if pt.Pm == pollingUnknown {
			pt.Pm = pollingRequestURI
		}
		// for 201 it's permissible for no headers to be returned
		if pt.resp.StatusCode == http.StatusCreated {
			if ao, err := getURLFromAsyncOpHeader(pt.resp); err != nil {
						pt.Err = err
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		// note the absence of the "final GET" mechanism for PATCH
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
						pt.Err = err
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			if ao == "" {
				if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
						 pt.Err = err
					return err
				} else if lh == "" {
								pt.Err = errors.New("didn't get any suitable polling URLs in 202 response")
					return pt.Err
				} else {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
			}
		}
		return nil
	}
	
	func (pt pollingTrackerPatch) checkForErrors() error {
		return pt.baseCheckForErrors()
	}
	
	func (pt pollingTrackerPatch) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusCreated
	}
	
	// POST
	
	type pollingTrackerPost struct {
		pollingTrackerBase
	}
	
	func (pt *pollingTrackerPost) updatePollingMethod() error {
		// 201 requires Location header
		if pt.resp.StatusCode == http.StatusCreated {
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil {
						pt.Err = err
				return err
			} else if lh == "" {
						pt.Err = errors.New("missing Location header in 201 response")
				return pt.Err
			} else {
				pt.URI = lh
				pt.FinalGetURI = lh
				pt.Pm = pollingLocation
			}
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
						pt.Err = err
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			// if the Location header is invalid and we already have a polling URL
			// then we don't care if the Location header URL is malformed.
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil && pt.URI == "" {
						pt.Err = err
				return err
			} else if lh != "" {
				if ao == "" {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
				// when both headers are returned we use the value in the Location header for the final GET
				pt.FinalGetURI = lh
			}
			// make sure a polling URL was found
			if pt.URI == "" {
						pt.Err = errors.New("didn't get any suitable polling URLs in 202 response")
				return pt.Err
			}
		}
		return nil
	}
	
	func (pt pollingTrackerPost) checkForErrors() error {
		return pt.baseCheckForErrors()
	}
	
	func (pt pollingTrackerPost) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusNoContent
	}
	
	// PUT
	
	type pollingTrackerPut struct {
		pollingTrackerBase
	}

	func (pt *pollingTrackerPut) updatePollingMethod() error {
		// by default we can use the original URL for polling and final GET
		if pt.URI == "" {
			pt.URI = pt.resp.Request.URL.String()
		}
		if pt.FinalGetURI == "" {
			pt.FinalGetURI = pt.resp.Request.URL.String()
		}
		if pt.Pm == pollingUnknown {
			pt.Pm = pollingRequestURI
		}
		// for 201 it's permissible for no headers to be returned
		if pt.resp.StatusCode == http.StatusCreated {
			if ao, err := getURLFromAsyncOpHeader(pt.resp); err != nil {
						pt.Err = err
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
		}
		// for 202 prefer the Azure-AsyncOperation header but fall back to Location if necessary
		if pt.resp.StatusCode == http.StatusAccepted {
			ao, err := getURLFromAsyncOpHeader(pt.resp)
			if err != nil {
						pt.Err = err
				return err
			} else if ao != "" {
				pt.URI = ao
				pt.Pm = pollingAsyncOperation
			}
			// if the Location header is invalid and we already have a polling URL
			// then we don't care if the Location header URL is malformed.
			if lh, err := getURLFromLocationHeader(pt.resp); err != nil && pt.URI == "" {
						pt.Err = err
				return err
			} else if lh != "" {
				if ao == "" {
					pt.URI = lh
					pt.Pm = pollingLocation
				}
			}
			// make sure a polling URL was found
			if pt.URI == "" {
						pt.Err = errors.New("didn't get any suitable polling URLs in 202 response")
				return pt.Err
			}
		}
		return nil
	}
	
	func (pt pollingTrackerPut) checkForErrors() error {
		err := pt.baseCheckForErrors()
		if err != nil {
					pt.Err = err
			return err
		}
		// if there are no LRO headers then the body cannot be empty
		ao, err := getURLFromAsyncOpHeader(pt.resp)
		if err != nil {
			pt.Err = err
			return err
		}
		lh, err := getURLFromLocationHeader(pt.resp)
		if err != nil {
			pt.Err = err
			return err
		}
		if ao == "" && lh == "" && len(pt.rawBody) == 0 {
			pt.Err = errors.New("the response did not contain a body")
			return pt.Err
		}
		return nil
	}
	
	func (pt pollingTrackerPut) provisioningStateApplicable() bool {
		return pt.resp.StatusCode == http.StatusOK || pt.resp.StatusCode == http.StatusCreated
	}
	
	// creates a polling tracker based on the verb of the original request
	func createPollingTracker(pollerType string, resp *azcore.Response, errorHandler methodErrorHandler) (pollingTracker, error) {
		var pt pollingTracker
		switch strings.ToUpper(resp.Request.Method) {
		case http.MethodDelete:
			pt = &pollingTrackerDelete{pollingTrackerBase: pollingTrackerBase{PollerType: pollerType, resp: resp, errorHandler: errorHandler}}
		case http.MethodPatch:
			pt = &pollingTrackerPatch{pollingTrackerBase: pollingTrackerBase{PollerType: pollerType, resp: resp, errorHandler: errorHandler}}
		case http.MethodPost:
			pt = &pollingTrackerPost{pollingTrackerBase: pollingTrackerBase{PollerType: pollerType, resp: resp, errorHandler: errorHandler}}
		case http.MethodPut:
			pt = &pollingTrackerPut{pollingTrackerBase: pollingTrackerBase{PollerType: pollerType, resp: resp, errorHandler: errorHandler}}
		default:
			return nil, fmt.Errorf("unsupported HTTP method %s", resp.Request.Method)
		}
		if err := pt.initializeState(); err != nil {
			return pt, err
		}
		// this initializes the polling header values, we do this during creation in case the
		// initial response send us invalid values; this way the API call will return a non-nil
		// error (not doing this means the error shows up in Future.Done)
		return pt, pt.updatePollingMethod()
	}

// creates a polling tracker from a resume token
func resumePollingTracker(pollerType string, token string, errorHandler methodErrorHandler) (pollingTracker, error) {
	// unmarshal into JSON object to determine the tracker type
	obj := map[string]interface{}{}
	err := json.Unmarshal([]byte(token), &obj)
	if err != nil {
		return nil, err
  }
  if obj["pollerType"] != pollerType {
		return nil, fmt.Errorf("cannot resume from this poller type")
	}
	if obj["method"] == nil {
		return nil, fmt.Errorf("token is missing 'method' property")
	}
	var pt pollingTracker
	method := obj["method"].(string)
	switch strings.ToUpper(method) {
	case http.MethodDelete:
		pt = &pollingTrackerDelete{pollingTrackerBase: pollingTrackerBase{errorHandler: errorHandler}}
	case http.MethodPatch:
		pt = &pollingTrackerPatch{pollingTrackerBase: pollingTrackerBase{errorHandler: errorHandler}}
	case http.MethodPost:
		pt = &pollingTrackerPost{pollingTrackerBase: pollingTrackerBase{errorHandler: errorHandler}}
	case http.MethodPut:
		pt = &pollingTrackerPut{pollingTrackerBase: pollingTrackerBase{errorHandler: errorHandler}}
	default:
		return nil, fmt.Errorf("unsupported method '%s'", method)
	}
	// now unmarshal into the tracker
	err = json.Unmarshal([]byte(token), &pt)
	if err != nil {
		return nil, err
	}
	return pt, nil
}

	// gets the polling URL from the Azure-AsyncOperation header.
	// ensures the URL is well-formed and absolute.
	func getURLFromAsyncOpHeader(resp *azcore.Response) (string, error) {
		s := resp.Header.Get(http.CanonicalHeaderKey(headerAsyncOperation))
		if s == "" {
			return "", nil
		}
		if !isValidURL(s) {
			return "", fmt.Errorf("invalid polling URL '%s'", s)
		}
		return s, nil
	}
	
	// gets the polling URL from the Location header.
	// ensures the URL is well-formed and absolute.
	func getURLFromLocationHeader(resp *azcore.Response) (string, error) {
		s := resp.Header.Get(http.CanonicalHeaderKey(headerLocation))
		if s == "" {
			return "", nil
		}
		if !isValidURL(s) {
			return "", fmt.Errorf("invalid polling URL '%s'", s)
		}
		return s, nil
	}
	
	// verify that the URL is valid and absolute
	func isValidURL(s string) bool {
		u, err := url.Parse(s)
		return err == nil && u.IsAbs()
	}
	
	// pollingMethodType defines a type used for enumerating polling mechanisms.
	type pollingMethodType string
	
	const (
		// pollingAsyncOperation indicates the polling method uses the Azure-AsyncOperation header.
		pollingAsyncOperation pollingMethodType = "AsyncOperation"
	
		// pollingLocation indicates the polling method uses the Location header.
		pollingLocation pollingMethodType = "Location"
	
		// pollingRequestURI indicates the polling method uses the original request URI.
		pollingRequestURI pollingMethodType = "RequestURI"
	
		// pollingUnknown indicates an unknown polling method and is the default value.
		pollingUnknown pollingMethodType = ""
	)
	`;
  return text;
}

