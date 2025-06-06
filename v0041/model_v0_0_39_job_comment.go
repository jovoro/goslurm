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

// checks if the V0039JobComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039JobComment{}

// V0039JobComment struct for V0039JobComment
type V0039JobComment struct {
	Administrator *string `json:"administrator,omitempty"`
	Job *string `json:"job,omitempty"`
	System *string `json:"system,omitempty"`
}

// NewV0039JobComment instantiates a new V0039JobComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039JobComment() *V0039JobComment {
	this := V0039JobComment{}
	return &this
}

// NewV0039JobCommentWithDefaults instantiates a new V0039JobComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039JobCommentWithDefaults() *V0039JobComment {
	this := V0039JobComment{}
	return &this
}

// GetAdministrator returns the Administrator field value if set, zero value otherwise.
func (o *V0039JobComment) GetAdministrator() string {
	if o == nil || IsNil(o.Administrator) {
		var ret string
		return ret
	}
	return *o.Administrator
}

// GetAdministratorOk returns a tuple with the Administrator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobComment) GetAdministratorOk() (*string, bool) {
	if o == nil || IsNil(o.Administrator) {
		return nil, false
	}
	return o.Administrator, true
}

// HasAdministrator returns a boolean if a field has been set.
func (o *V0039JobComment) HasAdministrator() bool {
	if o != nil && !IsNil(o.Administrator) {
		return true
	}

	return false
}

// SetAdministrator gets a reference to the given string and assigns it to the Administrator field.
func (o *V0039JobComment) SetAdministrator(v string) {
	o.Administrator = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *V0039JobComment) GetJob() string {
	if o == nil || IsNil(o.Job) {
		var ret string
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobComment) GetJobOk() (*string, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *V0039JobComment) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given string and assigns it to the Job field.
func (o *V0039JobComment) SetJob(v string) {
	o.Job = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *V0039JobComment) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039JobComment) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *V0039JobComment) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *V0039JobComment) SetSystem(v string) {
	o.System = &v
}

func (o V0039JobComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039JobComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Administrator) {
		toSerialize["administrator"] = o.Administrator
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	return toSerialize, nil
}

type NullableV0039JobComment struct {
	value *V0039JobComment
	isSet bool
}

func (v NullableV0039JobComment) Get() *V0039JobComment {
	return v.value
}

func (v *NullableV0039JobComment) Set(val *V0039JobComment) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039JobComment) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039JobComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039JobComment(val *V0039JobComment) *NullableV0039JobComment {
	return &NullableV0039JobComment{value: val, isSet: true}
}

func (v NullableV0039JobComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039JobComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


