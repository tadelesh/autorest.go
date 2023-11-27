/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import * as go from '../gocodemodel/gocodemodel';
import { values } from '@azure-tools/linq';
import { capitalize, comment, uncapitalize } from '@azure-tools/codegen';
import { ImportManager } from './imports';

// variable to be used to determine comment length when calling comment from @azure-tools
export const commentLength = 120;

export const dateFormat = '2006-01-02';
export const datetimeRFC3339Format = 'time.RFC3339Nano';
export const datetimeRFC1123Format = 'time.RFC1123';
export const timeRFC3339Format = '15:04:05.999999999Z07:00';

// returns the common source-file preamble (license comment, package name etc)
export function contentPreamble(codeModel: go.CodeModel, packageName?: string): string {
  if (!packageName) {
    packageName = codeModel.packageName;
  }
  const headerText = comment(codeModel.options.headerText, '// ');
  let text = '//go:build go1.18\n// +build go1.18\n\n';
  // ensure tools recognize the file as generated according to
  // https://pkg.go.dev/cmd/go#hdr-Generate_Go_files_by_processing_source
  text += headerText.replace(/^\/\/ Code generated .*\.$/m, '$& DO NOT EDIT.');
  if (!text.match(/^\/\/ Code generated .* DO NOT EDIT\.$/m)) {
    text += '\n// Code generated by @autorest/go. DO NOT EDIT.';
  }
  text += `\n\npackage ${packageName}\n\n`;
  return text;
}

// used to sort strings in ascending order
export function sortAscending(a: string, b: string): number {
  return a < b ? -1 : a > b ? 1 : 0;
}

export function isParameter(param: go.Parameter | go.ParameterGroup): param is go.Parameter {
  return (<go.ParameterGroup>param).groupName === undefined;
}

export function isParameterGroup(param: go.Parameter | go.ParameterGroup): param is go.ParameterGroup {
  return (<go.ParameterGroup>param).groupName !== undefined;
}

export function isRequiredParameter(param: go.Parameter): boolean {
  // parameters with a client-side default value are always optional
  if (go.isClientSideDefault(param.paramType)) {
    return false;
  }
  return param.paramType === 'required';
}

export function isLiteralParameter(param: go.Parameter): boolean {
  if (go.isClientSideDefault(param.paramType)) {
    return false;
  }
  return param.paramType === 'literal';
}

// returns the type name with possible * prefix
export function formatParameterTypeName(param: go.Parameter | go.ParameterGroup, pkgName?: string): string {
  let typeName: string;
  if (isParameterGroup(param)) {
    typeName = param.groupName;
    if (pkgName) {
      typeName = `${pkgName}.${typeName}`;
    }
    if (param.required) {
      return typeName;
    }
  } else {
    typeName = go.getTypeDeclaration(param.type, pkgName);
    if (isRequiredParameter(param) || (param.location === 'client' && go.isClientSideDefault(param.paramType))) {
      // client parameters with default values aren't emitted as pointer-to-type
      return typeName;
    }
  }
  return `*${typeName}`;
}

// sorts parameters by their required state, ordering required before optional
export function sortParametersByRequired(a: go.Parameter | go.ParameterGroup, b: go.Parameter | go.ParameterGroup): number {
  let aRequired = false;
  let bRequired = false;

  if (isParameter(a)) {
    aRequired = isRequiredParameter(a);
  } else {
    aRequired = a.required;
  }

  if (isParameter(b)) {
    bRequired = isRequiredParameter(b);
  } else {
    bRequired = b.required;
  }

  if (aRequired === bRequired) {
    return 0;
  } else if (aRequired && !bRequired) {
    return -1;
  }
  return 1;
}

// returns the parameters for the internal request creator method.
// e.g. "i int, s string"
export function getCreateRequestParametersSig(method: go.Method | go.NextPageMethod): string {
  const methodParams = getMethodParameters(method);
  const params = new Array<string>();
  params.push('ctx context.Context');
  for (const methodParam of values(methodParams)) {
    params.push(`${uncapitalize(methodParam.paramName)} ${formatParameterTypeName(methodParam)}`);
  }
  return params.join(', ');
}

// returns the parameter names for an operation (excludes the param types).
// e.g. "i, s"
export function getCreateRequestParameters(method: go.Method): string {
  // split param list into individual params
  const reqParams = getCreateRequestParametersSig(method).split(',');
  // keep the parameter names from the name/type tuples
  for (let i = 0; i < reqParams.length; ++i) {
    reqParams[i] = reqParams[i].trim().split(' ')[0];
  }
  return reqParams.join(', ');
}

