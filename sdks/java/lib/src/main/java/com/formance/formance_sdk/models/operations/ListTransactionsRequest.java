/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.operations;

import com.fasterxml.jackson.annotation.JsonProperty;
import com.formance.formance_sdk.utils.SpeakeasyMetadata;
import java.time.OffsetDateTime;


public class ListTransactionsRequest {
    /**
     * Filter transactions with postings involving given account, either as source or destination (regular expression placed between ^ and $).
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=account")
    public String account;

    public ListTransactionsRequest withAccount(String account) {
        this.account = account;
        return this;
    }
    
    /**
     * Pagination cursor, will return transactions after given txid (in descending order).
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=after")
    public String after;

    public ListTransactionsRequest withAfter(String after) {
        this.after = after;
        return this;
    }
    
    /**
     * Parameter used in pagination requests. Maximum page size is set to 15.
     * Set to the value of next for the next page of results.
     * Set to the value of previous for the previous page of results.
     * No other parameters can be set when this parameter is set.
     * 
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=cursor")
    public String cursor;

    public ListTransactionsRequest withCursor(String cursor) {
        this.cursor = cursor;
        return this;
    }
    
    /**
     * Filter transactions with postings involving given account at destination (regular expression placed between ^ and $).
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=destination")
    public String destination;

    public ListTransactionsRequest withDestination(String destination) {
        this.destination = destination;
        return this;
    }
    
    /**
     * Filter transactions that occurred before this timestamp.
     * The format is RFC3339 and is exclusive (for example, "2023-01-02T15:04:01Z" excludes the first second of 4th minute).
     * 
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=endTime")
    public OffsetDateTime endTime;

    public ListTransactionsRequest withEndTime(OffsetDateTime endTime) {
        this.endTime = endTime;
        return this;
    }
    
    /**
     * Name of the ledger.
     */
    @SpeakeasyMetadata("pathParam:style=simple,explode=false,name=ledger")
    public String ledger;

    public ListTransactionsRequest withLedger(String ledger) {
        this.ledger = ledger;
        return this;
    }
    
    /**
     * Filter transactions by metadata key value pairs. Nested objects can be used as seen in the example below.
     */
    @SpeakeasyMetadata("queryParam:style=deepObject,explode=true,name=metadata")
    public ListTransactionsMetadata metadata;

    public ListTransactionsRequest withMetadata(ListTransactionsMetadata metadata) {
        this.metadata = metadata;
        return this;
    }
    
    /**
     * The maximum number of results to return per page.
     * 
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=pageSize")
    public Long pageSize;

    public ListTransactionsRequest withPageSize(Long pageSize) {
        this.pageSize = pageSize;
        return this;
    }
    
    /**
     * Find transactions by reference field.
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=reference")
    public String reference;

    public ListTransactionsRequest withReference(String reference) {
        this.reference = reference;
        return this;
    }
    
    /**
     * Filter transactions with postings involving given account at source (regular expression placed between ^ and $).
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=source")
    public String source;

    public ListTransactionsRequest withSource(String source) {
        this.source = source;
        return this;
    }
    
    /**
     * Filter transactions that occurred after this timestamp.
     * The format is RFC3339 and is inclusive (for example, "2023-01-02T15:04:01Z" includes the first second of 4th minute).
     * 
     */
    @SpeakeasyMetadata("queryParam:style=form,explode=true,name=startTime")
    public OffsetDateTime startTime;

    public ListTransactionsRequest withStartTime(OffsetDateTime startTime) {
        this.startTime = startTime;
        return this;
    }
    
    public ListTransactionsRequest(@JsonProperty("ledger") String ledger) {
        this.ledger = ledger;
  }
}
