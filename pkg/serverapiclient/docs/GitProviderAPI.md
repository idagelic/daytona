# \GitProviderAPI

All URIs are relative to *http://localhost:3000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGitContext**](GitProviderAPI.md#GetGitContext) | **Get** /gitprovider/context/{gitUrl} | Get Git context
[**GetGitProvider**](GitProviderAPI.md#GetGitProvider) | **Get** /gitprovider/{gitProviderId} | Get Git provider
[**GetGitUser**](GitProviderAPI.md#GetGitUser) | **Get** /gitprovider/{gitProviderId}/user | Get Git context
[**GetGitUsernameFromToken**](GitProviderAPI.md#GetGitUsernameFromToken) | **Get** /gitprovider/username-from-token | Get username from token
[**GetNamespaces**](GitProviderAPI.md#GetNamespaces) | **Get** /gitprovider/{gitProviderId}/namespaces | Get Git namespaces
[**GetRepoBranches**](GitProviderAPI.md#GetRepoBranches) | **Get** /gitprovider/repositories/branches | Get Git repository branches
[**GetRepoPRs**](GitProviderAPI.md#GetRepoPRs) | **Get** /gitprovider/repositories/pull-requests | Get Git repository PRs
[**GetRepositories**](GitProviderAPI.md#GetRepositories) | **Get** /gitprovider/{gitProviderId}/{namespaceId}/repositories | Get Git repositories
[**ListGitProviders**](GitProviderAPI.md#ListGitProviders) | **Get** /gitprovider | List git providers
[**RemoveGitProvider**](GitProviderAPI.md#RemoveGitProvider) | **Delete** /gitprovider/{gitProviderId} | Remove Git provider
[**SetGitProvider**](GitProviderAPI.md#SetGitProvider) | **Post** /gitprovider | Set Git provider



## GetGitContext

> Repository GetGitContext(ctx, gitUrl).Execute()

Get Git context



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitUrl := "gitUrl_example" // string | Git URL

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetGitContext(context.Background(), gitUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetGitContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitContext`: Repository
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetGitContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitUrl** | **string** | Git URL | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitProvider

> GitProvider GetGitProvider(ctx, gitProviderId).Execute()

Get Git provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitProviderId := "gitProviderId_example" // string | Git provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetGitProvider(context.Background(), gitProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetGitProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitProvider`: GitProvider
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetGitProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitProviderId** | **string** | Git provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitProvider**](GitProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitUser

> GitUser GetGitUser(ctx, gitProviderId).Execute()

Get Git context



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitProviderId := "gitProviderId_example" // string | Git Provider Id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetGitUser(context.Background(), gitProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetGitUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitUser`: GitUser
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetGitUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitProviderId** | **string** | Git Provider Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitUser**](GitUser.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGitUsernameFromToken

> string GetGitUsernameFromToken(ctx).GitProviderData(gitProviderData).Execute()

Get username from token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitProviderData := *openapiclient.NewGitProvider() // GitProvider | Git provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetGitUsernameFromToken(context.Background()).GitProviderData(gitProviderData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetGitUsernameFromToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGitUsernameFromToken`: string
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetGitUsernameFromToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGitUsernameFromTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitProviderData** | [**GitProvider**](GitProvider.md) | Git provider | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNamespaces

> []GitNamespace GetNamespaces(ctx, gitProviderId).Execute()

Get Git namespaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitProviderId := "gitProviderId_example" // string | Git provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetNamespaces(context.Background(), gitProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetNamespaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNamespaces`: []GitNamespace
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetNamespaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitProviderId** | **string** | Git provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNamespacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GitNamespace**](GitNamespace.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepoBranches

> []GitBranch GetRepoBranches(ctx).Artifacts(artifacts).Execute()

Get Git repository branches



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	artifacts := *openapiclient.NewGetRepoArtifactsRequest() // GetRepoArtifactsRequest | Repository artifacts request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetRepoBranches(context.Background()).Artifacts(artifacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetRepoBranches``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepoBranches`: []GitBranch
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetRepoBranches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepoBranchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifacts** | [**GetRepoArtifactsRequest**](GetRepoArtifactsRequest.md) | Repository artifacts request | 

### Return type

[**[]GitBranch**](GitBranch.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepoPRs

> []GitPullRequest GetRepoPRs(ctx).Artifacts(artifacts).Execute()

Get Git repository PRs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	artifacts := *openapiclient.NewGetRepoArtifactsRequest() // GetRepoArtifactsRequest | Repository artifacts request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetRepoPRs(context.Background()).Artifacts(artifacts).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetRepoPRs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepoPRs`: []GitPullRequest
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetRepoPRs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRepoPRsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **artifacts** | [**GetRepoArtifactsRequest**](GetRepoArtifactsRequest.md) | Repository artifacts request | 

### Return type

[**[]GitPullRequest**](GitPullRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepositories

> []Repository GetRepositories(ctx, gitProviderId, namespaceId).Execute()

Get Git repositories



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitProviderId := "gitProviderId_example" // string | Git provider
	namespaceId := "namespaceId_example" // string | Namespace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.GetRepositories(context.Background(), gitProviderId, namespaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.GetRepositories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepositories`: []Repository
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.GetRepositories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitProviderId** | **string** | Git provider | 
**namespaceId** | **string** | Namespace | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]Repository**](Repository.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGitProviders

> []GitProvider ListGitProviders(ctx).Execute()

List git providers



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitProviderAPI.ListGitProviders(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.ListGitProviders``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGitProviders`: []GitProvider
	fmt.Fprintf(os.Stdout, "Response from `GitProviderAPI.ListGitProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGitProvidersRequest struct via the builder pattern


### Return type

[**[]GitProvider**](GitProvider.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGitProvider

> RemoveGitProvider(ctx, gitProviderId).Execute()

Remove Git provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitProviderId := "gitProviderId_example" // string | Git provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitProviderAPI.RemoveGitProvider(context.Background(), gitProviderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.RemoveGitProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**gitProviderId** | **string** | Git provider | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGitProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetGitProvider

> SetGitProvider(ctx).GitProviderData(gitProviderData).Execute()

Set Git provider



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/serverapiclient"
)

func main() {
	gitProviderData := *openapiclient.NewGitProvider() // GitProvider | Git provider

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GitProviderAPI.SetGitProvider(context.Background()).GitProviderData(gitProviderData).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitProviderAPI.SetGitProvider``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetGitProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gitProviderData** | [**GitProvider**](GitProvider.md) | Git provider | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

