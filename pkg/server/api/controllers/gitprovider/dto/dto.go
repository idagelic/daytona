// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package dto

import (
	"github.com/daytonaio/daytona/pkg/types"
)

type GetRepoArtifactsRequest struct {
	Repository    types.GitRepository
	GitProviderId string
	NamespaceId   string
} //	@name	GetRepoArtifactsRequest
