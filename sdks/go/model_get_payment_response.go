/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions />

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// GetPaymentResponse struct for GetPaymentResponse
type GetPaymentResponse struct {
	Data Payment `json:"data"`
}

// NewGetPaymentResponse instantiates a new GetPaymentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPaymentResponse(data Payment) *GetPaymentResponse {
	this := GetPaymentResponse{}
	this.Data = data
	return &this
}

// NewGetPaymentResponseWithDefaults instantiates a new GetPaymentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPaymentResponseWithDefaults() *GetPaymentResponse {
	this := GetPaymentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetPaymentResponse) GetData() Payment {
	if o == nil {
		var ret Payment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetPaymentResponse) GetDataOk() (*Payment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetPaymentResponse) SetData(v Payment) {
	o.Data = v
}

func (o GetPaymentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableGetPaymentResponse struct {
	value *GetPaymentResponse
	isSet bool
}

func (v NullableGetPaymentResponse) Get() *GetPaymentResponse {
	return v.value
}

func (v *NullableGetPaymentResponse) Set(val *GetPaymentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPaymentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPaymentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPaymentResponse(val *GetPaymentResponse) *NullableGetPaymentResponse {
	return &NullableGetPaymentResponse{value: val, isSet: true}
}

func (v NullableGetPaymentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPaymentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
