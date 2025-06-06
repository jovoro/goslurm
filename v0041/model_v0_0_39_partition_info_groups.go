/*
Slurm REST API

API to access and control Slurm

API version: Slurm-24.05.5&openapi/slurmdbd&openapi/dbv0.0.39&openapi/slurmctld&openapi/v0.0.39
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v0041

import (
	"encoding/json"
)

// checks if the V0039PartitionInfoGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039PartitionInfoGroups{}

// V0039PartitionInfoGroups struct for V0039PartitionInfoGroups
type V0039PartitionInfoGroups struct {
	Allowed *string `json:"allowed,omitempty"`
}

// NewV0039PartitionInfoGroups instantiates a new V0039PartitionInfoGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039PartitionInfoGroups() *V0039PartitionInfoGroups {
	this := V0039PartitionInfoGroups{}
	return &this
}

// NewV0039PartitionInfoGroupsWithDefaults instantiates a new V0039PartitionInfoGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039PartitionInfoGroupsWithDefaults() *V0039PartitionInfoGroups {
	this := V0039PartitionInfoGroups{}
	return &this
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *V0039PartitionInfoGroups) GetAllowed() string {
	if o == nil || IsNil(o.Allowed) {
		var ret string
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionInfoGroups) GetAllowedOk() (*string, bool) {
	if o == nil || IsNil(o.Allowed) {
		return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *V0039PartitionInfoGroups) HasAllowed() bool {
	if o != nil && !IsNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given string and assigns it to the Allowed field.
func (o *V0039PartitionInfoGroups) SetAllowed(v string) {
	o.Allowed = &v
}

func (o V0039PartitionInfoGroups) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039PartitionInfoGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return toSerialize, nil
}

type NullableV0039PartitionInfoGroups struct {
	value *V0039PartitionInfoGroups
	isSet bool
}

func (v NullableV0039PartitionInfoGroups) Get() *V0039PartitionInfoGroups {
	return v.value
}

func (v *NullableV0039PartitionInfoGroups) Set(val *V0039PartitionInfoGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039PartitionInfoGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039PartitionInfoGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039PartitionInfoGroups(val *V0039PartitionInfoGroups) *NullableV0039PartitionInfoGroups {
	return &NullableV0039PartitionInfoGroups{value: val, isSet: true}
}

func (v NullableV0039PartitionInfoGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039PartitionInfoGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


