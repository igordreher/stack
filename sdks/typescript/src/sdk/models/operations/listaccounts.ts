/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import * as shared from "../shared";
import { AxiosResponse } from "axios";

/**
 * Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
 *
 * @remarks
 *
 */
export enum ListAccountsBalanceOperator {
    Gte = "gte",
    Lte = "lte",
    Gt = "gt",
    Lt = "lt",
    E = "e",
    Ne = "ne",
}

/**
 * Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
 */
export class ListAccountsMetadata extends SpeakeasyBase {}

export class ListAccountsRequest extends SpeakeasyBase {
    /**
     * Filter accounts by address pattern (regular expression placed between ^ and $).
     */
    @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=address" })
    address?: string;

    /**
     * Pagination cursor, will return accounts after given address, in descending order.
     */
    @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=after" })
    after?: string;

    /**
     * Filter accounts by their balance (default operator is gte)
     */
    @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=balance" })
    balance?: number;

    /**
     * Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
     *
     * @remarks
     *
     */
    @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=balanceOperator" })
    balanceOperator?: ListAccountsBalanceOperator;

    /**
     * Parameter used in pagination requests. Maximum page size is set to 15.
     *
     * @remarks
     * Set to the value of next for the next page of results.
     * Set to the value of previous for the previous page of results.
     * No other parameters can be set when this parameter is set.
     *
     */
    @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=cursor" })
    cursor?: string;

    /**
     * Name of the ledger.
     */
    @SpeakeasyMetadata({ data: "pathParam, style=simple;explode=false;name=ledger" })
    ledger: string;

    /**
     * Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
     */
    @SpeakeasyMetadata({ data: "queryParam, style=deepObject;explode=true;name=metadata" })
    metadata?: ListAccountsMetadata;

    /**
     * The maximum number of results to return per page.
     *
     * @remarks
     *
     */
    @SpeakeasyMetadata({ data: "queryParam, style=form;explode=true;name=pageSize" })
    pageSize?: number;
}

export class ListAccountsResponse extends SpeakeasyBase {
    /**
     * OK
     */
    @SpeakeasyMetadata()
    accountsCursorResponse?: shared.AccountsCursorResponse;

    @SpeakeasyMetadata()
    contentType: string;

    /**
     * Error
     */
    @SpeakeasyMetadata()
    errorResponse?: shared.ErrorResponse;

    @SpeakeasyMetadata()
    statusCode: number;

    @SpeakeasyMetadata()
    rawResponse?: AxiosResponse;
}
