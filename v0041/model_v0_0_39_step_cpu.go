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

// checks if the V0039StepCPU type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepCPU{}

// V0039StepCPU struct for V0039StepCPU
type V0039StepCPU struct {
	RequestedFrequency *V0039StepCPURequestedFrequency `json:"requested_frequency,omitempty"`
	Governor *string `json:"governor,omitempty"`
}

// NewV0039StepCPU instantiates a new V0039StepCPU object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepCPU() *V0039StepCPU {
	this := V0039StepCPU{}
	return &this
}

// NewV0039StepCPUWithDefaults instantiates a new V0039StepCPU object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepCPUWithDefaults() *V0039StepCPU {
	this := V0039StepCPU{}
	return &this
}

// GetRequestedFrequency returns the RequestedFrequency field value if set, zero value otherwise.
func (o *V0039StepCPU) GetRequestedFrequency() V0039StepCPURequestedFrequency {
	if o == nil || IsNil(o.RequestedFrequency) {
		var ret V0039StepCPURequestedFrequency
		return ret
	}
	return *o.RequestedFrequency
}

// GetRequestedFrequencyOk returns a tuple with the RequestedFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepCPU) GetRequestedFrequencyOk() (*V0039StepCPURequestedFrequency, bool) {
	if o == nil || IsNil(o.RequestedFrequency) {
		return nil, false
	}
	return o.RequestedFrequency, true
}

// HasRequestedFrequency returns a boolean if a field has been set.
func (o *V0039StepCPU) HasRequestedFrequency() bool {
	if o != nil && !IsNil(o.RequestedFrequency) {
		return true
	}

	return false
}

// SetRequestedFrequency gets a reference to the given V0039StepCPURequestedFrequency and assigns it to the RequestedFrequency field.
func (o *V0039StepCPU) SetRequestedFrequency(v V0039StepCPURequestedFrequency) {
	o.RequestedFrequency = &v
}

// GetGovernor returns the Governor field value if set, zero value otherwise.
func (o *V0039StepCPU) GetGovernor() string {
	if o == nil || IsNil(o.Governor) {
		var ret string
		return ret
	}
	return *o.Governor
}

// GetGovernorOk returns a tuple with the Governor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepCPU) GetGovernorOk() (*string, bool) {
	if o == nil || IsNil(o.Governor) {
		return nil, false
	}
	return o.Governor, true
}

// HasGovernor returns a boolean if a field has been set.
func (o *V0039StepCPU) HasGovernor() bool {
	if o != nil && !IsNil(o.Governor) {
		return true
	}

	return false
}

// SetGovernor gets a reference to the given string and assigns it to the Governor field.
func (o *V0039StepCPU) SetGovernor(v string) {
	o.Governor = &v
}

func (o V0039StepCPU) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepCPU) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RequestedFrequency) {
		toSerialize["requested_frequency"] = o.RequestedFrequency
	}
	if !IsNil(o.Governor) {
		toSerialize["governor"] = o.Governor
	}
	return toSerialize, nil
}

type NullableV0039StepCPU struct {
	value *V0039StepCPU
	isSet bool
}

func (v NullableV0039StepCPU) Get() *V0039StepCPU {
	return v.value
}

func (v *NullableV0039StepCPU) Set(val *V0039StepCPU) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepCPU) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepCPU) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepCPU(val *V0039StepCPU) *NullableV0039StepCPU {
	return &NullableV0039StepCPU{value: val, isSet: true}
}

func (v NullableV0039StepCPU) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepCPU) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


