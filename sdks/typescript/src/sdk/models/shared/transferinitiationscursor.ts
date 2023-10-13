/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { TransferInitiation } from "./transferinitiation";
import { Expose, Type } from "class-transformer";

export class TransferInitiationsCursorCursor extends SpeakeasyBase {
  @SpeakeasyMetadata({ elemType: TransferInitiation })
  @Expose({ name: "data" })
  @Type(() => TransferInitiation)
  data: TransferInitiation[];

  @SpeakeasyMetadata()
  @Expose({ name: "hasMore" })
  hasMore: boolean;

  @SpeakeasyMetadata()
  @Expose({ name: "next" })
  next?: string;

  @SpeakeasyMetadata()
  @Expose({ name: "pageSize" })
  pageSize: number;

  @SpeakeasyMetadata()
  @Expose({ name: "previous" })
  previous?: string;
}

/**
 * OK
 */
export class TransferInitiationsCursor extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "cursor" })
  @Type(() => TransferInitiationsCursorCursor)
  cursor: TransferInitiationsCursorCursor;
}
