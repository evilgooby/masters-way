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
  CustomErrorsNoRightToChangeDayReportError,
  SchemasCommentPopulatedResponse,
  SchemasCreateCommentPayload,
  SchemasUpdateCommentPayload,
} from '../models/index';
import {
    CustomErrorsNoRightToChangeDayReportErrorFromJSON,
    CustomErrorsNoRightToChangeDayReportErrorToJSON,
    SchemasCommentPopulatedResponseFromJSON,
    SchemasCommentPopulatedResponseToJSON,
    SchemasCreateCommentPayloadFromJSON,
    SchemasCreateCommentPayloadToJSON,
    SchemasUpdateCommentPayloadFromJSON,
    SchemasUpdateCommentPayloadToJSON,
} from '../models/index';

export interface CreateCommentRequest {
    request: SchemasCreateCommentPayload;
}

export interface DeleteCommentRequest {
    commentId: string;
}

export interface UpdateCommentRequest {
    commentId: string;
    request: SchemasUpdateCommentPayload;
}

/**
 * 
 */
export class CommentApi extends runtime.BaseAPI {

    /**
     * Create a new comment
     */
    async createCommentRaw(requestParameters: CreateCommentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasCommentPopulatedResponse>> {
        if (requestParameters.request === null || requestParameters.request === undefined) {
            throw new runtime.RequiredError('request','Required parameter requestParameters.request was null or undefined when calling createComment.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/comments`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: SchemasCreateCommentPayloadToJSON(requestParameters.request),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasCommentPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Create a new comment
     */
    async createComment(requestParameters: CreateCommentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasCommentPopulatedResponse> {
        const response = await this.createCommentRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Delete comment by UUID
     */
    async deleteCommentRaw(requestParameters: DeleteCommentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        if (requestParameters.commentId === null || requestParameters.commentId === undefined) {
            throw new runtime.RequiredError('commentId','Required parameter requestParameters.commentId was null or undefined when calling deleteComment.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/comments/{commentId}`.replace(`{${"commentId"}}`, encodeURIComponent(String(requestParameters.commentId))),
            method: 'DELETE',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * Delete comment by UUID
     */
    async deleteComment(requestParameters: DeleteCommentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.deleteCommentRaw(requestParameters, initOverrides);
    }

    /**
     * Update comment by UUID
     */
    async updateCommentRaw(requestParameters: UpdateCommentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<SchemasCommentPopulatedResponse>> {
        if (requestParameters.commentId === null || requestParameters.commentId === undefined) {
            throw new runtime.RequiredError('commentId','Required parameter requestParameters.commentId was null or undefined when calling updateComment.');
        }

        if (requestParameters.request === null || requestParameters.request === undefined) {
            throw new runtime.RequiredError('request','Required parameter requestParameters.request was null or undefined when calling updateComment.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/comments/{commentId}`.replace(`{${"commentId"}}`, encodeURIComponent(String(requestParameters.commentId))),
            method: 'PATCH',
            headers: headerParameters,
            query: queryParameters,
            body: SchemasUpdateCommentPayloadToJSON(requestParameters.request),
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => SchemasCommentPopulatedResponseFromJSON(jsonValue));
    }

    /**
     * Update comment by UUID
     */
    async updateComment(requestParameters: UpdateCommentRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<SchemasCommentPopulatedResponse> {
        const response = await this.updateCommentRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
