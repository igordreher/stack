// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type TaskBankingCircleDescriptor struct {
	Key  *string `json:"key,omitempty"`
	Name *string `json:"name,omitempty"`
}

type TaskBankingCircle struct {
	ConnectorID string                      `json:"connectorID"`
	CreatedAt   time.Time                   `json:"createdAt"`
	Descriptor  TaskBankingCircleDescriptor `json:"descriptor"`
	Error       *string                     `json:"error,omitempty"`
	ID          string                      `json:"id"`
	State       map[string]interface{}      `json:"state"`
	Status      PaymentStatus               `json:"status"`
	UpdatedAt   time.Time                   `json:"updatedAt"`
}
