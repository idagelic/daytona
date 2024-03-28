# GetRepoArtifactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitProviderId** | Pointer to **string** |  | [optional] 
**NamespaceId** | Pointer to **string** |  | [optional] 
**Repository** | Pointer to [**Repository**](Repository.md) |  | [optional] 

## Methods

### NewGetRepoArtifactsRequest

`func NewGetRepoArtifactsRequest() *GetRepoArtifactsRequest`

NewGetRepoArtifactsRequest instantiates a new GetRepoArtifactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRepoArtifactsRequestWithDefaults

`func NewGetRepoArtifactsRequestWithDefaults() *GetRepoArtifactsRequest`

NewGetRepoArtifactsRequestWithDefaults instantiates a new GetRepoArtifactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitProviderId

`func (o *GetRepoArtifactsRequest) GetGitProviderId() string`

GetGitProviderId returns the GitProviderId field if non-nil, zero value otherwise.

### GetGitProviderIdOk

`func (o *GetRepoArtifactsRequest) GetGitProviderIdOk() (*string, bool)`

GetGitProviderIdOk returns a tuple with the GitProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitProviderId

`func (o *GetRepoArtifactsRequest) SetGitProviderId(v string)`

SetGitProviderId sets GitProviderId field to given value.

### HasGitProviderId

`func (o *GetRepoArtifactsRequest) HasGitProviderId() bool`

HasGitProviderId returns a boolean if a field has been set.

### GetNamespaceId

`func (o *GetRepoArtifactsRequest) GetNamespaceId() string`

GetNamespaceId returns the NamespaceId field if non-nil, zero value otherwise.

### GetNamespaceIdOk

`func (o *GetRepoArtifactsRequest) GetNamespaceIdOk() (*string, bool)`

GetNamespaceIdOk returns a tuple with the NamespaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaceId

`func (o *GetRepoArtifactsRequest) SetNamespaceId(v string)`

SetNamespaceId sets NamespaceId field to given value.

### HasNamespaceId

`func (o *GetRepoArtifactsRequest) HasNamespaceId() bool`

HasNamespaceId returns a boolean if a field has been set.

### GetRepository

`func (o *GetRepoArtifactsRequest) GetRepository() Repository`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *GetRepoArtifactsRequest) GetRepositoryOk() (*Repository, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *GetRepoArtifactsRequest) SetRepository(v Repository)`

SetRepository sets Repository field to given value.

### HasRepository

`func (o *GetRepoArtifactsRequest) HasRepository() bool`

HasRepository returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


