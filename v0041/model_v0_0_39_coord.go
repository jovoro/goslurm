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
	"bytes"
	"fmt"
)

// checks if the V0039Coord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0039Coord{}

// V0039Coord struct for V0039Coord
type V0039Coord struct {
	Name string `json:"name"`
	Direct *bool `json:"direct,omitempty"`
}

type _V0039Coord V0039Coord

// NewV0039Coord instantiates a new V0039Coord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0039Coord(name string) *V0039Coord {
	this := V0039Coord{}
	this.Name = name
	return &this
}

// NewV0039CoordWithDefaults instantiates a new V0039Coord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0039CoordWithDefaults() *V0039Coord {
	this := V0039Coord{}
	return &this
}

// GetName returns the Name field value
func (o *V0039Coord) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V0039Coord) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V0039Coord) SetName(v string) {
	o.Name = v
}

// GetDirect returns the Direct field value if set, zero value otherwise.
func (o *V0039Coord) GetDirect() bool {
	if o == nil || IsNil(o.Direct) {
		var ret bool
		return ret
	}
	return *o.Direct
}

// GetDirectOk returns a tuple with the Direct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0039Coord) GetDirectOk() (*bool, bool) {
	if o == nil || IsNil(o.Direct) {
		return nil, false
	}
	return o.Direct, true
}

// HasDirect returns a boolean if a field has been set.
func (o *V0039Coord) HasDirect() bool {
	if o != nil && !IsNil(o.Direct) {
		return true
	}

	return false
}

// SetDirect gets a reference to the given bool and assigns it to the Direct field.
func (o *V0039Coord) SetDirect(v bool) {
	o.Direct = &v
}

func (o V0039Coord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0039Coord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Direct) {
		toSerialize["direct"] = o.Direct
	}
	return toSerialize, nil
}

func (o *V0039Coord) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV0039Coord := _V0039Coord{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV0039Coord)

	if err != nil {
		return err
	}

	*o = V0039Coord(varV0039Coord)

	return err
}

type NullableV0039Coord struct {
	value *V0039Coord
	isSet bool
}

func (v NullableV0039Coord) Get() *V0039Coord {
	return v.value
}

func (v *NullableV0039Coord) Set(val *V0039Coord) {
	v.value = val
	v.isSet = true
}

func (v NullableV0039Coord) IsSet() bool {
	return v.isSet
}

func (v *NullableV0039Coord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0039Coord(val *V0039Coord) *NullableV0039Coord {
	return &NullableV0039Coord{value: val, isSet: true}
}

func (v NullableV0039Coord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0039Coord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


