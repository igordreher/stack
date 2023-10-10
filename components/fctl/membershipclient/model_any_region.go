/*
Membership API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package membershipclient

import (
	"encoding/json"
	"time"
)

// checks if the AnyRegion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnyRegion{}

// AnyRegion struct for AnyRegion
type AnyRegion struct {
	Id string `json:"id"`
	BaseUrl string `json:"baseUrl"`
	CreatedAt string `json:"createdAt"`
	Active bool `json:"active"`
	LastPing *time.Time `json:"lastPing,omitempty"`
	Name string `json:"name"`
	ClientID *string `json:"clientID,omitempty"`
	OrganizationID *string `json:"organizationID,omitempty"`
	Creator *User `json:"creator,omitempty"`
	Production *bool `json:"production,omitempty"`
	Public bool `json:"public"`
	Version *string `json:"version,omitempty"`
}

// NewAnyRegion instantiates a new AnyRegion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnyRegion(id string, baseUrl string, createdAt string, active bool, name string, public bool) *AnyRegion {
	this := AnyRegion{}
	this.Id = id
	this.BaseUrl = baseUrl
	this.CreatedAt = createdAt
	this.Active = active
	this.Name = name
	this.Public = public
	return &this
}

// NewAnyRegionWithDefaults instantiates a new AnyRegion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnyRegionWithDefaults() *AnyRegion {
	this := AnyRegion{}
	return &this
}

// GetId returns the Id field value
func (o *AnyRegion) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AnyRegion) SetId(v string) {
	o.Id = v
}

// GetBaseUrl returns the BaseUrl field value
func (o *AnyRegion) GetBaseUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BaseUrl
}

// GetBaseUrlOk returns a tuple with the BaseUrl field value
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetBaseUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseUrl, true
}

// SetBaseUrl sets field value
func (o *AnyRegion) SetBaseUrl(v string) {
	o.BaseUrl = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *AnyRegion) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *AnyRegion) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetActive returns the Active field value
func (o *AnyRegion) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *AnyRegion) SetActive(v bool) {
	o.Active = v
}

// GetLastPing returns the LastPing field value if set, zero value otherwise.
func (o *AnyRegion) GetLastPing() time.Time {
	if o == nil || IsNil(o.LastPing) {
		var ret time.Time
		return ret
	}
	return *o.LastPing
}

// GetLastPingOk returns a tuple with the LastPing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetLastPingOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPing) {
		return nil, false
	}
	return o.LastPing, true
}

// HasLastPing returns a boolean if a field has been set.
func (o *AnyRegion) HasLastPing() bool {
	if o != nil && !IsNil(o.LastPing) {
		return true
	}

	return false
}

// SetLastPing gets a reference to the given time.Time and assigns it to the LastPing field.
func (o *AnyRegion) SetLastPing(v time.Time) {
	o.LastPing = &v
}

// GetName returns the Name field value
func (o *AnyRegion) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AnyRegion) SetName(v string) {
	o.Name = v
}

// GetClientID returns the ClientID field value if set, zero value otherwise.
func (o *AnyRegion) GetClientID() string {
	if o == nil || IsNil(o.ClientID) {
		var ret string
		return ret
	}
	return *o.ClientID
}

// GetClientIDOk returns a tuple with the ClientID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetClientIDOk() (*string, bool) {
	if o == nil || IsNil(o.ClientID) {
		return nil, false
	}
	return o.ClientID, true
}

// HasClientID returns a boolean if a field has been set.
func (o *AnyRegion) HasClientID() bool {
	if o != nil && !IsNil(o.ClientID) {
		return true
	}

	return false
}

// SetClientID gets a reference to the given string and assigns it to the ClientID field.
func (o *AnyRegion) SetClientID(v string) {
	o.ClientID = &v
}

// GetOrganizationID returns the OrganizationID field value if set, zero value otherwise.
func (o *AnyRegion) GetOrganizationID() string {
	if o == nil || IsNil(o.OrganizationID) {
		var ret string
		return ret
	}
	return *o.OrganizationID
}

// GetOrganizationIDOk returns a tuple with the OrganizationID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetOrganizationIDOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationID) {
		return nil, false
	}
	return o.OrganizationID, true
}

// HasOrganizationID returns a boolean if a field has been set.
func (o *AnyRegion) HasOrganizationID() bool {
	if o != nil && !IsNil(o.OrganizationID) {
		return true
	}

	return false
}

// SetOrganizationID gets a reference to the given string and assigns it to the OrganizationID field.
func (o *AnyRegion) SetOrganizationID(v string) {
	o.OrganizationID = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *AnyRegion) GetCreator() User {
	if o == nil || IsNil(o.Creator) {
		var ret User
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetCreatorOk() (*User, bool) {
	if o == nil || IsNil(o.Creator) {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *AnyRegion) HasCreator() bool {
	if o != nil && !IsNil(o.Creator) {
		return true
	}

	return false
}

// SetCreator gets a reference to the given User and assigns it to the Creator field.
func (o *AnyRegion) SetCreator(v User) {
	o.Creator = &v
}

// GetProduction returns the Production field value if set, zero value otherwise.
func (o *AnyRegion) GetProduction() bool {
	if o == nil || IsNil(o.Production) {
		var ret bool
		return ret
	}
	return *o.Production
}

// GetProductionOk returns a tuple with the Production field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetProductionOk() (*bool, bool) {
	if o == nil || IsNil(o.Production) {
		return nil, false
	}
	return o.Production, true
}

// HasProduction returns a boolean if a field has been set.
func (o *AnyRegion) HasProduction() bool {
	if o != nil && !IsNil(o.Production) {
		return true
	}

	return false
}

// SetProduction gets a reference to the given bool and assigns it to the Production field.
func (o *AnyRegion) SetProduction(v bool) {
	o.Production = &v
}

// GetPublic returns the Public field value
func (o *AnyRegion) GetPublic() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Public
}

// GetPublicOk returns a tuple with the Public field value
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetPublicOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Public, true
}

// SetPublic sets field value
func (o *AnyRegion) SetPublic(v bool) {
	o.Public = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AnyRegion) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnyRegion) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AnyRegion) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AnyRegion) SetVersion(v string) {
	o.Version = &v
}

func (o AnyRegion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnyRegion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["baseUrl"] = o.BaseUrl
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["active"] = o.Active
	if !IsNil(o.LastPing) {
		toSerialize["lastPing"] = o.LastPing
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.ClientID) {
		toSerialize["clientID"] = o.ClientID
	}
	if !IsNil(o.OrganizationID) {
		toSerialize["organizationID"] = o.OrganizationID
	}
	if !IsNil(o.Creator) {
		toSerialize["creator"] = o.Creator
	}
	if !IsNil(o.Production) {
		toSerialize["production"] = o.Production
	}
	toSerialize["public"] = o.Public
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableAnyRegion struct {
	value *AnyRegion
	isSet bool
}

func (v NullableAnyRegion) Get() *AnyRegion {
	return v.value
}

func (v *NullableAnyRegion) Set(val *AnyRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableAnyRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableAnyRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnyRegion(val *AnyRegion) *NullableAnyRegion {
	return &NullableAnyRegion{value: val, isSet: true}
}

func (v NullableAnyRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnyRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


