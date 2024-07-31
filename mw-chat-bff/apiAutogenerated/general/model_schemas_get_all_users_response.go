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

// checks if the SchemasGetAllUsersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasGetAllUsersResponse{}

// SchemasGetAllUsersResponse struct for SchemasGetAllUsersResponse
type SchemasGetAllUsersResponse struct {
	Size int32 `json:"size"`
	Users []SchemasUserPlainResponseWithInfo `json:"users"`
}

type _SchemasGetAllUsersResponse SchemasGetAllUsersResponse

// NewSchemasGetAllUsersResponse instantiates a new SchemasGetAllUsersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasGetAllUsersResponse(size int32, users []SchemasUserPlainResponseWithInfo) *SchemasGetAllUsersResponse {
	this := SchemasGetAllUsersResponse{}
	this.Size = size
	this.Users = users
	return &this
}

// NewSchemasGetAllUsersResponseWithDefaults instantiates a new SchemasGetAllUsersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasGetAllUsersResponseWithDefaults() *SchemasGetAllUsersResponse {
	this := SchemasGetAllUsersResponse{}
	return &this
}

// GetSize returns the Size field value
func (o *SchemasGetAllUsersResponse) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *SchemasGetAllUsersResponse) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *SchemasGetAllUsersResponse) SetSize(v int32) {
	o.Size = v
}

// GetUsers returns the Users field value
func (o *SchemasGetAllUsersResponse) GetUsers() []SchemasUserPlainResponseWithInfo {
	if o == nil {
		var ret []SchemasUserPlainResponseWithInfo
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *SchemasGetAllUsersResponse) GetUsersOk() ([]SchemasUserPlainResponseWithInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *SchemasGetAllUsersResponse) SetUsers(v []SchemasUserPlainResponseWithInfo) {
	o.Users = v
}

func (o SchemasGetAllUsersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasGetAllUsersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *SchemasGetAllUsersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
		"users",
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

	varSchemasGetAllUsersResponse := _SchemasGetAllUsersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasGetAllUsersResponse)

	if err != nil {
		return err
	}

	*o = SchemasGetAllUsersResponse(varSchemasGetAllUsersResponse)

	return err
}

type NullableSchemasGetAllUsersResponse struct {
	value *SchemasGetAllUsersResponse
	isSet bool
}

func (v NullableSchemasGetAllUsersResponse) Get() *SchemasGetAllUsersResponse {
	return v.value
}

func (v *NullableSchemasGetAllUsersResponse) Set(val *SchemasGetAllUsersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasGetAllUsersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasGetAllUsersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasGetAllUsersResponse(val *SchemasGetAllUsersResponse) *NullableSchemasGetAllUsersResponse {
	return &NullableSchemasGetAllUsersResponse{value: val, isSet: true}
}

func (v NullableSchemasGetAllUsersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasGetAllUsersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


