/* 
 * Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.
 */

package com.formance.formance_sdk.models.shared;

import com.fasterxml.jackson.annotation.JsonInclude.Include;
import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;


public class WorkflowInstanceHistoryStageOutput {
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("CreateTransaction")
    public ActivityCreateTransactionOutput createTransaction;

    public WorkflowInstanceHistoryStageOutput withCreateTransaction(ActivityCreateTransactionOutput createTransaction) {
        this.createTransaction = createTransaction;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("DebitWallet")
    public ActivityDebitWalletOutput debitWallet;

    public WorkflowInstanceHistoryStageOutput withDebitWallet(ActivityDebitWalletOutput debitWallet) {
        this.debitWallet = debitWallet;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("GetAccount")
    public ActivityGetAccountOutput getAccount;

    public WorkflowInstanceHistoryStageOutput withGetAccount(ActivityGetAccountOutput getAccount) {
        this.getAccount = getAccount;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("GetPayment")
    public ActivityGetPaymentOutput getPayment;

    public WorkflowInstanceHistoryStageOutput withGetPayment(ActivityGetPaymentOutput getPayment) {
        this.getPayment = getPayment;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("GetWallet")
    public ActivityGetWalletOutput getWallet;

    public WorkflowInstanceHistoryStageOutput withGetWallet(ActivityGetWalletOutput getWallet) {
        this.getWallet = getWallet;
        return this;
    }
    
    @JsonInclude(Include.NON_ABSENT)
    @JsonProperty("RevertTransaction")
    public ActivityRevertTransactionOutput revertTransaction;

    public WorkflowInstanceHistoryStageOutput withRevertTransaction(ActivityRevertTransactionOutput revertTransaction) {
        this.revertTransaction = revertTransaction;
        return this;
    }
    
    public WorkflowInstanceHistoryStageOutput(){}
}
