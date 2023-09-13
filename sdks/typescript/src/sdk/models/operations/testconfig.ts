/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";

export class TestConfigRequest extends SpeakeasyBase {
    /**
     * Config ID
     */
    @SpeakeasyMetadata({ data: "pathParam, style=simple;explode=false;name=id" })
    id: string;
}

export class TestConfigResponse extends SpeakeasyBase {
    /**
     * OK
     */
    @SpeakeasyMetadata()
    attemptResponse?: shared.AttemptResponse;

    @SpeakeasyMetadata()
    contentType: string;

    @SpeakeasyMetadata()
    statusCode: number;

    @SpeakeasyMetadata()
    rawResponse?: AxiosResponse;

    /**
     * Error
     */
    @SpeakeasyMetadata()
    webhooksErrorResponse?: shared.WebhooksErrorResponse;
}
