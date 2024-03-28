// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package gitprovider

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server/api/controllers/gitprovider/dto"
	"github.com/daytonaio/daytona/pkg/server/config"
	"github.com/gin-gonic/gin"
)

// GetRepoPRs 			godoc
//
//	@Tags			gitProvider
//	@Summary		Get Git repository PRs
//	@Description	Get Git repository PRs
//	@Param			artifacts	body	GetRepoArtifactsRequest	true	"Repository artifacts request"
//	@Produce		json
//	@Success		200	{array}	GitPullRequest
//	@Router			/gitprovider/repositories/pull-requests [get]
//
//	@id				GetRepoPRs
func GetRepoPRs(ctx *gin.Context) {
	var req dto.GetRepoArtifactsRequest
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	c, err := config.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get credentials: %s", err.Error()))
		return
	}

	gitProvider := gitprovider.GetGitProvider(req.GitProviderId, c.GitProviders)
	if gitProvider == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	response, err := gitProvider.GetRepoPRs(req.Repository, req.NamespaceId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get pull requests: %s", err.Error()))
		return
	}

	ctx.JSON(200, response)
}
