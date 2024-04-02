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

// checks if the GitRepository type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GitRepository{}

// GitRepository struct for GitRepository
type GitRepository struct {
	Branch *string `json:"branch,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Path *string `json:"path,omitempty"`
	PrNumber *int32 `json:"prNumber,omitempty"`
	Sha *string `json:"sha,omitempty"`
	Source *string `json:"source,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewGitRepository instantiates a new GitRepository object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGitRepository() *GitRepository {
	this := GitRepository{}
	return &this
}

// NewGitRepositoryWithDefaults instantiates a new GitRepository object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitRepositoryWithDefaults() *GitRepository {
	this := GitRepository{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *GitRepository) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *GitRepository) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *GitRepository) SetBranch(v string) {
	o.Branch = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GitRepository) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GitRepository) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GitRepository) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GitRepository) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GitRepository) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GitRepository) SetName(v string) {
	o.Name = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GitRepository) GetOwner() string {
	if o == nil || IsNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetOwnerOk() (*string, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *GitRepository) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *GitRepository) SetOwner(v string) {
	o.Owner = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GitRepository) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GitRepository) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GitRepository) SetPath(v string) {
	o.Path = &v
}

// GetPrNumber returns the PrNumber field value if set, zero value otherwise.
func (o *GitRepository) GetPrNumber() int32 {
	if o == nil || IsNil(o.PrNumber) {
		var ret int32
		return ret
	}
	return *o.PrNumber
}

// GetPrNumberOk returns a tuple with the PrNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetPrNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PrNumber) {
		return nil, false
	}
	return o.PrNumber, true
}

// HasPrNumber returns a boolean if a field has been set.
func (o *GitRepository) HasPrNumber() bool {
	if o != nil && !IsNil(o.PrNumber) {
		return true
	}

	return false
}

// SetPrNumber gets a reference to the given int32 and assigns it to the PrNumber field.
func (o *GitRepository) SetPrNumber(v int32) {
	o.PrNumber = &v
}

// GetSha returns the Sha field value if set, zero value otherwise.
func (o *GitRepository) GetSha() string {
	if o == nil || IsNil(o.Sha) {
		var ret string
		return ret
	}
	return *o.Sha
}

// GetShaOk returns a tuple with the Sha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetShaOk() (*string, bool) {
	if o == nil || IsNil(o.Sha) {
		return nil, false
	}
	return o.Sha, true
}

// HasSha returns a boolean if a field has been set.
func (o *GitRepository) HasSha() bool {
	if o != nil && !IsNil(o.Sha) {
		return true
	}

	return false
}

// SetSha gets a reference to the given string and assigns it to the Sha field.
func (o *GitRepository) SetSha(v string) {
	o.Sha = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GitRepository) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GitRepository) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *GitRepository) SetSource(v string) {
	o.Source = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *GitRepository) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GitRepository) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *GitRepository) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *GitRepository) SetUrl(v string) {
	o.Url = &v
}

func (o GitRepository) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GitRepository) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.PrNumber) {
		toSerialize["prNumber"] = o.PrNumber
	}
	if !IsNil(o.Sha) {
		toSerialize["sha"] = o.Sha
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableGitRepository struct {
	value *GitRepository
	isSet bool
}

func (v NullableGitRepository) Get() *GitRepository {
	return v.value
}

func (v *NullableGitRepository) Set(val *GitRepository) {
	v.value = val
	v.isSet = true
}

func (v NullableGitRepository) IsSet() bool {
	return v.isSet
}

func (v *NullableGitRepository) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGitRepository(val *GitRepository) *NullableGitRepository {
	return &NullableGitRepository{value: val, isSet: true}
}

func (v NullableGitRepository) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGitRepository) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


