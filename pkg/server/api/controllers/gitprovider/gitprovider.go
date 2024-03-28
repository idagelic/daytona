// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package gitprovider

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server/config"
	"github.com/daytonaio/daytona/pkg/types"
	"github.com/gin-gonic/gin"
)

// ListGitProviders		godoc
//
//	@Tags			gitProvider
//	@Summary		List git providers
//	@Description	List git providers
//	@Produce		json
//	@Success		200	{array}	types.GitProvider
//	@Router			/gitprovider [get]
//
//	@id				ListGitProviders
func ListGitProviders(ctx *gin.Context) {
	c, err := config.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get config: %s", err.Error()))
		return
	}

	ctx.JSON(200, c.GitProviders)
}

// GetGitProvider 			godoc
//
//	@Tags			gitProvider
//	@Summary		Get Git provider
//	@Description	Get Git provider
//	@Produce		json
//	@Param			gitProviderId	path		string	true	"Git provider"
//	@Success		200				{object}	types.GitProvider
//	@Router			/gitprovider/{gitProviderId} [get]
//
//	@id				GetGitProvider
func GetGitProvider(ctx *gin.Context) {
	var gitProvider types.GitProvider

	gitProviderId := ctx.Param("gitProviderId")

	c, err := config.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get config: %s", err.Error()))
		return
	}

	for _, provider := range c.GitProviders {
		if provider.Id == gitProviderId {
			gitProvider = provider
			break
		}
	}

	ctx.JSON(200, gitProvider)
}

// SetGitProvider 			godoc
//
//	@Tags			gitProvider
//	@Summary		Set Git provider
//	@Description	Set Git provider
//	@Param			gitProviderData	body	types.GitProvider	true	"Git provider"
//	@Produce		json
//	@Success		200
//	@Router			/gitprovider [post]
//
//	@id				SetGitProvider
func SetGitProvider(ctx *gin.Context) {
	var gitProviderData types.GitProvider
	var providerExists bool

	err := ctx.BindJSON(&gitProviderData)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid request body: %s", err.Error()))
		return
	}

	c, err := config.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get credentials: %s", err.Error()))
		return
	}

	for i, provider := range c.GitProviders {
		if provider.Id == gitProviderData.Id {
			c.GitProviders[i].Token = gitProviderData.Token
			c.GitProviders[i].Username = gitProviderData.Username
			c.GitProviders[i].BaseApiUrl = gitProviderData.BaseApiUrl
			providerExists = true
			break
		}
	}

	if !providerExists {
		c.GitProviders = append(c.GitProviders, types.GitProvider{
			Id:    gitProviderData.Id,
			Token: gitProviderData.Token,
		})
	}

	ctx.JSON(200, nil)
}

// RemoveGitProvider 			godoc
//
//	@Tags			gitProvider
//	@Summary		Remove Git provider
//	@Description	Remove Git provider
//	@Param			gitProviderId	path	string	true	"Git provider"
//	@Produce		json
//	@Success		200
//	@Router			/gitprovider/{gitProviderId} [delete]
//
//	@id				RemoveGitProvider
func RemoveGitProvider(ctx *gin.Context) {
	gitProviderId := ctx.Param("gitProviderId")

	c, err := config.GetConfig()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to get config: %s", err.Error()))
		return
	}

	var newProviders []types.GitProvider
	for _, provider := range c.GitProviders {
		if provider.Id != gitProviderId {
			newProviders = append(newProviders, provider)
		}
	}

	c.GitProviders = newProviders
	err = config.Save(c)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to save config: %s", err.Error()))
		return
	}

	ctx.JSON(200, nil)
}
