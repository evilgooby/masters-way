/*


No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SchemasUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUserResponse{}

// SchemasUserResponse struct for SchemasUserResponse
type SchemasUserResponse struct {
	Role string `json:"role"`
	UserId string `json:"userId"`
}

type _SchemasUserResponse SchemasUserResponse

// NewSchemasUserResponse instantiates a new SchemasUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUserResponse(role string, userId string) *SchemasUserResponse {
	this := SchemasUserResponse{}
	this.Role = role
	this.UserId = userId
	return &this
}

// NewSchemasUserResponseWithDefaults instantiates a new SchemasUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUserResponseWithDefaults() *SchemasUserResponse {
	this := SchemasUserResponse{}
	return &this
}

// GetRole returns the Role field value
func (o *SchemasUserResponse) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SchemasUserResponse) SetRole(v string) {
	o.Role = v
}

// GetUserId returns the UserId field value
func (o *SchemasUserResponse) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *SchemasUserResponse) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *SchemasUserResponse) SetUserId(v string) {
	o.UserId = v
}

func (o SchemasUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["userId"] = o.UserId
	return toSerialize, nil
}

func (o *SchemasUserResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"role",
		"userId",
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

	varSchemasUserResponse := _SchemasUserResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasUserResponse)

	if err != nil {
		return err
	}

	*o = SchemasUserResponse(varSchemasUserResponse)

	return err
}

type NullableSchemasUserResponse struct {
	value *SchemasUserResponse
	isSet bool
}

func (v NullableSchemasUserResponse) Get() *SchemasUserResponse {
	return v.value
}

func (v *NullableSchemasUserResponse) Set(val *SchemasUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUserResponse(val *SchemasUserResponse) *NullableSchemasUserResponse {
	return &NullableSchemasUserResponse{value: val, isSet: true}
}

func (v NullableSchemasUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


