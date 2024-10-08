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


import * as runtime from '../runtime';
import type {
  SchemasCreateWayCollectionWay,
  SchemasWayCollectionWayResponse,
} from '../models/index';
import {
    SchemasCreateWayCollectionWayFromJSON,
    SchemasCreateWayCollectionWayToJSON,
    SchemasWayCollectionWayResponseFromJSON,
    SchemasWayCollectionWayResponseToJSON,
} from '../models/index';

export interface CreateWayCollectionWayRequest {
    request: SchemasCreateWayCollectionWay;
}

export interface DeleteWayCollectionWayRequest {
    wayCollectionId: string;
    wayId: string;
}

/**
 * 
 */
export class WayCollectionWayApi extends runtime.BaseAPI {

    /**
     * Create a new wayCollectionWay
     */
    async createWayCollectionWayRaw(requestParameters: CreateWayCollectionWayRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasWayCollectionWayResponse>> {
        if (requestParameters.request === null || requestParameters.request === undefined) {
            throw new runtime.RequiredError('request','Required parameter requestParameters.request was null or undefined when calling createWayCollectionWay.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/wayCollectionWays`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SchemasCreateWayCollectionWayToJSON(requestParameters.request),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasWayCollectionWayResponseFromJSON(jsonValue));
    }

    /**
     * Create a new wayCollectionWay
     */
    async createWayCollectionWay(requestParameters: CreateWayCollectionWayRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasWayCollectionWayResponse> {
        const response = await this.createWayCollectionWayRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete wayCollectionWay by UUID
     */
    async deleteWayCollectionWayRaw(requestParameters: DeleteWayCollectionWayRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.wayCollectionId === null || requestParameters.wayCollectionId === undefined) {
            throw new runtime.RequiredError('wayCollectionId','Required parameter requestParameters.wayCollectionId was null or undefined when calling deleteWayCollectionWay.');
        }

        if (requestParameters.wayId === null || requestParameters.wayId === undefined) {
            throw new runtime.RequiredError('wayId','Required parameter requestParameters.wayId was null or undefined when calling deleteWayCollectionWay.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/wayCollectionWays/{wayId}/{wayCollectionId}`.replace(`{${"wayCollectionId"}}`, encodeURIComponent(String(requestParameters.wayCollectionId))).replace(`{${"wayId"}}`, encodeURIComponent(String(requestParameters.wayId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete wayCollectionWay by UUID
     */
    async deleteWayCollectionWay(requestParameters: DeleteWayCollectionWayRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteWayCollectionWayRaw(requestParameters, initOverrides);
    }

}
