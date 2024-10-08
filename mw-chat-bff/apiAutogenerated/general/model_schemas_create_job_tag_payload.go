/*
Masters way general API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasCreateJobTagPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCreateJobTagPayload{}

// SchemasCreateJobTagPayload struct for SchemasCreateJobTagPayload
type SchemasCreateJobTagPayload struct {
	Color string
	Description string
	Name string
	WayUuid string
}

type _SchemasCreateJobTagPayload SchemasCreateJobTagPayload

// NewSchemasCreateJobTagPayload instantiates a new SchemasCreateJobTagPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCreateJobTagPayload(color string, description string, name string, wayUuid string) *SchemasCreateJobTagPayload {
	this := SchemasCreateJobTagPayload{}
	this.Color = color
	this.Description = description
	this.Name = name
	this.WayUuid = wayUuid
	return &this
}

// NewSchemasCreateJobTagPayloadWithDefaults instantiates a new SchemasCreateJobTagPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCreateJobTagPayloadWithDefaults() *SchemasCreateJobTagPayload {
	this := SchemasCreateJobTagPayload{}
	return &this
}

// GetColor returns the Color field value
func (o *SchemasCreateJobTagPayload) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobTagPayload) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *SchemasCreateJobTagPayload) SetColor(v string) {
	o.Color = v
}

// GetDescription returns the Description field value
func (o *SchemasCreateJobTagPayload) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobTagPayload) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemasCreateJobTagPayload) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *SchemasCreateJobTagPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobTagPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasCreateJobTagPayload) SetName(v string) {
	o.Name = v
}

// GetWayUuid returns the WayUuid field value
func (o *SchemasCreateJobTagPayload) GetWayUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WayUuid
}

// GetWayUuidOk returns a tuple with the WayUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCreateJobTagPayload) GetWayUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WayUuid, true
}

// SetWayUuid sets field value
func (o *SchemasCreateJobTagPayload) SetWayUuid(v string) {
	o.WayUuid = v
}

func (o SchemasCreateJobTagPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCreateJobTagPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["color"] = o.Color
	toSerialize["description"] = o.Description
	toSerialize["name"] = o.Name
	toSerialize["wayUuid"] = o.WayUuid
	return toSerialize, nil
}

func (o *SchemasCreateJobTagPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"color",
		"description",
		"name",
		"wayUuid",
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

	varSchemasCreateJobTagPayload := _SchemasCreateJobTagPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCreateJobTagPayload)

	if err != nil {
		return err
	}

	*o = SchemasCreateJobTagPayload(varSchemasCreateJobTagPayload)

	return err
}

type NullableSchemasCreateJobTagPayload struct {
	value *SchemasCreateJobTagPayload
	isSet bool
}

func (v NullableSchemasCreateJobTagPayload) Get() *SchemasCreateJobTagPayload {
	return v.value
}

func (v *NullableSchemasCreateJobTagPayload) Set(val *SchemasCreateJobTagPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCreateJobTagPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCreateJobTagPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCreateJobTagPayload(val *SchemasCreateJobTagPayload) *NullableSchemasCreateJobTagPayload {
	return &NullableSchemasCreateJobTagPayload{value: val, isSet: true}
}

func (v NullableSchemasCreateJobTagPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCreateJobTagPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