// returns the complete collection of method parameters
export function getMethodParameters(method: go.Method | go.NextPageMethod, paramsFilter?: (p: Array<go.Parameter>) => Array<go.Parameter>): Array<go.Parameter | go.ParameterGroup> {
  const params = new Array<go.Parameter>();
  const paramGroups = new Array<go.ParameterGroup>();
  let methodParams = method.parameters;
  if (paramsFilter) {
    methodParams = paramsFilter(methodParams);
  }
  for (const param of values(methodParams)) {
    if (param.location === 'client') {
      // client params are passed via the receiver
      // must check before param group as client params can be grouped
      continue;
    } else if (param.group) {
      // param groups will be added after individual params
      if (!paramGroups.includes(param.group)) {
        paramGroups.push(param.group);
      }
    } else if (go.isLiteralValue(param.type)) {
      // don't generate a parameter for a constant
      // NOTE: this check must come last as non-required optional constants
      // in header/query params get dumped into the optional params group
      continue;
    } else {
      params.push(param);
    }
  }
  // move global optional params to the end of the slice
  params.sort(sortParametersByRequired);
  // add any parameter groups.  optional groups go last
  paramGroups.sort((a: go.ParameterGroup, b: go.ParameterGroup) => {
    if (a.required === b.required) {
      return 0;
    }
    if (a.required && !b.required) {
      return -1;
    }
    return 1;
  });
  // add the optional param group last if it's not already in the list.
  if (go.isMethod(method)) {
    if (!values(paramGroups).any(gp => { return gp.groupName === method.optionalParamsGroup.groupName; })) {
      paramGroups.push(method.optionalParamsGroup);
    }
  }
  const combined = new Array<go.Parameter | go.ParameterGroup>();
  for (const param of params) {
    combined.push(param);
  }
  for (const paramGroup of paramGroups) {
    combined.push(paramGroup);
  }
  return combined;
}

// returns the fully-qualified parameter name.  this is usually just the name
// but will include the client or optional param group name prefix as required.
export function getParamName(param: go.Parameter): string {
  let paramName = param.paramName;
  // must check paramGroup first as client params can also be grouped
  if (param.group) {
    paramName = `${uncapitalize(param.group.paramName)}.${capitalize(paramName)}`;
  }
  if (param.location === 'client') {
    paramName = `client.${paramName}`;
  }
  // client parameters with default values aren't emitted as pointer-to-type
  if (!isRequiredParameter(param) && !(param.location === 'client' && go.isClientSideDefault(param.paramType)) && !(isParameter(param) && param.byValue)) {
    paramName = `*${paramName}`;
  }
  return paramName;
}

export function formatParamValue(param: go.FormBodyParameter | go.HeaderParameter | go.PathParameter | go.QueryParameter, imports: ImportManager): string {
  let paramName = getParamName(param);
  if (go.isFormBodyCollectionParameter(param) || go.isHeaderCollectionParameter(param) || go.isPathCollectionParameter(param) || go.isQueryCollectionParameter(param)) {
    if (param.collectionFormat === 'multi') {
      throw new Error('multi collection format should have been previously handled');
    }
    const separator = getDelimiterForCollectionFormat(param.collectionFormat);
    if (go.isPrimitiveType(param.type.elementType) && param.type.elementType.typeName === 'string') {
      imports.add('strings');
      return `strings.Join(${paramName}, "${separator}")`;
    } else {
      imports.add('fmt');
      imports.add('strings');
      return `strings.Join(strings.Fields(strings.Trim(fmt.Sprint(${paramName}), "[]")), "${separator}")`;
    }
  } else if (go.isTimeType(param.type) && param.type.dateTimeFormat !== 'timeUnix') {
    // for most time types we call methods on time.Time which is why we remove the dereference.
    // however, for unix time, we cast to our unixTime helper first so we must keep the dereference.
    if (!isRequiredParameter(param) && paramName[0] === '*') {
      // remove the dereference
      paramName = paramName.substring(1);
    }
  }
  return formatValue(paramName, param.type, imports);
}

export function getDelimiterForCollectionFormat(cf: go.CollectionFormat): string {
  switch (cf) {
    case 'csv':
      return ',';
    case 'pipes':
      return '|';
    case 'ssv':
      return ' ';
    case 'tsv':
      return '\\t';
    default:
      throw new Error(`unhandled CollectionFormat ${cf}`);
  }
}

