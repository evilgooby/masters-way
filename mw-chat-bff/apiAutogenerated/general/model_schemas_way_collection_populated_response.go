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

// checks if the SchemasWayCollectionPopulatedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasWayCollectionPopulatedResponse{}

// SchemasWayCollectionPopulatedResponse struct for SchemasWayCollectionPopulatedResponse
type SchemasWayCollectionPopulatedResponse struct {
	CreatedAt string `json:"createdAt"`
	Name string `json:"name"`
	OwnerUuid string `json:"ownerUuid"`
	// should be removed after separation custom collections and default pseudocollections
	Type string `json:"type"`
	UpdatedAt string `json:"updatedAt"`
	Uuid string `json:"uuid"`
	Ways []SchemasWayPlainResponse `json:"ways"`
}

type _SchemasWayCollectionPopulatedResponse SchemasWayCollectionPopulatedResponse

// NewSchemasWayCollectionPopulatedResponse instantiates a new SchemasWayCollectionPopulatedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasWayCollectionPopulatedResponse(createdAt string, name string, ownerUuid string, type_ string, updatedAt string, uuid string, ways []SchemasWayPlainResponse) *SchemasWayCollectionPopulatedResponse {
	this := SchemasWayCollectionPopulatedResponse{}
	this.CreatedAt = createdAt
	this.Name = name
	this.OwnerUuid = ownerUuid
	this.Type = type_
	this.UpdatedAt = updatedAt
	this.Uuid = uuid
	this.Ways = ways
	return &this
}

// NewSchemasWayCollectionPopulatedResponseWithDefaults instantiates a new SchemasWayCollectionPopulatedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasWayCollectionPopulatedResponseWithDefaults() *SchemasWayCollectionPopulatedResponse {
	this := SchemasWayCollectionPopulatedResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *SchemasWayCollectionPopulatedResponse) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SchemasWayCollectionPopulatedResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SchemasWayCollectionPopulatedResponse) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetName returns the Name field value
func (o *SchemasWayCollectionPopulatedResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasWayCollectionPopulatedResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasWayCollectionPopulatedResponse) SetName(v string) {
	o.Name = v
}

// GetOwnerUuid returns the OwnerUuid field value
func (o *SchemasWayCollectionPopulatedResponse) GetOwnerUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OwnerUuid
}

// GetOwnerUuidOk returns a tuple with the OwnerUuid field value
// and a boolean to check if the value has been set.
func (o *SchemasWayCollectionPopulatedResponse) GetOwnerUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnerUuid, true
}

// SetOwnerUuid sets field value
func (o *SchemasWayCollectionPopulatedResponse) SetOwnerUuid(v string) {
	o.OwnerUuid = v
}

// GetType returns the Type field value
func (o *SchemasWayCollectionPopulatedResponse) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SchemasWayCollectionPopulatedResponse) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SchemasWayCollectionPopulatedResponse) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SchemasWayCollectionPopulatedResponse) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SchemasWayCollectionPopulatedResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SchemasWayCollectionPopulatedResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetUuid returns the Uuid field value
func (o *SchemasWayCollectionPopulatedResponse) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SchemasWayCollectionPopulatedResponse) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SchemasWayCollectionPopulatedResponse) SetUuid(v string) {
	o.Uuid = v
}

// GetWays returns the Ways field value
func (o *SchemasWayCollectionPopulatedResponse) GetWays() []SchemasWayPlainResponse {
	if o == nil {
		var ret []SchemasWayPlainResponse
		return ret
	}

	return o.Ways
}

// GetWaysOk returns a tuple with the Ways field value
// and a boolean to check if the value has been set.
func (o *SchemasWayCollectionPopulatedResponse) GetWaysOk() ([]SchemasWayPlainResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ways, true
}

// SetWays sets field value
func (o *SchemasWayCollectionPopulatedResponse) SetWays(v []SchemasWayPlainResponse) {
	o.Ways = v
}

func (o SchemasWayCollectionPopulatedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasWayCollectionPopulatedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["name"] = o.Name
	toSerialize["ownerUuid"] = o.OwnerUuid
	toSerialize["type"] = o.Type
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["uuid"] = o.Uuid
	toSerialize["ways"] = o.Ways
	return toSerialize, nil
}

func (o *SchemasWayCollectionPopulatedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"name",
		"ownerUuid",
		"type",
		"updatedAt",
		"uuid",
		"ways",
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

	varSchemasWayCollectionPopulatedResponse := _SchemasWayCollectionPopulatedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasWayCollectionPopulatedResponse)

	if err != nil {
		return err
	}

	*o = SchemasWayCollectionPopulatedResponse(varSchemasWayCollectionPopulatedResponse)

	return err
}

type NullableSchemasWayCollectionPopulatedResponse struct {
	value *SchemasWayCollectionPopulatedResponse
	isSet bool
}

func (v NullableSchemasWayCollectionPopulatedResponse) Get() *SchemasWayCollectionPopulatedResponse {
	return v.value
}

func (v *NullableSchemasWayCollectionPopulatedResponse) Set(val *SchemasWayCollectionPopulatedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasWayCollectionPopulatedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasWayCollectionPopulatedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasWayCollectionPopulatedResponse(val *SchemasWayCollectionPopulatedResponse) *NullableSchemasWayCollectionPopulatedResponse {
	return &NullableSchemasWayCollectionPopulatedResponse{value: val, isSet: true}
}

func (v NullableSchemasWayCollectionPopulatedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasWayCollectionPopulatedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


