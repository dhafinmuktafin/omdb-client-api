package external_api

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/types"
	"net/http"
	"time"

	circuit "github.com/eapache/go-resiliency/breaker"
)

type ExternalAPI struct {
	Config         *config.Config
	Client         *http.Client
	CircuitBreaker map[int]*circuit.Breaker
}

func New(cfg *config.Config, pool *http.Client) *ExternalAPI {
	breakers := map[int]*circuit.Breaker{}
	for _, v := range types.ExtAPIBreakerList {
		breakers[v] = circuit.New(cfg.CircuitBreaker.ErrorThreshold, cfg.CircuitBreaker.SuccessThreshold, cfg.CircuitBreaker.Timeout*time.Second)
	}
	return &ExternalAPI{
		Config:         cfg,
		Client:         pool,
		CircuitBreaker: breakers,
	}
}
