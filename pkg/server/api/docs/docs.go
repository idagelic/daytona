// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/gitprovider": {
            "get": {
                "description": "List git providers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "List git providers",
                "operationId": "ListGitProviders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GitProvider"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Set Git provider",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Set Git provider",
                "operationId": "SetGitProvider",
                "parameters": [
                    {
                        "description": "Git provider",
                        "name": "gitProviderData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GitProvider"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/gitprovider/context/{gitUrl}": {
            "get": {
                "description": "Get Git context",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get Git context",
                "operationId": "GetGitContext",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Git URL",
                        "name": "gitUrl",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GitRepository"
                        }
                    }
                }
            }
        },
        "/gitprovider/repositories/branches": {
            "get": {
                "description": "Get Git repository branches",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get Git repository branches",
                "operationId": "GetRepoBranches",
                "parameters": [
                    {
                        "description": "Repository artifacts request",
                        "name": "artifacts",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GetRepoArtifactsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GitBranch"
                            }
                        }
                    }
                }
            }
        },
        "/gitprovider/repositories/pull-requests": {
            "get": {
                "description": "Get Git repository PRs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get Git repository PRs",
                "operationId": "GetRepoPRs",
                "parameters": [
                    {
                        "description": "Repository artifacts request",
                        "name": "artifacts",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GetRepoArtifactsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GitPullRequest"
                            }
                        }
                    }
                }
            }
        },
        "/gitprovider/username-from-token": {
            "get": {
                "description": "Get username from token",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get username from token",
                "operationId": "GetGitUsernameFromToken",
                "parameters": [
                    {
                        "description": "Git provider",
                        "name": "gitProviderData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/GitProvider"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/gitprovider/{gitProviderId}": {
            "get": {
                "description": "Get Git provider",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get Git provider",
                "operationId": "GetGitProvider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Git provider",
                        "name": "gitProviderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GitProvider"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove Git provider",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Remove Git provider",
                "operationId": "RemoveGitProvider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Git provider",
                        "name": "gitProviderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/gitprovider/{gitProviderId}/namespaces": {
            "get": {
                "description": "Get Git namespaces",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get Git namespaces",
                "operationId": "GetNamespaces",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Git provider",
                        "name": "gitProviderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GitNamespace"
                            }
                        }
                    }
                }
            }
        },
        "/gitprovider/{gitProviderId}/user": {
            "get": {
                "description": "Get Git context",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get Git context",
                "operationId": "GetGitUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Git Provider Id",
                        "name": "gitProviderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GitUser"
                        }
                    }
                }
            }
        },
        "/gitprovider/{gitProviderId}/{namespaceId}/repositories": {
            "get": {
                "description": "Get Git repositories",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "gitProvider"
                ],
                "summary": "Get Git repositories",
                "operationId": "GetRepositories",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Git provider",
                        "name": "gitProviderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Namespace",
                        "name": "namespaceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/GitRepository"
                            }
                        }
                    }
                }
            }
        },
        "/provider": {
            "get": {
                "description": "List providers",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "List providers",
                "operationId": "ListProviders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Provider"
                            }
                        }
                    }
                }
            }
        },
        "/provider/install": {
            "post": {
                "description": "Install a provider",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "Install a provider",
                "operationId": "InstallProvider",
                "parameters": [
                    {
                        "description": "Provider to install",
                        "name": "provider",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/InstallProviderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/provider/{provider}/target-manifest": {
            "get": {
                "description": "Get provider target manifest",
                "tags": [
                    "provider"
                ],
                "summary": "Get provider target manifest",
                "operationId": "GetTargetManifest",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider name",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ProviderTargetManifest"
                        }
                    }
                }
            }
        },
        "/provider/{provider}/uninstall": {
            "post": {
                "description": "Uninstall a provider",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "provider"
                ],
                "summary": "Uninstall a provider",
                "operationId": "UninstallProvider",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider to uninstall",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/server/config": {
            "get": {
                "description": "Get the server configuration",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Get the server configuration",
                "operationId": "GetConfig",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ServerConfig"
                        }
                    }
                }
            },
            "post": {
                "description": "Set the server configuration",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Set the server configuration",
                "operationId": "SetConfig",
                "parameters": [
                    {
                        "description": "Server configuration",
                        "name": "config",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ServerConfig"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/ServerConfig"
                        }
                    }
                }
            }
        },
        "/server/network-key": {
            "post": {
                "description": "Generate a new authentication key",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "server"
                ],
                "summary": "Generate a new authentication key",
                "operationId": "GenerateNetworkKey",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/NetworkKey"
                        }
                    }
                }
            }
        },
        "/target": {
            "get": {
                "description": "List targets",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "target"
                ],
                "summary": "List targets",
                "operationId": "ListTargets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/ProviderTarget"
                            }
                        }
                    }
                }
            },
            "put": {
                "description": "Set a target",
                "tags": [
                    "target"
                ],
                "summary": "Set a target",
                "operationId": "SetTarget",
                "parameters": [
                    {
                        "description": "Target to set",
                        "name": "target",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ProviderTarget"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    }
                }
            }
        },
        "/target/{target}": {
            "delete": {
                "description": "Remove a target",
                "tags": [
                    "target"
                ],
                "summary": "Remove a target",
                "operationId": "RemoveTarget",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Target name",
                        "name": "target",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    }
                }
            }
        },
        "/workspace": {
            "get": {
                "description": "List workspaces",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "List workspaces",
                "operationId": "ListWorkspaces",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "Verbose",
                        "name": "verbose",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/WorkspaceDTO"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Create a workspace",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "Create a workspace",
                "operationId": "CreateWorkspace",
                "parameters": [
                    {
                        "description": "Create workspace",
                        "name": "workspace",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateWorkspace"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Workspace"
                        }
                    }
                }
            }
        },
        "/workspace/{workspaceId}": {
            "get": {
                "description": "Get workspace info",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "workspace"
                ],
                "summary": "Get workspace info",
                "operationId": "GetWorkspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID or Name",
                        "name": "workspaceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/WorkspaceDTO"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove workspace",
                "tags": [
                    "workspace"
                ],
                "summary": "Remove workspace",
                "operationId": "RemoveWorkspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID",
                        "name": "workspaceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/workspace/{workspaceId}/start": {
            "post": {
                "description": "Start workspace",
                "tags": [
                    "workspace"
                ],
                "summary": "Start workspace",
                "operationId": "StartWorkspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID or Name",
                        "name": "workspaceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/workspace/{workspaceId}/stop": {
            "post": {
                "description": "Stop workspace",
                "tags": [
                    "workspace"
                ],
                "summary": "Stop workspace",
                "operationId": "StopWorkspace",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID or Name",
                        "name": "workspaceId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/workspace/{workspaceId}/{projectId}/start": {
            "post": {
                "description": "Start project",
                "tags": [
                    "workspace"
                ],
                "summary": "Start project",
                "operationId": "StartProject",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID or Name",
                        "name": "workspaceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/workspace/{workspaceId}/{projectId}/stop": {
            "post": {
                "description": "Stop project",
                "tags": [
                    "workspace"
                ],
                "summary": "Stop project",
                "operationId": "StopProject",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Workspace ID or Name",
                        "name": "workspaceId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Project ID",
                        "name": "projectId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateWorkspace": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "repositories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GitRepository"
                    }
                },
                "target": {
                    "type": "string"
                }
            }
        },
        "FRPSConfig": {
            "type": "object",
            "properties": {
                "domain": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "protocol": {
                    "type": "string"
                }
            }
        },
        "GetRepoArtifactsRequest": {
            "type": "object",
            "properties": {
                "gitProviderId": {
                    "type": "string"
                },
                "namespaceId": {
                    "type": "string"
                },
                "repository": {
                    "$ref": "#/definitions/GitRepository"
                }
            }
        },
        "GitBranch": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "sha": {
                    "type": "string"
                }
            }
        },
        "GitNamespace": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "GitProvider": {
            "type": "object",
            "properties": {
                "baseApiUrl": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "GitPullRequest": {
            "type": "object",
            "properties": {
                "branch": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "GitRepository": {
            "type": "object",
            "properties": {
                "branch": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "prNumber": {
                    "type": "integer"
                },
                "sha": {
                    "type": "string"
                },
                "source": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "GitUser": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "InstallProviderRequest": {
            "type": "object",
            "properties": {
                "downloadUrls": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "NetworkKey": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                }
            }
        },
        "Project": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "repository": {
                    "$ref": "#/definitions/GitRepository"
                },
                "target": {
                    "type": "string"
                },
                "workspaceId": {
                    "type": "string"
                }
            }
        },
        "ProjectInfo": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "finished": {
                    "type": "string"
                },
                "isRunning": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "providerMetadata": {
                    "type": "string"
                },
                "started": {
                    "type": "string"
                },
                "workspaceId": {
                    "type": "string"
                }
            }
        },
        "Provider": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "ProviderTarget": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "options": {
                    "description": "JSON encoded map of options",
                    "type": "string"
                },
                "providerInfo": {
                    "$ref": "#/definitions/provider.ProviderInfo"
                }
            }
        },
        "ProviderTargetManifest": {
            "type": "object",
            "additionalProperties": {
                "$ref": "#/definitions/provider.ProviderTargetProperty"
            }
        },
        "ServerConfig": {
            "type": "object",
            "properties": {
                "apiPort": {
                    "type": "integer"
                },
                "binariesPath": {
                    "type": "string"
                },
                "frps": {
                    "$ref": "#/definitions/FRPSConfig"
                },
                "gitProviders": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/GitProvider"
                    }
                },
                "headscalePort": {
                    "type": "integer"
                },
                "id": {
                    "type": "string"
                },
                "providersDir": {
                    "type": "string"
                },
                "registryUrl": {
                    "type": "string"
                },
                "serverDownloadUrl": {
                    "type": "string"
                },
                "targetsFilePath": {
                    "type": "string"
                }
            }
        },
        "Workspace": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Project"
                    }
                },
                "target": {
                    "type": "string"
                }
            }
        },
        "WorkspaceDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "info": {
                    "$ref": "#/definitions/WorkspaceInfo"
                },
                "name": {
                    "type": "string"
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Project"
                    }
                },
                "target": {
                    "type": "string"
                }
            }
        },
        "WorkspaceInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/ProjectInfo"
                    }
                },
                "providerMetadata": {
                    "type": "string"
                }
            }
        },
        "provider.ProviderInfo": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "provider.ProviderTargetProperty": {
            "type": "object",
            "properties": {
                "defaultValue": {
                    "description": "DefaultValue is converted into the appropriate type based on the Type\nIf the property is a FilePath, the DefaultValue is a path to a directory",
                    "type": "string"
                },
                "disabledPredicate": {
                    "description": "A regex string matched with the name of the target to determine if the property should be disabled\nIf the regex matches the target name, the property will be disabled\nE.g. \"^local$\" will disable the property for the local target",
                    "type": "string"
                },
                "inputMasked": {
                    "type": "boolean"
                },
                "options": {
                    "description": "Options is only used if the Type is ProviderTargetPropertyTypeOption",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "type": {
                    "$ref": "#/definitions/provider.ProviderTargetPropertyType"
                }
            }
        },
        "provider.ProviderTargetPropertyType": {
            "type": "string",
            "enum": [
                "string",
                "option",
                "boolean",
                "int",
                "float",
                "file-path"
            ],
            "x-enum-varnames": [
                "ProviderTargetPropertyTypeString",
                "ProviderTargetPropertyTypeOption",
                "ProviderTargetPropertyTypeBoolean",
                "ProviderTargetPropertyTypeInt",
                "ProviderTargetPropertyTypeFloat",
                "ProviderTargetPropertyTypeFilePath"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Daytona Server API",
	Description:      "Daytona Server API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