export function formatValue(paramName: string, type: go.PossibleType, imports: ImportManager, defef?: boolean): string {
  // callers don't have enough context to know if paramName needs to be
  // deferenced so we track that here when specified. note that not all
  // cases will require paramName to be dereferenced.
  let star = '';
  if (defef === true) {
    star = '*';
  }
  if (go.isLiteralValue(type)) {
    // cannot use formatLiteralValue() since all values are treated as strings
    return `"${type.literal}"`;
  } else if (go.isBytesType(type)) {
    // ByteArray is a base-64 encoded value in string format
    imports.add('encoding/base64');
    let byteFormat = 'Std';
    if (type.encoding === 'URL') {
      byteFormat = 'RawURL';
    }
    return `base64.${byteFormat}Encoding.EncodeToString(${paramName})`;
  } else if (go.isPrimitiveType(type)) {
    if (type.typeName === 'bool') {
      imports.add('strconv');
      return `strconv.FormatBool(${star}${paramName})`;
    } else if (type.typeName === 'int32') {
      imports.add('strconv');
      return `strconv.FormatInt(int64(${star}${paramName}), 10)`;
    } else if (type.typeName === 'int64') {
      imports.add('strconv');
      return `strconv.FormatInt(${star}${paramName}, 10)`;
    } else if (type.typeName === 'float32') {
      imports.add('strconv');
      return `strconv.FormatFloat(float64(${star}${paramName}), 'f', -1, 32)`;
    } else if (type.typeName === 'float64') {
      imports.add('strconv');
      return `strconv.FormatFloat(${star}${paramName}, 'f', -1, 64)`;
    }
  } else if (go.isTimeType(type)) {
    if (type.dateTimeFormat === 'dateType') {
      return `${paramName}.Format("${dateFormat}")`;
    } else if (type.dateTimeFormat === 'timeUnix') {
      return `timeUnix(${star}${paramName}).String()`;
    } else if (type.dateTimeFormat === 'timeRFC3339') {
      return `timeRFC3339(${star}${paramName}).String()`;
    } else {
      imports.add('time');
      let format = datetimeRFC3339Format;
      if (type.dateTimeFormat === 'dateTimeRFC1123') {
        format = datetimeRFC1123Format;
      }
      return `${paramName}.Format(${format})`;
    }
  } else if (go.isConstantType(type)) {
    if (type.type === 'string') {
      return `string(${star}${paramName})`;
    }
    imports.add('fmt');
    return `fmt.Sprintf("%v", ${star}${paramName})`;
  }
  return `${star}${paramName}`;
}

// returns the clientDefaultValue of the specified param.
// this is usually the value in quotes (i.e. a string) however
// it could also be a constant.
export function formatLiteralValue(value: go.LiteralValue): string {
  if (go.isConstantType(value.type)) {
    return (<go.ConstantValue>value.literal).valueName;
  } else if (go.isPrimitiveType(value.type)) {
    switch (value.type.typeName) {
      case 'float32':
        return `float32(${value.literal})`;
      case 'float64':
        return `float64(${value.literal})`;
      case 'int32':
        return `int32(${value.literal})`;
      case 'int64':
        return `int64(${value.literal})`;
      case 'string':
        if (value.literal[0] === '"') {
          // string is already quoted
          return value.literal;
        }
        return `"${value.literal}"`;
      default:
        return value.literal;
    }
  } else if (go.isTimeType(value.type)) {
    return `"${value.literal}"`;
  }
  return value.literal;
}

// returns true if at least one of the responses has a schema
export function hasSchemaResponse(method: go.Method): boolean {
  const result = method.responseEnvelope.result;
  if (!result) {
    return false;
  }
  return go.isAnyResult(result) || go.isMonomorphicResult(result) || go.isPolymorphicResult(result) || go.isModelResult(result);
}

// returns the name of the response field within the response envelope
export function getResultFieldName(method: go.Method): string {
  const result = method.responseEnvelope.result;
  if (!result) {
    throw new Error(`missing result for method ${method.methodName}`);
  } else if (go.isAnyResult(result) || go.isBinaryResult(result) || go.isHeadAsBooleanResult(result) || go.isMonomorphicResult(result)) {
    return result.fieldName;
  } else if (go.isPolymorphicResult(result)) {
    return result.interfaceType.name;
  } else if (go.isModelResult(result)) {
    return result.modelType.name;
  } else {
    throw new Error(`unhandled result type for method ${method.client.clientName}.${method.methodName}`);
  }
}

export function formatStatusCodes(statusCodes: Array<number>): string {
  const asHTTPStatus = new Array<string>();
  for (const rawCode of statusCodes) {
    asHTTPStatus.push(formatStatusCode(rawCode));
  }
  return asHTTPStatus.join(', ');
}

