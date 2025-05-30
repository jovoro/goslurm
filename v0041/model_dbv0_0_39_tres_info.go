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

// checks if the Dbv0039TresInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dbv0039TresInfo{}

// Dbv0039TresInfo struct for Dbv0039TresInfo
type Dbv0039TresInfo struct {
	Meta *Dbv0039Meta `json:"meta,omitempty"`
	// Slurm errors
	Errors []Dbv0039Error `json:"errors,omitempty"`
	// Slurm warnings
	Warnings []Dbv0039Warning `json:"warnings,omitempty"`
	Tres []V0039Tres `json:"tres,omitempty"`
}

// NewDbv0039TresInfo instantiates a new Dbv0039TresInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDbv0039TresInfo() *Dbv0039TresInfo {
	this := Dbv0039TresInfo{}
	return &this
}

// NewDbv0039TresInfoWithDefaults instantiates a new Dbv0039TresInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDbv0039TresInfoWithDefaults() *Dbv0039TresInfo {
	this := Dbv0039TresInfo{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Dbv0039TresInfo) GetMeta() Dbv0039Meta {
	if o == nil || IsNil(o.Meta) {
		var ret Dbv0039Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039TresInfo) GetMetaOk() (*Dbv0039Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Dbv0039TresInfo) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given Dbv0039Meta and assigns it to the Meta field.
func (o *Dbv0039TresInfo) SetMeta(v Dbv0039Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *Dbv0039TresInfo) GetErrors() []Dbv0039Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Dbv0039Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039TresInfo) GetErrorsOk() ([]Dbv0039Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *Dbv0039TresInfo) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Dbv0039Error and assigns it to the Errors field.
func (o *Dbv0039TresInfo) SetErrors(v []Dbv0039Error) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *Dbv0039TresInfo) GetWarnings() []Dbv0039Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Dbv0039Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039TresInfo) GetWarningsOk() ([]Dbv0039Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *Dbv0039TresInfo) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Dbv0039Warning and assigns it to the Warnings field.
func (o *Dbv0039TresInfo) SetWarnings(v []Dbv0039Warning) {
	o.Warnings = v
}

// GetTres returns the Tres field value if set, zero value otherwise.
func (o *Dbv0039TresInfo) GetTres() []V0039Tres {
	if o == nil || IsNil(o.Tres) {
		var ret []V0039Tres
		return ret
	}
	return o.Tres
}

// GetTresOk returns a tuple with the Tres field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dbv0039TresInfo) GetTresOk() ([]V0039Tres, bool) {
	if o == nil || IsNil(o.Tres) {
		return nil, false
	}
	return o.Tres, true
}

// HasTres returns a boolean if a field has been set.
func (o *Dbv0039TresInfo) HasTres() bool {
	if o != nil && !IsNil(o.Tres) {
		return true
	}

	return false
}

// SetTres gets a reference to the given []V0039Tres and assigns it to the Tres field.
func (o *Dbv0039TresInfo) SetTres(v []V0039Tres) {
	o.Tres = v
}

func (o Dbv0039TresInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dbv0039TresInfo) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Tres) {
		toSerialize["tres"] = o.Tres
	}
	return toSerialize, nil
}

type NullableDbv0039TresInfo struct {
	value *Dbv0039TresInfo
	isSet bool
}

func (v NullableDbv0039TresInfo) Get() *Dbv0039TresInfo {
	return v.value
}

func (v *NullableDbv0039TresInfo) Set(val *Dbv0039TresInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDbv0039TresInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDbv0039TresInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDbv0039TresInfo(val *Dbv0039TresInfo) *NullableDbv0039TresInfo {
	return &NullableDbv0039TresInfo{value: val, isSet: true}
}

func (v NullableDbv0039TresInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDbv0039TresInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


