/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SchemasUpdateMetricPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUpdateMetricPayload{}

// SchemasUpdateMetricPayload struct for SchemasUpdateMetricPayload
type SchemasUpdateMetricPayload struct {
	Description *string
	EstimationTime *int32
	IsDone *bool
}

// NewSchemasUpdateMetricPayload instantiates a new SchemasUpdateMetricPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUpdateMetricPayload() *SchemasUpdateMetricPayload {
	this := SchemasUpdateMetricPayload{}
	return &this
}

// NewSchemasUpdateMetricPayloadWithDefaults instantiates a new SchemasUpdateMetricPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUpdateMetricPayloadWithDefaults() *SchemasUpdateMetricPayload {
	this := SchemasUpdateMetricPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemasUpdateMetricPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateMetricPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemasUpdateMetricPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemasUpdateMetricPayload) SetDescription(v string) {
	o.Description = &v
}

// GetEstimationTime returns the EstimationTime field value if set, zero value otherwise.
func (o *SchemasUpdateMetricPayload) GetEstimationTime() int32 {
	if o == nil || IsNil(o.EstimationTime) {
		var ret int32
		return ret
	}
	return *o.EstimationTime
}

// GetEstimationTimeOk returns a tuple with the EstimationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateMetricPayload) GetEstimationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimationTime) {
		return nil, false
	}
	return o.EstimationTime, true
}

// HasEstimationTime returns a boolean if a field has been set.
func (o *SchemasUpdateMetricPayload) HasEstimationTime() bool {
	if o != nil && !IsNil(o.EstimationTime) {
		return true
	}

	return false
}

// SetEstimationTime gets a reference to the given int32 and assigns it to the EstimationTime field.
func (o *SchemasUpdateMetricPayload) SetEstimationTime(v int32) {
	o.EstimationTime = &v
}

// GetIsDone returns the IsDone field value if set, zero value otherwise.
func (o *SchemasUpdateMetricPayload) GetIsDone() bool {
	if o == nil || IsNil(o.IsDone) {
		var ret bool
		return ret
	}
	return *o.IsDone
}

// GetIsDoneOk returns a tuple with the IsDone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemasUpdateMetricPayload) GetIsDoneOk() (*bool, bool) {
	if o == nil || IsNil(o.IsDone) {
		return nil, false
	}
	return o.IsDone, true
}

// HasIsDone returns a boolean if a field has been set.
func (o *SchemasUpdateMetricPayload) HasIsDone() bool {
	if o != nil && !IsNil(o.IsDone) {
		return true
	}

	return false
}

// SetIsDone gets a reference to the given bool and assigns it to the IsDone field.
func (o *SchemasUpdateMetricPayload) SetIsDone(v bool) {
	o.IsDone = &v
}

func (o SchemasUpdateMetricPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUpdateMetricPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EstimationTime) {
		toSerialize["estimationTime"] = o.EstimationTime
	}
	if !IsNil(o.IsDone) {
		toSerialize["isDone"] = o.IsDone
	}
	return toSerialize, nil
}

type NullableSchemasUpdateMetricPayload struct {
	value *SchemasUpdateMetricPayload
	isSet bool
}

func (v NullableSchemasUpdateMetricPayload) Get() *SchemasUpdateMetricPayload {
	return v.value
}

func (v *NullableSchemasUpdateMetricPayload) Set(val *SchemasUpdateMetricPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUpdateMetricPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUpdateMetricPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUpdateMetricPayload(val *SchemasUpdateMetricPayload) *NullableSchemasUpdateMetricPayload {
	return &NullableSchemasUpdateMetricPayload{value: val, isSet: true}
}

func (v NullableSchemasUpdateMetricPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUpdateMetricPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


