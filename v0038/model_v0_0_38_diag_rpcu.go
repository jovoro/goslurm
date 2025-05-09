/*
Slurm Rest API RO

API to access Slurm. Only GET requests are implemented.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v0038

import (
	"encoding/json"
)

// checks if the V0038DiagRpcu type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038DiagRpcu{}

// V0038DiagRpcu struct for V0038DiagRpcu
type V0038DiagRpcu struct {
	// user
	User *string `json:"user,omitempty"`
	// user id
	UserId *int32 `json:"user_id,omitempty"`
	// rpc count
	Count *int32 `json:"count,omitempty"`
	// average time
	AverageTime *int32 `json:"average_time,omitempty"`
	// total time
	TotalTime *int32 `json:"total_time,omitempty"`
}

// NewV0038DiagRpcu instantiates a new V0038DiagRpcu object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038DiagRpcu() *V0038DiagRpcu {
	this := V0038DiagRpcu{}
	return &this
}

// NewV0038DiagRpcuWithDefaults instantiates a new V0038DiagRpcu object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038DiagRpcuWithDefaults() *V0038DiagRpcu {
	this := V0038DiagRpcu{}
	return &this
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *V0038DiagRpcu) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038DiagRpcu) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *V0038DiagRpcu) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *V0038DiagRpcu) SetUser(v string) {
	o.User = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *V0038DiagRpcu) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038DiagRpcu) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *V0038DiagRpcu) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *V0038DiagRpcu) SetUserId(v int32) {
	o.UserId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V0038DiagRpcu) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038DiagRpcu) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V0038DiagRpcu) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V0038DiagRpcu) SetCount(v int32) {
	o.Count = &v
}

// GetAverageTime returns the AverageTime field value if set, zero value otherwise.
func (o *V0038DiagRpcu) GetAverageTime() int32 {
	if o == nil || IsNil(o.AverageTime) {
		var ret int32
		return ret
	}
	return *o.AverageTime
}

// GetAverageTimeOk returns a tuple with the AverageTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038DiagRpcu) GetAverageTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AverageTime) {
		return nil, false
	}
	return o.AverageTime, true
}

// HasAverageTime returns a boolean if a field has been set.
func (o *V0038DiagRpcu) HasAverageTime() bool {
	if o != nil && !IsNil(o.AverageTime) {
		return true
	}

	return false
}

// SetAverageTime gets a reference to the given int32 and assigns it to the AverageTime field.
func (o *V0038DiagRpcu) SetAverageTime(v int32) {
	o.AverageTime = &v
}

// GetTotalTime returns the TotalTime field value if set, zero value otherwise.
func (o *V0038DiagRpcu) GetTotalTime() int32 {
	if o == nil || IsNil(o.TotalTime) {
		var ret int32
		return ret
	}
	return *o.TotalTime
}

// GetTotalTimeOk returns a tuple with the TotalTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038DiagRpcu) GetTotalTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalTime) {
		return nil, false
	}
	return o.TotalTime, true
}

// HasTotalTime returns a boolean if a field has been set.
func (o *V0038DiagRpcu) HasTotalTime() bool {
	if o != nil && !IsNil(o.TotalTime) {
		return true
	}

	return false
}

// SetTotalTime gets a reference to the given int32 and assigns it to the TotalTime field.
func (o *V0038DiagRpcu) SetTotalTime(v int32) {
	o.TotalTime = &v
}

func (o V0038DiagRpcu) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038DiagRpcu) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.AverageTime) {
		toSerialize["average_time"] = o.AverageTime
	}
	if !IsNil(o.TotalTime) {
		toSerialize["total_time"] = o.TotalTime
	}
	return toSerialize, nil
}

type NullableV0038DiagRpcu struct {
	value *V0038DiagRpcu
	isSet bool
}

func (v NullableV0038DiagRpcu) Get() *V0038DiagRpcu {
	return v.value
}

func (v *NullableV0038DiagRpcu) Set(val *V0038DiagRpcu) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038DiagRpcu) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038DiagRpcu) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038DiagRpcu(val *V0038DiagRpcu) *NullableV0038DiagRpcu {
	return &NullableV0038DiagRpcu{value: val, isSet: true}
}

func (v NullableV0038DiagRpcu) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038DiagRpcu) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


