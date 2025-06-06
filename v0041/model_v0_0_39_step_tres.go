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

// checks if the V0039StepTres type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039StepTres{}

// V0039StepTres struct for V0039StepTres
type V0039StepTres struct {
	Requested *V0039StepTresRequested `json:"requested,omitempty"`
	Consumed *V0039StepTresConsumed `json:"consumed,omitempty"`
	Allocated []V0039Tres `json:"allocated,omitempty"`
}

// NewV0039StepTres instantiates a new V0039StepTres object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039StepTres() *V0039StepTres {
	this := V0039StepTres{}
	return &this
}

// NewV0039StepTresWithDefaults instantiates a new V0039StepTres object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039StepTresWithDefaults() *V0039StepTres {
	this := V0039StepTres{}
	return &this
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *V0039StepTres) GetRequested() V0039StepTresRequested {
	if o == nil || IsNil(o.Requested) {
		var ret V0039StepTresRequested
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTres) GetRequestedOk() (*V0039StepTresRequested, bool) {
	if o == nil || IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *V0039StepTres) HasRequested() bool {
	if o != nil && !IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given V0039StepTresRequested and assigns it to the Requested field.
func (o *V0039StepTres) SetRequested(v V0039StepTresRequested) {
	o.Requested = &v
}

// GetConsumed returns the Consumed field value if set, zero value otherwise.
func (o *V0039StepTres) GetConsumed() V0039StepTresConsumed {
	if o == nil || IsNil(o.Consumed) {
		var ret V0039StepTresConsumed
		return ret
	}
	return *o.Consumed
}

// GetConsumedOk returns a tuple with the Consumed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTres) GetConsumedOk() (*V0039StepTresConsumed, bool) {
	if o == nil || IsNil(o.Consumed) {
		return nil, false
	}
	return o.Consumed, true
}

// HasConsumed returns a boolean if a field has been set.
func (o *V0039StepTres) HasConsumed() bool {
	if o != nil && !IsNil(o.Consumed) {
		return true
	}

	return false
}

// SetConsumed gets a reference to the given V0039StepTresConsumed and assigns it to the Consumed field.
func (o *V0039StepTres) SetConsumed(v V0039StepTresConsumed) {
	o.Consumed = &v
}

// GetAllocated returns the Allocated field value if set, zero value otherwise.
func (o *V0039StepTres) GetAllocated() []V0039Tres {
	if o == nil || IsNil(o.Allocated) {
		var ret []V0039Tres
		return ret
	}
	return o.Allocated
}

// GetAllocatedOk returns a tuple with the Allocated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039StepTres) GetAllocatedOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Allocated) {
		return nil, false
	}
	return o.Allocated, true
}

// HasAllocated returns a boolean if a field has been set.
func (o *V0039StepTres) HasAllocated() bool {
	if o != nil && !IsNil(o.Allocated) {
		return true
	}

	return false
}

// SetAllocated gets a reference to the given []V0039Tres and assigns it to the Allocated field.
func (o *V0039StepTres) SetAllocated(v []V0039Tres) {
	o.Allocated = v
}

func (o V0039StepTres) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039StepTres) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	if !IsNil(o.Consumed) {
		toSerialize["consumed"] = o.Consumed
	}
	if !IsNil(o.Allocated) {
		toSerialize["allocated"] = o.Allocated
	}
	return toSerialize, nil
}

type NullableV0039StepTres struct {
	value *V0039StepTres
	isSet bool
}

func (v NullableV0039StepTres) Get() *V0039StepTres {
	return v.value
}

func (v *NullableV0039StepTres) Set(val *V0039StepTres) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039StepTres) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039StepTres) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039StepTres(val *V0039StepTres) *NullableV0039StepTres {
	return &NullableV0039StepTres{value: val, isSet: true}
}

func (v NullableV0039StepTres) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039StepTres) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


