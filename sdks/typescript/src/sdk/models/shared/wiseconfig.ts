/*
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

import { SpeakeasyBase, SpeakeasyMetadata } from "../../../internal/utils";
import { Expose } from "class-transformer";

export class WiseConfig extends SpeakeasyBase {
  @SpeakeasyMetadata()
  @Expose({ name: "apiKey" })
  apiKey: string;

  @SpeakeasyMetadata()
  @Expose({ name: "name" })
  name: string;

  /**
   * The frequency at which the connector will try to fetch new BalanceTransaction objects from Wise API.
   *
   * @remarks
   *
   */
  @SpeakeasyMetadata()
  @Expose({ name: "pollingPeriod" })
  pollingPeriod?: string;
}
