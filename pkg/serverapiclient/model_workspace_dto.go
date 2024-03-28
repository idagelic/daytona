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

// checks if the WorkspaceDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkspaceDTO{}

// WorkspaceDTO struct for WorkspaceDTO
type WorkspaceDTO struct {
	Id *string `json:"id,omitempty"`
	Info *WorkspaceInfo `json:"info,omitempty"`
	Name *string `json:"name,omitempty"`
	Projects []Project `json:"projects,omitempty"`
	Target *string `json:"target,omitempty"`
}

// NewWorkspaceDTO instantiates a new WorkspaceDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkspaceDTO() *WorkspaceDTO {
	this := WorkspaceDTO{}
	return &this
}

// NewWorkspaceDTOWithDefaults instantiates a new WorkspaceDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkspaceDTOWithDefaults() *WorkspaceDTO {
	this := WorkspaceDTO{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *WorkspaceDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *WorkspaceDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *WorkspaceDTO) SetId(v string) {
	o.Id = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *WorkspaceDTO) GetInfo() WorkspaceInfo {
	if o == nil || IsNil(o.Info) {
		var ret WorkspaceInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceDTO) GetInfoOk() (*WorkspaceInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *WorkspaceDTO) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given WorkspaceInfo and assigns it to the Info field.
func (o *WorkspaceDTO) SetInfo(v WorkspaceInfo) {
	o.Info = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkspaceDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkspaceDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkspaceDTO) SetName(v string) {
	o.Name = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *WorkspaceDTO) GetProjects() []Project {
	if o == nil || IsNil(o.Projects) {
		var ret []Project
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceDTO) GetProjectsOk() ([]Project, bool) {
	if o == nil || IsNil(o.Projects) {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *WorkspaceDTO) HasProjects() bool {
	if o != nil && !IsNil(o.Projects) {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []Project and assigns it to the Projects field.
func (o *WorkspaceDTO) SetProjects(v []Project) {
	o.Projects = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *WorkspaceDTO) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkspaceDTO) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *WorkspaceDTO) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *WorkspaceDTO) SetTarget(v string) {
	o.Target = &v
}

func (o WorkspaceDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkspaceDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Projects) {
		toSerialize["projects"] = o.Projects
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

type NullableWorkspaceDTO struct {
	value *WorkspaceDTO
	isSet bool
}

func (v NullableWorkspaceDTO) Get() *WorkspaceDTO {
	return v.value
}

func (v *NullableWorkspaceDTO) Set(val *WorkspaceDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkspaceDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkspaceDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkspaceDTO(val *WorkspaceDTO) *NullableWorkspaceDTO {
	return &NullableWorkspaceDTO{value: val, isSet: true}
}

func (v NullableWorkspaceDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkspaceDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


