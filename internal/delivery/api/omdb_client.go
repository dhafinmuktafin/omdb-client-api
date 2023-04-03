package api

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/delivery/util"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
)

type OmdbClientAPI struct {
	Middleware *util.HttpMiddleware
	UseCase    usecase.OmdbClientUseCase
	Config     *config.Config
}

func New(middleware *util.HttpMiddleware, useCase usecase.OmdbClientUseCase, cfg *config.Config) *OmdbClientAPI {
	omdbClientAPI := &OmdbClientAPI{
		Middleware: middleware,
		UseCase:    useCase,
		Config:     cfg,
	}

	omdbClientAPI.RegisterRoutes()
	return omdbClientAPI
}

func (c *OmdbClientAPI) RegisterRoutes() {
	//register routes here

	//CARD
	c.Middleware.GET("/health", c.HealthCheck)
	c.Middleware.GET("/omdb-client/search", c.GetSearchMovieByKeyword)
	c.Middleware.GET("/omdb-client/movie/{id}", c.GetMovieDetailByMovieId)

}
