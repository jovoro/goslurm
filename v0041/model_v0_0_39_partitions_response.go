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

// checks if the V0039PartitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039PartitionsResponse{}

// V0039PartitionsResponse struct for V0039PartitionsResponse
type V0039PartitionsResponse struct {
	Meta *V0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []V0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []V0039Warning `json:"warnings,omitempty"`
	Partitions []V0039PartitionInfo `json:"partitions,omitempty"`
}

// NewV0039PartitionsResponse instantiates a new V0039PartitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039PartitionsResponse() *V0039PartitionsResponse {
	this := V0039PartitionsResponse{}
	return &this
}

// NewV0039PartitionsResponseWithDefaults instantiates a new V0039PartitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039PartitionsResponseWithDefaults() *V0039PartitionsResponse {
	this := V0039PartitionsResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0039PartitionsResponse) GetMeta() V0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionsResponse) GetMetaOk() (*V0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0039PartitionsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0039Meta and assigns it to the Meta field.
func (o *V0039PartitionsResponse) SetMeta(v V0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0039PartitionsResponse) GetErrors() []V0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionsResponse) GetErrorsOk() ([]V0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0039PartitionsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0039Error and assigns it to the Errors field.
func (o *V0039PartitionsResponse) SetErrors(v []V0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *V0039PartitionsResponse) GetWarnings() []V0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []V0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionsResponse) GetWarningsOk() ([]V0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *V0039PartitionsResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []V0039Warning and assigns it to the Warnings field.
func (o *V0039PartitionsResponse) SetWarnings(v []V0039Warning) {
	o.Warnings = v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *V0039PartitionsResponse) GetPartitions() []V0039PartitionInfo {
	if o == nil || IsNil(o.Partitions) {
		var ret []V0039PartitionInfo
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039PartitionsResponse) GetPartitionsOk() ([]V0039PartitionInfo, bool) {
	if o == nil || IsNil(o.Partitions) {
		return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *V0039PartitionsResponse) HasPartitions() bool {
	if o != nil && !IsNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []V0039PartitionInfo and assigns it to the Partitions field.
func (o *V0039PartitionsResponse) SetPartitions(v []V0039PartitionInfo) {
	o.Partitions = v
}

func (o V0039PartitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039PartitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !IsNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	return toSerialize, nil
}

type NullableV0039PartitionsResponse struct {
	value *V0039PartitionsResponse
	isSet bool
}

func (v NullableV0039PartitionsResponse) Get() *V0039PartitionsResponse {
	return v.value
}

func (v *NullableV0039PartitionsResponse) Set(val *V0039PartitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039PartitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039PartitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039PartitionsResponse(val *V0039PartitionsResponse) *NullableV0039PartitionsResponse {
	return &NullableV0039PartitionsResponse{value: val, isSet: true}
}

func (v NullableV0039PartitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039PartitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


