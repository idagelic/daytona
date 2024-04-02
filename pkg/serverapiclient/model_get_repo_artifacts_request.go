/*
Daytona Server API

Daytona Server API

API version: 0.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serverapiclient

import (
	"encoding/json"
)

// checks if the GetRepoArtifactsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRepoArtifactsRequest{}

// GetRepoArtifactsRequest struct for GetRepoArtifactsRequest
type GetRepoArtifactsRequest struct {
	GitProviderId *string        `json:"gitProviderId,omitempty"`
	NamespaceId   *string        `json:"namespaceId,omitempty"`
	Repository    *GitRepository `json:"repository,omitempty"`
}

// NewGetRepoArtifactsRequest instantiates a new GetRepoArtifactsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRepoArtifactsRequest() *GetRepoArtifactsRequest {
	this := GetRepoArtifactsRequest{}
	return &this
}

// NewGetRepoArtifactsRequestWithDefaults instantiates a new GetRepoArtifactsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRepoArtifactsRequestWithDefaults() *GetRepoArtifactsRequest {
	this := GetRepoArtifactsRequest{}
	return &this
}

// GetGitProviderId returns the GitProviderId field value if set, zero value otherwise.
func (o *GetRepoArtifactsRequest) GetGitProviderId() string {
	if o == nil || IsNil(o.GitProviderId) {
		var ret string
		return ret
	}
	return *o.GitProviderId
}

// GetGitProviderIdOk returns a tuple with the GitProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRepoArtifactsRequest) GetGitProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.GitProviderId) {
		return nil, false
	}
	return o.GitProviderId, true
}

// HasGitProviderId returns a boolean if a field has been set.
func (o *GetRepoArtifactsRequest) HasGitProviderId() bool {
	if o != nil && !IsNil(o.GitProviderId) {
		return true
	}

	return false
}

// SetGitProviderId gets a reference to the given string and assigns it to the GitProviderId field.
func (o *GetRepoArtifactsRequest) SetGitProviderId(v string) {
	o.GitProviderId = &v
}

// GetNamespaceId returns the NamespaceId field value if set, zero value otherwise.
func (o *GetRepoArtifactsRequest) GetNamespaceId() string {
	if o == nil || IsNil(o.NamespaceId) {
		var ret string
		return ret
	}
	return *o.NamespaceId
}

// GetNamespaceIdOk returns a tuple with the NamespaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRepoArtifactsRequest) GetNamespaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.NamespaceId) {
		return nil, false
	}
	return o.NamespaceId, true
}

// HasNamespaceId returns a boolean if a field has been set.
func (o *GetRepoArtifactsRequest) HasNamespaceId() bool {
	if o != nil && !IsNil(o.NamespaceId) {
		return true
	}

	return false
}

// SetNamespaceId gets a reference to the given string and assigns it to the NamespaceId field.
func (o *GetRepoArtifactsRequest) SetNamespaceId(v string) {
	o.NamespaceId = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GetRepoArtifactsRequest) GetRepository() GitRepository {
	if o == nil || IsNil(o.Repository) {
		var ret GitRepository
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRepoArtifactsRequest) GetRepositoryOk() (*GitRepository, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GetRepoArtifactsRequest) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given GitRepository and assigns it to the Repository field.
func (o *GetRepoArtifactsRequest) SetRepository(v GitRepository) {
	o.Repository = &v
}

func (o GetRepoArtifactsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRepoArtifactsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GitProviderId) {
		toSerialize["gitProviderId"] = o.GitProviderId
	}
	if !IsNil(o.NamespaceId) {
		toSerialize["namespaceId"] = o.NamespaceId
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	return toSerialize, nil
}

type NullableGetRepoArtifactsRequest struct {
	value *GetRepoArtifactsRequest
	isSet bool
}

func (v NullableGetRepoArtifactsRequest) Get() *GetRepoArtifactsRequest {
	return v.value
}

func (v *NullableGetRepoArtifactsRequest) Set(val *GetRepoArtifactsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRepoArtifactsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRepoArtifactsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRepoArtifactsRequest(val *GetRepoArtifactsRequest) *NullableGetRepoArtifactsRequest {
	return &NullableGetRepoArtifactsRequest{value: val, isSet: true}
}

func (v NullableGetRepoArtifactsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRepoArtifactsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
