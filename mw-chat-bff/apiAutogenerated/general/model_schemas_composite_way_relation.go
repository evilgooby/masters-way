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

// checks if the SchemasCompositeWayRelation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasCompositeWayRelation{}

// SchemasCompositeWayRelation struct for SchemasCompositeWayRelation
type SchemasCompositeWayRelation struct {
	ChildWayUuid string
	ParentWayUuid string
}

type _SchemasCompositeWayRelation SchemasCompositeWayRelation

// NewSchemasCompositeWayRelation instantiates a new SchemasCompositeWayRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasCompositeWayRelation(childWayUuid string, parentWayUuid string) *SchemasCompositeWayRelation {
	this := SchemasCompositeWayRelation{}
	this.ChildWayUuid = childWayUuid
	this.ParentWayUuid = parentWayUuid
	return &this
}

// NewSchemasCompositeWayRelationWithDefaults instantiates a new SchemasCompositeWayRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasCompositeWayRelationWithDefaults() *SchemasCompositeWayRelation {
	this := SchemasCompositeWayRelation{}
	return &this
}

// GetChildWayUuid returns the ChildWayUuid field value
func (o *SchemasCompositeWayRelation) GetChildWayUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChildWayUuid
}

// GetChildWayUuidOk returns a tuple with the ChildWayUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCompositeWayRelation) GetChildWayUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChildWayUuid, true
}

// SetChildWayUuid sets field value
func (o *SchemasCompositeWayRelation) SetChildWayUuid(v string) {
	o.ChildWayUuid = v
}

// GetParentWayUuid returns the ParentWayUuid field value
func (o *SchemasCompositeWayRelation) GetParentWayUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentWayUuid
}

// GetParentWayUuidOk returns a tuple with the ParentWayUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasCompositeWayRelation) GetParentWayUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentWayUuid, true
}

// SetParentWayUuid sets field value
func (o *SchemasCompositeWayRelation) SetParentWayUuid(v string) {
	o.ParentWayUuid = v
}

func (o SchemasCompositeWayRelation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasCompositeWayRelation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["childWayUuid"] = o.ChildWayUuid
	toSerialize["parentWayUuid"] = o.ParentWayUuid
	return toSerialize, nil
}

func (o *SchemasCompositeWayRelation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"childWayUuid",
		"parentWayUuid",
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

	varSchemasCompositeWayRelation := _SchemasCompositeWayRelation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasCompositeWayRelation)

	if err != nil {
		return err
	}

	*o = SchemasCompositeWayRelation(varSchemasCompositeWayRelation)

	return err
}

type NullableSchemasCompositeWayRelation struct {
	value *SchemasCompositeWayRelation
	isSet bool
}

func (v NullableSchemasCompositeWayRelation) Get() *SchemasCompositeWayRelation {
	return v.value
}

func (v *NullableSchemasCompositeWayRelation) Set(val *SchemasCompositeWayRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasCompositeWayRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasCompositeWayRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasCompositeWayRelation(val *SchemasCompositeWayRelation) *NullableSchemasCompositeWayRelation {
	return &NullableSchemasCompositeWayRelation{value: val, isSet: true}
}

func (v NullableSchemasCompositeWayRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasCompositeWayRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


