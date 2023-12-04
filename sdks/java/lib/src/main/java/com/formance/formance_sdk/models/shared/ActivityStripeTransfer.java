/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

public class ActivityStripeTransfer {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("amount")
    public Long amount;

    public ActivityStripeTransfer withAmount(Long amount) {
        this.amount = amount;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("asset")
    public String asset;

    public ActivityStripeTransfer withAsset(String asset) {
        this.asset = asset;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("connectorID")
    public String connectorID;

    public ActivityStripeTransfer withConnectorID(String connectorID) {
        this.connectorID = connectorID;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("destination")
    public String destination;

    public ActivityStripeTransfer withDestination(String destination) {
        this.destination = destination;
        return this;
    }
    
    /**
     * A set of key/value pairs that you can attach to a transfer object.
     * It can be useful for storing additional information about the transfer in a structured format.
     * 
     */
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("metadata")
    public java.util.Map<String, Object> metadata;

    public ActivityStripeTransfer withMetadata(java.util.Map<String, Object> metadata) {
        this.metadata = metadata;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("waitingValidation")
    public Boolean waitingValidation;

    public ActivityStripeTransfer withWaitingValidation(Boolean waitingValidation) {
        this.waitingValidation = waitingValidation;
        return this;
    }
    
    public ActivityStripeTransfer(){}
}
