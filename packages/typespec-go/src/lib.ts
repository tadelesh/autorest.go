/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Microsoft Corporation. All rights reserved.
 *  Licensed under the MIT License. See License.txt in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

import { createTypeSpecLibrary, JSONSchemaType } from '@typespec/compiler';

export interface GoEmitterOptions {
  'azcore-version'?: string;
  'disallow-unknown-fields'?: boolean;
  'file-prefix'?: string;
  'generate-fakes'?: boolean;
  'inject-spans'?: boolean;
  'module'?: string;
  'module-version'?: string;
  'package-name'?: string;
  'rawjson-as-bytes'?: boolean;
  'slice-elements-byval'?: boolean;
  'single-client'?: boolean;
}

const EmitterOptionsSchema: JSONSchemaType<GoEmitterOptions> = {
  type: 'object',
  additionalProperties: true,
  properties: {
    'azcore-version': { type: 'string', nullable: true },
    'disallow-unknown-fields': { type: 'boolean', nullable: true },
    'file-prefix': { type: 'string', nullable: true },
    'generate-fakes': { type: 'boolean', nullable: true },
    'inject-spans': { type: 'boolean', nullable: true },
    'module': { type: 'string', nullable: true },
    'module-version': { type: 'string', nullable: true },
    'package-name': { type: 'string', nullable: true },
    'rawjson-as-bytes': { type: 'boolean', nullable: true },
    'slice-elements-byval': { type: 'boolean', nullable: true },
    'single-client': { type: 'boolean', nullable: true }
  },
  required: [],
};

const libDef = {
  name: '@azure-tools/typespec-go',
  diagnostics: {},
  emitter: {
    options: <JSONSchemaType<GoEmitterOptions>>EmitterOptionsSchema,
  },
} as const;

export const $lib = createTypeSpecLibrary(libDef);
export const { reportDiagnostic, createStateSymbol, getTracer } = $lib;
