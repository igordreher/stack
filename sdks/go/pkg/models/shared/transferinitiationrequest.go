// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"math/big"
	"time"
)

type TransferInitiationRequestType string

const (
	TransferInitiationRequestTypeTransfer TransferInitiationRequestType = "TRANSFER"
	TransferInitiationRequestTypePayout   TransferInitiationRequestType = "PAYOUT"
)

func (e TransferInitiationRequestType) ToPointer() *TransferInitiationRequestType {
	return &e
}

func (e *TransferInitiationRequestType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TRANSFER":
		fallthrough
	case "PAYOUT":
		*e = TransferInitiationRequestType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TransferInitiationRequestType: %v", v)
	}
}

type TransferInitiationRequest struct {
	Amount               *big.Int                      `json:"amount"`
	Asset                string                        `json:"asset"`
	CreatedAt            time.Time                     `json:"createdAt"`
	Description          string                        `json:"description"`
	DestinationAccountID string                        `json:"destinationAccountID"`
	Provider             Connector                     `json:"provider"`
	Reference            string                        `json:"reference"`
	SourceAccountID      string                        `json:"sourceAccountID"`
	Type                 TransferInitiationRequestType `json:"type"`
	Validated            bool                          `json:"validated"`
}
