package dto

import (
	"github.com/daytonaio/daytona/pkg/types"
)

type GetRepoArtifactsRequest struct {
	Repository    types.GitRepository
	GitProviderId string
	NamespaceId   string
} //	@name	GetRepoArtifactsRequest
