/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonProperty;

/**
 * ConnectorsConfigsResponse - OK
 */

public class ConnectorsConfigsResponse {
    @JsonProperty("data")
    public ConnectorsConfigsResponseData data;

    public ConnectorsConfigsResponse withData(ConnectorsConfigsResponseData data) {
        this.data = data;
        return this;
    }
    
    public ConnectorsConfigsResponse(@JsonProperty("data") ConnectorsConfigsResponseData data) {
        this.data = data;
  }
}
