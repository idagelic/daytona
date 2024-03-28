package dto

import (
	"github.com/daytonaio/daytona/pkg/types"
)

type GetRepoArtifactsRequest struct {
	Repository    types.Repository
	GitProviderId string
	NamespaceId   string
} //	@name	GetRepoArtifactsRequest
