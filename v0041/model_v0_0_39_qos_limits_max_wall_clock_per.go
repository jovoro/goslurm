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

// checks if the V0039QosLimitsMaxWallClockPer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039QosLimitsMaxWallClockPer{}

// V0039QosLimitsMaxWallClockPer struct for V0039QosLimitsMaxWallClockPer
type V0039QosLimitsMaxWallClockPer struct {
	Qos *V0039Uint32NoVal `json:"qos,omitempty"`
	Job *V0039Uint32NoVal `json:"job,omitempty"`
}

// NewV0039QosLimitsMaxWallClockPer instantiates a new V0039QosLimitsMaxWallClockPer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039QosLimitsMaxWallClockPer() *V0039QosLimitsMaxWallClockPer {
	this := V0039QosLimitsMaxWallClockPer{}
	return &this
}

// NewV0039QosLimitsMaxWallClockPerWithDefaults instantiates a new V0039QosLimitsMaxWallClockPer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039QosLimitsMaxWallClockPerWithDefaults() *V0039QosLimitsMaxWallClockPer {
	this := V0039QosLimitsMaxWallClockPer{}
	return &this
}

// GetQos returns the Qos field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxWallClockPer) GetQos() V0039Uint32NoVal {
	if o == nil || IsNil(o.Qos) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Qos
}

// GetQosOk returns a tuple with the Qos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxWallClockPer) GetQosOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Qos) {
		return nil, false
	}
	return o.Qos, true
}

// HasQos returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxWallClockPer) HasQos() bool {
	if o != nil && !IsNil(o.Qos) {
		return true
	}

	return false
}

// SetQos gets a reference to the given V0039Uint32NoVal and assigns it to the Qos field.
func (o *V0039QosLimitsMaxWallClockPer) SetQos(v V0039Uint32NoVal) {
	o.Qos = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0039QosLimitsMaxWallClockPer) GetJob() V0039Uint32NoVal {
	if o == nil || IsNil(o.Job) {
		var ret V0039Uint32NoVal
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039QosLimitsMaxWallClockPer) GetJobOk() (*V0039Uint32NoVal, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0039QosLimitsMaxWallClockPer) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given V0039Uint32NoVal and assigns it to the Job field.
func (o *V0039QosLimitsMaxWallClockPer) SetJob(v V0039Uint32NoVal) {
	o.Job = &v
}

func (o V0039QosLimitsMaxWallClockPer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039QosLimitsMaxWallClockPer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Qos) {
		toSerialize["qos"] = o.Qos
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableV0039QosLimitsMaxWallClockPer struct {
	value *V0039QosLimitsMaxWallClockPer
	isSet bool
}

func (v NullableV0039QosLimitsMaxWallClockPer) Get() *V0039QosLimitsMaxWallClockPer {
	return v.value
}

func (v *NullableV0039QosLimitsMaxWallClockPer) Set(val *V0039QosLimitsMaxWallClockPer) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039QosLimitsMaxWallClockPer) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039QosLimitsMaxWallClockPer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039QosLimitsMaxWallClockPer(val *V0039QosLimitsMaxWallClockPer) *NullableV0039QosLimitsMaxWallClockPer {
	return &NullableV0039QosLimitsMaxWallClockPer{value: val, isSet: true}
}

func (v NullableV0039QosLimitsMaxWallClockPer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039QosLimitsMaxWallClockPer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


