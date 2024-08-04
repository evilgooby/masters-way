// @ts-nocheck
/* eslint-disable */
/**
 * 
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { SchemasUserResponse } from './SchemasUserResponse';
import {
    SchemasUserResponseFromJSON,
    SchemasUserResponseFromJSONTyped,
    SchemasUserResponseToJSON,
} from './SchemasUserResponse';

/**
 * 
 * @export
 * @interface SchemasRoomPreviewResponse
 */
export interface SchemasRoomPreviewResponse {
    /**
     * 
     * @type {string}
     * @memberof SchemasRoomPreviewResponse
     */
    imageUrl: string;
    /**
     * 
     * @type {boolean}
     * @memberof SchemasRoomPreviewResponse
     */
    isBlocked: boolean;
    /**
     * 
     * @type {string}
     * @memberof SchemasRoomPreviewResponse
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasRoomPreviewResponse
     */
    roomId: string;
    /**
     * 
     * @type {string}
     * @memberof SchemasRoomPreviewResponse
     */
    roomType: string;
    /**
     * 
     * @type {Array<SchemasUserResponse>}
     * @memberof SchemasRoomPreviewResponse
     */
    users: Array<SchemasUserResponse>;
}

/**
 * Check if a given object implements the SchemasRoomPreviewResponse interface.
 */
export function instanceOfSchemasRoomPreviewResponse(
    value: object
): boolean {
    let isInstance = true;
    isInstance = isInstance && "imageUrl" in value;
    isInstance = isInstance && "isBlocked" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "roomId" in value;
    isInstance = isInstance && "roomType" in value;
    isInstance = isInstance && "users" in value;

    return isInstance;
}

export function SchemasRoomPreviewResponseFromJSON(json: any): SchemasRoomPreviewResponse {
    return SchemasRoomPreviewResponseFromJSONTyped(json, false);
}

export function SchemasRoomPreviewResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): SchemasRoomPreviewResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'imageUrl': json['imageUrl'],
        'isBlocked': json['isBlocked'],
        'name': json['name'],
        'roomId': json['roomId'],
        'roomType': json['roomType'],
        'users': ((json['users'] as Array<any>).map(SchemasUserResponseFromJSON)),
    };
}


export function SchemasRoomPreviewResponseToJSON(value?: SchemasRoomPreviewResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'imageUrl': value.imageUrl,
        'isBlocked': value.isBlocked,
        'name': value.name,
        'roomId': value.roomId,
        'roomType': value.roomType,
        'users': ((value.users as Array<any>).map(SchemasUserResponseToJSON)),
    };
}