export function formatStatusCode(statusCode: number): string {
  switch (statusCode) {
    // 1xx
    case 100:
      return 'http.StatusContinue';
    case 101:
      return 'http.StatusSwitchingProtocols';
    case 102:
      return 'http.StatusProcessing';
    case 103:
      return 'http.StatusEarlyHints';
    // 2xx
    case 200:
      return 'http.StatusOK';
    case 201:
      return 'http.StatusCreated';
    case 202:
      return 'http.StatusAccepted';
    case 203:
      return 'http.StatusNonAuthoritativeInfo';
    case 204:
      return 'http.StatusNoContent';
    case 205:
      return 'http.StatusResetContent';
    case 206:
      return 'http.StatusPartialContent';
    case 207:
      return 'http.StatusMultiStatus';
    case 208:
      return 'http.StatusAlreadyReported';
    case 226:
      return 'http.StatusIMUsed';
    // 3xx
    case 300:
      return 'http.StatusMultipleChoices';
    case 301:
      return 'http.StatusMovedPermanently';
    case 302:
      return 'http.StatusFound';
    case 303:
      return 'http.StatusSeeOther';
    case 304:
      return 'http.StatusNotModified';
    case 305:
      return 'http.StatusUseProxy';
    case 307:
      return 'http.StatusTemporaryRedirect';
    // 4xx
    case 400:
      return 'http.StatusBadRequest';
    case 401:
      return 'http.StatusUnauthorized';
    case 402:
      return 'http.StatusPaymentRequired';
    case 403:
      return 'http.StatusForbidden';
    case 404:
      return 'http.StatusNotFound';
    case 405:
      return 'http.StatusMethodNotAllowed';
    case 406:
      return 'http.StatusNotAcceptable';
    case 407:
      return 'http.StatusProxyAuthRequired';
    case 408:
      return 'http.StatusRequestTimeout';
    case 409:
      return 'http.StatusConflict';
    case 410:
      return 'http.StatusGone';
    case 411:
      return 'http.StatusLengthRequired';
    case 412:
      return 'http.StatusPreconditionFailed';
    case 413:
      return 'http.StatusRequestEntityTooLarge';
    case 414:
      return 'http.StatusRequestURITooLong';
    case 415:
      return 'http.StatusUnsupportedMediaType';
    case 416:
      return 'http.StatusRequestedRangeNotSatisfiable';
    case 417:
      return 'http.StatusExpectationFailed';
    case 418:
      return 'http.StatusTeapot';
    case 421:
      return 'http.StatusMisdirectedRequest';
    case 422:
      return 'http.StatusUnprocessableEntity';
    case 423:
      return 'http.StatusLocked';
    case 424:
      return 'http.StatusFailedDependency';
    case 425:
      return 'http.StatusTooEarly';
    case 426:
      return 'http.StatusUpgradeRequired';
    case 428:
      return 'http.StatusPreconditionRequired';
    case 429:
      return 'http.StatusTooManyRequests';
    case 431:
      return 'http.StatusRequestHeaderFieldsTooLarge';
    case 451:
      return 'http.StatusUnavailableForLegalReasons';
    // 5xx
    case 500:
      return 'http.StatusInternalServerError';
    case 501:
      return 'http.StatusNotImplemented';
    case 502:
      return 'http.StatusBadGateway';
    case 503:
      return 'http.StatusServiceUnavailable';
    case 504:
      return 'http.StatusGatewayTimeout ';
    case 505:
      return 'http.StatusHTTPVersionNotSupported';
    case 506:
      return 'http.StatusVariantAlsoNegotiates';
    case 507:
      return 'http.StatusInsufficientStorage';
    case 508:
      return 'http.StatusLoopDetected';
    case 510:
      return 'http.StatusNotExtended';
    case 511:
      return 'http.StatusNetworkAuthenticationRequired';
    default:
      throw new Error(`unhandled status code ${statusCode}`);
  }
}

export function formatCommentAsBulletItem(description: string): string {
  // first create the comment block. note that it can be multi-line depending on length:
  //
  // some comment first line
  // and it finishes here.
  description = comment(description, '//', undefined, commentLength);

  // transform the above to look like this:
  //
  //   - some comment first line
  //     and it finishes here.
  const chunks = description.split('\n');
  for (let i = 0; i < chunks.length; ++i) {
    if (i === 0) {
      chunks[i] = chunks[i].replace('// ', '//   - ');
    } else {
      chunks[i] = chunks[i].replace('// ', '//     ');
    }
  }
  return chunks.join('\n');
}

export function getParentImport(codeModel: go.CodeModel): string {
  const clientPkg = codeModel.packageName;
  if (codeModel.options.module) {
    return codeModel.options.module;
  } else if (codeModel.options.containingModule) {
    return codeModel.options.containingModule + '/' + clientPkg;
  } else {
    throw new Error('unable to determine containing module for fakes. specify either the module or containing-module switch');
  }
}
