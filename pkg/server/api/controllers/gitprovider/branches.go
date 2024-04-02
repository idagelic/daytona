// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package gitprovider

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server/config"
	"github.com/gin-gonic/gin"
)

// GetRepoBranches 			godoc
//
//	@Tags			gitProvider
//	@Summary		Get Git repository branches
//	@Description	Get Git repository branches
//	@Param			gitProviderId	path	string	true	"Git provider"
//	@Param			namespaceId		path	string	true	"Namespace"
//	@Param			repositoryId	path	string	true	"Repository"
//	@Produce		json
//	@Success		200	{array}	GitBranch
//	@Router			/gitprovider/repositories/branches [get]
//
//	@id				GetRepoBranches
func GetRepoBranches(ctx *gin.Context) {
	gitProviderId := ctx.Param("gitProviderId")
	namespaceId := ctx.Param("namespaceId")
	repositoryId := ctx.Param("repositoryId")

	c, err := config.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get credentials: %s", err.Error()))
		return
	}

	gitProvider := gitprovider.GetGitProvider(gitProviderId, c.GitProviders)
	if gitProvider == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	response, err := gitProvider.GetRepoBranches(repositoryId, namespaceId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get branches: %s", err.Error()))
		return
	}

	ctx.JSON(200, response)
}
