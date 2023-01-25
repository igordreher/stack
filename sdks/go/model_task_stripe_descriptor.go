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

// TaskStripeDescriptor struct for TaskStripeDescriptor
type TaskStripeDescriptor struct {
	Name string `json:"name"`
	Main *bool `json:"main,omitempty"`
	Account *string `json:"account,omitempty"`
}

// NewTaskStripeDescriptor instantiates a new TaskStripeDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskStripeDescriptor(name string) *TaskStripeDescriptor {
	this := TaskStripeDescriptor{}
	this.Name = name
	return &this
}

// NewTaskStripeDescriptorWithDefaults instantiates a new TaskStripeDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskStripeDescriptorWithDefaults() *TaskStripeDescriptor {
	this := TaskStripeDescriptor{}
	return &this
}

// GetName returns the Name field value
func (o *TaskStripeDescriptor) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TaskStripeDescriptor) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TaskStripeDescriptor) SetName(v string) {
	o.Name = v
}

// GetMain returns the Main field value if set, zero value otherwise.
func (o *TaskStripeDescriptor) GetMain() bool {
	if o == nil || isNil(o.Main) {
		var ret bool
		return ret
	}
	return *o.Main
}

// GetMainOk returns a tuple with the Main field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStripeDescriptor) GetMainOk() (*bool, bool) {
	if o == nil || isNil(o.Main) {
    return nil, false
	}
	return o.Main, true
}

// HasMain returns a boolean if a field has been set.
func (o *TaskStripeDescriptor) HasMain() bool {
	if o != nil && !isNil(o.Main) {
		return true
	}

	return false
}

// SetMain gets a reference to the given bool and assigns it to the Main field.
func (o *TaskStripeDescriptor) SetMain(v bool) {
	o.Main = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *TaskStripeDescriptor) GetAccount() string {
	if o == nil || isNil(o.Account) {
		var ret string
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskStripeDescriptor) GetAccountOk() (*string, bool) {
	if o == nil || isNil(o.Account) {
    return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *TaskStripeDescriptor) HasAccount() bool {
	if o != nil && !isNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given string and assigns it to the Account field.
func (o *TaskStripeDescriptor) SetAccount(v string) {
	o.Account = &v
}

func (o TaskStripeDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Main) {
		toSerialize["main"] = o.Main
	}
	if !isNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	return json.Marshal(toSerialize)
}

type NullableTaskStripeDescriptor struct {
	value *TaskStripeDescriptor
	isSet bool
}

func (v NullableTaskStripeDescriptor) Get() *TaskStripeDescriptor {
	return v.value
}

func (v *NullableTaskStripeDescriptor) Set(val *TaskStripeDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskStripeDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskStripeDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskStripeDescriptor(val *TaskStripeDescriptor) *NullableTaskStripeDescriptor {
	return &NullableTaskStripeDescriptor{value: val, isSet: true}
}

func (v NullableTaskStripeDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskStripeDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


