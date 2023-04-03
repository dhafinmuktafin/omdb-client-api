package delivery

import (
	omdbClientAPI "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/api"
	"time"

	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/usecase"
	"github.com/gorilla/mux"

	httpUtil "github.com/dhafinmuktafin/omdb-client-api/internal/delivery/util"
)

type OmdbClientDelivery struct {
	OmdbUseCase   usecase.OmdbClientUseCase
	Config        *config.Config
	Middleware    *httpUtil.HttpMiddleware
	RouterWrapper struct {
		OmdbClientRouter *omdbClientAPI.OmdbClientAPI
	}
}

func Init(cfg *config.Config, router *mux.Router, useCase usecase.UseCase) *OmdbClientDelivery {
	cd := &OmdbClientDelivery{
		OmdbUseCase: useCase.OmdbClientUseCase,
		Config:      cfg,
		Middleware:  httpUtil.NewMiddleware(router, 5*time.Second),
	}
	return cd
}
