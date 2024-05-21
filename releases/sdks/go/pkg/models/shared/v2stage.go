// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/formancehq/formance-sdk-go/v2/pkg/utils"
)

type V2StageType string

const (
	V2StageTypeV2StageSend      V2StageType = "V2StageSend"
	V2StageTypeV2StageDelay     V2StageType = "V2StageDelay"
	V2StageTypeV2StageWaitEvent V2StageType = "V2StageWaitEvent"
)

type V2Stage struct {
	V2StageSend      *V2StageSend
	V2StageDelay     *V2StageDelay
	V2StageWaitEvent *V2StageWaitEvent

	Type V2StageType
}

func CreateV2StageV2StageSend(v2StageSend V2StageSend) V2Stage {
	typ := V2StageTypeV2StageSend

	return V2Stage{
		V2StageSend: &v2StageSend,
		Type:        typ,
	}
}

func CreateV2StageV2StageDelay(v2StageDelay V2StageDelay) V2Stage {
	typ := V2StageTypeV2StageDelay

	return V2Stage{
		V2StageDelay: &v2StageDelay,
		Type:         typ,
	}
}

func CreateV2StageV2StageWaitEvent(v2StageWaitEvent V2StageWaitEvent) V2Stage {
	typ := V2StageTypeV2StageWaitEvent

	return V2Stage{
		V2StageWaitEvent: &v2StageWaitEvent,
		Type:             typ,
	}
}

func (u *V2Stage) UnmarshalJSON(data []byte) error {

	var v2StageWaitEvent V2StageWaitEvent = V2StageWaitEvent{}
	if err := utils.UnmarshalJSON(data, &v2StageWaitEvent, "", true, true); err == nil {
		u.V2StageWaitEvent = &v2StageWaitEvent
		u.Type = V2StageTypeV2StageWaitEvent
		return nil
	}

	var v2StageDelay V2StageDelay = V2StageDelay{}
	if err := utils.UnmarshalJSON(data, &v2StageDelay, "", true, true); err == nil {
		u.V2StageDelay = &v2StageDelay
		u.Type = V2StageTypeV2StageDelay
		return nil
	}

	var v2StageSend V2StageSend = V2StageSend{}
	if err := utils.UnmarshalJSON(data, &v2StageSend, "", true, true); err == nil {
		u.V2StageSend = &v2StageSend
		u.Type = V2StageTypeV2StageSend
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u V2Stage) MarshalJSON() ([]byte, error) {
	if u.V2StageSend != nil {
		return utils.MarshalJSON(u.V2StageSend, "", true)
	}

	if u.V2StageDelay != nil {
		return utils.MarshalJSON(u.V2StageDelay, "", true)
	}

	if u.V2StageWaitEvent != nil {
		return utils.MarshalJSON(u.V2StageWaitEvent, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
