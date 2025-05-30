# V0041OpenapiPingArrayRespPingsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hostname** | Pointer to **string** | Target for ping | [optional] 
**Pinged** | Pointer to **string** | Ping result | [optional] 
**Latency** | Pointer to **int64** | Number of microseconds it took to successfully ping or timeout | [optional] 
**Mode** | Pointer to **string** | The operating mode of the responding slurmctld | [optional] 

## Methods

### NewV0041OpenapiPingArrayRespPingsInner

`func NewV0041OpenapiPingArrayRespPingsInner() *V0041OpenapiPingArrayRespPingsInner`

NewV0041OpenapiPingArrayRespPingsInner instantiates a new V0041OpenapiPingArrayRespPingsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV0041OpenapiPingArrayRespPingsInnerWithDefaults

`func NewV0041OpenapiPingArrayRespPingsInnerWithDefaults() *V0041OpenapiPingArrayRespPingsInner`

NewV0041OpenapiPingArrayRespPingsInnerWithDefaults instantiates a new V0041OpenapiPingArrayRespPingsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostname

`func (o *V0041OpenapiPingArrayRespPingsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V0041OpenapiPingArrayRespPingsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V0041OpenapiPingArrayRespPingsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V0041OpenapiPingArrayRespPingsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetPinged

`func (o *V0041OpenapiPingArrayRespPingsInner) GetPinged() string`

GetPinged returns the Pinged field if non-nil, zero value otherwise.

### GetPingedOk

`func (o *V0041OpenapiPingArrayRespPingsInner) GetPingedOk() (*string, bool)`

GetPingedOk returns a tuple with the Pinged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinged

`func (o *V0041OpenapiPingArrayRespPingsInner) SetPinged(v string)`

SetPinged sets Pinged field to given value.

### HasPinged

`func (o *V0041OpenapiPingArrayRespPingsInner) HasPinged() bool`

HasPinged returns a boolean if a field has been set.

### GetLatency

`func (o *V0041OpenapiPingArrayRespPingsInner) GetLatency() int64`

GetLatency returns the Latency field if non-nil, zero value otherwise.

### GetLatencyOk

`func (o *V0041OpenapiPingArrayRespPingsInner) GetLatencyOk() (*int64, bool)`

GetLatencyOk returns a tuple with the Latency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatency

`func (o *V0041OpenapiPingArrayRespPingsInner) SetLatency(v int64)`

SetLatency sets Latency field to given value.

### HasLatency

`func (o *V0041OpenapiPingArrayRespPingsInner) HasLatency() bool`

HasLatency returns a boolean if a field has been set.

### GetMode

`func (o *V0041OpenapiPingArrayRespPingsInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *V0041OpenapiPingArrayRespPingsInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *V0041OpenapiPingArrayRespPingsInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *V0041OpenapiPingArrayRespPingsInner) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


