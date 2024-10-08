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

// checks if the SchemasDefaultWayCollections type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasDefaultWayCollections{}

// SchemasDefaultWayCollections struct for SchemasDefaultWayCollections
type SchemasDefaultWayCollections struct {
	Favorite SchemasWayCollectionPopulatedResponse
	Mentoring SchemasWayCollectionPopulatedResponse
	Own SchemasWayCollectionPopulatedResponse
}

type _SchemasDefaultWayCollections SchemasDefaultWayCollections

// NewSchemasDefaultWayCollections instantiates a new SchemasDefaultWayCollections object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasDefaultWayCollections(favorite SchemasWayCollectionPopulatedResponse, mentoring SchemasWayCollectionPopulatedResponse, own SchemasWayCollectionPopulatedResponse) *SchemasDefaultWayCollections {
	this := SchemasDefaultWayCollections{}
	this.Favorite = favorite
	this.Mentoring = mentoring
	this.Own = own
	return &this
}

// NewSchemasDefaultWayCollectionsWithDefaults instantiates a new SchemasDefaultWayCollections object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasDefaultWayCollectionsWithDefaults() *SchemasDefaultWayCollections {
	this := SchemasDefaultWayCollections{}
	return &this
}

// GetFavorite returns the Favorite field value
func (o *SchemasDefaultWayCollections) GetFavorite() SchemasWayCollectionPopulatedResponse {
	if o == nil {
		var ret SchemasWayCollectionPopulatedResponse
		return ret
	}

	return o.Favorite
}

// GetFavoriteOk returns a tuple with the Favorite field value
// and a boolean to check if the value has been set.
func (o *SchemasDefaultWayCollections) GetFavoriteOk() (*SchemasWayCollectionPopulatedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Favorite, true
}

// SetFavorite sets field value
func (o *SchemasDefaultWayCollections) SetFavorite(v SchemasWayCollectionPopulatedResponse) {
	o.Favorite = v
}

// GetMentoring returns the Mentoring field value
func (o *SchemasDefaultWayCollections) GetMentoring() SchemasWayCollectionPopulatedResponse {
	if o == nil {
		var ret SchemasWayCollectionPopulatedResponse
		return ret
	}

	return o.Mentoring
}

// GetMentoringOk returns a tuple with the Mentoring field value
// and a boolean to check if the value has been set.
func (o *SchemasDefaultWayCollections) GetMentoringOk() (*SchemasWayCollectionPopulatedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mentoring, true
}

// SetMentoring sets field value
func (o *SchemasDefaultWayCollections) SetMentoring(v SchemasWayCollectionPopulatedResponse) {
	o.Mentoring = v
}

// GetOwn returns the Own field value
func (o *SchemasDefaultWayCollections) GetOwn() SchemasWayCollectionPopulatedResponse {
	if o == nil {
		var ret SchemasWayCollectionPopulatedResponse
		return ret
	}

	return o.Own
}

// GetOwnOk returns a tuple with the Own field value
// and a boolean to check if the value has been set.
func (o *SchemasDefaultWayCollections) GetOwnOk() (*SchemasWayCollectionPopulatedResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Own, true
}

// SetOwn sets field value
func (o *SchemasDefaultWayCollections) SetOwn(v SchemasWayCollectionPopulatedResponse) {
	o.Own = v
}

func (o SchemasDefaultWayCollections) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasDefaultWayCollections) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["favorite"] = o.Favorite
	toSerialize["mentoring"] = o.Mentoring
	toSerialize["own"] = o.Own
	return toSerialize, nil
}

func (o *SchemasDefaultWayCollections) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"favorite",
		"mentoring",
		"own",
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

	varSchemasDefaultWayCollections := _SchemasDefaultWayCollections{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasDefaultWayCollections)

	if err != nil {
		return err
	}

	*o = SchemasDefaultWayCollections(varSchemasDefaultWayCollections)

	return err
}

type NullableSchemasDefaultWayCollections struct {
	value *SchemasDefaultWayCollections
	isSet bool
}

func (v NullableSchemasDefaultWayCollections) Get() *SchemasDefaultWayCollections {
	return v.value
}

func (v *NullableSchemasDefaultWayCollections) Set(val *SchemasDefaultWayCollections) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasDefaultWayCollections) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasDefaultWayCollections) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasDefaultWayCollections(val *SchemasDefaultWayCollections) *NullableSchemasDefaultWayCollections {
	return &NullableSchemasDefaultWayCollections{value: val, isSet: true}
}

func (v NullableSchemasDefaultWayCollections) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasDefaultWayCollections) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


