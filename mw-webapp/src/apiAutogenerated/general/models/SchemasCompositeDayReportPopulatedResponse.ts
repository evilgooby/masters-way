// @ts-nocheck
/* eslint-disable */
/**
 * Masters way general API
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { SchemasCommentPopulatedResponse } from './SchemasCommentPopulatedResponse';
import {
    SchemasCommentPopulatedResponseFromJSON,
    SchemasCommentPopulatedResponseFromJSONTyped,
    SchemasCommentPopulatedResponseToJSON,
} from './SchemasCommentPopulatedResponse';
import type { SchemasDayReportsCompositionParticipants } from './SchemasDayReportsCompositionParticipants';
import {
    SchemasDayReportsCompositionParticipantsFromJSON,
    SchemasDayReportsCompositionParticipantsFromJSONTyped,
    SchemasDayReportsCompositionParticipantsToJSON,
} from './SchemasDayReportsCompositionParticipants';
import type { SchemasJobDonePopulatedResponse } from './SchemasJobDonePopulatedResponse';
import {
    SchemasJobDonePopulatedResponseFromJSON,
    SchemasJobDonePopulatedResponseFromJSONTyped,
    SchemasJobDonePopulatedResponseToJSON,
} from './SchemasJobDonePopulatedResponse';
import type { SchemasPlanPopulatedResponse } from './SchemasPlanPopulatedResponse';
import {
    SchemasPlanPopulatedResponseFromJSON,
    SchemasPlanPopulatedResponseFromJSONTyped,
    SchemasPlanPopulatedResponseToJSON,
} from './SchemasPlanPopulatedResponse';
import type { SchemasProblemPopulatedResponse } from './SchemasProblemPopulatedResponse';
import {
    SchemasProblemPopulatedResponseFromJSON,
    SchemasProblemPopulatedResponseFromJSONTyped,
    SchemasProblemPopulatedResponseToJSON,
} from './SchemasProblemPopulatedResponse';

/**
 * 
 * @export
 * @interface SchemasCompositeDayReportPopulatedResponse
 */
export interface SchemasCompositeDayReportPopulatedResponse {
    /**
     * 
     * @type {Array<SchemasCommentPopulatedResponse>}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    comments: Array<SchemasCommentPopulatedResponse>;
    /**
     * 
     * @type {Array<SchemasDayReportsCompositionParticipants>}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    compositionParticipants: Array<SchemasDayReportsCompositionParticipants>;
    /**
     * Calculated by - just date
     * @type {string}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    createdAt: string;
    /**
     * 
     * @type {Array<SchemasJobDonePopulatedResponse>}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    jobsDone: Array<SchemasJobDonePopulatedResponse>;
    /**
     * 
     * @type {Array<SchemasPlanPopulatedResponse>}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    plans: Array<SchemasPlanPopulatedResponse>;
    /**
     * 
     * @type {Array<SchemasProblemPopulatedResponse>}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    problems: Array<SchemasProblemPopulatedResponse>;
    /**
     * Calculated by - just last date
     * @type {string}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    updatedAt: string;
    /**
     * Always generated
     * @type {string}
     * @memberof SchemasCompositeDayReportPopulatedResponse
     */
    uuid: string;
}

/**
 * Check if a given object implements the SchemasCompositeDayReportPopulatedResponse interface.
 */
export function instanceOfSchemasCompositeDayReportPopulatedResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "comments" in value;
    isInstance = isInstance && "compositionParticipants" in value;
    isInstance = isInstance && "createdAt" in value;
    isInstance = isInstance && "jobsDone" in value;
    isInstance = isInstance && "plans" in value;
    isInstance = isInstance && "problems" in value;
    isInstance = isInstance && "updatedAt" in value;
    isInstance = isInstance && "uuid" in value;

    return isInstance;
}

export function SchemasCompositeDayReportPopulatedResponseFromJSON(json: any): SchemasCompositeDayReportPopulatedResponse {
    return SchemasCompositeDayReportPopulatedResponseFromJSONTyped(json, false);
}

export function SchemasCompositeDayReportPopulatedResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasCompositeDayReportPopulatedResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'comments': ((json['comments'] as Array<any>).map(SchemasCommentPopulatedResponseFromJSON)),
        'compositionParticipants': ((json['compositionParticipants'] as Array<any>).map(SchemasDayReportsCompositionParticipantsFromJSON)),
        'createdAt': json['createdAt'],
        'jobsDone': ((json['jobsDone'] as Array<any>).map(SchemasJobDonePopulatedResponseFromJSON)),
        'plans': ((json['plans'] as Array<any>).map(SchemasPlanPopulatedResponseFromJSON)),
        'problems': ((json['problems'] as Array<any>).map(SchemasProblemPopulatedResponseFromJSON)),
        'updatedAt': json['updatedAt'],
        'uuid': json['uuid'],
    };
}


export function SchemasCompositeDayReportPopulatedResponseToJSON(value?: SchemasCompositeDayReportPopulatedResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'comments': ((value.comments as Array<any>).map(SchemasCommentPopulatedResponseToJSON)),
        'compositionParticipants': ((value.compositionParticipants as Array<any>).map(SchemasDayReportsCompositionParticipantsToJSON)),
        'createdAt': value.createdAt,
        'jobsDone': ((value.jobsDone as Array<any>).map(SchemasJobDonePopulatedResponseToJSON)),
        'plans': ((value.plans as Array<any>).map(SchemasPlanPopulatedResponseToJSON)),
        'problems': ((value.problems as Array<any>).map(SchemasProblemPopulatedResponseToJSON)),
        'updatedAt': value.updatedAt,
        'uuid': value.uuid,
    };
}

