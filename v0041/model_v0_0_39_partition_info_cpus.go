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

// checks if the V0039PartitionInfoCpus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039PartitionInfoCpus{}

// V0039PartitionInfoCpus struct for V0039PartitionInfoCpus
type V0039PartitionInfoCpus struct {
	TaskBinding *int32 `json:"task_binding,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewV0039PartitionInfoCpus instantiates a new V0039PartitionInfoCpus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039PartitionInfoCpus() *V0039PartitionInfoCpus {
	this := V0039PartitionInfoCpus{}
	return &this
}

// NewV0039PartitionInfoCpusWithDefaults instantiates a new V0039PartitionInfoCpus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039PartitionInfoCpusWithDefaults() *V0039PartitionInfoCpus {
	this := V0039PartitionInfoCpus{}
	return &this
}

// GetTaskBinding returns the TaskBinding field value if set, zero value otherwise.
func (o *V0039PartitionInfoCpus) GetTaskBinding() int32 {
	if o == nil || IsNil(o.TaskBinding) {
		var ret int32
		return ret
	}
	return *o.TaskBinding
}

// GetTaskBindingOk returns a tuple with the TaskBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionInfoCpus) GetTaskBindingOk() (*int32, bool) {
	if o == nil || IsNil(o.TaskBinding) {
		return nil, false
	}
	return o.TaskBinding, true
}

// HasTaskBinding returns a boolean if a field has been set.
func (o *V0039PartitionInfoCpus) HasTaskBinding() bool {
	if o != nil && !IsNil(o.TaskBinding) {
		return true
	}

	return false
}

// SetTaskBinding gets a reference to the given int32 and assigns it to the TaskBinding field.
func (o *V0039PartitionInfoCpus) SetTaskBinding(v int32) {
	o.TaskBinding = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *V0039PartitionInfoCpus) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionInfoCpus) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *V0039PartitionInfoCpus) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *V0039PartitionInfoCpus) SetTotal(v int32) {
	o.Total = &v
}

func (o V0039PartitionInfoCpus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039PartitionInfoCpus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaskBinding) {
		toSerialize["task_binding"] = o.TaskBinding
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableV0039PartitionInfoCpus struct {
	value *V0039PartitionInfoCpus
	isSet bool
}

func (v NullableV0039PartitionInfoCpus) Get() *V0039PartitionInfoCpus {
	return v.value
}

func (v *NullableV0039PartitionInfoCpus) Set(val *V0039PartitionInfoCpus) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039PartitionInfoCpus) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039PartitionInfoCpus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039PartitionInfoCpus(val *V0039PartitionInfoCpus) *NullableV0039PartitionInfoCpus {
	return &NullableV0039PartitionInfoCpus{value: val, isSet: true}
}

func (v NullableV0039PartitionInfoCpus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039PartitionInfoCpus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


