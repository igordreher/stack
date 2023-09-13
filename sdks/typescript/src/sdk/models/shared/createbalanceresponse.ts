/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Balance } from "./balance";
import { Expose, Type } from "class-transformer";

/**
 * Created balance
 */
export class CreateBalanceResponse extends SpeakeasyBase {
    @SpeakeasyMetadata()
    @Expose({ name: "data" })
    @Type(() => Balance)
    data: Balance;
}
