/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.databind.annotation.JsonDeserialize;
import com.fasterxml.jackson.databind.annotation.JsonSerialize;
import com.formance.formance_sdk.utils.DateTimeDeserializer;
import com.formance.formance_sdk.utils.DateTimeSerializer;
import java.time.OffsetDateTime;

public class TaskStripe {
    @JsonProperty("connectorID")
    public String connectorID;

    public TaskStripe withConnectorID(String connectorID) {
        this.connectorID = connectorID;
        return this;
    }
    
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("createdAt")
    public OffsetDateTime createdAt;

    public TaskStripe withCreatedAt(OffsetDateTime createdAt) {
        this.createdAt = createdAt;
        return this;
    }
    
    @JsonProperty("descriptor")
    public TaskStripeDescriptor descriptor;

    public TaskStripe withDescriptor(TaskStripeDescriptor descriptor) {
        this.descriptor = descriptor;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("error")
    public String error;

    public TaskStripe withError(String error) {
        this.error = error;
        return this;
    }
    
    @JsonProperty("id")
    public String id;

    public TaskStripe withId(String id) {
        this.id = id;
        return this;
    }
    
    @JsonProperty("state")
    public java.util.Map<String, Object> state;

    public TaskStripe withState(java.util.Map<String, Object> state) {
        this.state = state;
        return this;
    }
    
    @JsonProperty("status")
    public PaymentStatus status;

    public TaskStripe withStatus(PaymentStatus status) {
        this.status = status;
        return this;
    }
    
    @JsonSerialize(using = DateTimeSerializer.class)
    @JsonDeserialize(using = DateTimeDeserializer.class)
    @JsonProperty("updatedAt")
    public OffsetDateTime updatedAt;

    public TaskStripe withUpdatedAt(OffsetDateTime updatedAt) {
        this.updatedAt = updatedAt;
        return this;
    }
    
    public TaskStripe(@JsonProperty("connectorID") String connectorID, @JsonProperty("createdAt") OffsetDateTime createdAt, @JsonProperty("descriptor") TaskStripeDescriptor descriptor, @JsonProperty("id") String id, @JsonProperty("state") java.util.Map<String, Object> state, @JsonProperty("status") PaymentStatus status, @JsonProperty("updatedAt") OffsetDateTime updatedAt) {
        this.connectorID = connectorID;
        this.createdAt = createdAt;
        this.descriptor = descriptor;
        this.id = id;
        this.state = state;
        this.status = status;
        this.updatedAt = updatedAt;
  }
}
