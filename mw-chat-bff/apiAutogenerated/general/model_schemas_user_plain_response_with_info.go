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

// checks if the SchemasUserPlainResponseWithInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemasUserPlainResponseWithInfo{}

// SchemasUserPlainResponseWithInfo struct for SchemasUserPlainResponseWithInfo
type SchemasUserPlainResponseWithInfo struct {
	CreatedAt string
	Description string
	Email string
	FavoriteForUsers int32
	FavoriteWays int32
	ImageUrl string
	IsMentor bool
	MentoringWays int32
	Name string
	OwnWays int32
	Tags []SchemasUserTagResponse
	Uuid string
}

type _SchemasUserPlainResponseWithInfo SchemasUserPlainResponseWithInfo

// NewSchemasUserPlainResponseWithInfo instantiates a new SchemasUserPlainResponseWithInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemasUserPlainResponseWithInfo(createdAt string, description string, email string, favoriteForUsers int32, favoriteWays int32, imageUrl string, isMentor bool, mentoringWays int32, name string, ownWays int32, tags []SchemasUserTagResponse, uuid string) *SchemasUserPlainResponseWithInfo {
	this := SchemasUserPlainResponseWithInfo{}
	this.CreatedAt = createdAt
	this.Description = description
	this.Email = email
	this.FavoriteForUsers = favoriteForUsers
	this.FavoriteWays = favoriteWays
	this.ImageUrl = imageUrl
	this.IsMentor = isMentor
	this.MentoringWays = mentoringWays
	this.Name = name
	this.OwnWays = ownWays
	this.Tags = tags
	this.Uuid = uuid
	return &this
}

// NewSchemasUserPlainResponseWithInfoWithDefaults instantiates a new SchemasUserPlainResponseWithInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemasUserPlainResponseWithInfoWithDefaults() *SchemasUserPlainResponseWithInfo {
	this := SchemasUserPlainResponseWithInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *SchemasUserPlainResponseWithInfo) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SchemasUserPlainResponseWithInfo) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetDescription returns the Description field value
func (o *SchemasUserPlainResponseWithInfo) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *SchemasUserPlainResponseWithInfo) SetDescription(v string) {
	o.Description = v
}

// GetEmail returns the Email field value
func (o *SchemasUserPlainResponseWithInfo) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SchemasUserPlainResponseWithInfo) SetEmail(v string) {
	o.Email = v
}

// GetFavoriteForUsers returns the FavoriteForUsers field value
func (o *SchemasUserPlainResponseWithInfo) GetFavoriteForUsers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FavoriteForUsers
}

// GetFavoriteForUsersOk returns a tuple with the FavoriteForUsers field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetFavoriteForUsersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoriteForUsers, true
}

// SetFavoriteForUsers sets field value
func (o *SchemasUserPlainResponseWithInfo) SetFavoriteForUsers(v int32) {
	o.FavoriteForUsers = v
}

// GetFavoriteWays returns the FavoriteWays field value
func (o *SchemasUserPlainResponseWithInfo) GetFavoriteWays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FavoriteWays
}

// GetFavoriteWaysOk returns a tuple with the FavoriteWays field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetFavoriteWaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FavoriteWays, true
}

// SetFavoriteWays sets field value
func (o *SchemasUserPlainResponseWithInfo) SetFavoriteWays(v int32) {
	o.FavoriteWays = v
}

// GetImageUrl returns the ImageUrl field value
func (o *SchemasUserPlainResponseWithInfo) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *SchemasUserPlainResponseWithInfo) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetIsMentor returns the IsMentor field value
func (o *SchemasUserPlainResponseWithInfo) GetIsMentor() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMentor
}

// GetIsMentorOk returns a tuple with the IsMentor field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetIsMentorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMentor, true
}

// SetIsMentor sets field value
func (o *SchemasUserPlainResponseWithInfo) SetIsMentor(v bool) {
	o.IsMentor = v
}

// GetMentoringWays returns the MentoringWays field value
func (o *SchemasUserPlainResponseWithInfo) GetMentoringWays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MentoringWays
}

// GetMentoringWaysOk returns a tuple with the MentoringWays field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetMentoringWaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MentoringWays, true
}

// SetMentoringWays sets field value
func (o *SchemasUserPlainResponseWithInfo) SetMentoringWays(v int32) {
	o.MentoringWays = v
}

// GetName returns the Name field value
func (o *SchemasUserPlainResponseWithInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemasUserPlainResponseWithInfo) SetName(v string) {
	o.Name = v
}

// GetOwnWays returns the OwnWays field value
func (o *SchemasUserPlainResponseWithInfo) GetOwnWays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OwnWays
}

// GetOwnWaysOk returns a tuple with the OwnWays field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetOwnWaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OwnWays, true
}

// SetOwnWays sets field value
func (o *SchemasUserPlainResponseWithInfo) SetOwnWays(v int32) {
	o.OwnWays = v
}

// GetTags returns the Tags field value
func (o *SchemasUserPlainResponseWithInfo) GetTags() []SchemasUserTagResponse {
	if o == nil {
		var ret []SchemasUserTagResponse
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetTagsOk() ([]SchemasUserTagResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *SchemasUserPlainResponseWithInfo) SetTags(v []SchemasUserTagResponse) {
	o.Tags = v
}

// GetUuid returns the Uuid field value
func (o *SchemasUserPlainResponseWithInfo) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SchemasUserPlainResponseWithInfo) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SchemasUserPlainResponseWithInfo) SetUuid(v string) {
	o.Uuid = v
}

func (o SchemasUserPlainResponseWithInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemasUserPlainResponseWithInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["description"] = o.Description
	toSerialize["email"] = o.Email
	toSerialize["favoriteForUsers"] = o.FavoriteForUsers
	toSerialize["favoriteWays"] = o.FavoriteWays
	toSerialize["imageUrl"] = o.ImageUrl
	toSerialize["isMentor"] = o.IsMentor
	toSerialize["mentoringWays"] = o.MentoringWays
	toSerialize["name"] = o.Name
	toSerialize["ownWays"] = o.OwnWays
	toSerialize["tags"] = o.Tags
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *SchemasUserPlainResponseWithInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"description",
		"email",
		"favoriteForUsers",
		"favoriteWays",
		"imageUrl",
		"isMentor",
		"mentoringWays",
		"name",
		"ownWays",
		"tags",
		"uuid",
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

	varSchemasUserPlainResponseWithInfo := _SchemasUserPlainResponseWithInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSchemasUserPlainResponseWithInfo)

	if err != nil {
		return err
	}

	*o = SchemasUserPlainResponseWithInfo(varSchemasUserPlainResponseWithInfo)

	return err
}

type NullableSchemasUserPlainResponseWithInfo struct {
	value *SchemasUserPlainResponseWithInfo
	isSet bool
}

func (v NullableSchemasUserPlainResponseWithInfo) Get() *SchemasUserPlainResponseWithInfo {
	return v.value
}

func (v *NullableSchemasUserPlainResponseWithInfo) Set(val *SchemasUserPlainResponseWithInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemasUserPlainResponseWithInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemasUserPlainResponseWithInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemasUserPlainResponseWithInfo(val *SchemasUserPlainResponseWithInfo) *NullableSchemasUserPlainResponseWithInfo {
	return &NullableSchemasUserPlainResponseWithInfo{value: val, isSet: true}
}

func (v NullableSchemasUserPlainResponseWithInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemasUserPlainResponseWithInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


