package omdbClient

import (
	"github.com/dhafinmuktafin/omdb-client-api/internal/model/core/config"
	"github.com/dhafinmuktafin/omdb-client-api/internal/repository"
)

//OmdbClient is a struct
type OmdbClient struct {
	Config    *config.Config
	MySQLRepo repository.MySQLRepository
	ExtAPI    repository.ExternalAPIRepo
}

//New returns struct of OmdbClient
func New(cfg *config.Config, repo *repository.RepositoryTypes) *OmdbClient {
	return &OmdbClient{
		Config:    cfg,
		MySQLRepo: repo.MySQL,
		ExtAPI:    repo.ExtAPI,
	}
}
